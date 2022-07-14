package corelib

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

var typeColumnNames = make(map[reflect.Type]string)
var typeTableNames = make(map[reflect.Type]string)

func FetchById[T any](id int) *T {
	u := new(T)
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	data := StrutForScan(u)

	query := fmt.Sprintf("SELECT %s from %s where id=%d", columnsNames, tableName, id)
	err2 := _db.QueryRow(query).Scan(data...)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err2)
	}

	return u
}

func FetchByIds[T any](ids []int) []*T {
	tableName, columnsNames := GetTableAndColumnsNames[T]()

	idstr := JoinInts(ids, ",")
	query := fmt.Sprintf("SELECT %s from %s where id in(%s)", columnsNames, tableName, idstr)

	return Query[T](query)
}

func Exec[T any](query string, params ...interface{}) (sql.Result, error) {
	return _db.Exec(query, params...)
}

func FindOne[T any](where WhereQuery) *T {
	u := new(T)
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	data := StrutForScan(u)
	whereSql, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s from %s where %s", columnsNames,
		tableName, whereSql)
	err2 := _db.QueryRow(query, params...).Scan(data...)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err2)
	}

	return u
}

func Find[T any](where WhereQuery) []*T {
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	whereSql, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s from %s %s", columnsNames,
		tableName, whereSql)

	return Query[T](query, params...)
}

func Query[T any](query string, params ...interface{}) []*T {
	var result []*T
	var u *T

	rows, err2 := _db.Query(query, params...)

	if err2 != nil {
		log.Fatal(err2)
	}

	for rows.Next() {
		u = new(T)
		data := StrutForScan(u)
		rows.Scan(data...)
		result = append(result, u)
	}

	return result
}

func GetTableAndColumnsNames[T any]() (tableName string, joinedColumnNames string) {
	var o *T
	t := reflect.TypeOf(o)
	joinedColumnNames, ok := typeColumnNames[t]
	if ok {
		tableName = typeTableNames[t]
		return
	}

	o = new(T)
	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			columnNames = append(columnNames, f.GetColumnName())
			if tableName == "" {
				tableName = f.GetTableType().GetTableName()
			}
		}
	}

	joinedColumnNames = strings.Join(columnNames, ",")
	typeTableNames[t] = tableName
	typeColumnNames[t] = joinedColumnNames

	return
}

func StrutForScan[T any](u *T) (pointers []interface{}) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]interface{}, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			pointers = append(pointers, f.GetValPointer())
		}
	}
	return
}
