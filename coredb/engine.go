package coredb

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var _dbp DBProvider

type DBMode int

const (
	// DBModeRead allow read from slave db
	DBModeRead DBMode = iota
	// DBModeRead set write to master db
	DBModeWrite
	// DBModeReadFromWrite is aka force read from master db
	DBModeReadFromWrite
)

type DBProvider func(dbname string, mode DBMode) *sql.DB

// Setup default db provider for all db ops
func Setup(dbp DBProvider) {
	_dbp = dbp
}

var typeColumnNames = make(map[reflect.Type]string)

func getDB(dbname string, mode DBMode) *sql.DB {
	if _dbp == nil {
		panic("coredb DBProvider hasn't setup")
	}
	db := _dbp(dbname, mode)
	if db != nil {
		return db
	}
	panic(fmt.Sprintf("Can't get db for %s %v", dbname, mode))
}

// FetchByPK returns a row of T type with given primary key value
func FetchByPK[T any](dbname string, tableName string, pkName []string, val ...any) *T {
	sql := "WHERE `" + pkName[0] + "` = ?"
	for _, name := range pkName[1:] {
		sql += " AND `" + name + "` = ?"
	}
	w := NewWhere(sql, val...)
	return FindOne[T](dbname, tableName, w)
}

// FetchByPKs returns rows of T type with given primary key values
func FetchByPKs[T any](dbname string, tableName string, pkName string, vals []any) []*T {
	if len(vals) == 0 {
		return make([]*T, 0)
	}

	query := fmt.Sprintf("WHERE `%s` IN (%s)", pkName, GetParamPlaceHolder(len(vals)))
	w := NewWhere(query, vals...)

	result, _ := Find[T](dbname, tableName, w)
	return result
}

// Exec given query with given db info & params
func Exec(dbname string, query string, params ...any) (sql.Result, error) {
	mydb := getDB(dbname, DBModeWrite)
	return mydb.Exec(query, params...)
}

// FindOne returns a row from given table type with where query
func FindOne[T any](dbname string, tableName string, where WhereQuery) *T {
	u := new(T)
	columnsNames := GetColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	mydb := getDB(dbname, DBModeRead)
	err2 := mydb.QueryRow(query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 != sql.ErrNoRows {
			log.Fatal(err2)
		}

		return nil
	}

	return u
}

// Find returns rows from given table type with where query
func Find[T any](dbname string, tableName string, where WhereQuery) ([]*T, error) {
	columnsNames := GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return Query[T](dbname, query, params...)
}

// QueryInt single int result by query, handy for count(*) querys
func QueryInt(dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeRead)
	mydb.QueryRow(query, params...).Scan(&result)
	return
}

// Query rows from given table type with where query & params
func Query[T any](dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeRead)
	rows, err := mydb.Query(query, params...)
	if err != nil {
		return
	}

	var u *T
	for rows.Next() {
		u = new(T)
		data := StrutForScan(u)
		err = rows.Scan(data...)
		if err != nil {
			return
		}
		result = append(result, u)
	}

	return
}

// GetColumnsNames returns column names joined by `,` of given type
func GetColumnsNames[T any]() (joinedColumnNames string) {
	var o *T
	t := reflect.TypeOf(o)
	joinedColumnNames, ok := typeColumnNames[t]
	if ok {
		return
	}

	o = new(T)
	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			columnNames = append(columnNames, "`"+f.GetColumnName()+"`")
		}
	}

	joinedColumnNames = strings.Join(columnNames, ",")
	typeColumnNames[t] = joinedColumnNames

	return
}

// StrutForScan returns value pointers of given obj
func StrutForScan[T any](u *T) (pointers []any) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]any, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			pointers = append(pointers, f.GetValPointer())
		}
	}
	return
}
