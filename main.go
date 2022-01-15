package main

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/user"
)

type User struct {
	user.Id
	user.Name
	user.Email
	user.PhoneNumbers
	user.Created
}

var types = make(map[reflect.Type]bool)

func Print[T any, PT corelib.PointerType[T]](s PT) PT {
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
	println(t.GetName())

	u := corelib.ExecScalar[struct {
		user.Email
	}]()
	Print(u)

	println("Query:")
	users := corelib.Query[SimpleUser]()
	for _, user := range users {
		Print(user)
	}
}
