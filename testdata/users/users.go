// Code generated by gola 0.0.2; DO NOT EDIT.

package users

import (
	"database/sql"
	"strings"

	"github.com/olachat/gola/coredb"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

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
	// user hobby enum('swimming','running','singing')
	HobbyNoDefault
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	Sports
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	Sports2
	// user sports set('swim','tennis','basketball','football','squash','badminton')
	SportsNoDefault
}

type UserTable struct{}

func (*UserTable) GetTableName() string {
	return "users"
}

var table *UserTable

// FetchUserByPKs returns a row from users table with given primary key value
func FetchUserByPK(val int) *User {
	return coredb.FetchByPK[User](val, "id", _db)
}

// FetchByPKs returns a row with selected fields from users table with given primary key value
func FetchByPK[T any](val int) *T {
	return coredb.FetchByPK[T](val, "id", _db)
}

// FetchUserByPKs returns rows with from users table with given primary key values
func FetchUserByPKs(vals ...int) []*User {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[User](pks, "id", _db)
}

// FetchByPKs returns rows with selected fields from users table with given primary key values
func FetchByPKs[T any](vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](pks, "id", _db)
}

// FindOneUser returns a row from users table with arbitary where query
// whereSQL must start with "where ..."
func FindOneUser(whereSQL string, params ...any) *User {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[User](w, _db)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM users "+whereSQL, _db, params...)
}

// FindOne returns a row with selected fields from users table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, _db)
}

// FindUser returns rows from users table with arbitary where query
// whereSQL must start with "where ..."
func FindUser(whereSQL string, params ...any) ([]*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[User](w, _db)
}

// Find returns rows with selected fields from users table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, _db)
}

// Column types
type UserHobby string

const (
	UserHobbySwimming UserHobby = "swimming"
	UserHobbyRunning  UserHobby = "running"
	UserHobbySinging  UserHobby = "singing"
)

type UserHobbyNoDefault string

const (
	UserHobbyNoDefaultSwimming UserHobbyNoDefault = "swimming"
	UserHobbyNoDefaultRunning  UserHobbyNoDefault = "running"
	UserHobbyNoDefaultSinging  UserHobbyNoDefault = "singing"
)

type UserSports string

const (
	UserSportsSwim       UserSports = "swim"
	UserSportsTennis     UserSports = "tennis"
	UserSportsBasketball UserSports = "basketball"
	UserSportsFootball   UserSports = "football"
	UserSportsSquash     UserSports = "squash"
	UserSportsBadminton  UserSports = "badminton"
)

type UserSports2 string

const (
	UserSports2Swim       UserSports2 = "swim"
	UserSports2Tennis     UserSports2 = "tennis"
	UserSports2Basketball UserSports2 = "basketball"
	UserSports2Football   UserSports2 = "football"
	UserSports2Squash     UserSports2 = "squash"
	UserSports2Badminton  UserSports2 = "badminton"
)

type UserSportsNoDefault string

const (
	UserSportsNoDefaultSwim       UserSportsNoDefault = "swim"
	UserSportsNoDefaultTennis     UserSportsNoDefault = "tennis"
	UserSportsNoDefaultBasketball UserSportsNoDefault = "basketball"
	UserSportsNoDefaultFootball   UserSportsNoDefault = "football"
	UserSportsNoDefaultSquash     UserSportsNoDefault = "squash"
	UserSportsNoDefaultBadminton  UserSportsNoDefault = "badminton"
)

// Id field
//
type Id struct {
	_updated bool
	val      int
}

func (c *Id) GetId() int {
	return c.val
}

func (c *Id) SetId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) IsUpdated() bool {
	return c._updated
}

func (c *Id) IsPrimaryKey() bool {
	return true
}

func (c *Id) GetValPointer() any {
	return &c.val
}

func (c *Id) GetTableType() coredb.TableType {
	return table
}

// Name field
// Name
type Name struct {
	_updated bool
	val      string
}

func (c *Name) GetName() string {
	return c.val
}

func (c *Name) SetName(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Name) GetColumnName() string {
	return "name"
}

func (c *Name) IsUpdated() bool {
	return c._updated
}

func (c *Name) IsPrimaryKey() bool {
	return false
}

func (c *Name) GetValPointer() any {
	return &c.val
}

func (c *Name) GetTableType() coredb.TableType {
	return table
}

// Email field
// Email address
type Email struct {
	_updated bool
	val      string
}

func (c *Email) GetEmail() string {
	return c.val
}

func (c *Email) SetEmail(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Email) GetColumnName() string {
	return "email"
}

