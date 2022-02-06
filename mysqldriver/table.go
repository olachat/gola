package mysqldriver

// modified from: https://github.com/volatiletech/sqlboiler/blob/v4.6.0/drivers/sqlboiler-mysql/driver/interface.go

import (
	"sort"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	"github.com/volatiletech/strmangle"
)

func Tables(c drivers.Constructor, schema string, whitelist, blacklist []string) ([]drivers.Table, error) {
	var err error

	names, err := c.TableNames(schema, whitelist, blacklist)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get table names")
	}

	sort.Strings(names)

	var tables []drivers.Table
	for _, name := range names {
		t := drivers.Table{
			Name: name,
		}

		if t.Columns, err = c.Columns(schema, name, whitelist, blacklist); err != nil {
			return nil, errors.Wrapf(err, "unable to fetch table column info (%s)", name)
		}

		for i, col := range t.Columns {
			t.Columns[i] = c.TranslateColumnType(col)
		}

		if t.PKey, err = c.PrimaryKeyInfo(schema, name); err != nil {
			return nil, errors.Wrapf(err, "unable to fetch table pkey info (%s)", name)
		}

		if t.FKeys, err = c.ForeignKeyInfo(schema, name); err != nil {
			return nil, errors.Wrapf(err, "unable to fetch table fkey info (%s)", name)
		}

		filterForeignKeys(&t, whitelist, blacklist)

		setIsJoinTable(&t)

		tables = append(tables, t)
	}

	// Relationships have a dependency on foreign key nullability.
	for i := range tables {
		tbl := &tables[i]
		setForeignKeyConstraints(tbl, tables)
	}

	return tables, nil
}

// setIsJoinTable if there are:
// A composite primary key involving two columns
// Both primary key columns are also foreign keys
func setIsJoinTable(t *drivers.Table) {
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
func filterForeignKeys(t *drivers.Table, whitelist, blacklist []string) {
	var fkeys []drivers.ForeignKey
	for _, fkey := range t.FKeys {
		if (len(whitelist) == 0 || strmangle.SetInclude(fkey.ForeignTable, whitelist)) &&
			(len(blacklist) == 0 || !strmangle.SetInclude(fkey.ForeignTable, blacklist)) {
			fkeys = append(fkeys, fkey)
		}
	}
	t.FKeys = fkeys
}

func setForeignKeyConstraints(t *drivers.Table, tables []drivers.Table) {
	for i, fkey := range t.FKeys {
		localColumn := t.GetColumn(fkey.Column)
		foreignTable := drivers.GetTable(tables, fkey.ForeignTable)
		foreignColumn := foreignTable.GetColumn(fkey.ForeignColumn)

		t.FKeys[i].Nullable = localColumn.Nullable
		t.FKeys[i].Unique = localColumn.Unique
		t.FKeys[i].ForeignColumnNullable = foreignColumn.Nullable
		t.FKeys[i].ForeignColumnUnique = foreignColumn.Unique
	}
}
