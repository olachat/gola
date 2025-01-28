package txengine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/olachat/gola/v2/coredb"
)

type (
	TypedTx[T any] sql.Tx
	Tx             sql.Tx
)

// WithTypedTx converts *sql.Tx to *txengine.TypedTx[T].
//
// With *txengine.TypedTx[T], you have access to FindOne[T], Find[T] and Query[T]
func WithTypedTx[T any](tx *sql.Tx) *TypedTx[T] {
	return (*TypedTx[T])(tx)
}

// WithTx converts *sql.Tx to *txengine.Tx.
//
// With *txengine.Tx, you have access to Exec and QueryInt
func WithTx(tx *sql.Tx) *Tx {
	return (*Tx)(tx)
}

// RunTransaction runs a transaction
func RunTransaction(ctx context.Context, dbName string, fn func(ctx context.Context, sqlTx *sql.Tx) error) (err error) {
	tx, err := coredb.BeginTx(ctx, dbName, &coredb.DefaultTxOpts)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	err = runTransaction(ctx, tx, nil, fn)
	return
}

// RunTxWithRetry runs a transaction with retry logic on failure.
func RunTxWithRetry(ctx context.Context, dbName string, retryConfig coredb.RetryConfig, fn func(ctx context.Context, sqlTx *sql.Tx) error) (err error) {
	// Set defaults for invalid config
	if retryConfig.MaxRetries <= 0 {
		retryConfig.MaxRetries = coredb.DefaultRetryConfig.MaxRetries
	}

	if retryConfig.InitialBackoff <= 0 {
		retryConfig.InitialBackoff = coredb.DefaultRetryConfig.InitialBackoff
	}

	nonRetryableErrorFunc := retryConfig.IsNonRetryableErrorFunc
	if nonRetryableErrorFunc == nil {
		nonRetryableErrorFunc = coredb.IsNonRetryableError
	}

	var resultErr error
	retryCount := 0
	currentBackoff := retryConfig.InitialBackoff

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled during retry: %w", ctx.Err())
		default:
			resultErr = RunTransaction(ctx, dbName, fn)
			if resultErr == nil {
				return nil // Success!
			}

			if nonRetryableErrorFunc(resultErr) {
				log.Printf("Non-retryable error: %v", resultErr)
				return resultErr // Fail immediately for non-retryable errors
			}

			retryCount++
			if retryCount > retryConfig.MaxRetries {
				log.Printf("Max retries (%d) exceeded, last error: %v", retryConfig.MaxRetries, resultErr)
				return fmt.Errorf("max retries exceeded, last error: %w", resultErr)

			}

			delay := currentBackoff
			log.Printf("Retrying attempt %d with delay %v. Last error: %v", retryCount, delay, resultErr)
			time.Sleep(delay)
			currentBackoff *= 2
		}
	}
}

func runTransaction(ctx context.Context, tx *sql.Tx, conn *sql.Conn, fn func(ctx context.Context, sqlTx *sql.Tx) error) (err error) {
	if tx == nil && conn == nil {
		return errors.New("wrong usage. tx and conn cannot both be nil")
	}
	if tx == nil {
		tx, err = conn.BeginTx(ctx, &coredb.DefaultTxOpts)
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
	}
	defer func() {
		//nolint:gocritic
		if r := recover(); r != nil {
			errRollback := tx.Rollback()
			var ok bool
			errPanic, ok := r.(error)
			if !ok {
				errPanic = fmt.Errorf("%v", r)
			}
			err = errors.Join(err, errPanic, errRollback)
		} else if err != nil {
			errRollback := tx.Rollback()
			err = errors.Join(err, errRollback)
		} else {
			errCommit := tx.Commit()
			err = errors.Join(err, errCommit)
		}
	}()

	err = fn(ctx, tx)

	return
}

const lockTimeoutBuffer = 5 * time.Millisecond

// RunTransactionWithLock runs a transaction with a lock for durationInSec seconds
func RunTransactionWithLock(ctx context.Context, dbName string, lock string, durationInSec int, fn func(ctx context.Context, sqlTx *sql.Tx) error) (err error) {
	connCtx, cancel := context.WithTimeout(ctx, time.Duration(durationInSec)*time.Second+lockTimeoutBuffer)
	defer cancel()

	conn, err := coredb.Conn(connCtx, dbName, coredb.DBModeWrite)
	if err != nil {
		return fmt.Errorf("fail to get db connection: %w", err)
	}

	defer func() {
		if conn != nil {
			errCloseConn := conn.Close()
			if errCloseConn != nil {
				log.Printf("fail to close db connection: %#v", errCloseConn)
				err = errors.Join(err, errCloseConn)
			}
		}
	}()

	{
		var res int
		err = conn.QueryRowContext(ctx, "select get_lock(?,?)", lock, durationInSec).Scan(&res)
		if err != nil {
			return fmt.Errorf("get_lock failed: %w", err)
		}
		if res != 1 {
			return newLockError(lock, durationInSec)
		}
	}

	defer func() {
		var res int
		errRelease := conn.QueryRowContext(ctx, "select release_lock(?)", lock).Scan(&res)
		if errRelease != nil {
			err = errors.Join(err, fmt.Errorf("release_lock failed: %w", errRelease))
			return
		}
		if res != 1 {
			err = errors.Join(err, newReleaseLockError(lock, durationInSec))
		}
	}()

	err = runTransaction(ctx, nil, conn, fn)
	return
}

func newLockError(lock string, durationInSec int) error {
	return fmt.Errorf("fail to acquire lock: %s, durationInSec: %d", lock, durationInSec)
}

func newReleaseLockError(lock string, durationInSec int) error {
	return fmt.Errorf("fail to release lock: %s, durationInSec: %d", lock, durationInSec)
}

// FindOne returns a row from given table type with where query.
// If no rows found, *T will be nil. No error will be returned.
func (o *TypedTx[T]) FindOne(ctx context.Context, tableName string, whereSQL string, params ...any) (*T, error) {
	u := new(T)
	columnsNames := coredb.GetColumnsNames[T]()
	data := coredb.StrutForScan(u)
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	err2 := (*sql.Tx)(o).QueryRowContext(ctx, query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err2
	}

	return u, nil
}

// Exec given query with given db info & params
func (o *Tx) Exec(ctx context.Context, query string, params ...any) (sql.Result, error) {
	return (*sql.Tx)(o).ExecContext(ctx, query, params...)
}

// Find returns rows from given table type with where query
func (o *TypedTx[T]) Find(ctx context.Context, tableName string, whereSQL string, params ...any) ([]*T, error) {
	columnsNames := coredb.GetColumnsNames[T]()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return o.Query(ctx, query, params...)
}

// QueryInt single int result by query, handy for count(*) querys
func (o *Tx) QueryInt(ctx context.Context, query string, params ...any) (result int, err error) {
	err = (*sql.Tx)(o).QueryRowContext(ctx, query, params...).Scan(&result)
	return
}

// Query rows from given table type with where query & params
func (o *TypedTx[T]) Query(ctx context.Context, query string, params ...any) (result []*T, err error) {
	var rows *sql.Rows
	rows, err = (*sql.Tx)(o).QueryContext(ctx, query, params...)
	if err != nil {
		return
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Printf("Gola: TypedTx.Query: failed to close rows: %v", errClose)
		}
	}()

	var u *T
	for rows.Next() {
		u = new(T)
		data := coredb.StrutForScan(u)
		err = rows.Scan(data...)
		if err != nil {
			return
		}
		result = append(result, u)
	}

	if rows.Err() != nil {
		err = rows.Err()
	}

	return
}
