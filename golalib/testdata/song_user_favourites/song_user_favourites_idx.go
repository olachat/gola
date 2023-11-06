// Code generated by gola 0.1.1; DO NOT EDIT.

package song_user_favourites

import (
	"fmt"
	"strings"

	"github.com/olachat/gola/v2/coredb"
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
			q.orders[i] = "`user_id` asc"
		case UserIdDesc:
			q.orders[i] = "`user_id` desc"
		case SongIdAsc:
			q.orders[i] = "`song_id` asc"
		case SongIdDesc:
			q.orders[i] = "`song_id` desc"
		case RemarkAsc:
			q.orders[i] = "`remark` asc"
		case RemarkDesc:
			q.orders[i] = "`remark` desc"
		case IsFavouriteAsc:
			q.orders[i] = "`is_favourite` asc"
		case IsFavouriteDesc:
			q.orders[i] = "`is_favourite` desc"
		case CreatedAtAsc:
			q.orders[i] = "`created_at` asc"
		case CreatedAtDesc:
			q.orders[i] = "`created_at` desc"
		case UpdatedAtAsc:
			q.orders[i] = "`updated_at` asc"
		case UpdatedAtDesc:
			q.orders[i] = "`updated_at` desc"
		}
	}
	return q
}

func (q *idxQuery[T]) All() []*T {
	result, _ := coredb.Find[T](DBName, TableName, q)
	return result
}

func (q *idxQuery[T]) Limit(offset, limit int) []*T {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, limit)
	result, _ := coredb.Find[T](DBName, TableName, q)
	return result
}

func (q *idxQuery[T]) AllFromMaster() []*T {
	result, _ := coredb.FindFromMaster[T](DBName, TableName, q)
	return result
}

func (q *idxQuery[T]) LimitFromMaster(offset, limit int) []*T {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, limit)
	result, _ := coredb.FindFromMaster[T](DBName, TableName, q)
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
	AndIsFavouriteEQ(val bool) orderReadQuery[T]
	AndIsFavouriteIN(vals ...bool) orderReadQuery[T]
	AndSongIdEQ(val uint) orderReadQuery[T]
	AndSongIdIN(vals ...uint) orderReadQuery[T]
	orderReadQuery[T]
}

type idxQuery1[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery1[T]) AndIsFavouriteEQ(val bool) orderReadQuery[T] {
	q.whereSql += " and `is_favourite` = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery1[T]) AndIsFavouriteIN(vals ...bool) orderReadQuery[T] {
	q.whereSql += " and `is_favourite` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

func (q *idxQuery1[T]) AndSongIdEQ(val uint) orderReadQuery[T] {
	q.whereSql += " and `song_id` = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery1[T]) AndSongIdIN(vals ...uint) orderReadQuery[T] {
	q.whereSql += " and `song_id` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

// Find methods

// Select returns rows from `song_user_favourites` table with index awared query
func Select() iQuery[SongUserFavourite] {
	return new(idxQuery[SongUserFavourite])
}

// SelectFields returns rows with selected fields from `song_user_favourites` table with index awared query
func SelectFields[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereUserIdEQ(val uint) iQuery1[T] {
	q.whereSql += " where `user_id` = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...uint) iQuery1[T] {
	q.whereSql = " where `user_id` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
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
