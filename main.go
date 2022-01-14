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

type PointerType[B any] interface {
	*B
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

func PrintString(s string) {
	fmt.Println(s)
}

func StrutForScan(u any) (columnNames []string, pointers []any) {
	val := reflect.ValueOf(u).Elem()
	pointers = make([]any, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(user.Column); ok {
			pointers[i] = f.GetPointer()
			columnNames = append(columnNames, f.GetColumnName())
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

	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	type SimpleUser struct {
		user.Name
		user.Email
	}
	u := &SimpleUser{}

	columnNames, data := StrutForScan(u)
	names := strings.Join(columnNames, ",")

	err2 := db.QueryRow("SELECT " + names + " from user where id=1").Scan(data...)

	if err2 != nil {
		log.Fatal(err2)
	}

	Print(u)
}
