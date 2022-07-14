package corelib

// RowStruct defines interface of an ORM row struct
type RowStruct interface {
	GetColumnNames() string
	GetPointers() []any
}

// Ops defines operation in a where query
type Ops int

const (
	// OpInit is default
	OpInit Ops = iota
	// OpEqual is =
	OpEqual
	// OpIn is in
	OpIn
	// OpGreater is >
	OpGreater
	// OpSmaller is <
	OpSmaller
	// OpRange is < ? <
	OpRange
)
