package mysqldriver

import (
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

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

func (m *MySQLDriver) SetIndexAndKey(dbinfo *drivers.DBInfo) (err error) {
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
				t.PKey = &drivers.PrimaryKey{}
				t.PKey.Name = rd.Field
				t.PKey.Columns = []string{rd.Field}
			}
		}
	}

	return nil
}
