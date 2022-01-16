package mysqldriver

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/drivers"
)

type RowDesc struct {
	Field, Type, Null, Key, Default, Extra string
}

func (m *MySQLDriver) GetIndexAndKey(db *sql.DB, dbinfo *drivers.DBInfo) (err error) {
	for i := range dbinfo.Tables {
		t := &dbinfo.Tables[i]
		var tableDesc []*RowDesc
		rows, err := db.Query("desc " + t.Name)
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
