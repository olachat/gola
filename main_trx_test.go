package main

import (
	"testing"
	"database/sql"
	"fmt"
	"context"
	"errors"
	"github.com/olachat/gola/util/mysql_util"
	_ "github.com/go-sql-driver/mysql"
)

/*
create database testdb;
use testdb;
CREATE TABLE `squareNum` (
	`number` int(11) unsigned NOT NULL DEFAULT '0',
	`squareNumber` int(11) unsigned NOT NULL DEFAULT '0'
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
*/

func TestTransaction(t *testing.T) {

	ctx := context.Background()

	//fmt.Println("-----------")

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/testdb")
	if err != nil {

		fmt.Println("-----------")

		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for inserting data
	//stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	//i := 2
	//_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}

	// Prepare statement for reading data
	//stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//defer stmtOut.Close()
	//
	//var squareNum int // we "scan" the result in here
	//
	//// Query the square-number of 2
	//err = stmtOut.QueryRow(2).Scan(&squareNum) // WHERE number = 2
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Printf("-----------The square number of 2 is: %d\n", squareNum)

	err = mysql_util.RunMultipleAsTrx(ctx, db, []mysql_util.TrxAction{

		// load pAccount
		func(ctx context.Context, trx *sql.Tx) (stopAndCommit bool, errInternal error) {

			fmt.Println("-----------start")

			rows, err := trx.Query("SELECT * FROM squareNum")
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			count := 0
			for rows.Next() {
				count++
			}

			fmt.Printf("-----------count 1: %d\n", count)

			stmtIns, err := trx.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

			//i := 3
			_, err = trx.Exec("INSERT INTO squareNum VALUES( 3, 9 )") // Insert tuples (i, i^2)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			fmt.Println("-----------end")

			//panic(errors.New("update error"))

			return false, errors.New("insert error")
		},
	})

	//fmt.Println("-----------2")

	if err != nil {
		print(err)
	}
}

