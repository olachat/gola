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

	"github.com/olachat/gola/v2/mysqldriver"
	"github.com/olachat/gola/v2/mysqlparser"
	"github.com/olachat/gola/v2/ormtpl"
	"github.com/olachat/gola/v2/structs"
)

func GenWithParser(config mysqlparser.MySQLParserConfig, output string) int {
	m := &mysqlparser.MySQLParser{}
	db, err := m.Assemble(config)
	if err != nil {
		panic(err)
	}

	if output == "" {
		output = "temp"
	}

	if !strings.HasPrefix(output, "/") {
		// output folder is relative path
		wd, err := os.Getwd()
		if err != nil {
			wd = "."
		}
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
				err = os.Mkdir(expectedFileFolder, os.ModePerm)
				if err != nil && os.IsNotExist(err) {
					println("Failed to create folder, please ensure " + output[:len(output)-1] + " exists")
					return 1
				}
				needMkdir = false
			}

			ioutil.WriteFile(output+path, data, 0644)
		}
	}

	files := genPackage(db)
	for path, data := range files {
		ioutil.WriteFile(output+path, data, 0644)
	}

	fmt.Printf("code generated in %s\n", output[:len(output)-1])
	return 0
}

/*
Run gola to perform code gen

`output`: output folder path
*/
func Run(config mysqldriver.DBConfig, output string) int {
	m := &mysqldriver.MySQLDriver{}
	db, err := m.Assemble(config)
	if err != nil {
		panic(err)
	}

	if output == "" {
		output = "temp"
	}

	if !strings.HasPrefix(output, "/") {
		// output folder is relative path
		wd, err := os.Getwd()
		if err != nil {
			wd = "."
		}
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
				err = os.Mkdir(expectedFileFolder, os.ModePerm)
				if err != nil && os.IsNotExist(err) {
					println("Failed to create folder, please ensure " + output[:len(output)-1] + " exists")
					return 1
				}
				needMkdir = false
			}

			ioutil.WriteFile(output+path, data, 0644)
		}
	}

	files := genPackage(db)
	for path, data := range files {
		ioutil.WriteFile(output+path, data, 0644)
	}

	fmt.Printf("code generated in %s\n", output[:len(output)-1])
	return 0
}

func genTPL(t ormtpl.TplStruct, tplName string) []byte {
	buf := bytes.NewBufferString("")
	t.SetVersion(VERSION)
	err := ormtpl.GetTpl(tplName).Execute(buf, t)
	if err != nil {
		panic(t.GetName() + " " + tplName +
			" genTpl error:\n" + err.Error())
	}
	return buf.Bytes()
}

func genPackage(db *structs.DBInfo) map[string][]byte {
	files := make(map[string][]byte)

	genFiles := map[string]string{
		"02_package.gogo": db.Schema + "_goladb.go",
	}

	for genTpl, genPath := range genFiles {
		data, err := formatBuffer(genTPL(db, genTpl))
		if err != nil {
			panic(db.Schema + " db code error:\n" + err.Error())
		}
		files[genPath] = data
	}

	return files
}

func genORM(t *structs.Table) map[string][]byte {
	files := make(map[string][]byte)

	tableFolder := t.Name + string(filepath.Separator)

	genFiles := map[string]string{
		"00_struct.gogo":         tableFolder + t.Name + ".go",
		"00_struct_ctx.gogo":     tableFolder + t.Name + "_ctx.go",
		"01_struct_idx.gogo":     tableFolder + t.Name + "_idx.go",
		"01_struct_idx_ctx.gogo": tableFolder + t.Name + "_idx_ctx.go",
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
