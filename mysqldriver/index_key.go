package mysqldriver

import "github.com/olachat/gola/structs"

func (m *MySQLDriver) SetIndexAndKey(dbinfo *structs.DBInfo) (err error) {
	for i := range dbinfo.Tables {
		t := &dbinfo.Tables[i]
		var tableDesc []*structs.RowDesc
		rows, err := m.conn.Query("desc " + t.Name)
		if err != nil {
			return err
		}

		for rows.Next() {
			rd := new(structs.RowDesc)
			rows.Scan(&rd.Field, &rd.Type, &rd.Null, &rd.Key, &rd.Default, &rd.Extra)
			tableDesc = append(tableDesc, rd)
		}

		for _, rd := range tableDesc {
			if rd.Key == "PRI" && t.PKey == nil {
				t.PKey = &structs.PrimaryKey{}
				t.PKey.Name = rd.Field
				t.PKey.Columns = []string{rd.Field}
			}
		}

		var indexDesc []*structs.IndexDesc
		rows, err = m.conn.Query("show index from " + t.Name)
		if err != nil {
			return err
		}
		for rows.Next() {
			id := new(structs.IndexDesc)
			rows.Scan(&id.Table, &id.Non_unique, &id.Key_name, &id.Seq_in_index, &id.Column_name, &id.Collation, &id.Cardinality,
				&id.Sub_part, &id.Packed, &id.Null, &id.Index_type, &id.Comment, &id.Index_comment, &id.Visible, &id.Expression)
			indexDesc = append(indexDesc, id)
		}
	}

	return nil
}
