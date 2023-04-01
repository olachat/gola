// Code generated by gola 0.1.1; DO NOT EDIT.

package gifts_nn_with_default

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"

	"time"
)

const DBName string = "testdata"
const TableName string = "gifts_nn_with_default"

// GiftsNnWithDefault represents `gifts_nn_with_default` table
type GiftsNnWithDefault struct {
	//  int(10) unsigned
	Id `json:"id"`
	// gift name varchar(100)
	Name `json:"name"`
	// is free gift tinyint(1)
	IsFree `json:"is_free"`
	//  smallint(6)
	GiftCount `json:"gift_count"`
	//  enum('','freebie','sovenir','membership')
	GiftType `json:"gift_type"`
	//  bigint(20)
	CreateTime `json:"create_time"`
	//  float unsigned
	Discount `json:"discount"`
	//  double unsigned
	Price `json:"price"`
	//  varchar(128)
	Remark `json:"remark"`
	//  varbinary(255) binary
	Manifest `json:"manifest"`
	//  text
	Description `json:"description"`
	//  timestamp
	UpdateTime `json:"update_time"`
	//  timestamp
	UpdateTime2 `json:"update_time2"`
	//  set('orchard','vivo','sentosa','changi')
	Branches `json:"branches"`
}

type withPK interface {
	GetId() uint
}

