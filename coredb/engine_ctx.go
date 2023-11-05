package coredb

import (
	"context"
	"database/sql"
	"fmt"
)

// FetchByPKCtx returns a row of T type with given primary key value
func FetchByPKCtx[T any](ctx context.Context, dbname string, tableName string, pkName []string, val ...any) (*T, error) {
	sql := "WHERE `" + pkName[0] + "` = ?"
	for _, name := range pkName[1:] {
		sql += " AND `" + name + "` = ?"
	}
	w := NewWhere(sql, val...)
	return FindOneCtx[T](ctx, dbname, tableName, w)
}

// FetchByPKsCtx returns rows of T type with given primary key values
func FetchByPKsCtx[T any](ctx context.Context, dbname string, tableName string, pkName string, vals []any) ([]*T, error) {
	if len(vals) == 0 {
		return make([]*T, 0), nil
	}

	query := fmt.Sprintf("WHERE `%s` IN (%s)", pkName, GetParamPlaceHolder(len(vals)))
	w := NewWhere(query, vals...)

	return FindCtx[T](ctx, dbname, tableName, w)
}

// FetchByPKFromMasterCtx returns a row of T type with given primary key value
func FetchByPKFromMasterCtx[T any](ctx context.Context, dbname string, tableName string, pkName []string, val ...any) (*T, error) {
	sql := "WHERE `" + pkName[0] + "` = ?"
	for _, name := range pkName[1:] {
		sql += " AND `" + name + "` = ?"
	}
	w := NewWhere(sql, val...)
	return FindOneFromMasterCtx[T](ctx, dbname, tableName, w)
}

// FetchByPKsFromMasterCtx returns rows of T type with given primary key values
func FetchByPKsFromMasterCtx[T any](ctx context.Context, dbname string, tableName string, pkName string, vals []any) ([]*T, error) {
	if len(vals) == 0 {
		return make([]*T, 0), nil
	}

	query := fmt.Sprintf("WHERE `%s` IN (%s)", pkName, GetParamPlaceHolder(len(vals)))
	w := NewWhere(query, vals...)

	return FindFromMasterCtx[T](ctx, dbname, tableName, w)
}

// ExecCtx given query with given db info & params
func ExecCtx(ctx context.Context, dbname string, query string, params ...any) (sql.Result, error) {
	mydb := getDB(dbname, DBModeWrite)
	return mydb.ExecContext(ctx, query, params...)
}

// FindOneCtx returns a row from given table type with where query.
// If no rows found, *T will be nil. No error will be returned.
func FindOneCtx[T any](ctx context.Context, dbname string, tableName string, where WhereQuery) (*T, error) {
	u := new(T)
	columnsNames := GetColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	mydb := getDB(dbname, DBModeRead)
	err2 := mydb.QueryRowContext(ctx, query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err2
	}

	return u, nil
}

// FindCtx returns rows from given table type with where query
func FindCtx[T any](ctx context.Context, dbname string, tableName string, where WhereQuery) ([]*T, error) {
	columnsNames := GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return QueryCtx[T](ctx, dbname, query, params...)
}

// FindOneFromMasterCtx using master DB returns a row from given table type with where query
// If no rows found, *T will be nil. No error will be returned.
func FindOneFromMasterCtx[T any](ctx context.Context, dbname string, tableName string, where WhereQuery) (*T, error) {
	u := new(T)
	columnsNames := GetColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	mydb := getDB(dbname, DBModeReadFromWrite)
	err2 := mydb.QueryRowContext(ctx, query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err2
	}

	return u, nil
}

// FindFromMasterCtx using master DB returns rows from given table type with where query
func FindFromMasterCtx[T any](ctx context.Context, dbname string, tableName string, where WhereQuery) ([]*T, error) {
	columnsNames := GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return QueryFromMasterCtx[T](ctx, dbname, query, params...)
}

// QueryIntCtx single int result by query, handy for count(*) querys
func QueryIntCtx(ctx context.Context, dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeRead)
	mydb.QueryRowContext(ctx, query, params...).Scan(&result)
	return
}

// QueryIntFromMasterCtx single int result by query, handy for count(*) querys
func QueryIntFromMasterCtx(ctx context.Context, dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	mydb.QueryRowContext(ctx, query, params...).Scan(&result)
	return
}

// QueryCtx rows from given table type with where query & params
func QueryCtx[T any](ctx context.Context, dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeRead)
	rows, err := mydb.QueryContext(ctx, query, params...)
	if err != nil {
		return
	}

	var u *T
	for rows.Next() {
		u = new(T)
		data := StrutForScan(u)
		err = rows.Scan(data...)
		if err != nil {
			return
		}
		result = append(result, u)
	}

	return
}

// QueryCtx rows from master DB from given table type with where query & params
func QueryFromMasterCtx[T any](ctx context.Context, dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	rows, err := mydb.QueryContext(ctx, query, params...)
	if err != nil {
		return
	}

	var u *T
	for rows.Next() {
		u = new(T)
		data := StrutForScan(u)
		err = rows.Scan(data...)
		if err != nil {
			return
		}
		result = append(result, u)
	}

	return
}
