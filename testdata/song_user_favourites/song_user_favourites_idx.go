// Code generated by gola 0.0.2; DO NOT EDIT.

package song_user_favourites

import (
	"fmt"
	"strings"

	"github.com/olachat/gola/coredb"
)

type orderBy int

type idxQuery[T any] struct {
	whereSql    string
	limitSql    string
	orders      []string
	whereParams []any
}

// order by enum & interface
const (
	UserIdAsc orderBy = iota
	UserIdDesc
	SongIdAsc
	SongIdDesc
	RemarkAsc
	RemarkDesc
	IsFavouriteAsc
	IsFavouriteDesc
	CreatedAtAsc
	CreatedAtDesc
	UpdatedAtAsc
	UpdatedAtDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case UserIdAsc:
			q.orders[i] = "user_id asc"
		case UserIdDesc:
			q.orders[i] = "user_id desc"
		case SongIdAsc:
			q.orders[i] = "song_id asc"
		case SongIdDesc:
			q.orders[i] = "song_id desc"
		case RemarkAsc:
			q.orders[i] = "remark asc"
		case RemarkDesc:
			q.orders[i] = "remark desc"
		case IsFavouriteAsc:
			q.orders[i] = "is_favourite asc"
		case IsFavouriteDesc:
			q.orders[i] = "is_favourite desc"
		case CreatedAtAsc:
			q.orders[i] = "created_at asc"
		case CreatedAtDesc:
			q.orders[i] = "created_at desc"
		case UpdatedAtAsc:
			q.orders[i] = "updated_at asc"
		case UpdatedAtDesc:
			q.orders[i] = "updated_at desc"
		}
	}
	return q
}

func (q *idxQuery[T]) All() []*T {
	result, _ := coredb.Find[T](q, _db)
	return result
}

func (q *idxQuery[T]) Limit(offset, count int) []*T {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, count)
	result, _ := coredb.Find[T](q, _db)
	return result
}

type order[T any] interface {
	OrderBy(args ...orderBy) coredb.ReadQuery[T]
}

type orderReadQuery[T any] interface {
	order[T]
	coredb.ReadQuery[T]
}

type iQuery[T any] interface {
	WhereUserIdEQ(val uint) iQuery1[T]
	WhereUserIdIN(vals ...uint) iQuery1[T]
	orderReadQuery[T]
}
type iQuery1[T any] interface {
	AndIsFavouriteEQ(val int8) orderReadQuery[T]
	AndIsFavouriteIN(vals ...int8) orderReadQuery[T]
	orderReadQuery[T]
}

type idxQuery1[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery1[T]) AndIsFavouriteEQ(val int8) orderReadQuery[T] {
	q.whereSql = " and user_id = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery1[T]) AndIsFavouriteIN(vals ...int8) orderReadQuery[T] {
	q.whereSql += " and user_id in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

// Find methods
func SelectSongUserFavourite() iQuery[SongUserFavourite] {
	return new(idxQuery[SongUserFavourite])
}

func Select[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereUserIdEQ(val uint) iQuery1[T] {
	q.whereSql += " where user_id = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...uint) iQuery1[T] {
	q.whereSql = " where user_id in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) GetWhere() (whereSql string, params []any) {
	var orderSql string
	if len(q.orders) > 0 {
		orderSql = " order by " + strings.Join(q.orders, ",")
	}
	return q.whereSql + orderSql + q.limitSql, q.whereParams
}