// FetchByPK returns a row from `gifts_nn_with_default` table with given primary key value
func FetchByPK(val uint) *GiftsNnWithDefault {
	return coredb.FetchByPK[GiftsNnWithDefault](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from gifts_nn_with_default table with given primary key value
func FetchFieldsByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `gifts_nn_with_default` table with given primary key values
func FetchByPKs(vals ...uint) []*GiftsNnWithDefault {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[GiftsNnWithDefault](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `gifts_nn_with_default` table with given primary key values
func FetchFieldsByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `gifts_nn_with_default` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *GiftsNnWithDefault {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[GiftsNnWithDefault](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `gifts_nn_with_default` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `gifts_nn_with_default` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*GiftsNnWithDefault, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[GiftsNnWithDefault](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `gifts_nn_with_default` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `gifts_nn_with_default` "+whereSQL, params...)
}

// Column types
type GiftsNnWithDefaultGiftType string

const (
	GiftsNnWithDefaultGiftTypeEmpty      GiftsNnWithDefaultGiftType = ""
	GiftsNnWithDefaultGiftTypeFreebie    GiftsNnWithDefaultGiftType = "freebie"
	GiftsNnWithDefaultGiftTypeSovenir    GiftsNnWithDefaultGiftType = "sovenir"
	GiftsNnWithDefaultGiftTypeMembership GiftsNnWithDefaultGiftType = "membership"
)

type GiftsNnWithDefaultBranches string

const (
	GiftsNnWithDefaultBranchesOrchard GiftsNnWithDefaultBranches = "orchard"
	GiftsNnWithDefaultBranchesVivo    GiftsNnWithDefaultBranches = "vivo"
	GiftsNnWithDefaultBranchesSentosa GiftsNnWithDefaultBranches = "sentosa"
	GiftsNnWithDefaultBranchesChangi  GiftsNnWithDefaultBranches = "changi"
)

var GiftsNnWithDefaultBranchesList = []string{
	"orchard",
	"vivo",
	"sentosa",
	"changi",
}

// Id field
type Id struct {
	isAssigned bool
	val        uint
}

func (c *Id) GetId() uint {
	return c.val
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) GetValPointer() any {
	return &c.val
}

func (c *Id) getIdForDB() uint {
	return c.val
}

func (c *Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Id) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Name field
// gift name
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

func (c *Name) IsUpdated() bool {
	return c._updated
}

func (c *Name) resetUpdated() {
	c._updated = false
}

func (c *Name) GetColumnName() string {
	return "name"
}

func (c *Name) GetValPointer() any {
	return &c.val
}

func (c *Name) getNameForDB() string {
	return c.val
}

func (c *Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Name) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// IsFree field
// is free gift
type IsFree struct {
	_updated bool
	val      bool
}

func (c *IsFree) GetIsFree() bool {
	return c.val
}

func (c *IsFree) SetIsFree(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsFree) IsUpdated() bool {
	return c._updated
}

func (c *IsFree) resetUpdated() {
	c._updated = false
}

func (c *IsFree) GetColumnName() string {
	return "is_free"
}

func (c *IsFree) GetValPointer() any {
	return &c.val
}

func (c *IsFree) getIsFreeForDB() bool {
	return c.val
}

func (c *IsFree) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *IsFree) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// GiftCount field
type GiftCount struct {
	_updated bool
	val      int16
}

func (c *GiftCount) GetGiftCount() int16 {
	return c.val
}

func (c *GiftCount) SetGiftCount(val int16) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *GiftCount) IsUpdated() bool {
	return c._updated
}

func (c *GiftCount) resetUpdated() {
	c._updated = false
}

func (c *GiftCount) GetColumnName() string {
	return "gift_count"
}

func (c *GiftCount) GetValPointer() any {
	return &c.val
}

func (c *GiftCount) getGiftCountForDB() int16 {
	return c.val
}

func (c *GiftCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *GiftCount) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// GiftType field
type GiftType struct {
	_updated bool
	val      GiftsNnWithDefaultGiftType
}

func (c *GiftType) GetGiftType() GiftsNnWithDefaultGiftType {
	return c.val
}

func (c *GiftType) SetGiftType(val GiftsNnWithDefaultGiftType) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *GiftType) IsUpdated() bool {
	return c._updated
}

func (c *GiftType) resetUpdated() {
	c._updated = false
}

func (c *GiftType) GetColumnName() string {
	return "gift_type"
}

func (c *GiftType) GetValPointer() any {
	return &c.val
}

func (c *GiftType) getGiftTypeForDB() GiftsNnWithDefaultGiftType {
	return c.val
}

func (c *GiftType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *GiftType) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// CreateTime field
type CreateTime struct {
	_updated bool
	val      int64
}

func (c *CreateTime) GetCreateTime() int64 {
	return c.val
}

func (c *CreateTime) SetCreateTime(val int64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CreateTime) IsUpdated() bool {
	return c._updated
}

func (c *CreateTime) resetUpdated() {
	c._updated = false
}

func (c *CreateTime) GetColumnName() string {
	return "create_time"
}

func (c *CreateTime) GetValPointer() any {
	return &c.val
}

func (c *CreateTime) getCreateTimeForDB() int64 {
	return c.val
}

func (c *CreateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *CreateTime) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Discount field
type Discount struct {
	_updated bool
	val      float64
}

func (c *Discount) GetDiscount() float64 {
	return c.val
}

func (c *Discount) SetDiscount(val float64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Discount) IsUpdated() bool {
	return c._updated
}

func (c *Discount) resetUpdated() {
	c._updated = false
}

func (c *Discount) GetColumnName() string {
	return "discount"
}

func (c *Discount) GetValPointer() any {
	return &c.val
}

func (c *Discount) getDiscountForDB() float64 {
	return c.val
}

func (c *Discount) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Discount) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Price field
type Price struct {
	_updated bool
	val      float64
}

func (c *Price) GetPrice() float64 {
	return c.val
}

func (c *Price) SetPrice(val float64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Price) IsUpdated() bool {
	return c._updated
}

func (c *Price) resetUpdated() {
	c._updated = false
}

func (c *Price) GetColumnName() string {
	return "price"
}

func (c *Price) GetValPointer() any {
	return &c.val
}

func (c *Price) getPriceForDB() float64 {
	return c.val
}

func (c *Price) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Price) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Remark field
type Remark struct {
	_updated bool
	val      string
}

func (c *Remark) GetRemark() string {
	return c.val
}

func (c *Remark) SetRemark(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Remark) IsUpdated() bool {
	return c._updated
}

func (c *Remark) resetUpdated() {
	c._updated = false
}

func (c *Remark) GetColumnName() string {
	return "remark"
}

func (c *Remark) GetValPointer() any {
	return &c.val
}

func (c *Remark) getRemarkForDB() string {
	return c.val
}

func (c *Remark) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Remark) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Manifest field
type Manifest struct {
	_updated bool
	val      []byte
}

