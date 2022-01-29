package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/testdata"
	"github.com/olachat/gola/util/mysql_util"
	"testing"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/mysqldriver"
)

//var fixtures embed.FS
//var s *server.Server
//
var testDbPort2 int = 33068

//
//var testDBName string = "testdb"
//var testTables = []string{"users"}
//var testDataPath = "testdata" + string(filepath.Separator)

//var update = flag.Bool("update", false, "update generated files")

var testDBPort int = 33069
var testDBName string = "testdb"
var testTables = []string{"users"}
var tableName string = "users"

var db *sql.DB

func init() {
	corelib.Setup(fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))

	engine := sqle.NewDefault(gsql.NewDatabaseProvider(
		memory.NewDatabase(testDBName),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", testDBPort),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()

	connStr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDbPort2, "false")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	//create table
	query, _ := testdata.Fixtures.ReadFile(tableName + ".sql")
	db.Exec(string(query))

	//add data
	db.Exec(`
insert into users (name, email, created_at, updated_at) values
("John Doe", "john@doe.com", NOW(), NOW()),
("John Doe", "johnalt@doe.com", NOW(), NOW()),
("Jane Doe", "jane@doe.com", NOW(), NOW()),
("Evil Bob", "evilbob@gmail.com", NOW(), NOW())
	`)
}

func TestTransaction(t *testing.T) {

	ctx := context.Background()

	fmt.Println("-----------")

	err := mysql_util.RunMultipleAsTrx(ctx, db, []mysql_util.TrxAction{

		// load pAccount
		func(ctx context.Context, trx *sql.Tx) (stopAndCommit bool, errInternal error) {

			fmt.Println("-----------1")
			//insertStat := `
			//insert into users (name, email, created_at, updated_at) values
			//("John Doe", "john@doe.com", NOW(), NOW()),
			//("John Doe", "johnalt@doe.com", NOW(), NOW()),
			//("Jane Doe", "jane@doe.com", NOW(), NOW()),
			//("Evil Bob", "evilbob@gmail.com", NOW(), NOW())
			//	`

			updateStatement := `
			update users set (email) values
			("yan.feng@olaola.chat")
			where id = 1;
			)
			`

			//db.Exec(insertStat)
			exec, err := db.Exec(updateStatement)
			if err != nil {
				return false, err
			}

			fmt.Println(exec.LastInsertId())
			fmt.Println(exec.RowsAffected())

			panic(errors.New("update error"))

			return
		},
	})

	fmt.Println("-----------2")

	if err != nil {
		print(err)
	}
}
