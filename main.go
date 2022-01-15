package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/user"
)

const (
	testDBPort int    = 33066
	testDBName string = "testdb"
	tableName  string = "user"
)

type User struct {
	user.Id
	user.Name
	user.Email
	user.PhoneNumbers
	user.Created
}

var types = make(map[reflect.Type]bool)
var typeColumnNames = make(map[reflect.Type]string)

type PointerType[T any] interface {
	*T
}

func Print[T any, PT PointerType[T]](s PT) PT {
	t := reflect.TypeOf(s)
	flag := types[t]

	s1 := new(T)
	fmt.Printf("%v \n", s1)

	fmt.Printf("%v %v %v\n", s, t, flag)
	types[t] = true
	return s1
}

type SimpleUser struct {
	user.Name
	user.Email
}

func ExecScalar[T any, PT PointerType[T]]() PT {
	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	u := new(T)
	names := GetColumnsNames(u)
	data := StrutForScan(u)

	err2 := db.QueryRow("SELECT " + names + " from user where id=1").Scan(data...)

	if err2 != nil {
		log.Fatal(err2)
	}

	return u
}

func Query[T any, PT PointerType[T]]() []PT {
	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []PT

	var u *T
	names := GetColumnsNames(u)

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

func PrintString(s string) {
	fmt.Println("\n\n" + s)
}

func GetColumnsNames[T any, PT PointerType[T]](o PT) (joinedColumnNames string) {
	t := reflect.TypeOf(o)
	joinedColumnNames, ok := typeColumnNames[t]
	if ok {
		return joinedColumnNames
	}

	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(user.Column); ok {
			columnNames = append(columnNames, f.GetColumnName())
		}

	}

	joinedColumnNames = strings.Join(columnNames, ",")
	typeColumnNames[t] = joinedColumnNames

	return joinedColumnNames
}

func StrutForScan[T any, PT PointerType[T]](u PT) (pointers []interface{}) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]any, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(user.Column); ok {
			pointers[i] = f.GetPointer()
		}
	}
	return
}

func main() {
	t := &User{}
	t.Id = 1
	Print(t)
	t.SetName("piggy")
	t2 := Print(t)
	t2.SetName("bar")
	fmt.Printf("%v\n", t2)
	fmt.Printf("%v\n", t)

	for i := 0; i < 5; i++ {
		var q *struct {
			user.Id
		}
		Print(q)
	}
	for i := 0; i < 5; i++ {
		var q = new(struct {
			user.Id
		})
		q.Id = 1
		Print(user.Run())
	}
	PrintString(t.GetName())

	u := ExecScalar[SimpleUser]()
	Print(u)

	PrintString("Query:")
	users := Query[SimpleUser]()
	for _, user := range users {
		Print(user)
	}
}
