package coredb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type IsNonRetryableErrorFunc func(err error) bool

// RetryConfig encapsulates retry parameters.
type RetryConfig struct {
	MaxRetries              int
	InitialBackoff          time.Duration
	IsNonRetryableErrorFunc IsNonRetryableErrorFunc
}

// DefaultRetryConfig provides a reasonable default configuration
var DefaultRetryConfig = RetryConfig{
	MaxRetries:              5,
	InitialBackoff:          200 * time.Millisecond,
	IsNonRetryableErrorFunc: IsNonRetryableError,
}

// IsNonRetryableError checks if an error is non-retryable.
func IsNonRetryableError(err error) bool {
	if err == nil {
		return false
	}
	// Example (Replace with your database's non-retryable errors)

	// SQL specific errors that are not retryable
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}

	// Example: Invalid SQL syntax
	if strings.Contains(err.Error(), "syntax error") {
		return true
	}

	if strings.Contains(err.Error(), "1146") { // Table doesn't exists
		return true
	}
	if strings.Contains(err.Error(), "1064") { // No database selected
		return true
	}
	if strings.Contains(err.Error(), "1149") { // Invalid SQL statement
		return true
	}
	// Example: Authentication issues
	if strings.Contains(err.Error(), "Access denied") {
		return true
	}

	return false // Default is retryable
}

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

// ExecWithRetry executes a query with retry logic on failure.
func ExecWithRetry(ctx context.Context, dbname string, query string, retryConfig RetryConfig, params ...any) (sql.Result, error) {
	// Set defaults for invalid config
	if retryConfig.MaxRetries <= 0 {
		retryConfig.MaxRetries = DefaultRetryConfig.MaxRetries
	}

	if retryConfig.InitialBackoff <= 0 {
		retryConfig.InitialBackoff = DefaultRetryConfig.InitialBackoff
	}

	// Use the default if NonRetryableErrorFunc is nil
	nonRetryableErrorFunc := retryConfig.IsNonRetryableErrorFunc
	if nonRetryableErrorFunc == nil {
		nonRetryableErrorFunc = IsNonRetryableError
	}

	var result sql.Result
	var err error
	retryCount := 0
	currentBackoff := retryConfig.InitialBackoff

	for {
		select {
		case <-ctx.Done():
			return result, fmt.Errorf("context cancelled during retry: %w", ctx.Err())
		default:
			result, err = ExecCtx(ctx, dbname, query, params...)
			if err == nil {
				return result, nil // Success!
			}

			if nonRetryableErrorFunc(err) {
				return result, err // Fail immediately for non-retryable errors
			}

			retryCount++
			if retryCount > retryConfig.MaxRetries {
				log.Printf("Max retries (%d) exceeded for: %s, last error: %v", retryConfig.MaxRetries, query, err)
				return result, fmt.Errorf("max retries exceeded, last error: %w", err)
			}

			delay := currentBackoff
			log.Printf("Retrying attempt %d with delay %v. Last error: %v", retryCount, delay, err)
			time.Sleep(delay)
			currentBackoff *= 2
		}
	}
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
	err = mydb.QueryRowContext(ctx, query, params...).Scan(&result)
	return
}

// QueryIntFromMasterCtx single int result by query, handy for count(*) querys
func QueryIntFromMasterCtx(ctx context.Context, dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	err = mydb.QueryRowContext(ctx, query, params...).Scan(&result)
	return
}

// QueryCtx rows from given table type with where query & params
func QueryCtx[T any](ctx context.Context, dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeRead)
	var rows *sql.Rows
	rows, err = mydb.QueryContext(ctx, query, params...)
	if err != nil {
		return
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Printf("Gola: QueryCtx: failed to close rows: %v", errClose)
		}
	}()

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

	if rows.Err() != nil {
		err = rows.Err()
	}

	return
}

// QueryFromMasterCtx rows from master DB from given table type with where query & params
func QueryFromMasterCtx[T any](ctx context.Context, dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	var rows *sql.Rows
	rows, err = mydb.QueryContext(ctx, query, params...)
	if err != nil {
		return
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Printf("Gola: QueryFromMasterCtx: failed to close rows: %v", errClose)
		}
	}()

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

	if rows.Err() != nil {
		err = rows.Err()
	}

	return
}
