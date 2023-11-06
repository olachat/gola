// Code generated by gola 0.1.1; DO NOT EDIT.

package users

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `users` table with given primary key value
func FetchByPKCtx(ctx context.Context, val int) (*User, error) {
	return coredb.FetchByPKCtx[User](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from users table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val int) (*T, error) {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `users` table with given primary key values
func FetchByPKsCtx(ctx context.Context, vals ...int) ([]*User, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[User](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `users` table with given primary key values
func FetchFieldsByPKsCtx[T any](ctx context.Context, vals ...int) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) (*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[User](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[User](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `users` "+whereSQL, params...)
}

// FetchByPK returns a row from `users` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val int) (*User, error) {
	return coredb.FetchByPKFromMasterCtx[User](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from users table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val int) (*T, error) {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `users` table with given primary key values
func FetchByPKsFromMasterCtx(ctx context.Context, vals ...int) ([]*User, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[User](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `users` table with given primary key values
func FetchFieldsByPKsFromMasterCtx[T any](ctx context.Context, vals ...int) ([]*T, error) {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[User](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*User, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[User](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `users` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `users` "+whereSQL, params...)
}

// Insert User struct to `users` table
func (c *User) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithPK, c.getIdForDB(), c.getNameForDB(), c.getEmailForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB(), c.getFloatTypeForDB(), c.getDoubleTypeForDB(), c.getHobbyForDB(), c.getHobbyNoDefaultForDB(), c.getSportsForDB(), c.getSports2ForDB(), c.getSportsNoDefaultForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getNameForDB(), c.getEmailForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB(), c.getFloatTypeForDB(), c.getDoubleTypeForDB(), c.getHobbyForDB(), c.getHobbyNoDefaultForDB(), c.getSportsForDB(), c.getSports2ForDB(), c.getSportsNoDefaultForDB())
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

// Update User struct in `users` table
func (obj *User) UpdateCtx(ctx context.Context) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Name.IsUpdated() {
		updatedFields = append(updatedFields, "`name` = ?")
		params = append(params, obj.getNameForDB())
	}
	if obj.Email.IsUpdated() {
		updatedFields = append(updatedFields, "`email` = ?")
		params = append(params, obj.getEmailForDB())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.getCreatedAtForDB())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.getUpdatedAtForDB())
	}
	if obj.FloatType.IsUpdated() {
		updatedFields = append(updatedFields, "`float_type` = ?")
		params = append(params, obj.getFloatTypeForDB())
	}
	if obj.DoubleType.IsUpdated() {
		updatedFields = append(updatedFields, "`double_type` = ?")
		params = append(params, obj.getDoubleTypeForDB())
	}
	if obj.Hobby.IsUpdated() {
		updatedFields = append(updatedFields, "`hobby` = ?")
		params = append(params, obj.getHobbyForDB())
	}
	if obj.HobbyNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "`hobby_no_default` = ?")
		params = append(params, obj.getHobbyNoDefaultForDB())
	}
	if obj.Sports.IsUpdated() {
		updatedFields = append(updatedFields, "`sports` = ?")
		params = append(params, obj.getSportsForDB())
	}
	if obj.Sports2.IsUpdated() {
		updatedFields = append(updatedFields, "`sports2` = ?")
		params = append(params, obj.getSports2ForDB())
	}
	if obj.SportsNoDefault.IsUpdated() {
		updatedFields = append(updatedFields, "`sports_no_default` = ?")
		params = append(params, obj.getSportsNoDefaultForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `users` SET "
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

// Update User struct with given fields in `users` table
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
		case *Email:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`email` = ?")
				params = append(params, c.getEmailForDB())
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
		case *FloatType:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`float_type` = ?")
				params = append(params, c.getFloatTypeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *DoubleType:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`double_type` = ?")
				params = append(params, c.getDoubleTypeForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Hobby:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`hobby` = ?")
				params = append(params, c.getHobbyForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *HobbyNoDefault:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`hobby_no_default` = ?")
				params = append(params, c.getHobbyNoDefaultForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Sports:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports` = ?")
				params = append(params, c.getSportsForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Sports2:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports2` = ?")
				params = append(params, c.getSports2ForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *SportsNoDefault:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`sports_no_default` = ?")
				params = append(params, c.getSportsNoDefaultForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `users` SET "
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

// DeleteByPK delete a row from users table with given primary key value
func DeleteByPKCtx(ctx context.Context, val int) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val)
	return err
}
