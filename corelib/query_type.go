package corelib

// WhereQuery defines interface which support forming query query
type WhereQuery interface {
	GetWhere() (whereSQL string, params []any)
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
