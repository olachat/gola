// Code generated by gola 0.1.1; DO NOT EDIT.

package profile

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
	"github.com/olachat/gola/v2/coredb/txengine"
)

// InsertTx inserts Profile struct to `profile` table with transaction
func (c *Profile) InsertTx(ctx context.Context, tx *sql.Tx) error {
	var result sql.Result
	var err error
	result, err = txengine.WithTx(tx).Exec(ctx, insertWithoutPK, c.getUserIdForDB(), c.getLevelForDB(), c.getNickNameForDB())
	if err != nil {
		return err
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

// UpdateTx updates Profile struct in `profile` table with transaction
func (obj *Profile) UpdateTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Level.IsUpdated() {
		updatedFields = append(updatedFields, "`level` = ?")
		params = append(params, obj.getLevelForDB())
	}
	if obj.NickName.IsUpdated() {
		updatedFields = append(updatedFields, "`nick_name` = ?")
		params = append(params, obj.getNickNameForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `profile` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ?"
	params = append(params, obj.GetUserId())

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

// UpdateTx updates Profile struct with given fields in `profile` table with transaction
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
		case *Level:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`level` = ?")
				params = append(params, c.getLevelForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *NickName:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`nick_name` = ?")
				params = append(params, c.getNickNameForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `profile` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ?"
	params = append(params, obj.GetUserId())

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