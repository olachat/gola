package corelib

// ColumnType defines the generated type of a table column
type ColumnType interface {
	GetColumnName() string
	GetValPointer() interface{}
	IsPrimaryKey() bool
	GetTableType() TableType
}
