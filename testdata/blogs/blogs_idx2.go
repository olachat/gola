package blogs

import "github.com/olachat/gola/corelib"

type orderBy int

const (
	IdAsc orderBy = iota
	IdDesc
	CategoryIdAsc
	CategoryIdDesc
	TitleAsc
	TitleDesc
)

type Idx1[T any] struct {
}

func (i *Idx1[T]) WhereCountryEQ(country string) iQuery2[T] {
	return i
}

func (i *Idx1[T]) WhereCountryIN(country ...string) iQuery2[T] {
	return i
}

func (i *Idx1[T]) AndCategoryIdEQ(categoryId int) corelib.ReadQuery[T] {
	return i
}
func (i *Idx1[T]) AndCategoryIdIN(args ...int) corelib.ReadQuery[T] {
	return i
}

func (i *Idx1[T]) All() []*T {
	return nil
}

func (i *Idx1[T]) Limit(limit, offset int) []*T {
	return nil
}

func (i *Idx1[T]) OrderBy(args ...orderBy) corelib.ReadQuery[T] {
	return i
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
	return new(Idx1[T])
}
