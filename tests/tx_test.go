package tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/golalib/testdata/worker"
)

func ExampleNewTxProvider() {

	prov := coredb.NewTxProvider("newdb")
	err := prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		_, err := tx.Exec("truncate table worker")
		panicOnErr(err)

		var workers []*worker.Worker
		err = tx.Find(&workers, "worker", "where id > ?", 0)
		panicOnErr(err)
		mustEqual(0, len(workers))
		fmt.Println("no of workers:", len(workers)) // uncomment to run test
		// Output: no of workers: 0

		_, err = tx.Exec("insert into worker (name,age) values (?, ?)", "peter", 18)
		panicOnErr(err)

		_, err = tx.Exec("insert into worker (name,age) values (?, ?)", "john", 28)
		panicOnErr(err)
		return err
	})
	panicOnErr(err)

	err = prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		var workers []*worker.Worker
		err := tx.Find(&workers, "worker", "where id > ?", 0)
		panicOnErr(err)
		mustEqual(2, len(workers))
		mustEqual("peter", workers[0].GetName())
		mustEqual(18, workers[0].GetAge())
		mustEqual("john", workers[1].GetName())
		mustEqual(28, workers[1].GetAge())

		var w worker.Worker
		err = tx.FindOne(&w, "worker", "where id = ?", 1)
		panicOnErr(err)
		mustEqual("peter", w.GetName())
		mustEqual(18, w.GetAge())

		r, err := tx.QueryInt("select count(1) from worker")
		panicOnErr(err)
		mustEqual(2, r)

		var workers2 []*worker.Worker
		err = tx.Query(&workers2, "select * from worker where id > ?", 0)
		panicOnErr(err)
		mustEqual(2, len(workers2))
		mustEqual("peter", workers2[0].GetName())
		mustEqual(18, workers2[0].GetAge())
		mustEqual("john", workers2[1].GetName())
		mustEqual(28, workers2[1].GetAge())
		return nil
	})

	prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		_, err := tx.Exec("insert into worker (name,age) values (?, ?)", "winson", 19)
		panicOnErr(err)

		return errors.New("abort")
	})

	prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		var w []*worker.Worker
		err := tx.Find(&w, "worker", "where id > ?", 0)
		panicOnErr(err)
		mustEqual(2, len(w))
		mustEqual("peter", w[0].GetName())
		mustEqual(18, w[0].GetAge())
		mustEqual("john", w[1].GetName())
		mustEqual(28, w[1].GetAge())
		return nil
	})
	panicOnErr(err)

	prov2 := coredb.NewTxProvider("newdb")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err = prov2.Tx(ctx, func(tx coredb.TxContext) error {
		_, err := tx.Exec("insert into worker (name,age) values (?, ?)", "winson", 19)
		if err != nil {
			return err
		}
		var w []*worker.Worker
		time.Sleep(10 * time.Millisecond)
		err = tx.Find(&w, "worker", "where age = ?", 28)
		if err != nil {
			return err
		}
		return nil
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		panic(err)
	}

}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
func mustEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic(fmt.Sprintf("%v != %v", a, b))
	}
}

func open() (db *sql.DB, err error) {
	dsn := "root:123456@tcp(127.0.0.1:3307)/newdb"
	if !strings.Contains(dsn, "?parseTime=true") {
		dsn += "?parseTime=true"
	}

	maxIdle := 3.0

	maxOpen := 50.0

	maxLifetime := 30.0

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(maxIdle) * time.Second)
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	db.SetMaxOpenConns(int(maxOpen))
	return
}
