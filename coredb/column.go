package coredb

type ColumnVal[T any] struct {
	val T
}

func (c *ColumnVal[T]) GetValPointer() any {
	return &c.val
}

func (c *ColumnVal[T]) GetVal() T {
	return c.val
}
