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

func (c *Name) GetPointer() interface{} {
	return &c.val
}

func (c *Email) GetEmail() string {
	return c.val
}

func (c *Email) GetColumnName() string {
	return "email"
}

func (c *Email) GetPointer() interface{} {
	return &c.val
}

func Run() interface{} {
	var q struct {
		Email
	}
	return q
}

type Column interface {
	GetColumnName() string
	GetPointer() interface{}
}
