// Code generated by gola 0.0.1; DO NOT EDIT.

package users

import (
	"context"
	"github.com/olachat/gola/corelib"
)

// User represents users table
type User struct {
	//  int
	Id
	// Name varchar(255)
	Name
	// Email address varchar(255)
	Email
	// Created Timestamp int unsigned
	CreatedAt
	// Updated Timestamp int unsigned
	UpdatedAt
	// float float
	FloatType
	// double double
	DoubleType
	// user hobby enum('swimming','running','singing')
	Hobby
}

type UserTable struct{}

func (*UserTable) GetTableName(ctx context.Context) string {
	return "users"
}

var table *UserTable

// Fetch methods
func FetchUserById(ctx context.Context, id int) *User {
	return corelib.FetchById[User](ctx, id)
}

func FetchById[T any, PT corelib.PointerType[T]](ctx context.Context, id int) *T {
	return corelib.FetchById[T](ctx, id)
}

func FetchUserByIds(ctx context.Context, ids []int) []*User {
	return corelib.FetchByIds[User](ctx, ids)
}

func FetchByIds[T any, PT corelib.PointerType[T]](ctx context.Context, ids []int) []*T {
	return corelib.FetchByIds[T](ctx, ids)
}

// Column types
type UserHobby string

const (
	UserHobbySwimming UserHobby = "swimming"
	UserHobbyRunning  UserHobby = "running"
	UserHobbySinging  UserHobby = "singing"
)

// Id field
//
type Id struct {
	val int
}

func (c *Id) GetId(ctx context.Context) int {
	return c.val
}

func (c *Id) SetId(ctx context.Context, val int) {
	c.val = val
}

func (c *Id) GetColumnName(ctx context.Context) string {
	return "id"
}

func (c *Id) IsPrimaryKey(ctx context.Context) bool {
	return true
}

func (c *Id) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *Id) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// Name field
// Name
type Name struct {
	val string
}

func (c *Name) GetName(ctx context.Context) string {
	return c.val
}

func (c *Name) SetName(ctx context.Context, val string) {
	c.val = val
}

func (c *Name) GetColumnName(ctx context.Context) string {
	return "name"
}

func (c *Name) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *Name) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *Name) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// Email field
// Email address
type Email struct {
	val string
}

func (c *Email) GetEmail(ctx context.Context) string {
	return c.val
}

func (c *Email) SetEmail(ctx context.Context, val string) {
	c.val = val
}

func (c *Email) GetColumnName(ctx context.Context) string {
	return "email"
}

func (c *Email) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *Email) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *Email) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// CreatedAt field
// Created Timestamp
type CreatedAt struct {
	val uint
}

func (c *CreatedAt) GetCreatedAt(ctx context.Context) uint {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(ctx context.Context, val uint) {
	c.val = val
}

func (c *CreatedAt) GetColumnName(ctx context.Context) string {
	return "created_at"
}

func (c *CreatedAt) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *CreatedAt) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *CreatedAt) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// UpdatedAt field
// Updated Timestamp
type UpdatedAt struct {
	val uint
}

func (c *UpdatedAt) GetUpdatedAt(ctx context.Context) uint {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(ctx context.Context, val uint) {
	c.val = val
}

func (c *UpdatedAt) GetColumnName(ctx context.Context) string {
	return "updated_at"
}

func (c *UpdatedAt) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *UpdatedAt) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *UpdatedAt) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// FloatType field
// float
type FloatType struct {
	val float32
}

func (c *FloatType) GetFloatType(ctx context.Context) float32 {
	return c.val
}

func (c *FloatType) SetFloatType(ctx context.Context, val float32) {
	c.val = val
}

func (c *FloatType) GetColumnName(ctx context.Context) string {
	return "float_type"
}

func (c *FloatType) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *FloatType) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *FloatType) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// DoubleType field
// double
type DoubleType struct {
	val float64
}

func (c *DoubleType) GetDoubleType(ctx context.Context) float64 {
	return c.val
}

func (c *DoubleType) SetDoubleType(ctx context.Context, val float64) {
	c.val = val
}

func (c *DoubleType) GetColumnName(ctx context.Context) string {
	return "double_type"
}

func (c *DoubleType) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *DoubleType) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *DoubleType) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

// Hobby field
// user hobby
type Hobby struct {
	val UserHobby
}

func (c *Hobby) GetHobby(ctx context.Context) UserHobby {
	return c.val
}

func (c *Hobby) SetHobby(ctx context.Context, val UserHobby) {
	c.val = val
}

func (c *Hobby) GetColumnName(ctx context.Context) string {
	return "hobby"
}

func (c *Hobby) IsPrimaryKey(ctx context.Context) bool {
	return false
}

func (c *Hobby) GetValPointer(ctx context.Context) interface{} {
	return &c.val
}

func (c *Hobby) GetTableType(ctx context.Context) corelib.TableType {
	return table
}

func NewUser(ctx context.Context) *User {
	return &User{
		Id{},
		Name{val: ""},
		Email{val: ""},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
		FloatType{val: float32(0)},
		DoubleType{val: float64(0)},
		Hobby{val: "swimming"},
	}
}
