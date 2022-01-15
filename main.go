package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/corelib"
	"github.com/olachat/gola/user"
)

type SimpleUser struct {
	user.Name
	user.Email
}

func Print[T any, PT corelib.PointerType[T]](s PT) {
	fmt.Printf("%v\n", s)
}

func main() {
	println("ExecScalar:")
	u := user.FetchById[struct {
		user.Email
	}](1)
	Print(u)

	u2 := user.FetchById[user.User](1)
	Print(u2)

	println("\n\nQuery:")
	users := user.FetchByIds[SimpleUser]([]int{1, 2})
	for _, user := range users {
		Print(user)
	}

	users2 := user.FetchUserByIds([]int{3, 4})
	for _, user := range users2 {
		Print(user)
	}
}
