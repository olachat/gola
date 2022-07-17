package coredb

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

func NewWhere(whereSQL string, params ...any) WhereQuery {
	w := &whereQuery{
		whereSQL: whereSQL,
		params:   params,
	}

	return w
}

// ReadQuery defines interface which support reading multiple objects
type ReadQuery[T any] interface {
	All() []*T
	Limit(limit, offset int) []*T
}

// ReadOneQuery defines interface which support reading one object
type ReadOneQuery[T any] interface {
	One() *T
}