func (c *Email) IsUpdated() bool {
	return c._updated
}

func (c *Email) IsPrimaryKey() bool {
	return false
}

func (c *Email) GetValPointer() any {
	return &c.val
}

func (c *Email) GetTableType() coredb.TableType {
	return table
}

// CreatedAt field
// Created Timestamp
type CreatedAt struct {
	_updated bool
	val      uint
}

func (c *CreatedAt) GetCreatedAt() uint {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CreatedAt) GetColumnName() string {
	return "created_at"
}

func (c *CreatedAt) IsUpdated() bool {
	return c._updated
}

func (c *CreatedAt) IsPrimaryKey() bool {
	return false
}

func (c *CreatedAt) GetValPointer() any {
	return &c.val
}

func (c *CreatedAt) GetTableType() coredb.TableType {
	return table
}

// UpdatedAt field
// Updated Timestamp
type UpdatedAt struct {
	_updated bool
	val      uint
}

func (c *UpdatedAt) GetUpdatedAt() uint {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdatedAt) GetColumnName() string {
	return "updated_at"
}

func (c *UpdatedAt) IsUpdated() bool {
	return c._updated
}

func (c *UpdatedAt) IsPrimaryKey() bool {
	return false
}

func (c *UpdatedAt) GetValPointer() any {
	return &c.val
}

func (c *UpdatedAt) GetTableType() coredb.TableType {
	return table
}

// FloatType field
// float
type FloatType struct {
	_updated bool
	val      float32
}

func (c *FloatType) GetFloatType() float32 {
	return c.val
}

func (c *FloatType) SetFloatType(val float32) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *FloatType) GetColumnName() string {
	return "float_type"
}

func (c *FloatType) IsUpdated() bool {
	return c._updated
}

func (c *FloatType) IsPrimaryKey() bool {
	return false
}

func (c *FloatType) GetValPointer() any {
	return &c.val
}

func (c *FloatType) GetTableType() coredb.TableType {
	return table
}

// DoubleType field
// double
type DoubleType struct {
	_updated bool
	val      float64
}

func (c *DoubleType) GetDoubleType() float64 {
	return c.val
}

func (c *DoubleType) SetDoubleType(val float64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *DoubleType) GetColumnName() string {
	return "double_type"
}

func (c *DoubleType) IsUpdated() bool {
	return c._updated
}

func (c *DoubleType) IsPrimaryKey() bool {
	return false
}

func (c *DoubleType) GetValPointer() any {
	return &c.val
}

func (c *DoubleType) GetTableType() coredb.TableType {
	return table
}

// Hobby field
// user hobby
type Hobby struct {
	_updated bool
	val      UserHobby
}

func (c *Hobby) GetHobby() UserHobby {
	return c.val
}

func (c *Hobby) SetHobby(val UserHobby) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Hobby) GetColumnName() string {
	return "hobby"
}

func (c *Hobby) IsUpdated() bool {
	return c._updated
}

func (c *Hobby) IsPrimaryKey() bool {
	return false
}

func (c *Hobby) GetValPointer() any {
	return &c.val
}

func (c *Hobby) GetTableType() coredb.TableType {
	return table
}

// HobbyNoDefault field
// user hobby
type HobbyNoDefault struct {
	_updated bool
	val      UserHobbyNoDefault
}

func (c *HobbyNoDefault) GetHobbyNoDefault() UserHobbyNoDefault {
	return c.val
}

func (c *HobbyNoDefault) SetHobbyNoDefault(val UserHobbyNoDefault) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *HobbyNoDefault) GetColumnName() string {
	return "hobby_no_default"
}

func (c *HobbyNoDefault) IsUpdated() bool {
	return c._updated
}

func (c *HobbyNoDefault) IsPrimaryKey() bool {
	return false
}

func (c *HobbyNoDefault) GetValPointer() any {
	return &c.val
}

func (c *HobbyNoDefault) GetTableType() coredb.TableType {
	return table
}

// Sports field
// user sports
type Sports struct {
	_updated bool
	val      string
}

func (c *Sports) GetSports() []UserSports {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSports, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSports(strings.ToLower(s)))
	}
	return valSlice
}

func (c *Sports) SetSports(val []UserSports) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *Sports) GetColumnName() string {
	return "sports"
}

func (c *Sports) IsUpdated() bool {
	return c._updated
}

func (c *Sports) IsPrimaryKey() bool {
	return false
}

func (c *Sports) GetValPointer() any {
	return &c.val
}

