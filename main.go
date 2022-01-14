package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/user"
)

const (
	testDBPort int    = 33066
	testDBName string = "testdb"
	tableName  string = "mytable"
)

type User struct {
	user.Name
	user.Nick
	user.Age
}

var types = make(map[reflect.Type]bool)

func Print[T any](s T) T {
	t := reflect.TypeOf(s)
	flag := types[t]

	fmt.Printf("%v %v %v\n", s, t, flag)
	types[t] = true
	return s
}

func PrintString(s string) {
	fmt.Println(s)
}

func main() {
	t := &User{}
	t.Age = 1
	Print(t)
	t.SetAge(2)
	t.SetName("run")
	Print(t)

	for i := 0; i < 5; i++ {
		var q struct {
			user.Age
		}
		Print(q)
	}
	for i := 0; i < 5; i++ {
		var q struct {
			user.Age
		}
		q.Age = 1
		Print(user.Run())
	}
	PrintString(t.GetName())

	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string
	data := []any{
		&version,
	}

	err2 := db.QueryRow("SELECT VERSION()").Scan(data...)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}
