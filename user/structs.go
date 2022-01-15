package user

import (
	"time"

	"github.com/olachat/gola/corelib"
)

type User struct {
	Id
	Name
	Email
	PhoneNumbers
	Created
}

func FetchUserById(id int) *User {
	return corelib.FetchById[User](id)
}

func FetchById[T any, PT corelib.PointerType[T]](id int) *T {
	return corelib.FetchById[T](id)
}

func FetchUserByIds(ids []int) []*User {
	return corelib.FetchByIds[User](ids)
}

func FetchByIds[T any, PT corelib.PointerType[T]](ids []int) []*T {
	return corelib.FetchByIds[T](ids)
}

type UserTable struct{}

func (*UserTable) GetTableName() string {
	return "users"
}

var _ corelib.ColumnType = &Name{}

var table *UserTable

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

func (c *Name) GetTableType() corelib.TableType {
	return table
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

func (c *Email) GetTableType() corelib.TableType {
	return table
}
