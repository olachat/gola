package corelib

type WhereQuery interface {
	GetWhere() string
}

type ReadQuery[T any] interface {
	All() []*T
	Limit(limit, offset int) []*T
}
