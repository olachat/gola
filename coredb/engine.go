package coredb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var _dbp DBProvider

type DBMode int

const (
	// DBModeRead allow read from slave db
	DBModeRead DBMode = iota
	// DBModeRead set write to master db
	DBModeWrite
	// DBModeReadFromWrite is aka force read from master db
	DBModeReadFromWrite
)

type DBProvider func(dbname string, mode DBMode) *sql.DB

// Setup default db provider for all db ops
func Setup(dbp DBProvider) {
	_dbp = dbp
}

var (
	typeColumnNames     = make(map[reflect.Type]string)
	typeColumnNamesLock sync.RWMutex
)

func getDB(dbname string, mode DBMode) *sql.DB {
	if _dbp == nil {
		panic("coredb DBProvider hasn't setup")
	}
	db := _dbp(dbname, mode)
	if db != nil {
		return db
	}
	panic(fmt.Sprintf("Can't get db for %s %v", dbname, mode))
}

// FetchByPK returns a row of T type with given primary key value
//
// Deprecated: use the function with context
func FetchByPK[T any](dbname string, tableName string, pkName []string, val ...any) *T {
	sql := "WHERE `" + pkName[0] + "` = ?"
	for _, name := range pkName[1:] {
		sql += " AND `" + name + "` = ?"
	}
	w := NewWhere(sql, val...)
	return FindOne[T](dbname, tableName, w)
}

// FetchByPKs returns rows of T type with given primary key values
//
// Deprecated: use the function with context
func FetchByPKs[T any](dbname string, tableName string, pkName string, vals []any) []*T {
	if len(vals) == 0 {
		return make([]*T, 0)
	}

	query := fmt.Sprintf("WHERE `%s` IN (%s)", pkName, GetParamPlaceHolder(len(vals)))
	w := NewWhere(query, vals...)

	result, err := Find[T](dbname, tableName, w)
	if err != nil {
		panic("Find failled: " + err.Error())
	}
	return result
}

// FetchByPKFromMaster returns a row of T type with given primary key value
//
// Deprecated: use the function with context
func FetchByPKFromMaster[T any](dbname string, tableName string, pkName []string, val ...any) *T {
	sql := "WHERE `" + pkName[0] + "` = ?"
	for _, name := range pkName[1:] {
		sql += " AND `" + name + "` = ?"
	}
	w := NewWhere(sql, val...)
	return FindOneFromMaster[T](dbname, tableName, w)
}

// FetchByPKsFromMaster returns rows of T type with given primary key values
//
// Deprecated: use the function with context
func FetchByPKsFromMaster[T any](dbname string, tableName string, pkName string, vals []any) []*T {
	if len(vals) == 0 {
		return make([]*T, 0)
	}

	query := fmt.Sprintf("WHERE `%s` IN (%s)", pkName, GetParamPlaceHolder(len(vals)))
	w := NewWhere(query, vals...)

	result, err := FindFromMaster[T](dbname, tableName, w)
	if err != nil {
		panic("Find failled: " + err.Error())
	}
	return result
}

// Exec given query with given db info & params
//
// Deprecated: use the function with context
func Exec(dbname string, query string, params ...any) (sql.Result, error) {
	mydb := getDB(dbname, DBModeWrite)
	return mydb.Exec(query, params...)
}

// FindOne returns a row from given table type with where query
//
// Deprecated: use the function with context
func FindOne[T any](dbname string, tableName string, where WhereQuery) *T {
	u := new(T)
	columnsNames := GetColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	mydb := getDB(dbname, DBModeRead)
	err2 := mydb.QueryRow(query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 != sql.ErrNoRows {
			panic("QueryRow failed: " + err2.Error())
		}

		return nil
	}

	return u
}

// Find returns rows from given table type with where query
//
// Deprecated: use the function with context
func Find[T any](dbname string, tableName string, where WhereQuery) ([]*T, error) {
	columnsNames := GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return Query[T](dbname, query, params...)
}

