// Code generated by gola 0.1.1; DO NOT EDIT.

package blogs

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/coredb/txengine"
)

// InsertTx inserts Blog struct to `blogs` table with transaction
func (c *Blog) InsertTx(ctx context.Context, tx *sql.Tx) error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = txengine.WithTx(tx).Exec(ctx, insertWithPK, c.getIdForDB(), c.getUserIdForDB(), c.getSlugForDB(), c.getTitleForDB(), c.getCategoryIdForDB(), c.getIsPinnedForDB(), c.getIsVipForDB(), c.getCountryForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB(), c.getCountForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = txengine.WithTx(tx).Exec(ctx, insertWithoutPK, c.getUserIdForDB(), c.getSlugForDB(), c.getTitleForDB(), c.getCategoryIdForDB(), c.getIsPinnedForDB(), c.getIsVipForDB(), c.getCountryForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB(), c.getCountForDB())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = int(id)
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

// UpdateTx updates Blog struct in `blogs` table with transaction
func (obj *Blog) UpdateTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.UserId.IsUpdated() {
		updatedFields = append(updatedFields, "`user_id` = ?")
		params = append(params, obj.getUserIdForDB())
	}
	if obj.Slug.IsUpdated() {
		updatedFields = append(updatedFields, "`slug` = ?")
		params = append(params, obj.getSlugForDB())
	}
	if obj.Title.IsUpdated() {
		updatedFields = append(updatedFields, "`title` = ?")
		params = append(params, obj.getTitleForDB())
	}
	if obj.CategoryId.IsUpdated() {
		updatedFields = append(updatedFields, "`category_id` = ?")
		params = append(params, obj.getCategoryIdForDB())
	}
	if obj.IsPinned.IsUpdated() {
		updatedFields = append(updatedFields, "`is_pinned` = ?")
		params = append(params, obj.getIsPinnedForDB())
	}
	if obj.IsVip.IsUpdated() {
		updatedFields = append(updatedFields, "`is_vip` = ?")
		params = append(params, obj.getIsVipForDB())
	}
	if obj.Country.IsUpdated() {
		updatedFields = append(updatedFields, "`country` = ?")
		params = append(params, obj.getCountryForDB())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.getCreatedAtForDB())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.getUpdatedAtForDB())
	}
	if obj.Count_.IsUpdated() {
		updatedFields = append(updatedFields, "`count` = ?")
		params = append(params, obj.getCountForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := txengine.WithTx(tx).Exec(ctx, sql, params...)
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

// UpdateTx updates Blog struct with given fields in `blogs` table with transaction
func UpdateTx(ctx context.Context, tx *sql.Tx, obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *UserId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`user_id` = ?")
				params = append(params, c.getUserIdForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Slug:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`slug` = ?")
				params = append(params, c.getSlugForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Title:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`title` = ?")
				params = append(params, c.getTitleForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CategoryId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`category_id` = ?")
				params = append(params, c.getCategoryIdForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsPinned:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_pinned` = ?")
				params = append(params, c.getIsPinnedForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsVip:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_vip` = ?")
				params = append(params, c.getIsVipForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Country:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`country` = ?")
				params = append(params, c.getCountryForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`created_at` = ?")
				params = append(params, c.getCreatedAtForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`updated_at` = ?")
				params = append(params, c.getUpdatedAtForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Count_:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`count` = ?")
				params = append(params, c.getCountForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := txengine.WithTx(tx).Exec(ctx, sql, params...)
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
