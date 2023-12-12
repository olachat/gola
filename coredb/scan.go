package coredb

import (
	"database/sql"
	"reflect"
)

// An InvalidScanError describes an invalid argument passed to Scan.
type InvalidScanError struct {
	Type reflect.Type
}

func (e *InvalidScanError) Error() string {
	if e.Type == nil {
		return "coredb: target is nil"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "coredb: target must be a non-nil pointer, got " + e.Type.String()
	}
	return "coredb: nil " + e.Type.String() + ")"
}

// RowsToStructSlice converts the rows of a SQL query result into a slice of structs.
//
// It takes a pointer to a sql.Rows object as input.
// The function also uses a generic type T, which represents the type of the struct.
//
// The function returns a slice of pointers to T structs and an error.
func RowsToStructSlice[T any](rows *sql.Rows) (result []*T, err error) {
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

// RowToStruct converts a database row into a struct.
//
// It takes a pointer to a sql.Row and returns a pointer to the converted struct and an error.
func RowToStruct[T any](row *sql.Row) (result *T, err error) {
	result = new(T)
	data := StrutForScan(result)
	err = row.Scan(data...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return
}

// StrutForScan returns value pointers of given obj
func StrutForScan(u any) (pointers []any) {
	val := reflect.ValueOf(u)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		err := &InvalidScanError{reflect.TypeOf(u)}
		panic(err)
	}

	val = val.Elem()

	pointers = make([]any, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if f, ok := valueField.Addr().Interface().(ColumnType); ok {
			pointers = append(pointers, f.GetValPointer())
		}
	}
	return
}

func RowsToStructSliceReflect(rows *sql.Rows, out any) (err error) {
	sliceValue := reflect.ValueOf(out)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.IsNil() {
		return &InvalidScanError{Type: sliceValue.Type()}
	}
	sliceValue = sliceValue.Elem()
	if sliceValue.Kind() != reflect.Slice {
		return &InvalidScanError{Type: reflect.TypeOf(out)}
	}
	elementType := sliceValue.Type().Elem()
	if elementType.Kind() != reflect.Ptr {
		return &InvalidScanError{Type: reflect.TypeOf(out)}
	}
	elementType = elementType.Elem()

	for rows.Next() {
		v := reflect.New(elementType)
		data := StrutForScan(v.Interface())
		err = rows.Scan(data...)
		if err != nil {
			return
		}
		sliceValue.Set(reflect.Append(sliceValue, v))
	}
	return
}

func RowToStructReflect(row *sql.Row, v any) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidScanError{reflect.TypeOf(v)}
	}

	data := StrutForScan(v)
	err = row.Scan(data...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return
}
