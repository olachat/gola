package mysqldriver

// modified from: https://github.com/volatiletech/sqlboiler/blob/v4.6.0/drivers/sqlboiler-mysql/driver/mysql.go

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/structs"
	"github.com/pkg/errors"
)

// Assemble is more useful for calling into the library so you don't
// have to instantiate an empty type.
func Assemble(config DBConfig) (dbinfo *structs.DBInfo, err error) {
	driver := MySQLDriver{}
	return driver.Assemble(config)
}

// MySQLDriver holds the database connection string and a handle
// to the database connection.
type MySQLDriver struct {
	connStr string
	conn    *sql.DB
}

// Assemble all the information we need to provide back to the driver
func (m *MySQLDriver) Assemble(c DBConfig) (dbinfo *structs.DBInfo, err error) {
	defer func() {
		if r := recover(); r != nil && err == nil {
			dbinfo = nil
			err = r.(error)
		}
	}()

	m.connStr = MySQLBuildQueryString(c.user, c.pass, c.dbname, c.host, c.port, c.sslmode)
	m.conn, err = sql.Open("mysql", m.connStr)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboiler-mysql failed to connect to database")
	}

	defer func() {
		if e := m.conn.Close(); e != nil {
			dbinfo = nil
			err = e
		}
	}()

	dbinfo = &structs.DBInfo{}

	dbinfo.Schema = c.dbname
	dbinfo.Tables, err = structs.Tables(m, c.dbname, c.whitelist, c.blacklist)
	if err != nil {
		return nil, err
	}

	return dbinfo, err
}

// MySQLBuildQueryString builds a query string for MySQL.
func MySQLBuildQueryString(user, pass, dbname, host string, port int, sslmode string) string {
	config := mysql.NewConfig()

	config.User = user
	if len(pass) != 0 {
		config.Passwd = pass
	}
	config.DBName = dbname
	config.Net = "tcp"
	config.Addr = host
	if port == 0 {
		port = 3306
	}
	config.Addr += ":" + strconv.Itoa(port)
	config.TLSConfig = sslmode

	// MySQL is a bad, and by default reads date/datetime into a []byte
	// instead of a time.Time. Tell it to stop being a bad.
	config.ParseTime = true

	return config.FormatDSN()
}

// TableNames connects to the mysql database and
// retrieves all table names from the information_schema where the
// table schema is public.
func (m *MySQLDriver) TableNames(schema string, whitelist, blacklist []string) ([]string, error) {
	var names []string

	query := `select table_name from information_schema.tables where table_schema = ? and table_type = 'BASE TABLE'`
	args := []interface{}{schema}
	if len(whitelist) > 0 {
		tables := TablesFromList(whitelist)
		if len(tables) > 0 {
			query += fmt.Sprintf(" and table_name in (%s)", strings.Repeat(",?", len(tables))[1:])
			for _, w := range tables {
				args = append(args, w)
			}
		}
	} else if len(blacklist) > 0 {
		tables := TablesFromList(blacklist)
		if len(tables) > 0 {
			query += fmt.Sprintf(" and table_name not in (%s)", strings.Repeat(",?", len(tables))[1:])
			for _, b := range tables {
				args = append(args, b)
			}
		}
	}

	query += ` order by table_name;`

	rows, err := m.conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}

// Columns takes a table name and attempts to retrieve the table information
// from the database information_schema.columns. It retrieves the column names
// and column types and returns those as a []Column after TranslateColumnType()
// converts the SQL types to Go types, for example: "varchar" to "string"
func (m *MySQLDriver) Columns(schema string, table *structs.Table, tableName string, whitelist, blacklist []string) ([]structs.Column, error) {
	var columns []structs.Column
	args := []interface{}{tableName, schema}

	query := `
	select
	c.column_name,
	c.column_type,
	c.column_comment,
	if(c.data_type = 'enum', c.column_type, c.data_type),
	if(extra = 'auto_increment','auto_increment',
		if(version() like '%MariaDB%' and c.column_default = 'NULL', '',
		if(version() like '%MariaDB%' and c.data_type in ('varchar','char','binary','date','datetime','time'),
			replace(substring(c.column_default,2,length(c.column_default)-2),'\'\'','\''),
				c.column_default))),
	c.is_nullable = 'YES',
		0 as is_unique
	from information_schema.columns as c
	where table_name = ? and table_schema = ? and c.extra not like '%VIRTUAL%'`

	if len(whitelist) > 0 {
		cols := ColumnsFromList(whitelist, tableName)
		if len(cols) > 0 {
			query += fmt.Sprintf(" and c.column_name in (%s)", strings.Repeat(",?", len(cols))[1:])
			for _, w := range cols {
				args = append(args, w)
			}
		}
	} else if len(blacklist) > 0 {
		cols := ColumnsFromList(blacklist, tableName)
		if len(cols) > 0 {
			query += fmt.Sprintf(" and c.column_name not in (%s)", strings.Repeat(",?", len(cols))[1:])
			for _, w := range cols {
				args = append(args, w)
			}
		}
	}

	query += ` order by c.ordinal_position;`

	rows, err := m.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var colName, colFullType, colComment, colType string
		var nullable, unique bool
		var defaultValue *string
		if err := rows.Scan(&colName, &colFullType, &colComment, &colType, &defaultValue, &nullable, &unique); err != nil {
			return nil, errors.Wrapf(err, "unable to scan for table %s", tableName)
		}

		column := structs.Column{
			Name:       colName,
			Comment:    colComment,
			FullDBType: colFullType, // example: tinyint(1) instead of tinyint
			DBType:     colType,
			Nullable:   nullable,
			Unique:     unique,
		}

		if defaultValue != nil && *defaultValue != "NULL" {
			column.Default = *defaultValue
		}

		column.Comment = strings.ReplaceAll(column.Comment, "\r\n", " ")
		column.Comment = strings.ReplaceAll(column.Comment, "\n", " ")
		column.Comment = strings.ReplaceAll(column.Comment, "\"", "'")

		column.Table = table

		columns = append(columns, column)
	}

	return columns, nil
}

