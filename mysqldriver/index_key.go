package mysqldriver

/*
desc users;
+------------+--------------+------+------+---------+----------------+
| Field      | Type         | Null | Key  | Default | Extra          |
+------------+--------------+------+------+---------+----------------+
| id         | int          | NO   | PRI  |         | auto_increment |
| name       | varchar(255) | NO   |      | ""      |                |
| email      | varchar(255) | NO   | MUL  | ""      |                |
| created_at | int unsigned | NO   |      | "0"     |                |
| updated_at | int unsigned | NO   |      | "0"     |                |
+------------+--------------+------+------+---------+----------------+
*/

type RowDesc struct {
	Field, Type, Null, Key, Default, Extra string
}

/*
show index from users;
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment | Visible | Expression |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| users |          0 | email    |            1 | email       | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
| users |          1 | name     |            1 | name        | NULL      |           0 |     NULL | NULL   |      | BTREE      |         |               | YES     | NULL       |
+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
*/

type IndexDesc struct {
	Table, Key_name, Column_name, Collation, Sub_part, Packed, Null, Index_type, Comment, Index_comment, Visible, Expression string
	Non_unique, Seq_in_index, Cardinality                                                                                    int
}

func (m *MySQLDriver) SetIndexAndKey(dbinfo *DBInfo) (err error) {
	for i := range dbinfo.Tables {
		t := &dbinfo.Tables[i]
		var tableDesc []*RowDesc
		rows, err := m.conn.Query("desc " + t.Name)
		if err != nil {
			return err
		}

		for rows.Next() {
			rd := new(RowDesc)
			rows.Scan(&rd.Field, &rd.Type, &rd.Null, &rd.Key, &rd.Default, &rd.Extra)
			tableDesc = append(tableDesc, rd)
		}

		for _, rd := range tableDesc {
			if rd.Key == "PRI" && t.PKey == nil {
				t.PKey = &PrimaryKey{}
				t.PKey.Name = rd.Field
				t.PKey.Columns = []string{rd.Field}
			}
		}

		var indexDesc []*IndexDesc
		rows, err = m.conn.Query("show index from " + t.Name)
		if err != nil {
			return err
		}
		for rows.Next() {
			id := new(IndexDesc)
			rows.Scan(&id.Table, &id.Non_unique, &id.Key_name, &id.Seq_in_index, &id.Column_name, &id.Collation, &id.Cardinality,
				&id.Sub_part, &id.Packed, &id.Null, &id.Index_type, &id.Comment, &id.Index_comment, &id.Visible, &id.Expression)
			indexDesc = append(indexDesc, id)
		}
	}

	return nil
}
