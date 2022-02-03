package corelib

type PointerType[T any] interface {
	*T
}

type RowStruct interface {
	GetColumnNames() string
	GetPointers() []interface{}
}

type Ops int

const (
	OpInit    Ops = 0
	OpEqual   Ops = 1
	OpIn      Ops = 2
	OpGreater Ops = 3
	OpSmaller Ops = 4
	OpRange   Ops = 5
)
