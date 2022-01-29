package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/olachat/gola/mysqldriver"
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

//go:embed testdata
var fixtures embed.FS
var s *server.Server

var testDBPort1 int = 33067

//var testDBName string = "testdb"
//var testTables = []string{"users"}
var testDataPath = "testdata" + string(filepath.Separator)

var update = flag.Bool("update", false, "update generated files")

// init the database with tables based on .sql files in the testdb folder
func init() {
	engine := sqle.NewDefault(gsql.NewDatabaseProvider(
		memory.NewDatabase(testDBName),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", testDBPort1),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err = server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()

	connStr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort1, "false")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	for _, tableName := range testTables {
		query, _ := fixtures.ReadFile(testDataPath + tableName + ".sql")
		_, err = db.Exec(string(query))
		if err != nil {
			panic(err.Error())
		}
	}
}

type genMethod func(db *drivers.DBInfo, t drivers.Table) []byte

func testGen(t *testing.T, wd string, gen genMethod, db *drivers.DBInfo, table drivers.Table, extName string) {
	resultFile := gen(db, table)
	expectedFileFolder := testDataPath + table.Name + string(filepath.Separator)
	expectedFilePath := expectedFileFolder + table.Name + "." + extName

	if *update {
		os.Mkdir(expectedFileFolder, os.ModePerm)
		err := ioutil.WriteFile(expectedFilePath, resultFile, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		expectedFile, _ := fixtures.ReadFile(expectedFilePath)
		if diff := cmp.Diff(resultFile, expectedFile); diff != "" {
			t.Error("file different: ", expectedFilePath)
			fmt.Println(diff)
		}
	}
}

func TestCodeGen(t *testing.T) {
	var config drivers.Config = map[string]interface{}{
		"dbname":    testDBName,
		"whitelist": testTables,
		"host":      "localhost",
		"port":      testDBPort1,
		"user":      "root",
		"pass":      "",
		"sslmode":   "false",
	}

	m := &mysqldriver.MySQLDriver{}
	db, err := m.Assemble(config)
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	for _, table := range db.Tables {
		testGen(t, wd, genORM, db, table, "go")
	}
}
