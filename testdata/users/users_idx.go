// Code generated by gola 0.0.1; DO NOT EDIT.

package users

import "github.com/olachat/gola/corelib"

type orderBy int

type idxQuery[T any] struct {
}

// order by enum & interface
const (
	IdAsc orderBy = iota
	IdDesc
	NameAsc
	NameDesc
	EmailAsc
	EmailDesc
	CreatedAtAsc
	CreatedAtDesc
	UpdatedAtAsc
	UpdatedAtDesc
	FloatTypeAsc
	FloatTypeDesc
	DoubleTypeAsc
	DoubleTypeDesc
	HobbyAsc
	HobbyDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) corelib.ReadQuery[T] {
	return q
}

func (q *idxQuery[T]) All() []*T {
	return nil
}

func (q *idxQuery[T]) Limit(limit, offset int) []*T {
	return nil
}

type order[T any] interface {
	OrderBy(args ...orderBy) corelib.ReadQuery[T]
}

type iQuery[T any] interface {
	WhereEmailEQ(val string) *idxQuery1[T]
	WhereEmailIN(vals ...string) *idxQuery1[T]
	WhereNameEQ(val string) *idxQuery2[T]
	WhereNameIN(vals ...string) *idxQuery2[T]
	order[T]
	corelib.ReadQuery[T]
}

// Find methods
func Select[T any]() iQuery[T] {
	return new(idxQuery[T])
}
func (q *idxQuery[T]) WhereEmailEQ(val string) *idxQuery1[T] {
	return q
}

func (q *idxQuery[T]) WhereEmailIN(vals ...string) *idxQuery1[T] {
	return q
}

func (q *idxQuery[T]) WhereNameEQ(val string) *idxQuery2[T] {
	return q
}

func (q *idxQuery[T]) WhereNameIN(vals ...string) *idxQuery2[T] {
	return q
}
