package corelib

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var _connstr string

func Setup(connstr string) {
	_connstr = connstr
}

var typeColumnNames = make(map[reflect.Type]string)
var typeTableNames = make(map[reflect.Type]string)

func FetchById[T any, PT PointerType[T]](id int) PT {
	db, err := sql.Open("mysql", _connstr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	u := new(T)
	tableName, columnsNames := GetTableAndColumnsNames[T]()
	data := StrutForScan(u)

	query := fmt.Sprintf("SELECT %s from %s where id=%d", columnsNames, tableName, id)
	err2 := db.QueryRow(query).Scan(data...)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err2)
	}

	return u
}

func Query[T any, PT PointerType[T]]() []PT {
	db, err := sql.Open("mysql", _connstr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []PT

	var u *T
	tableName, columnsNames := GetTableAndColumnsNames[T]()

	rows, err2 := db.Query("SELECT " + columnsNames + " from " + tableName + " where id in (1,2)")

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

func GetTableAndColumnsNames[T any, PT PointerType[T]]() (tableName string, joinedColumnNames string) {
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

func StrutForScan[T any, PT PointerType[T]](u PT) (pointers []interface{}) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]any, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			pointers[i] = f.GetValPointer()
		}
	}
	return
}
