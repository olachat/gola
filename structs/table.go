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

func (t *Table) GetPrimaryKey() string {
	for _, c := range t.Columns {
		if c.IsPrimaryKey() {
			return c.GoName()
		}
	}

	return ""
}

func (t *Table) NonPrimaryColumns() []Column {
	result := make([]Column, 0, len(t.Columns))

	for _, c := range t.Columns {
		if !c.IsPrimaryKey() {
			result = append(result, c)
		}
	}

	return result
}

func (t *Table) GetIndexRoot() *idxNode {
	if t.idxRoot != nil {
		return t.idxRoot
	}

	root := &idxNode{
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
			node = node.GetChildren(item.Column_name)
		}
	}

	t.idxRoot = root
	return t.idxRoot
}

// GetIndexNodes returns all index nodes need customized interface
// i.e. has non-empty children nodes
func (t *Table) GetIndexNodes() []*idxNode {
	allNodes := t.GetIndexRoot().GetAllChildren()
	nodes := make([]*idxNode, 0, len(allNodes))
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
		if strings.Contains(strings.ToLower(c.SQLType()), "set") {
			packages[`"strings"`] = true
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
