package blogs

import "github.com/olachat/gola/corelib"

func (q *idxQuery[T]) AndCategoryIdEQ(categoryId int) corelib.ReadQuery[T] {
	return q
}
func (q *idxQuery[T]) AndCategoryIdIN(args ...int) corelib.ReadQuery[T] {
	return q
}

func (q *idxQuery[T]) All() []*T {
	return nil
}

func (q *idxQuery[T]) Limit(limit, offset int) []*T {
	return nil
}

func (q *idxQuery[T]) OrderBy(args ...orderBy) corelib.ReadQuery[T] {
	return q
}

type order[T any] interface {
	OrderBy(args ...orderBy) corelib.ReadQuery[T]
}

type iQuery1[T any] interface {
	WhereCountryEQ(country string) iQuery2[T]
	WhereCountryIN(args ...string) iQuery2[T]
	order[T]
	corelib.ReadQuery[T]
}

type iQuery2[T any] interface {
	AndCategoryIdEQ(categoryId int) corelib.ReadQuery[T]
	AndCategoryIdIN(args ...int) corelib.ReadQuery[T]
	order[T]
	corelib.ReadQuery[T]
}

// Find methods
func Select[T any]() iQuery1[T] {
	return new(idxQuery[T])
}
