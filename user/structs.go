package user

import "time"

type Id int
type Name struct {
	val string
}
type Email struct {
	val string
}
type PhoneNumbers string
type Created time.Time

func (c *Name) GetName() string {
	return c.val
}

func (c *Name) SetName(val string) {
	c.val = val
}

func (c *Name) GetColumnName() string {
	return "name"
}

func (c *Name) IsPrimaryKey() bool {
	return false
}

func (c *Name) GetValPointer() interface{} {
	return &c.val
}

func (c *Email) GetEmail() string {
	return c.val
}

func (c *Email) GetColumnName() string {
	return "email"
}

func (c *Email) GetValPointer() interface{} {
	return &c.val
}

func (c *Email) IsPrimaryKey() bool {
	return false
}

type E struct {
	Email
}

func Run() *E {
	var q *E
	return q
}
