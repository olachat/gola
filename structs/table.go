package structs

import (
	"sort"
	"strings"
)

type Table struct {
	dbinfo *DBInfo
	Name   string `json:"name"`
	// For dbs with real schemas, like Postgres.
	// Example value: "schema_name"."table_name"
	SchemaName string   `json:"schema_name"`
	Columns    []Column `json:"columns"`

	PKey  *PrimaryKey  `json:"p_key"`
	FKeys []ForeignKey `json:"f_keys"`

	IsJoinTable bool `json:"is_join_table"`
	Indexes     map[string][]*IndexDesc
	VERSION     string
	idxRoot     *idxNode
}

func (t *Table) Package() string {
	return t.Name
}

func (t *Table) ClassName() string {
	name := getGoName(t.Name)
	if strings.HasSuffix(name, "s") {
		return name[:len(name)-1]
	}

	return name
}

func (t *Table) GetIndexRoot() *idxNode {
	if t.idxRoot != nil {
		return t.idxRoot
	}

	root := &idxNode{
		ColName: "",
	}

	for _, items := range t.Indexes {
		node := root
		for _, item := range items {
			node = node.GetChildren(item.Column_name)
		}
	}

	t.idxRoot = root
	return t.idxRoot
}

func (t *Table) FirstIdxColumns() []*idxNode {
	cols := t.GetIndexRoot().Children

	for _, col := range cols {
		col.Column = t.GetColumn(col.ColName)
	}

	return cols
}

func (t *Table) Imports() string {
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
