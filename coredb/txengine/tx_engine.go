package txengine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/olachat/gola/v2/coredb"
)

type TypedTx[T any] sql.Tx
type Tx sql.Tx

func WithTypedTx[T any](tx *sql.Tx) *TypedTx[T] {
	return (*TypedTx[T])(tx)
}

func WithTx(tx *sql.Tx) *Tx {
	return (*Tx)(tx)
}

func StartTx(ctx context.Context, tx *sql.Tx, fn func(ctx context.Context, sqlTx *sql.Tx) error) (err error) {
	defer func() {
		//nolint:gocritic
		if r := recover(); r != nil {
			_ = tx.Rollback()
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		} else if err != nil {
			errRollback := tx.Rollback()
			if errors.Is(errRollback, sql.ErrTxDone) && ctx.Err() != nil {
				errRollback = nil
			}
			if errRollback != nil {
				err = fmt.Errorf("%v encountered. but rollback failed: %w", err, errRollback)
			}
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx, tx)

	return err
}

// FindOne returns a row from given table type with where query.
// If no rows found, *T will be nil. No error will be returned.
func (o *TypedTx[T]) FindOne(ctx context.Context, tableName string, where coredb.WhereQuery) (*T, error) {
	u := new(T)
	columnsNames := coredb.GetColumnsNames[T]()
	data := coredb.StrutForScan(u)
	whereSQL, params := where.GetWhere()
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
func (o *TypedTx[T]) Find(ctx context.Context, tableName string, where coredb.WhereQuery) ([]*T, error) {
	columnsNames := coredb.GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
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
	rows, err := (*sql.Tx)(o).QueryContext(ctx, query, params...)
	if err != nil {
		return
	}

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
