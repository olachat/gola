package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/testdata"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/olachat/gola/dolttpl"
	"github.com/olachat/gola/mysqldriver"
	"github.com/olachat/gola/structs"
	"github.com/volatiletech/sqlboiler/v4/drivers"

	"github.com/spf13/viper"
)

//var testDBPort int = 33069
//var testDBName string = "testdb"
//var testTables = []string{"users"}
//var tableName string = "users"

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

	connStr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort, "false")
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

func main() {
	viper.SetConfigName("gola")
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	configPaths := []string{wd}
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	viper.ReadInConfig()
	viper.AutomaticEnv()
	//driverName := "mysql"

	//var config drivers.Config = viper.GetStringMap(driverName)

	//var fixtures embed.FS
	//var s *server.Server
	//var testDataPath = "testdata" + string(filepath.Separator)

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

	output := config.DefaultString("output", "temp")
	gentype := config.DefaultString("gentype", "orm")

	if !strings.HasPrefix(output, "/") {
		output = wd + string(filepath.Separator) + output
	}

	if !strings.HasSuffix(output, string(filepath.Separator)) {
		output = output + string(filepath.Separator)
	}

	for _, t := range db.Tables {
		println(t.Name)
		println(gentype)
		switch gentype {
		case "orm":
			ioutil.WriteFile(output+t.Name+".go", genORM(db, t), 0644)
		}
	}
}

func genTPL(db *drivers.DBInfo, t drivers.Table, tplName string) []byte {
	buf := bytes.NewBufferString("")

	const VERSION string = "0.0.1"

	err := dolttpl.GetTpl(tplName).Execute(buf, structs.NewTableStruct(db, t, VERSION))
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func genORM(db *drivers.DBInfo, t drivers.Table) []byte {
	return formatBuffer(genTPL(db, t, "00_struct.gogo"))
}

var (
	rgxSyntaxError = regexp.MustCompile(`(\d+):\d+: `)
)

func formatBuffer(buf []byte) []byte {
	output, err := format.Source(buf)
	if err == nil {
		return output
	}

	matches := rgxSyntaxError.FindStringSubmatch(err.Error())
	if matches == nil {
		panic(errors.New("failed to format template: " + err.Error()))
	}

	lineNum, _ := strconv.Atoi(matches[1])
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	errBuf := &bytes.Buffer{}
	line := 1
	for ; scanner.Scan(); line++ {
		if delta := line - lineNum; delta < -5 || delta > 5 {
			continue
		}

		if line == lineNum {
			errBuf.WriteString(">>>> ")
		} else {
			fmt.Fprintf(errBuf, "% 4d ", line)
		}
		errBuf.Write(scanner.Bytes())
		errBuf.WriteByte('\n')
	}

	panic(fmt.Errorf("failed to format template\n\n%s", errBuf.Bytes()))
}
