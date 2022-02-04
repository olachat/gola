package mysql_util

import (
	"context"
	"database/sql"
	"fmt"
	"errors"
	//"github.com/davecgh/go-spew/spew"
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

	//fmt.Println("-----------1.5")

	if ctx == nil || db == nil || len(queryFuncs) == 0 {
		err = fmt.Errorf("mysql_util: RunAsTrx: ctx == nil || db == nil || len(queryFuncs) == 0: %v %v %v",
			ctx, db, queryFuncs)

		fmt.Println(ctx)
		fmt.Println(db)
		fmt.Println(len(queryFuncs))

		return
	}

	//fmt.Println("-----------1.6")

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	//fmt.Println("-----------1.7")

	defer func() {
		if r := recover(); r != nil {
			errRollback := trx.Rollback()

			fmt.Println("-----------roll back 1------")

			rows, err := trx.Query("SELECT * FROM squareNum")
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			if rows != nil {
				count := 0
				for rows.Next() {
					count++
				}

				fmt.Printf("-----------count 4: %d\n", count)
			}


			err = fmt.Errorf("mysql_util: RunAsTrx: rollback: panic=%#v, errRollback=%#v", r, errRollback)

			fmt.Println(err)
		} else if err != nil {
			errRollback := trx.Rollback()

			fmt.Println("-----------roll back 2------")

			rows, _ := trx.Query("SELECT * FROM squareNum")
			//if err != nil {
			//	panic(err.Error()) // proper error handling instead of panic in your app
			//}

			if rows != nil {
				count := 0
				for rows.Next() {
					count++
				}

				fmt.Printf("-----------count 3: %d\n", count)
			}

			err = fmt.Errorf("mysql_util: RunAsTrx: err is not nil: err=%w, errRollback=%#v", err, errRollback)

			fmt.Println(err)
		}
	}()

	stopAndCommit := false
	for _, queryFunc := range queryFuncs {

		//fmt.Println(key)

		stopAndCommit, err = queryFunc(ctx, trx)
		if err != nil {

			fmt.Println("-----------roll back------")

			rows, err := trx.Query("SELECT * FROM squareNum")
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			if rows != nil {
				count := 0
				for rows.Next() {
					count++
				}

				fmt.Printf("-----------count 2: %d\n", count)
			}

			//spew.Dump(rows)

			err = errors.New("intentional error")

			return err// any of error will trigger a Rollback
		}
		if stopAndCommit {
			break
		}
	}

	err = trx.Commit()
	return
}
