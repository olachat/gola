package blogs

import "github.com/olachat/gola/corelib"

func (q *idxQuery[T]) AndCategoryIdEQ(categoryId int) corelib.ReadQuery[T] {
	return q
}
func (q *idxQuery[T]) AndCategoryIdIN(args ...int) corelib.ReadQuery[T] {
	return q
}

type iQuery2[T any] interface {
	AndCategoryIdEQ(categoryId int) corelib.ReadQuery[T]
	AndCategoryIdIN(args ...int) corelib.ReadQuery[T]
	order[T]
	corelib.ReadQuery[T]
}
