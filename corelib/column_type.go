package corelib

type ColumnType interface {
	GetColumnName() string
	GetValPointer() interface{}
	IsPrimaryKey() bool
}
