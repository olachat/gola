package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	"github.com/olachat/gola/structs"
)

//go:embed testdata
var fixtures embed.FS
var s *server.Server
var testDBPort int = 33066
var testDBName string = "testdb"
var testTables = []string{"blogs", "users"}
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
		_, err = db.Exec(string(query))
		if err != nil {
			panic(err.Error())
		}
	}
}

func getDB() *structs.DBInfo {
	var config mysqldriver.Config = map[string]interface{}{
		"dbname":    testDBName,
		"whitelist": "blogs",
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
	return db
}

type genMethod func(db *structs.DBInfo, t *structs.Table) map[string][]byte

func testGen(t *testing.T, wd string, gen genMethod, db *structs.DBInfo, table *structs.Table) {
	resultFiles := gen(db, table)
	expectedFileFolder := testDataPath + table.Name + string(filepath.Separator)

	if *update {
		for path, data := range resultFiles {
			pos := strings.LastIndex(path, string(filepath.Separator))
			expectedFileFolder = testDataPath + path[0:pos]
			os.Mkdir(expectedFileFolder, os.ModePerm)
			err := ioutil.WriteFile(testDataPath+path, data, 0644)
			if err != nil {
				panic(err)
			}
		}
	} else {
		for path, data := range resultFiles {
			expectedFilePath := testDataPath + path
			expectedFile, _ := fixtures.ReadFile(expectedFilePath)
			if diff := cmp.Diff(expectedFile, data); diff != "" {
				t.Error("file different: ", expectedFilePath)
				fmt.Println(diff)
			}
		}
	}
}

func TestCodeGen(t *testing.T) {
	db := getDB()

	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	for _, table := range db.Tables {
		testGen(t, wd, genORM, db, table)
	}
}

func TestIdx(t *testing.T) {
	db := getDB()

	// KEY `user` (`user_id`),
	// KEY `country_cate` (`country`, `category_id`, `is_vip`),
	// KEY `cate_pinned` (`category_id`, `is_pinned`, `is_vip`),
	// KEY `user_pinned_cate` (`user_id`, `is_pinned`, `category_id`),
	// UNIQUE KEY `slug` (`slug`)

	for _, tb := range db.Tables {
		if tb.Name != "blogs" {
			continue
		}

		if len(tb.Indexes) != 5 {
			t.Error("Failed to parse blogs table's 5 indexes")
		}

		for idxName, data := range tb.Indexes {
			switch idxName {
			case "user":
				if len(data) != 1 && data[0].Column_name != "user_id" {
					t.Error("Failed to parse blogs.user index")
				}

				if data[0].Non_unique != 1 {
					t.Error("Failed to parse blogs.user index unique")
				}
			case "slug":
				if len(data) != 1 && data[0].Column_name != "slug" {
					t.Error("Failed to parse blogs.slug index")
				}

				if data[0].Non_unique != 0 {
					t.Error("Failed to parse blogs.slug index unique")
				}
			case "user_pinned_cate":
				if len(data) != 3 {
					t.Error("Failed to parse blogs.user_pinned_cate index")
				}
				if data[0].Column_name != "user_id" {
					t.Error("Failed to parse blogs.user_pinned_cate index user_id column")
				}
				if data[1].Column_name != "is_pinned" {
					t.Error("Failed to parse blogs.user_pinned_cate index is_pinned column")
				}
				if data[2].Column_name != "category_id" {
					t.Error("Failed to parse blogs.user_pinned_cate index category_id column")
				}

			}
		}
	}
}