func (c *Sports) GetTableType() coredb.TableType {
	return table
}

// Sports2 field
// user sports
type Sports2 struct {
	_updated bool
	val      string
}

func (c *Sports2) GetSports2() []UserSports2 {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSports2, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSports2(strings.ToLower(s)))
	}
	return valSlice
}

func (c *Sports2) SetSports2(val []UserSports2) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *Sports2) GetColumnName() string {
	return "sports2"
}

func (c *Sports2) IsUpdated() bool {
	return c._updated
}

func (c *Sports2) IsPrimaryKey() bool {
	return false
}

func (c *Sports2) GetValPointer() any {
	return &c.val
}

func (c *Sports2) GetTableType() coredb.TableType {
	return table
}

// SportsNoDefault field
// user sports
type SportsNoDefault struct {
	_updated bool
	val      string
}

func (c *SportsNoDefault) GetSportsNoDefault() []UserSportsNoDefault {
	strSlice := strings.Split(c.val, ",")
	valSlice := make([]UserSportsNoDefault, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, UserSportsNoDefault(strings.ToLower(s)))
	}
	return valSlice
}

func (c *SportsNoDefault) SetSportsNoDefault(val []UserSportsNoDefault) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	return true
}

func (c *SportsNoDefault) GetColumnName() string {
	return "sports_no_default"
}

func (c *SportsNoDefault) IsUpdated() bool {
	return c._updated
}

func (c *SportsNoDefault) IsPrimaryKey() bool {
	return false
}

func (c *SportsNoDefault) GetValPointer() any {
	return &c.val
}

func (c *SportsNoDefault) GetTableType() coredb.TableType {
	return table
}

func NewUser() *User {
	return &User{
		Id{},
		Name{val: ""},
		Email{val: ""},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
		FloatType{val: float32(0)},
		DoubleType{val: float64(0)},
		Hobby{val: "swimming"},
		HobbyNoDefault{},
		Sports{val: "swim,football"},
		Sports2{val: "swim,football"},
		SportsNoDefault{},
	}
}

func (c *User) Insert() error {
	sql := `INSERT INTO users (name, email, created_at, updated_at, float_type, double_type, hobby, hobby_no_default, sports, sports2, sports_no_default) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := coredb.Exec(sql, _db, c.GetName(), c.GetEmail(), c.GetCreatedAt(), c.GetUpdatedAt(), c.GetFloatType(), c.GetDoubleType(), c.GetHobby(), c.GetHobbyNoDefault(), c.GetSports(), c.GetSports2(), c.GetSportsNoDefault())

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.SetId(int(id))

	return nil
}

func (c *User) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if c.Name.IsUpdated() {
		updatedFields = append(updatedFields, "name = ?")
		params = append(params, c.GetName())
	}
	if c.Email.IsUpdated() {
		updatedFields = append(updatedFields, "email = ?")
		params = append(params, c.GetEmail())
	}
	if c.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "created_at = ?")
		params = append(params, c.GetCreatedAt())
	}
	if c.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "updated_at = ?")
		params = append(params, c.GetUpdatedAt())
	}
	if c.FloatType.IsUpdated() {
		updatedFields = append(updatedFields, "float_type = ?")
		params = append(params, c.GetFloatType())
	}
	if c.DoubleType.IsUpdated() {
		updatedFields = append(updatedFields, "double_type = ?")
		params = append(params, c.GetDoubleType())
	}
	if c.Hobby.IsUpdated() {
		updatedFields = append(updatedFields, "hobby = ?")
		params = append(params, c.GetHobby())
	}
	if c.HobbyNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "hobby_no_default = ?")
		params = append(params, c.GetHobbyNoDefault())
	}
	if c.Sports.IsUpdated() {
		updatedFields = append(updatedFields, "sports = ?")
		params = append(params, c.GetSports())
	}
	if c.Sports2.IsUpdated() {
		updatedFields = append(updatedFields, "sports2 = ?")
		params = append(params, c.GetSports2())
	}
	if c.SportsNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "sports_no_default = ?")
		params = append(params, c.GetSportsNoDefault())
	}

	sql := `UPDATE users SET `

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql = sql + strings.Join(updatedFields, ",") + " WHERE id = ?"
	params = append(params, c.GetId())

	_, err := coredb.Exec(sql, _db, params...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *User) Delete() error {
	sql := `DELETE FROM users WHERE id = ?`

	_, err := coredb.Exec(sql, _db, c.GetId())
	return err
}

func Update[T any](obj *T) (bool, error) {
	return coredb.Update(obj, _db)
}
