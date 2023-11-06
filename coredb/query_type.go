package coredb

import "context"

// WhereQuery defines interface which support forming query query
type WhereQuery interface {
	GetWhere() (whereSQL string, params []any)
}

type whereQuery struct {
	whereSQL string
	params   []any
}

func (w *whereQuery) GetWhere() (whereSQL string, params []any) {
	return w.whereSQL, w.params
}

// NewWhere returns WhereQuery with given whereSQL and params
func NewWhere(whereSQL string, params ...any) WhereQuery {
	w := &whereQuery{
		whereSQL: whereSQL,
		params:   params,
	}

	return w
}

// ReadQuery defines interface which support reading multiple objects
type ReadQuery[T any] interface {
	// Deprecated: use the function with context
	All() []*T
	// Deprecated: use the function with context
	Limit(offset, limit int) []*T
	// Deprecated: use the function with context
	AllFromMaster() []*T
	// Deprecated: use the function with context
	LimitFromMaster(offset, limit int) []*T

	AllCtx(context.Context) ([]*T, error)
	LimitCtx(ctx context.Context, offset, limit int) ([]*T, error)
	AllFromMasterCtx(context.Context) ([]*T, error)
	LimitFromMasterCtx(ctx context.Context, offset, limit int) ([]*T, error)
}

// ReadOneQuery defines interface which support reading one object
type ReadOneQuery[T any] interface {
	One() *T
}
