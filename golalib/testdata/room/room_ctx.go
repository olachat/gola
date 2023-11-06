// Code generated by gola 0.1.1; DO NOT EDIT.

package room

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `room` table with given primary key value
func FetchByPKCtx(ctx context.Context, val uint) (*Room, error) {
	return coredb.FetchByPKCtx[Room](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from room table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val uint) (*T, error) {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `room` table with given primary key values
func FetchByPKsCtx(ctx context.Context, vals ...uint) ([]*Room, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[Room](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `room` table with given primary key values
func FetchFieldsByPKsCtx[T any](ctx context.Context, vals ...uint) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) (*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[Room](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[Room](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `room` "+whereSQL, params...)
}

// FetchByPK returns a row from `room` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val uint) (*Room, error) {
	return coredb.FetchByPKFromMasterCtx[Room](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from room table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val uint) (*T, error) {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `room` table with given primary key values
func FetchByPKsFromMasterCtx(ctx context.Context, vals ...uint) ([]*Room, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[Room](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `room` table with given primary key values
func FetchFieldsByPKsFromMasterCtx[T any](ctx context.Context, vals ...uint) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[Room](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[Room](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `room` "+whereSQL, params...)
}

// Insert Room struct to `room` table
func (c *Room) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithPK, c.getIdForDB(), c.getGroupForDB(), c.getLangForDB(), c.getPriorityForDB(), c.getDeletedForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getGroupForDB(), c.getLangForDB(), c.getPriorityForDB(), c.getDeletedForDB())
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

// Update Room struct in `room` table
func (obj *Room) UpdateCtx(ctx context.Context) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Group.IsUpdated() {
		updatedFields = append(updatedFields, "`group` = ?")
		params = append(params, obj.getGroupForDB())
	}
	if obj.Lang.IsUpdated() {
		updatedFields = append(updatedFields, "`lang` = ?")
		params = append(params, obj.getLangForDB())
	}
	if obj.Priority.IsUpdated() {
		updatedFields = append(updatedFields, "`priority` = ?")
		params = append(params, obj.getPriorityForDB())
	}
	if obj.Deleted.IsUpdated() {
		updatedFields = append(updatedFields, "`deleted` = ?")
		params = append(params, obj.getDeletedForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `room` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.ExecCtx(ctx, DBName, sql, params...)
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

// Update Room struct with given fields in `room` table
func UpdateCtx(ctx context.Context, obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *Group:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`group` = ?")
				params = append(params, c.getGroupForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Lang:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`lang` = ?")
				params = append(params, c.getLangForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Priority:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`priority` = ?")
				params = append(params, c.getPriorityForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Deleted:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`deleted` = ?")
				params = append(params, c.getDeletedForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `room` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.ExecCtx(ctx, DBName, sql, params...)
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

// DeleteByPK delete a row from room table with given primary key value
func DeleteByPKCtx(ctx context.Context, val uint) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val)
	return err
}