func (c *Manifest) GetManifest() []byte {
	return c.val
}

func (c *Manifest) SetManifest(val []byte) bool {
	c._updated = true
	c.val = val
	return true
}

func (c *Manifest) IsUpdated() bool {
	return c._updated
}

func (c *Manifest) resetUpdated() {
	c._updated = false
}

func (c *Manifest) GetColumnName() string {
	return "manifest"
}

func (c *Manifest) GetValPointer() any {
	return &c.val
}

func (c *Manifest) getManifestForDB() []byte {
	return c.val
}

func (c *Manifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Manifest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Description field
type Description struct {
	_updated bool
	val      string
}

func (c *Description) GetDescription() string {
	return c.val
}

func (c *Description) SetDescription(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Description) IsUpdated() bool {
	return c._updated
}

func (c *Description) resetUpdated() {
	c._updated = false
}

func (c *Description) GetColumnName() string {
	return "description"
}

func (c *Description) GetValPointer() any {
	return &c.val
}

func (c *Description) getDescriptionForDB() string {
	return c.val
}

func (c *Description) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Description) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// UpdateTime field
type UpdateTime struct {
	_updated bool
	val      time.Time
}

func (c *UpdateTime) GetUpdateTime() time.Time {
	return c.val
}

func (c *UpdateTime) SetUpdateTime(val time.Time) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdateTime) IsUpdated() bool {
	return c._updated
}

func (c *UpdateTime) resetUpdated() {
	c._updated = false
}

func (c *UpdateTime) GetColumnName() string {
	return "update_time"
}

func (c *UpdateTime) GetValPointer() any {
	return &c.val
}

func (c *UpdateTime) getUpdateTimeForDB() time.Time {
	return c.val
}

func (c *UpdateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UpdateTime) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// UpdateTime2 field
type UpdateTime2 struct {
	_updated bool
	val      time.Time
}

func (c *UpdateTime2) GetUpdateTime2() time.Time {
	return c.val
}

func (c *UpdateTime2) SetUpdateTime2(val time.Time) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdateTime2) IsUpdated() bool {
	return c._updated
}

func (c *UpdateTime2) resetUpdated() {
	c._updated = false
}

func (c *UpdateTime2) GetColumnName() string {
	return "update_time2"
}

func (c *UpdateTime2) GetValPointer() any {
	return &c.val
}

func (c *UpdateTime2) getUpdateTime2ForDB() time.Time {
	return c.val
}

func (c *UpdateTime2) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UpdateTime2) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Branches field
type Branches struct {
	_updated bool
	val      string
}

func (c *Branches) GetBranches() []GiftsNnWithDefaultBranches {
	strSlice := strings.Split(c.val, ",")
	if len(strSlice) == 1 && !coredb.ValueInSet(GiftsNnWithDefaultBranchesList, strSlice[0]) {
		return []GiftsNnWithDefaultBranches{}
	}
	valSlice := make([]GiftsNnWithDefaultBranches, 0, len(strSlice))
	for _, s := range strSlice {
		valSlice = append(valSlice, GiftsNnWithDefaultBranches(strings.ToLower(s)))
	}
	return valSlice
}

func (c *Branches) SetBranches(val []GiftsNnWithDefaultBranches) bool {
	strSlice := make([]string, 0, len(val))
	for _, v := range val {
		strSlice = append(strSlice, string(v))
	}
	c.val = strings.Join(strSlice, ",")
	c._updated = true
	return true
}

func (c *Branches) IsUpdated() bool {
	return c._updated
}

func (c *Branches) resetUpdated() {
	c._updated = false
}

func (c *Branches) GetColumnName() string {
	return "branches"
}

func (c *Branches) GetValPointer() any {
	return &c.val
}

func (c *Branches) getBranchesForDB() string {
	return c.val
}