// PrimaryKeyInfo looks up the primary key for a table.
func (m *MySQLDriver) PrimaryKeyInfo(schema, tableName string) (*structs.PrimaryKey, error) {
	// dummy
	// actual implementation in SetIndexAndKey
	return nil, nil
}

// ForeignKeyInfo retrieves the foreign keys for a given table name.
func (m *MySQLDriver) ForeignKeyInfo(schema, tableName string) ([]structs.ForeignKey, error) {
	var fkeys []structs.ForeignKey

	query := `
	select constraint_name, table_name, column_name, referenced_table_name, referenced_column_name
	from information_schema.key_column_usage
	where table_schema = ? and referenced_table_schema = ? and table_name = ?
	order by constraint_name, table_name, column_name, referenced_table_name, referenced_column_name
	`

	var rows *sql.Rows
	var err error
	if rows, err = m.conn.Query(query, schema, schema, tableName); err != nil {
		return nil, err
	}

	for rows.Next() {
		var fkey structs.ForeignKey
		var sourceTable string

		fkey.Table = tableName
		err = rows.Scan(&fkey.Name, &sourceTable, &fkey.Column, &fkey.ForeignTable, &fkey.ForeignColumn)
		if err != nil {
			return nil, err
		}

		fkeys = append(fkeys, fkey)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fkeys, nil
}

var nullSimpleTypes = map[string]string{
	"float":            "null.Float32",
	"double":           "null.Float64",
	"double precision": "null.Float64",
	"real":             "null.Float64",
	"boolean":          "null.Bool",
	"bool":             "null.Bool",
	"date":             "null.Time",
	"datetime":         "null.Time",
	"timestamp":        "null.Time",
	"binary":           "null.Bytes",
	"varbinary":        "null.Bytes",
	"tinyblob":         "null.Bytes",
	"blob":             "null.Bytes",
	"mediumblob":       "null.Bytes",
	"longblob":         "null.Bytes",
	"numeric":          "types.NullDecimal",
	"decimal":          "types.NullDecimal",
	"dec":              "types.NullDecimal",
	"fixed":            "types.NullDecimal",
	"json":             "null.JSON",
}

var simpleTypes = map[string]string{
	"float":            "float32",
	"double":           "float64",
	"double precision": "float64",
	"real":             "float64",
	"boolean":          "bool",
	"bool":             "bool",
	"date":             "time",
	"datetime":         "time",
	"timestamp":        "time",
	"binary":           "bytes",
	"varbinary":        "bytes",
	"tinyblob":         "bytes",
	"blob":             "bytes",
	"mediumblob":       "bytes",
	"longblob":         "bytes",
	"numeric":          "types.Decimal",
	"decimal":          "types.Decimal",
	"dec":              "types.Decimal",
	"fixed":            "types.Decimal",
	"json":             "types.JSON",
}

var nullIntTypes = map[string]string{
	"tinyint":   "null.Int8",
	"smallint":  "null.Int16",
	"mediumint": "null.Int32",
	"int":       "null.Int",
	"integer":   "null.Int",
	"bigint":    "null.Int64",
}

var nullUintTypes = map[string]string{
	"tinyint":   "null.Uint8",
	"smallint":  "null.Uint16",
	"mediumint": "null.Uint32",
	"int":       "null.Uint",
	"integer":   "null.Uint",
	"bigint":    "null.Uint64",
}

var intTypes = map[string]string{
	"tinyint":   "int8",
	"smallint":  "int16",
	"mediumint": "int32",
	"int":       "int",
	"integer":   "int",
	"bigint":    "int64",
}

var uintTypes = map[string]string{
	"tinyint":   "uint8",
	"smallint":  "uint16",
	"mediumint": "uint32",
	"int":       "uint",
	"integer":   "uint",
	"bigint":    "uint64",
}

// TranslateColumnType converts mysql database types to Go types, for example
// "varchar" to "string" and "bigint" to "int64". It returns this parsed data
// as a Column object.
func (m *MySQLDriver) TranslateColumnType(c structs.Column) structs.Column {
	unsigned := strings.Contains(c.FullDBType, "unsigned")

	boolType, strType, intTypeMap, uintTypeMap, simpleTypeMap := "bool", "string", intTypes, uintTypes, simpleTypes
	if c.Nullable {
		boolType, strType, intTypeMap, uintTypeMap, simpleTypeMap = "null.Bool", "null.String", nullIntTypes, nullUintTypes, nullSimpleTypes
	}

	if c.DBType == "tinyint" {
		// map tinyint(1) to bool
		if c.FullDBType == "tinyint(1)" {
			c.Type = boolType
			return c
		}
	}

	if columnType, ok := intTypeMap[c.DBType]; ok {
		if unsigned {
			c.Type = uintTypeMap[c.DBType]
		} else {
			c.Type = columnType
		}

		return c
	}

	if columnType, ok := simpleTypeMap[c.DBType]; ok {
		c.Type = columnType
		return c
	}

	c.Type = strType
	return c
}
