package main

import "fmt"

type UserName string
type UserAge int

type User struct {
	UserName
	UserAge
}

func main() {
	t := &User{}
	fmt.Printf("%v\n", t)
}
