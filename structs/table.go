package structs

import (
	"sort"
	"strings"
)

// Table struct represent table information read from mysql
type Table struct {
	Name string `json:"name"`
	// For dbs with real schemas, like Postgres.
	// Example value: "schema_name"."table_name"
	SchemaName string   `json:"schema_name"`
	Columns    []Column `json:"columns"`

	PKey  *PrimaryKey  `json:"p_key"`
	FKeys []ForeignKey `json:"f_keys"`

	IsJoinTable bool `json:"is_join_table"`
	Indexes     map[string][]*IndexDesc
	VERSION     string
	idxRoot     *IdxNode
}

// Package returns the package name for the table
func (t *Table) Package() string {
	return t.Name
}

// ClassName returns the go struct(class) name for the table
func (t *Table) ClassName() string {
	name := getGoName(t.Name)
	if strings.HasSuffix(name, "s") {
		return name[:len(name)-1]
	}

	return name
}

// GetPrimaryKey returns the field name of the primary key
func (t *Table) GetPrimaryKey() string {
	for _, c := range t.Columns {
		if c.IsPrimaryKey() {
			return c.GoName()
		}
	}

	return ""
}

// GetPrimaryKeyType returns the go type of the primary key
func (t *Table) GetPrimaryKeyType() string {
	for _, c := range t.Columns {
		if c.IsPrimaryKey() {
			return c.GoType()
		}
	}

	return ""
}

// GetPrimaryKeyName returns the column name of the primary key
func (t *Table) GetPrimaryKeyName() string {
	for _, c := range t.Columns {
		if c.IsPrimaryKey() {
			return c.Name
		}
	}

	return ""
}

// NonPrimaryColumns returns all columns except primary key
func (t *Table) NonPrimaryColumns() []Column {
	result := make([]Column, 0, len(t.Columns))

	for _, c := range t.Columns {
		if !c.IsPrimaryKey() {
			result = append(result, c)
		}
	}

	return result
}

// GetIndexRoot returns the root index node
func (t *Table) GetIndexRoot() *IdxNode {
	if t.idxRoot != nil {
		return t.idxRoot
	}

	root := &IdxNode{
		ColName: "",
	}

	idxNames := make([]string, 0, len(t.Indexes))
	for idxName := range t.Indexes {
		idxNames = append(idxNames, idxName)
	}

	sort.Slice(idxNames, func(i, j int) bool {
		return idxNames[i] < idxNames[j]
	})

	for _, idxName := range idxNames {
		items := t.Indexes[idxName]
		node := root
		for _, item := range items {
			node = node.GetChildren(item.ColumnName)
		}
	}

	t.idxRoot = root
	return t.idxRoot
}

// GetIndexNodes returns all index nodes need customized interface
// i.e. has non-empty children nodes
func (t *Table) GetIndexNodes() []*IdxNode {
	allNodes := t.GetIndexRoot().GetAllChildren()
	nodes := make([]*IdxNode, 0, len(allNodes))
	for _, n := range allNodes {
		if len(n.Children) > 0 {
			n.Column = t.GetColumn(n.ColName)
			for _, c := range n.Children {
				c.Column = t.GetColumn(c.ColName)
			}
			nodes = append(nodes, n)
		}

	}
	return nodes
}

// FirstIdxColumns returns first column of all indexes
func (t *Table) FirstIdxColumns() []*IdxNode {
	cols := t.GetIndexRoot().Children

	for _, col := range cols {
		col.Column = t.GetColumn(col.ColName)
	}

	return cols
}

// Imports returns new packages needed to import
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
