package coredb

// TableType defines the generated type of a table
type TableType interface {
	GetTableName() string
}
