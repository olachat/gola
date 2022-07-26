// Code generated by gola 0.0.3; DO NOT EDIT.

package users

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
	HobbyNoDefaultAsc
	HobbyNoDefaultDesc
	SportsAsc
	SportsDesc
	Sports2Asc
	Sports2Desc
	SportsNoDefaultAsc
	SportsNoDefaultDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case IdAsc:
			q.orders[i] = "`id` asc"
		case IdDesc:
			q.orders[i] = "`id` desc"
		case NameAsc:
			q.orders[i] = "`name` asc"
		case NameDesc:
			q.orders[i] = "`name` desc"
		case EmailAsc:
			q.orders[i] = "`email` asc"
		case EmailDesc:
			q.orders[i] = "`email` desc"
		case CreatedAtAsc:
			q.orders[i] = "`created_at` asc"
		case CreatedAtDesc:
			q.orders[i] = "`created_at` desc"
		case UpdatedAtAsc:
			q.orders[i] = "`updated_at` asc"
		case UpdatedAtDesc:
			q.orders[i] = "`updated_at` desc"
		case FloatTypeAsc:
			q.orders[i] = "`float_type` asc"
		case FloatTypeDesc:
			q.orders[i] = "`float_type` desc"
		case DoubleTypeAsc:
			q.orders[i] = "`double_type` asc"
		case DoubleTypeDesc:
			q.orders[i] = "`double_type` desc"
		case HobbyAsc:
			q.orders[i] = "`hobby` asc"
		case HobbyDesc:
			q.orders[i] = "`hobby` desc"
		case HobbyNoDefaultAsc:
			q.orders[i] = "`hobby_no_default` asc"
		case HobbyNoDefaultDesc:
			q.orders[i] = "`hobby_no_default` desc"
		case SportsAsc:
			q.orders[i] = "`sports` asc"
		case SportsDesc:
			q.orders[i] = "`sports` desc"
		case Sports2Asc:
			q.orders[i] = "`sports2` asc"
		case Sports2Desc:
			q.orders[i] = "`sports2` desc"
		case SportsNoDefaultAsc:
			q.orders[i] = "`sports_no_default` asc"
		case SportsNoDefaultDesc:
			q.orders[i] = "`sports_no_default` desc"
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
	WhereEmailEQ(val string) orderReadQuery[T]
	WhereEmailIN(vals ...string) orderReadQuery[T]
	WhereNameEQ(val string) orderReadQuery[T]
	WhereNameIN(vals ...string) orderReadQuery[T]
	orderReadQuery[T]
}

// Find methods
func SelectUser() iQuery[User] {
	return new(idxQuery[User])
}

func Select[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereEmailEQ(val string) orderReadQuery[T] {
	q.whereSql += " where `email` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereEmailIN(vals ...string) orderReadQuery[T] {
	q.whereSql = " where `email` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q
}

func (q *idxQuery[T]) WhereNameEQ(val string) orderReadQuery[T] {
	q.whereSql += " where `name` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereNameIN(vals ...string) orderReadQuery[T] {
	q.whereSql = " where `name` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q
}

func (q *idxQuery[T]) GetWhere() (whereSql string, params []any) {
	var orderSql string
	if len(q.orders) > 0 {
		orderSql = " order by " + strings.Join(q.orders, ",")
	}
	return q.whereSql + orderSql + q.limitSql, q.whereParams
}