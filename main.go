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

	"github.com/olachat/gola/mysqldriver"
	"github.com/olachat/gola/ormtpl"
	"github.com/olachat/gola/structs"

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

	var config mysqldriver.Config = viper.GetStringMap(driverName)

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
		switch gentype {
		case "orm":
			files := genORM(t)
			needMkdir := true
			for path, data := range files {
				if needMkdir {
					pos := strings.LastIndex(path, string(filepath.Separator))
					expectedFileFolder := output + path[0:pos]
					os.Mkdir(expectedFileFolder, os.ModePerm)
					needMkdir = false
				}

				ioutil.WriteFile(output+path, data, 0644)
			}
		}
	}
}

func genTPL(t *structs.Table, tplName string) []byte {
	buf := bytes.NewBufferString("")
	t.VERSION = VERSION
	err := ormtpl.GetTpl(tplName).Execute(buf, t)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func genORM(t *structs.Table) map[string][]byte {
	files := make(map[string][]byte)

	tableFolder := t.Name + string(filepath.Separator)

	genFiles := map[string]string{
		"00_struct.gogo":     tableFolder + t.Name + ".go",
		"01_struct_idx.gogo": tableFolder + t.Name + "_idx.go",
	}

	for genTpl, genPath := range genFiles {
		data := formatBuffer(genTPL(t, genTpl))
		files[genPath] = data
	}

	return files
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
