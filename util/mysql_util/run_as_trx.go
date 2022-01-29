package mysql_util

import (
	"context"
	"database/sql"
	"fmt"
)

func RunAsTrx(ctx context.Context, db *sql.DB, queryFunc func(ctx context.Context, trx *sql.Tx) error) (err error) {

	if ctx == nil || db == nil || queryFunc == nil {
		err = fmt.Errorf("mysql_util: RunAsTrx: ctx == nil || db == nil || queryFunc == nil: %v %v %p",
			ctx, db, queryFunc)
		return
	}

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			errRollback := trx.Rollback()
			err = fmt.Errorf("mysql_util: RunAsTrx: rollback: panic=%#v, errRollback=%#v", r, errRollback)
		} else if err != nil {
			errRollback := trx.Rollback()
			err = fmt.Errorf("mysql_util: RunAsTrx: err is not nil: err=%#v, errRollback=%#v", err, errRollback)
		}
	}()

	err = queryFunc(ctx, trx)
	if err != nil {
		err = trx.Commit()
	}
	return
}

type TrxAction func(ctx context.Context, trx *sql.Tx) (commitAndStop bool, err error)

func RunMultipleAsTrx(ctx context.Context, db *sql.DB, queryFuncs []TrxAction) (err error) {

	if ctx == nil || db == nil || len(queryFuncs) == 0 {
		err = fmt.Errorf("mysql_util: RunAsTrx: ctx == nil || db == nil || len(queryFuncs) == 0: %v %v %v",
			ctx, db, queryFuncs)
		return
	}

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			errRollback := trx.Rollback()
			err = fmt.Errorf("mysql_util: RunAsTrx: rollback: panic=%#v, errRollback=%#v", r, errRollback)
		} else if err != nil {
			errRollback := trx.Rollback()
			err = fmt.Errorf("mysql_util: RunAsTrx: err is not nil: err=%w, errRollback=%#v", err, errRollback)
		}
	}()

	stopAndCommit := false
	for _, queryFunc := range queryFuncs {
		stopAndCommit, err = queryFunc(ctx, trx)
		if err != nil {
			return // any of error will trigger a Rollback
		}
		if stopAndCommit {
			break
		}
	}

	err = trx.Commit()
	return
}
