package main

import (
	"fmt"

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

type SimpleUser struct {
	user.Name
	user.Email
}

func Print[T any, PT corelib.PointerType[T]](s PT) {
	fmt.Printf("%v\n", s)
}

func main() {
	println("ExecScalar:")
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
