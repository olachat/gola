package mysqldriver

import (
	"sort"

	"github.com/olachat/gola/structs"
)

// SetIndexAndKey sets indexes & key information to given tables
func (m *MySQLDriver) SetIndexAndKey(tables []*structs.Table) (err error) {
	for _, t := range tables {
		var tableDesc []*structs.RowDesc
		rows, err := m.conn.Query("desc `" + t.Name + "`")
		if err != nil {
			return err
		}

		for rows.Next() {
			rd := new(structs.RowDesc)
			rows.Scan(&rd.Field, &rd.Type, &rd.Null, &rd.Key, &rd.Default, &rd.Extra)
			tableDesc = append(tableDesc, rd)
		}

		for _, rd := range tableDesc {
			if rd.Key == "PRI" {
				if t.PKey == nil {
					t.PKey = &structs.PrimaryKey{}
					t.PKey.Name = rd.Field
					t.PKey.Columns = []string{rd.Field}
				} else {
					t.PKey.Columns = append(t.PKey.Columns, rd.Field)
				}
			}
		}

		var indexDesc []*structs.IndexDesc
		rows, err = m.conn.Query("show index from `" + t.Name + "`")
		if err != nil {
			return err
		}
		for rows.Next() {
			id := new(structs.IndexDesc)
			rows.Scan(&id.Table, &id.NonUnique, &id.KeyName, &id.SeqInIndex, &id.ColumnName, &id.Collation, &id.Cardinality,
				&id.SubPart, &id.Packed, &id.Null, &id.IndexType, &id.Comment, &id.IndexComment, &id.Visible, &id.Expression)
			indexDesc = append(indexDesc, id)
		}
		t.Indexes = groupIndex(indexDesc)

		// Hack to handle table without primary key, but has only one unique key
		// Just consider that unique key as primary
		if t.PKey == nil {
			uniqueIdxCount := 0
			var uniqueIdx []*structs.IndexDesc
			var uniqueIdxName string
			for idxName, idx := range t.Indexes {
				isUnique := true
				for _, node := range idx {
					if node.NonUnique > 0 {
						isUnique = false
					}
				}
				if isUnique {
					uniqueIdxCount += 1
					uniqueIdx = idx
					uniqueIdxName = idxName
				}
			}

			if uniqueIdxCount == 1 {
				t.PKey = &structs.PrimaryKey{}
				t.PKey.Name = uniqueIdxName
				t.PKey.Columns = make([]string, len(uniqueIdx))
				for i, node := range uniqueIdx {
					t.PKey.Columns[i] = node.ColumnName
				}
			}
		}
	}

	return nil
}

func filterBy[T any](items []*T, isNeeded func(item *T) bool) []*T {
	result := make([]*T, 0, len(items))

	for _, item := range items {
		if isNeeded(item) {
			result = append(result, item)
		}
	}

	return result
}

func groupIndex(indexDesc []*structs.IndexDesc) map[string][]*structs.IndexDesc {
	data := make(map[string][]*structs.IndexDesc, 0)

	for _, idx := range indexDesc {
		key := idx.KeyName
		if _, ok := data[key]; !ok {
			data[key] = []*structs.IndexDesc{}
		}
	}

	for name := range data {
		items := filterBy(indexDesc, func(item *structs.IndexDesc) bool {
			return item.KeyName == name
		})

		sort.Slice(items, func(i, j int) bool {
			return items[i].SeqInIndex < items[j].SeqInIndex
		})

		data[name] = items

	}

	return data
}
