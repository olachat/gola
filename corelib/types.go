package corelib

type RowStruct interface {
	GetColumnNames() string
	GetPointers() []interface{}
}

type Ops int

const (
	OpInit Ops = iota
	OpEqual
	OpIn
	OpGreater
	OpSmaller
	OpRange
)
