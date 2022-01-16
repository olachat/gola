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

//go:embed testdata/*
var fixtures embed.FS
var s *server.Server
var testDBPort int = 33066
var testDBName string = "testdb"
var testTables = []string{"users"}
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
		Address:  fmt.Sprintf("localhost:%d", testDBPort),
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}
	var err error

	s, err = server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()

	connStr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort, "false")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	for _, tableName := range testTables {
		query, _ := fixtures.ReadFile(testDataPath + tableName + ".sql")

		stmt, err := db.Prepare(string(query))
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec()
		if err != nil {
			panic(err)
		}
	}
}

type genMethod func(db *drivers.DBInfo, t drivers.Table) []byte

func testGen(t *testing.T, wd string, gen genMethod, db *drivers.DBInfo, table drivers.Table, extName string) {
	resultFile := gen(db, table)
	if *update {
		ioutil.WriteFile(wd+string(filepath.Separator)+testDataPath+table.Name+"."+extName, resultFile, 0644)
	} else {
		expectedFile, _ := fixtures.ReadFile(testDataPath + table.Name + "." + extName)
		if diff := cmp.Diff(resultFile, expectedFile); diff != "" {
			t.Error("file different: ", table.Name+"."+extName)
			fmt.Println(diff)
		}
	}
}

func TestCodeGen(t *testing.T) {
	var config drivers.Config = map[string]interface{}{
		"dbname":    testDBName,
		"whitelist": testTables,
		"host":      "localhost",
		"port":      testDBPort,
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
