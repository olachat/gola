package structs

import (
	"sort"
	"strings"
)

func NewTableStruct(DBInfo *DBInfo, t Table, version string) *TableStruct {
	columns := make([]ColumnStruct, 0, len(t.Columns))
	ts := &TableStruct{DBInfo, t, nil, version}
	for _, c := range t.Columns {
		c.Comment = strings.ReplaceAll(c.Comment, "\r\n", " ")
		c.Comment = strings.ReplaceAll(c.Comment, "\n", " ")
		c.Comment = strings.ReplaceAll(c.Comment, "\"", "'")
		columns = append(columns, ColumnStruct{c, ts})
	}
	ts.sqlColumns = columns
	return ts
}

type TableStruct struct {
	dbinfo *DBInfo
	Table
	sqlColumns []ColumnStruct
	VERSION    string
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

func (t *TableStruct) SqlColumns() []ColumnStruct {
	return t.sqlColumns
}

func (t *TableStruct) Imports() string {
	packages := make(map[string]bool)
	for _, c := range t.sqlColumns {
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
