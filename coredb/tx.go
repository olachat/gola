package coredb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// BeginTx returns a custom db.Tx based on opts. This method exists for flexibility.
// Make sure you call Commit or Rollback on the returned Tx.
// Refer to https://go.dev/doc/database/execute-transactions on how to use the returned Tx.
func BeginTx(ctx context.Context, dbname string, opts *sql.TxOptions) (tx *sql.Tx, err error) {
	mydb := getDB(dbname, DBModeWrite)
	return mydb.BeginTx(ctx, opts)
}

// DefaultTxOpts is package variable with default transaction level
var DefaultTxOpts = sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}

// TxContext interface for DAO operations with context.
type TxContext interface {
	// Exec executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	Exec(query string, args ...any) (sql.Result, error)

	// Query executes a SELECT query and scans the resulting rows into the provided 'results' destination.
	// It accepts a SQL query and an optional list of parameters for placeholder substitution.
	// NOTE: results must be a pointer to a slice of struct pointers.
	Query(results any, query string, params ...any) (err error)

	// QueryInt executes a SELECT query expected to return a single integer value.
	// Commonly used for COUNT(*) operations or where the result is inherently an integer.
	// Multiple params for query placeholders are supported.
	QueryInt(query string, params ...any) (result int, err error)

	// FindOne fetches a single record from the database and populates 'result'.
	// It requires the name of the table, an optional WHERE clause ('whereSQL'), and
	// parameters to substitute into the WHERE clause's placeholders.
	// NOTE: result must be a non-nil pointer to a struct.
	FindOne(result any, tableName string, whereSQL string, params ...any) error

	// Find executes a SELECT query based on the given 'tableName' and 'whereSQL',
	// placing all matching records into the 'results' slice.
	// Parameters for the WHERE clause's placeholders can be passed with 'params'.
	// NOTE: results must be a pointer to a slice of struct pointers.
	Find(results any, tableName string, whereSQL string, params ...any) error
}

// tx represents transaction with context as inner object.
type tx struct {
	ctx context.Context //nolint:containedctx
	Tx  *sql.Tx
}

// Exec executes query with params.
func (t *tx) Exec(query string, params ...any) (sql.Result, error) {
	return t.Tx.ExecContext(t.ctx, query, params...)
}

// Query loads data from db.
func (t *tx) Query(results any, query string, params ...any) error {
	rows, err := t.Tx.QueryContext(t.ctx, query, params...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return RowsToStructSliceReflect(rows, results)
}

func (t *tx) QueryInt(query string, params ...any) (result int, err error) {
	err = t.Tx.QueryRowContext(t.ctx, query, params...).Scan(&result)
	return
}

func (t *tx) FindOne(result any, tableName string, whereSQL string, params ...any) error {
	columnsNames := GetColumnsNamesReflect(result)
	data := StrutForScan(result)
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	err2 := t.Tx.QueryRowContext(t.ctx, query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 == sql.ErrNoRows {
			return nil
		}
		return err2
	}

	return nil
}

func (t *tx) Find(results any, tableName string, whereSQL string, params ...any) error {
	columnsNames := GetColumnsNamesReflect(results)
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	return t.Query(results, query, params...)
}

// Commit this transaction.
func (t *tx) Commit() error {
	return t.Tx.Commit()
}

// Rollback cancel this transaction.
func (t *tx) Rollback() error {
	return t.Tx.Rollback()
}

// Connector for sql database.
type Connector interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// TxProvider ...
type TxProvider struct {
	conn Connector
}

// NewTxProvider ...
func NewTxProvider(dbname string) *TxProvider {
	mydb := getDB(dbname, DBModeWrite)
	return &TxProvider{
		conn: mydb,
	}
}

// acquireWithOpts transaction from db
func (t *TxProvider) acquireWithOpts(ctx context.Context, opts *sql.TxOptions) (*tx, error) {
	trx, err := t.conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &tx{
		ctx: ctx,
		Tx:  trx,
	}, nil
}

// TxWithOpts ...
func (t *TxProvider) TxWithOpts(ctx context.Context, fn func(TxContext) error, opts *sql.TxOptions) (err error) {
	var trx *tx
	trx, err = t.acquireWithOpts(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		//nolint:gocritic
		if r := recover(); r != nil {
			_ = trx.Rollback()
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		} else if err != nil {
			errRollback := trx.Rollback()
			if errors.Is(errRollback, sql.ErrTxDone) && ctx.Err() != nil {
				errRollback = nil
			}
			if errRollback != nil {
				err = fmt.Errorf("%v encountered. but rollback failed: %w", err, errRollback)
			}
		} else {
			err = trx.Commit()
		}
	}()

	err = fn(trx)

	return err
}

// Tx runs fn in transaction.
func (t *TxProvider) Tx(ctx context.Context, fn func(TxContext) error) error {
	return t.TxWithOpts(ctx, fn, &DefaultTxOpts)
}