func (c *Branches) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Branches) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// New return new *GiftsNnWithDefault with default values
func New() *GiftsNnWithDefault {
	return &GiftsNnWithDefault{
		Id{},
		Name{val: "gift for you"},
		IsFree{val: true},
		GiftCount{val: int16(1)},
		GiftType{val: "membership"},
		CreateTime{val: int64(999)},
		Discount{val: float64(0.1)},
		Price{val: float64(5.0)},
		Remark{val: "hope you like it"},
		Manifest{val: []byte("manifest data")},
		Description{},
		UpdateTime{val: coredb.MustParseTime("2023-01-19 03:14:07.999999")},
		UpdateTime2{val: time.Now()},
		Branches{val: "sentosa,changi"},
	}
}

// NewWithPK takes "id"
// and returns new *GiftsNnWithDefault with given PK
func NewWithPK(val uint) *GiftsNnWithDefault {
	c := &GiftsNnWithDefault{
		Id{},
		Name{val: "gift for you"},
		IsFree{val: true},
		GiftCount{val: int16(1)},
		GiftType{val: "membership"},
		CreateTime{val: int64(999)},
		Discount{val: float64(0.1)},
		Price{val: float64(5.0)},
		Remark{val: "hope you like it"},
		Manifest{val: []byte("manifest data")},
		Description{},
		UpdateTime{val: coredb.MustParseTime("2023-01-19 03:14:07.999999")},
		UpdateTime2{val: time.Now()},
		Branches{val: "sentosa,changi"},
	}
	c.Id.val = val
	c.Id.isAssigned = true
	return c
}

const insertWithoutPK string = "INSERT INTO `gifts_nn_with_default` (`name`, `is_free`, `gift_count`, `gift_type`, `create_time`, `discount`, `price`, `remark`, `manifest`, `description`, `update_time`, `update_time2`, `branches`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
const insertWithPK string = "INSERT INTO `gifts_nn_with_default` (`id`, `name`, `is_free`, `gift_count`, `gift_type`, `create_time`, `discount`, `price`, `remark`, `manifest`, `description`, `update_time`, `update_time2`, `branches`) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

// Insert GiftsNnWithDefault struct to `gifts_nn_with_default` table
func (c *GiftsNnWithDefault) Insert() error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.Exec(DBName, insertWithPK, c.getIdForDB(), c.getNameForDB(), c.getIsFreeForDB(), c.getGiftCountForDB(), c.getGiftTypeForDB(), c.getCreateTimeForDB(), c.getDiscountForDB(), c.getPriceForDB(), c.getRemarkForDB(), c.getManifestForDB(), c.getDescriptionForDB(), c.getUpdateTimeForDB(), c.getUpdateTime2ForDB(), c.getBranchesForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.Exec(DBName, insertWithoutPK, c.getNameForDB(), c.getIsFreeForDB(), c.getGiftCountForDB(), c.getGiftTypeForDB(), c.getCreateTimeForDB(), c.getDiscountForDB(), c.getPriceForDB(), c.getRemarkForDB(), c.getManifestForDB(), c.getDescriptionForDB(), c.getUpdateTimeForDB(), c.getUpdateTime2ForDB(), c.getBranchesForDB())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = uint(id)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return coredb.ErrAvoidInsert
	}

	c.resetUpdated()
	return nil
}

func (c *GiftsNnWithDefault) resetUpdated() {
	c.Name.resetUpdated()
	c.IsFree.resetUpdated()
	c.GiftCount.resetUpdated()
	c.GiftType.resetUpdated()
	c.CreateTime.resetUpdated()
	c.Discount.resetUpdated()
	c.Price.resetUpdated()
	c.Remark.resetUpdated()
	c.Manifest.resetUpdated()
	c.Description.resetUpdated()
	c.UpdateTime.resetUpdated()
	c.UpdateTime2.resetUpdated()
	c.Branches.resetUpdated()
}

