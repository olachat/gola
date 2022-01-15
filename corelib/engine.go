package corelib

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	TestDBPort int    = 33066
	TestDBName string = "testdb"
	TableName  string = "user"
)

var typeColumnNames = make(map[reflect.Type]string)

func ExecScalar[T any, PT PointerType[T]]() PT {
	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", TestDBPort, TestDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	u := new(T)
	names := GetColumnsNames[T]()
	data := StrutForScan(u)

	err2 := db.QueryRow("SELECT " + names + " from user where id=1").Scan(data...)

	if err2 != nil {
		log.Fatal(err2)
	}

	return u
}

func Query[T any, PT PointerType[T]]() []PT {
	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", TestDBPort, TestDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []PT

	var u *T
	names := GetColumnsNames[T]()

	rows, err2 := db.Query("SELECT " + names + " from user where id in (1,2)")

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

func GetColumnsNames[T any, PT PointerType[T]]() (joinedColumnNames string) {
	var o *T
	t := reflect.TypeOf(o)
	joinedColumnNames, ok := typeColumnNames[t]
	if ok {
		return joinedColumnNames
	}

	o = new(T)
	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			columnNames = append(columnNames, f.GetColumnName())
		}

	}

	joinedColumnNames = strings.Join(columnNames, ",")
	typeColumnNames[t] = joinedColumnNames
	println("NewTypeColumns: " + joinedColumnNames)

	return joinedColumnNames
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
