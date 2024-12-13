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
		err = rows.Close()
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

	return
}
