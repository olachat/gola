package golalib

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dolthub/go-mysql-server/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/olachat/gola/v2/mysqlparser"
	"github.com/olachat/gola/v2/ormtpl"
	"github.com/olachat/gola/v2/structs"
)

//go:embed testdata
var fixtures embed.FS
var s *server.Server
var testDBPort int = 33066
var testDBName string = "testdata"
var testTables = []string{
	"blogs",
	"users", "songs", "song_user_favourites",
	"profile", "account",
	"room",
	"gifts", "gifts_nn",
	"gifts_with_default", "gifts_nn_with_default",
}
var testDataPath = "testdata" + string(filepath.Separator)

var update = flag.Bool("update", false, "update generated files")

func getDB() *structs.DBInfo {
	c := mysqlparser.MySQLParserConfig{}
	c.DbName = testDBName
	c.TableCreateSQLs = make([]mysqlparser.TableCreateSQL, 0, len(testTables))
	for _, tableName := range testTables {
		query, _ := fixtures.ReadFile(testDataPath + tableName + ".sql")
		tableSQL := mysqlparser.TableCreateSQL{
			Table:     tableName,
			CreateSQL: string(query),
		}
		c.TableCreateSQLs = append(c.TableCreateSQLs, tableSQL)
	}
	m := &mysqlparser.MySQLParser{}
	dbInfo, err := m.Assemble(c)
	if err != nil {
		panic(err.Error())
	}
	return dbInfo
}

type genMethod func(t ormtpl.TplStruct) map[string][]byte

func testGen(t *testing.T, wd string, gen genMethod, data ormtpl.TplStruct) {
	resultFiles := gen(data)

	if *update {
		for path, data := range resultFiles {
			pos := strings.LastIndex(path, string(filepath.Separator))
			if pos > 0 {
				expectedFileFolder := testDataPath + path[0:pos]
				os.Mkdir(expectedFileFolder, os.ModePerm)
			}

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
	// db := getDB_old()
	// pp.Println(db)
	// isSame := reflect.DeepEqual(db, db2)
	// if !isSame {
	// 	pp.Println(db)
	// 	pp.Println(db2)
	// }
	// return

	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	for _, table := range db.Tables {
		testGen(t, wd, func(t ormtpl.TplStruct) map[string][]byte {
			return genORM(t.(*structs.Table))
		}, table)
	}

	testGen(t, wd, func(t ormtpl.TplStruct) map[string][]byte {
		return genPackage(t.(*structs.DBInfo))
	}, db)
}

func getOne[T any](objs []T, filter func(obj T) bool) T {
	for _, obj := range objs {
		if filter(obj) {
			return obj
		}
	}

	return *new(T)
}

func TestIdx(t *testing.T) {
	db := getDB()

	// KEY `user` (`user_id`),
	// KEY `country_cate` (`country`, `category_id`, `is_vip`),
	// KEY `cate_pinned` (`category_id`, `is_pinned`, `is_vip`),
	// KEY `user_pinned_cate` (`user_id`, `is_pinned`, `category_id`),
	// UNIQUE KEY `slug` (`slug`)

	tb := getOne(db.Tables, func(tb *structs.Table) bool {
		return tb.Name == "blogs"
	})
	if len(tb.Indexes) != 7 {
		t.Error("Failed to parse blogs table's 7 indexes")
	}

	data := tb.Indexes["user"]
	if len(data) != 1 && data[0].ColumnName != "user_id" {
		t.Error("Failed to parse blogs.user index")
	}
	if data[0].NonUnique != 1 {
		t.Error("Failed to parse blogs.user index unique")
	}

	data = tb.Indexes["slug"]
	if len(data) != 1 && data[0].ColumnName != "slug" {
		t.Error("Failed to parse blogs.slug index")
	}
	if data[0].NonUnique != 0 {
		t.Error("Failed to parse blogs.slug index unique")
	}

	data = tb.Indexes["user_pinned_cate"]
	if len(data) != 3 {
		t.Error("Failed to parse blogs.user_pinned_cate index")
	}
	if data[0].ColumnName != "user_id" || data[1].ColumnName != "is_pinned" || data[2].ColumnName != "category_id" {
		t.Error("Failed to parse blogs index column names")
	}
}

func TestIdx2(t *testing.T) {
	db := getDB()

	for _, tb := range db.Tables {
		if tb.Name != "blogs" {
			continue
		}

		println(tb.GetIndexRoot().String(""))
		nodes := tb.GetIndexNodes()

		for _, n := range nodes {
			fmt.Printf("%s[%d] %s\n", n.GoName(), n.Order, n.InterfaceName())
		}

	}
}
