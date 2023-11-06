// Code generated by gola 0.1.1; DO NOT EDIT.

package gifts

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `gifts` table with given primary key value
func FetchByPKCtx(ctx context.Context, val uint) (*Gift, error) {
	return coredb.FetchByPKCtx[Gift](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from gifts table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val uint) (*T, error) {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `gifts` table with given primary key values
func FetchByPKsCtx(ctx context.Context, vals ...uint) ([]*Gift, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[Gift](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `gifts` table with given primary key values
func FetchFieldsByPKsCtx[T any](ctx context.Context, vals ...uint) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) (*Gift, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[Gift](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*Gift, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[Gift](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `gifts` "+whereSQL, params...)
}

// FetchByPK returns a row from `gifts` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val uint) (*Gift, error) {
	return coredb.FetchByPKFromMasterCtx[Gift](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from gifts table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val uint) (*T, error) {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `gifts` table with given primary key values
func FetchByPKsFromMasterCtx(ctx context.Context, vals ...uint) ([]*Gift, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[Gift](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `gifts` table with given primary key values
func FetchFieldsByPKsFromMasterCtx[T any](ctx context.Context, vals ...uint) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (*Gift, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[Gift](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*Gift, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[Gift](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `gifts` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `gifts` "+whereSQL, params...)
}

// Insert Gift struct to `gifts` table
func (c *Gift) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithPK, c.getIdForDB(), c.getNameForDB(), c.getIsFreeForDB(), c.getGiftCountForDB(), c.getGiftTypeForDB(), c.getCreateTimeForDB(), c.getDiscountForDB(), c.getPriceForDB(), c.getRemarkForDB(), c.getManifestForDB(), c.getDescriptionForDB(), c.getUpdateTimeForDB(), c.getBranchesForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getNameForDB(), c.getIsFreeForDB(), c.getGiftCountForDB(), c.getGiftTypeForDB(), c.getCreateTimeForDB(), c.getDiscountForDB(), c.getPriceForDB(), c.getRemarkForDB(), c.getManifestForDB(), c.getDescriptionForDB(), c.getUpdateTimeForDB(), c.getBranchesForDB())
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

// Update Gift struct in `gifts` table
func (obj *Gift) UpdateCtx(ctx context.Context) (bool, error) {
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
	if obj.Branches.IsUpdated() {
		updatedFields = append(updatedFields, "`branches` = ?")
		params = append(params, obj.getBranchesForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `gifts` SET "
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

// Update Gift struct with given fields in `gifts` table
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

	sql := "UPDATE `gifts` SET "
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

// DeleteByPK delete a row from gifts table with given primary key value
func DeleteByPKCtx(ctx context.Context, val uint) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val)
	return err
}
