package corelib

type WhereQuery interface {
	GetWhere() (whereSql string, params []interface{})
}

type ReadQuery[T any] interface {
	All() []*T
	Limit(limit, offset int) []*T
}

type ReadOneQuery[T any] interface {
	One() *T
}