// Update GiftsNnWithDefault struct in `gifts_nn_with_default` table
func (obj *GiftsNnWithDefault) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Name.IsUpdated() {
		updatedFields = append(updatedFields, "`name` = ?")
		params = append(params, obj.getNameForDB())
	}
	if obj.IsFree.IsUpdated() {
		updatedFields = append(updatedFields, "`is_free` = ?")
		params = append(params, obj.getIsFreeForDB())
	}
	if obj.GiftCount.IsUpdated() {
		updatedFields = append(updatedFields, "`gift_count` = ?")
		params = append(params, obj.getGiftCountForDB())
	}
	if obj.GiftType.IsUpdated() {
		updatedFields = append(updatedFields, "`gift_type` = ?")
		params = append(params, obj.getGiftTypeForDB())
	}
	if obj.CreateTime.IsUpdated() {
		updatedFields = append(updatedFields, "`create_time` = ?")
		params = append(params, obj.getCreateTimeForDB())
	}
	if obj.Discount.IsUpdated() {
		updatedFields = append(updatedFields, "`discount` = ?")
		params = append(params, obj.getDiscountForDB())
	}
	if obj.Price.IsUpdated() {
		updatedFields = append(updatedFields, "`price` = ?")
		params = append(params, obj.getPriceForDB())
	}
	if obj.Remark.IsUpdated() {
		updatedFields = append(updatedFields, "`remark` = ?")
		params = append(params, obj.getRemarkForDB())
	}
	if obj.Manifest.IsUpdated() {
		updatedFields = append(updatedFields, "`manifest` = ?")
		params = append(params, obj.getManifestForDB())
	}
	if obj.Description.IsUpdated() {
		updatedFields = append(updatedFields, "`description` = ?")
		params = append(params, obj.getDescriptionForDB())
	}
	if obj.UpdateTime.IsUpdated() {
		updatedFields = append(updatedFields, "`update_time` = ?")
		params = append(params, obj.getUpdateTimeForDB())
	}
	if obj.UpdateTime2.IsUpdated() {
		updatedFields = append(updatedFields, "`update_time2` = ?")
		params = append(params, obj.getUpdateTime2ForDB())
	}
	if obj.Branches.IsUpdated() {
		updatedFields = append(updatedFields, "`branches` = ?")
		params = append(params, obj.getBranchesForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `gifts_nn_with_default` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.Exec(DBName, sql, params...)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, coredb.ErrAvoidUpdate
	}

	obj.resetUpdated()
	return true, nil
}

// Update GiftsNnWithDefault struct with given fields in `gifts_nn_with_default` table
func Update(obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *Name:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`name` = ?")
				params = append(params, c.getNameForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsFree:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_free` = ?")
				params = append(params, c.getIsFreeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *GiftCount:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`gift_count` = ?")
				params = append(params, c.getGiftCountForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *GiftType:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`gift_type` = ?")
				params = append(params, c.getGiftTypeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreateTime:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`create_time` = ?")
				params = append(params, c.getCreateTimeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Discount:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`discount` = ?")
				params = append(params, c.getDiscountForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Price:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`price` = ?")
				params = append(params, c.getPriceForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Remark:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`remark` = ?")
				params = append(params, c.getRemarkForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Manifest:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`manifest` = ?")
				params = append(params, c.getManifestForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Description:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`description` = ?")
				params = append(params, c.getDescriptionForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdateTime:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`update_time` = ?")
				params = append(params, c.getUpdateTimeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdateTime2:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`update_time2` = ?")
				params = append(params, c.getUpdateTime2ForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Branches:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`branches` = ?")
				params = append(params, c.getBranchesForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `gifts_nn_with_default` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.Exec(DBName, sql, params...)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, coredb.ErrAvoidUpdate
	}

	for _, f := range resetFuncs {
		f()
	}
	return true, nil
}

const deleteSql string = "DELETE FROM `gifts_nn_with_default` WHERE `id` = ?"

// DeleteByPK delete a row from gifts_nn_with_default table with given primary key value
func DeleteByPK(val uint) error {
	_, err := coredb.Exec(DBName, deleteSql, val)
	return err
}
