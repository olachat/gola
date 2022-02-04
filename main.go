package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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
	driverName := "mysql"

	var config drivers.Config = viper.GetStringMap(driverName)

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
