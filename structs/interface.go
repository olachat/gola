package structs

// modified from: https://github.com/volatiletech/sqlboiler/blob/v4.6.0/drivers/sqlboiler-mysql/driver/interface.go

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

// These constants are used in the config map passed into the driver
const (
	ConfigBlacklist = "blacklist"
	ConfigWhitelist = "whitelist"
	ConfigSchema    = "schema"

	ConfigUser    = "user"
	ConfigPass    = "pass"
	ConfigHost    = "host"
	ConfigPort    = "port"
	ConfigDBName  = "dbname"
	ConfigSSLMode = "sslmode"
)

// DBInfo represents information about a database
type DBInfo struct {
	Schema string   `json:"schema"`
	Tables []*Table `json:"tables"`
}

// Constructor breaks down the functionality required to implement a driver
// such that the drivers.Tables method can be used to reduce duplication in driver
// implementations.
type Constructor interface {
	TableNames(schema string, whitelist, blacklist []string) ([]string, error)
	Columns(schema string, table *Table, tableName string, whitelist, blacklist []string) ([]Column, error)
	SetIndexAndKey(tables []*Table) (err error)
	ForeignKeyInfo(schema, tableName string) ([]ForeignKey, error)

	// TranslateColumnType takes a Database column type and returns a go column type.
	TranslateColumnType(Column) Column
}

// GetTable by name. Panics if not found (for use in templates mostly).
func GetTable(tables []*Table, name string) (tbl *Table) {
	for _, t := range tables {
		if t.Name == name {
			return t
		}
	}

	panic(fmt.Sprintf("could not find table name: %s", name))
}

// GetColumn by name. Panics if not found (for use in templates mostly).
func (t *Table) GetColumn(name string) (col Column) {
	for _, c := range t.Columns {
		if c.Name == name {
			return c
		}
	}

	panic(fmt.Sprintf("could not find column name: %s", name))
}

// Tables returns tables information of given schema
func Tables(c Constructor, schema string, whitelist, blacklist []string) ([]*Table, error) {
	var err error

	names, err := c.TableNames(schema, whitelist, blacklist)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get table names")
	}

	sort.Strings(names)

	var tables []*Table
	for _, name := range names {
		t := &Table{
			Name: name,
		}

		if t.Columns, err = c.Columns(schema, t, name, whitelist, blacklist); err != nil {
			return nil, errors.Wrapf(err, "unable to fetch table column info (%s)", name)
		}

		for i, col := range t.Columns {
			t.Columns[i] = c.TranslateColumnType(col)
		}

		if t.FKeys, err = c.ForeignKeyInfo(schema, name); err != nil {
			return nil, errors.Wrapf(err, "unable to fetch table fkey info (%s)", name)
		}

		filterForeignKeys(t, whitelist, blacklist)

		setIsJoinTable(t)

		tables = append(tables, t)
	}

	// Relationships have a dependency on foreign key nullability.
	for i := range tables {
		tbl := tables[i]
		setForeignKeyConstraints(tbl, tables)
	}

	c.SetIndexAndKey(tables)
	return tables, nil
}

// setIsJoinTable if there are:
// A composite primary key involving two columns
// Both primary key columns are also foreign keys
func setIsJoinTable(t *Table) {
	if t.PKey == nil || len(t.PKey.Columns) != 2 || len(t.FKeys) < 2 || len(t.Columns) > 2 {
		return
	}

	for _, c := range t.PKey.Columns {
		found := false
		for _, f := range t.FKeys {
			if c == f.Column {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	t.IsJoinTable = true
}

// filterForeignKeys filter FK whose ForeignTable is not in whitelist or in blacklist
func filterForeignKeys(t *Table, whitelist, blacklist []string) {
	var fkeys []ForeignKey
	for _, fkey := range t.FKeys {
		if (len(whitelist) == 0 || setInclude(fkey.ForeignTable, whitelist)) &&
			(len(blacklist) == 0 || !setInclude(fkey.ForeignTable, blacklist)) {
			fkeys = append(fkeys, fkey)
		}
	}
	t.FKeys = fkeys
}

func setForeignKeyConstraints(t *Table, tables []*Table) {
	for i, fkey := range t.FKeys {
		localColumn := t.GetColumn(fkey.Column)
		foreignTable := GetTable(tables, fkey.ForeignTable)
		foreignColumn := foreignTable.GetColumn(fkey.ForeignColumn)

		t.FKeys[i].Nullable = localColumn.Nullable
		t.FKeys[i].Unique = localColumn.Unique
		t.FKeys[i].ForeignColumnNullable = foreignColumn.Nullable
		t.FKeys[i].ForeignColumnUnique = foreignColumn.Unique
	}
}
