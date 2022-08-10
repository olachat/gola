package golalib

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
)

/*
	Run gola to perform code gen

`output`: output file path
*/
func Run(config mysqldriver.DBConfig, output string) {
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	m := &mysqldriver.MySQLDriver{}
	db, err := m.Assemble(config)
	if err != nil {
		panic(err)
	}

	if output == "" {
		output = "temp"
	}

	if !strings.HasPrefix(output, "/") {
		output = wd + string(filepath.Separator) + output
	}

	if !strings.HasSuffix(output, string(filepath.Separator)) {
		output = output + string(filepath.Separator)
	}

	for _, t := range db.Tables {
		if len(t.GetPKColumns()) == 0 {
			println(t.Name + " doesn't have primay key")
			continue
		}

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

func genTPL(t *structs.Table, tplName string) []byte {
	buf := bytes.NewBufferString("")
	t.VERSION = VERSION
	err := ormtpl.GetTpl(tplName).Execute(buf, t)
	if err != nil {
		panic(t.Name + " " + tplName +
			" genTpl error:\n" + err.Error())
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
		data, err := formatBuffer(genTPL(t, genTpl))
		if err != nil {
			panic(t.Name + " code error:\n" + err.Error())
		}
		files[genPath] = data
	}

	return files
}

var (
	rgxSyntaxError = regexp.MustCompile(`(\d+):\d+: `)
)

func formatBuffer(buf []byte) ([]byte, error) {
	output, err := format.Source(buf)
	if err == nil {
		return output, nil
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

	return nil, fmt.Errorf("failed to format template\n\n%s", errBuf.Bytes())
}
