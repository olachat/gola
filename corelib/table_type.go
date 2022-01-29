package corelib

import "context"

type TableType interface {
	GetTableName(ctx context.Context) string
}
