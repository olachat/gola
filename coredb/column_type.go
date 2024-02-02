package coredb

// ColumnType defines the generated type of a table column
type ColumnValPointer interface {
	GetValPointer() any
}
type ColumnNamer interface {
	GetColumnName() string
}
