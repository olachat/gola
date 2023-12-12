package tests

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/golalib/testdata/worker"
)

func TestBeginTx(t *testing.T) {
	as := assert.New(t)

	prov := coredb.NewTxProvider("newdb")
	err := prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		_, err := tx.Exec("truncate table worker")
		as.Nil(err)

		var workers []*worker.Worker
		err = tx.Find(&workers, "worker", "where id > ?", 0)
		as.Nil(err)
		as.Equal(0, len(workers))

		_, err = tx.Exec("insert into worker (name,age) values (?, ?)", "peter", 18)
		as.Nil(err)

		_, err = tx.Exec("insert into worker (name,age) values (?, ?)", "john", 28)
		as.Nil(err)
		return err
	})
	as.Nil(err)

	err = prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		var workers []*worker.Worker
		err := tx.Find(&workers, "worker", "where id > ?", 0)
		as.Nil(err)
		as.Equal(2, len(workers))
		as.Equal("peter", workers[0].GetName())
		as.Equal(18, workers[0].GetAge())
		as.Equal("john", workers[1].GetName())
		as.Equal(28, workers[1].GetAge())

		var w worker.Worker
		err = tx.FindOne(&w, "worker", "where id = ?", 1)
		as.Nil(err)
		as.Equal("peter", w.GetName())
		as.Equal(18, w.GetAge())

		r, err := tx.QueryInt("select count(1) from worker")
		as.Nil(err)
		as.Equal(2, r)

		var workers2 []*worker.Worker
		err = tx.Query(&workers2, "select * from worker where id > ?", 0)
		as.Nil(err)
		as.Equal(2, len(workers2))
		as.Equal("peter", workers2[0].GetName())
		as.Equal(18, workers2[0].GetAge())
		as.Equal("john", workers2[1].GetName())
		as.Equal(28, workers2[1].GetAge())
		return nil
	})

	prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		_, err := tx.Exec("insert into worker (name,age) values (?, ?)", "winson", 19)
		as.Nil(err)

		return errors.New("abort")
	})

	prov.Tx(context.Background(), func(tx coredb.TxContext) error {
		var w []*worker.Worker
		err := tx.Find(&w, "worker", "where id > ?", 0)
		as.Nil(err)
		as.Equal(2, len(w))
		as.Equal("peter", w[0].GetName())
		as.Equal(18, w[0].GetAge())
		as.Equal("john", w[1].GetName())
		as.Equal(28, w[1].GetAge())
		return nil
	})
	as.Nil(err)

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
