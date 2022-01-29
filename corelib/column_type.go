package corelib

import "context"

type ColumnType interface {
	GetColumnName(ctx context.Context) string
	GetValPointer(ctx context.Context) interface{}
	IsPrimaryKey(ctx context.Context) bool
	GetTableType(ctx context.Context) TableType
}