// FindOneFromMaster using master DB returns a row from given table type with where query
//
// Deprecated: use the function with context
func FindOneFromMaster[T any](dbname string, tableName string, where WhereQuery) *T {
	u := new(T)
	columnsNames := GetColumnsNames[T]()
	data := StrutForScan(u)
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)
	mydb := getDB(dbname, DBModeReadFromWrite)
	err2 := mydb.QueryRow(query, params...).Scan(data...)

	if err2 != nil {
		// It's on purpose the hide the error
		// But should re-consider later
		if err2 != sql.ErrNoRows {
			panic("QueryRow failed: " + err2.Error())
		}

		return nil
	}

	return u
}

// FindFromMaster using master DB returns rows from given table type with where query
//
// Deprecated: use the function with context
func FindFromMaster[T any](dbname string, tableName string, where WhereQuery) ([]*T, error) {
	columnsNames := GetColumnsNames[T]()
	whereSQL, params := where.GetWhere()
	query := fmt.Sprintf("SELECT %s FROM `%s` %s", columnsNames,
		tableName, whereSQL)

	return QueryFromMaster[T](dbname, query, params...)
}

// QueryInt single int result by query, handy for count(*) querys
//
// Deprecated: use the function with context
func QueryInt(dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeRead)
	err = mydb.QueryRow(query, params...).Scan(&result)
	return
}

// QueryIntFromMaster single int result by query, handy for count(*) querys
//
// Deprecated: use the function with context
func QueryIntFromMaster(dbname string, query string, params ...any) (result int, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	err = mydb.QueryRow(query, params...).Scan(&result)
	return
}

// Query rows from given table type with where query & params
//
// Deprecated: use the function with context
func Query[T any](dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeRead)
	var rows *sql.Rows
	rows, err = mydb.Query(query, params...)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
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

	return
}

// QueryFromMaster rows from master DB from given table type with where query & params
//
// Deprecated: use the function with context
func QueryFromMaster[T any](dbname string, query string, params ...any) (result []*T, err error) {
	mydb := getDB(dbname, DBModeReadFromWrite)
	var rows *sql.Rows
	rows, err = mydb.Query(query, params...)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
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

	return
}

// GetColumnsNames returns column names joined by `,` of given type
func GetColumnsNames[T any]() (joinedColumnNames string) {
	var o *T
	t := reflect.TypeOf(o)
	typeColumnNamesLock.RLock()
	var ok bool
	joinedColumnNames, ok = typeColumnNames[t]
	typeColumnNamesLock.RUnlock()
	if ok {
		return
	}

	o = new(T)
	var columnNames []string
	val := reflect.ValueOf(o).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnNamer); ok {
			columnNames = append(columnNames, "`"+f.GetColumnName()+"`")
		}
	}

	joinedColumnNames = strings.Join(columnNames, ",")

	typeColumnNamesLock.Lock()
	typeColumnNames[t] = joinedColumnNames
	typeColumnNamesLock.Unlock()

	return
}

// GetColumnsNamesReflect returns column names joined by `,` of given type
func GetColumnsNamesReflect(o any) (joinedColumnNames string) {
	t := reflect.TypeOf(o)
	elemType := t.Elem()
	switch t.Kind() {
	case reflect.Ptr:
		// 是指针，尝试获取其指向的元素类型
		if elemType.Kind() == reflect.Slice {
			elemType = elemType.Elem().Elem() // 如果指针指向切片，获取切片元素的类型
		}
	case reflect.Slice:
		// 是切片，获取切片元素的类型
		if elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem() // 切片元素是指针，获取其指向的类型
		}
	default:
		// 既不是指针也不是切片，返回错误
		panic("coredb: o is neither a pointer nor a slice")
	}

	typeColumnNamesLock.RLock()
	var ok bool
	joinedColumnNames, ok = typeColumnNames[elemType]
	typeColumnNamesLock.RUnlock()
	if ok {
		return
	}

	var columnNames []string
	val := reflect.New(elemType).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnNamer); ok {
			columnNames = append(columnNames, "`"+f.GetColumnName()+"`")
		}
	}

	joinedColumnNames = strings.Join(columnNames, ",")

	typeColumnNamesLock.Lock()
	typeColumnNames[t] = joinedColumnNames
	typeColumnNamesLock.Unlock()

	return
}
