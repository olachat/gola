package corelib

type PointerType[T any] interface {
	*T
}

type RowStruct interface {
	GetColumnNames() string
	GetPointers() []interface{}
}
