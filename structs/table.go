package structs

import (
	"sort"
	"strings"
)

func NewTableStruct(DBInfo *DBInfo, t Table, version string) *TableStruct {
	ts := &TableStruct{DBInfo, t, version}
	columns := make([]Column, len(t.Columns))
	for i, c := range t.Columns {
		c.Comment = strings.ReplaceAll(c.Comment, "\r\n", " ")
		c.Comment = strings.ReplaceAll(c.Comment, "\n", " ")
		c.Comment = strings.ReplaceAll(c.Comment, "\"", "'")
		c.table = ts
		columns[i] = c
	}

	ts.Columns = columns
	return ts
}

type TableStruct struct {
	dbinfo *DBInfo
	Table
	VERSION string
}

func (t *TableStruct) Package() string {
	return t.Name
}

func (t *TableStruct) ClassName() string {
	name := getGoName(t.Name)
	if strings.HasSuffix(name, "s") {
		return name[:len(name)-1]
	}

	return name
}

func (t *TableStruct) Imports() string {
	packages := make(map[string]bool)
	for _, c := range t.Columns {
		if strings.Contains(c.SQLType(), "Time") {
			packages[`"time"`] = true
		}
	}

	imports := []string{}
	for p := range packages {
		imports = append(imports, p)
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] > imports[j]
	})

	return strings.Join(imports, "\n")
}
