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
	t.Id = 1
	Print(t)
	t.Name = "run"
	Print(t)

	for i := 0; i < 5; i++ {
		var q struct {
			user.Id
		}
		Print(q)
	}
	for i := 0; i < 5; i++ {
		var q struct {
			user.Id
		}
		q.Id = 1
		Print(user.Run())
	}
	PrintString(string(t.Name))

	db, err := sql.Open("mysql", fmt.Sprintf("root:@tcp(127.0.0.1:%d)/%s", testDBPort, testDBName))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	data := []any{
		&t.Id,
		&t.Name,
		&t.Email,
	}

	err2 := db.QueryRow("SELECT id, name, email from user where id=1").Scan(data...)

	if err2 != nil {
		log.Fatal(err2)
	}

	Print(t)
}
