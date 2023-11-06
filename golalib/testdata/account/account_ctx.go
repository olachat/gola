// Code generated by gola 0.1.1; DO NOT EDIT.

package account

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `account` table with given primary key value
func FetchByPKCtx(ctx context.Context, val PK) (*Account, error) {
	return coredb.FetchByPKCtx[Account](ctx, DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FetchFieldsByPK returns a row with selected fields from account table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val PK) (*T, error) {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FindOne returns a row from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) (*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[Account](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[Account](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `account` "+whereSQL, params...)
}

// FetchByPK returns a row from `account` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val PK) (*Account, error) {
	return coredb.FetchByPKFromMasterCtx[Account](ctx, DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FetchFieldsByPK returns a row with selected fields from account table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val PK) (*T, error) {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FindOne returns a row from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[Account](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[Account](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `account` "+whereSQL, params...)
}

// Insert Account struct to `account` table
func (c *Account) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getUserIdForDB(), c.getTypeForDB(), c.getCountryCodeForDB(), c.getMoneyForDB())
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

// Update Account struct in `account` table
func (obj *Account) UpdateCtx(ctx context.Context) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Type.IsUpdated() {
		updatedFields = append(updatedFields, "`type` = ?")
		params = append(params, obj.getTypeForDB())
	}
	if obj.Money.IsUpdated() {
		updatedFields = append(updatedFields, "`money` = ?")
		params = append(params, obj.getMoneyForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `account` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `country_code` = ?"
	params = append(params, obj.GetUserId(), obj.GetCountryCode())

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

// Update Account struct with given fields in `account` table
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
		case *Type:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`type` = ?")
				params = append(params, c.getTypeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Money:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`money` = ?")
				params = append(params, c.getMoneyForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `account` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `country_code` = ?"
	params = append(params, obj.GetUserId(), obj.GetCountryCode())

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

// DeleteByPK delete a row from account table with given primary key value
func DeleteByPKCtx(ctx context.Context, val PK) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val.UserId, val.CountryCode)
	return err
}
