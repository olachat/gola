package coredb

// ColumnType defines the generated type of a table column
type ColumnType interface {
	GetColumnName() string
	GetValPointer() any
	IsPrimaryKey() bool
	GetTableType() TableType
}
