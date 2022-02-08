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

func (t *Table) FirstIdxColumns() []Column {
	firstCols := make(map[string]bool)

	for _, items := range t.Indexes {
		firstCols[items[0].Column_name] = true
	}

	result := make([]Column, len(firstCols))
	i := 0
	for colName := range firstCols {
		result[i] = t.GetColumn(colName)
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
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
