package user

import "time"

type Id int
type Name string
type Email string
type PhoneNumbers string
type Created time.Time

func Run() interface{} {
	var q struct {
		Email
	}
	return q
}
