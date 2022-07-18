// Code generated by gola 0.0.2; DO NOT EDIT.

package song_user_favourites

import (
	"database/sql"
	"strings"

	"github.com/olachat/gola/coredb"

	"time"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

// SongUserFavourite represents song_user_favourites table
type SongUserFavourite struct {
	// User ID int unsigned
	UserId
	// Song ID int unsigned
	SongId
	// Is favourite tinyint
	IsFavourite
	// Create Time timestamp
	CreatedAt
	// Last Update Time timestamp
	UpdatedAt
}

type SongUserFavouriteTable struct{}

func (*SongUserFavouriteTable) GetTableName() string {
	return "song_user_favourites"
}

var table *SongUserFavouriteTable

// FetchSongUserFavouriteByPKs returns a row from song_user_favourites table with given primary key value
func FetchSongUserFavouriteByPK(val uint) *SongUserFavourite {
	return coredb.FetchByPK[SongUserFavourite](val, "user_id", _db)
}

// FetchByPKs returns a row with selected fields from song_user_favourites table with given primary key value
func FetchByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](val, "user_id", _db)
}

// FetchSongUserFavouriteByPKs returns rows with from song_user_favourites table with given primary key values
func FetchSongUserFavouriteByPKs(vals ...uint) []*SongUserFavourite {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[SongUserFavourite](pks, "user_id", _db)
}

// FetchByPKs returns rows with selected fields from song_user_favourites table with given primary key values
func FetchByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](pks, "user_id", _db)
}

// FindOneSongUserFavourite returns a row from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindOneSongUserFavourite(whereSQL string, params ...any) *SongUserFavourite {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[SongUserFavourite](w, _db)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM song_user_favourites "+whereSQL, _db, params...)
}

// FindOne returns a row with selected fields from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, _db)
}

// FindSongUserFavourite returns rows from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func FindSongUserFavourite(whereSQL string, params ...any) ([]*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[SongUserFavourite](w, _db)
}

// Find returns rows with selected fields from song_user_favourites table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, _db)
}

// Column types

// UserId field
// User ID
type UserId struct {
	_updated bool
	val      uint
}

func (c *UserId) GetUserId() uint {
	return c.val
}

func (c *UserId) SetUserId(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) IsUpdated() bool {
	return c._updated
}

func (c *UserId) IsPrimaryKey() bool {
	return true
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) GetTableType() coredb.TableType {
	return table
}

// SongId field
// Song ID
type SongId struct {
	_updated bool
	val      uint
}

func (c *SongId) GetSongId() uint {
	return c.val
}

func (c *SongId) SetSongId(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *SongId) GetColumnName() string {
	return "song_id"
}

func (c *SongId) IsUpdated() bool {
	return c._updated
}

func (c *SongId) IsPrimaryKey() bool {
	return false
}

func (c *SongId) GetValPointer() any {
	return &c.val
}

func (c *SongId) GetTableType() coredb.TableType {
	return table
}

// IsFavourite field
// Is favourite
type IsFavourite struct {
	_updated bool
	val      int8
}

func (c *IsFavourite) GetIsFavourite() int8 {
	return c.val
}

func (c *IsFavourite) SetIsFavourite(val int8) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsFavourite) GetColumnName() string {
	return "is_favourite"
}

func (c *IsFavourite) IsUpdated() bool {
	return c._updated
}

func (c *IsFavourite) IsPrimaryKey() bool {
	return false
}

func (c *IsFavourite) GetValPointer() any {
	return &c.val
}

func (c *IsFavourite) GetTableType() coredb.TableType {
	return table
}

// CreatedAt field
// Create Time
type CreatedAt struct {
	_updated bool
	val      time.Time
}

func (c *CreatedAt) GetCreatedAt() time.Time {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val time.Time) bool {
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
// Last Update Time
type UpdatedAt struct {
	_updated bool
	val      time.Time
}

func (c *UpdatedAt) GetUpdatedAt() time.Time {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val time.Time) bool {
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

func NewSongUserFavourite() *SongUserFavourite {
	return &SongUserFavourite{
		UserId{},
		SongId{},
		IsFavourite{val: int8(1)},
		CreatedAt{val: time.Now()},
		UpdatedAt{val: time.Now()},
	}
}

func (c *SongUserFavourite) Insert() error {
	sql := `INSERT INTO song_user_favourites (user_id, song_id, is_favourite, created_at, updated_at) values (?, ?, ?, ?, ?)`
	_, err := coredb.Exec(sql, _db, c.GetUserId(), c.GetSongId(), c.GetIsFavourite(), c.GetCreatedAt(), c.GetUpdatedAt())

	if err != nil {
		return err
	}

	return nil
}

func (c *SongUserFavourite) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if c.SongId.IsUpdated() {
		updatedFields = append(updatedFields, "song_id = ?")
		params = append(params, c.GetSongId())
	}
	if c.IsFavourite.IsUpdated() {
		updatedFields = append(updatedFields, "is_favourite = ?")
		params = append(params, c.GetIsFavourite())
	}
	if c.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "created_at = ?")
		params = append(params, c.GetCreatedAt())
	}
	if c.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "updated_at = ?")
		params = append(params, c.GetUpdatedAt())
	}

	sql := `UPDATE song_user_favourites SET `

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql = sql + strings.Join(updatedFields, ",") + " WHERE user_id = ?"
	params = append(params, c.GetUserId())

	_, err := coredb.Exec(sql, _db, params...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *SongUserFavourite) Delete() error {
	sql := `DELETE FROM song_user_favourites WHERE user_id = ?`

	_, err := coredb.Exec(sql, _db, c.GetUserId())
	return err
}

func Update[T any](obj *T) (bool, error) {
	return coredb.Update(obj, _db)
}
