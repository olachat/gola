// Code generated by gola 0.1.0; DO NOT EDIT.

package profile

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
	LevelAsc
	LevelDesc
	NickNameAsc
	NickNameDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case UserIdAsc:
			q.orders[i] = "`user_id` asc"
		case UserIdDesc:
			q.orders[i] = "`user_id` desc"
		case LevelAsc:
			q.orders[i] = "`level` asc"
		case LevelDesc:
			q.orders[i] = "`level` desc"
		case NickNameAsc:
			q.orders[i] = "`nick_name` asc"
		case NickNameDesc:
			q.orders[i] = "`nick_name` desc"
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

type order[T any] interface {
	OrderBy(args ...orderBy) coredb.ReadQuery[T]
}

type orderReadQuery[T any] interface {
	order[T]
	coredb.ReadQuery[T]
}

type iQuery[T any] interface {
	WhereUserIdEQ(val int) orderReadQuery[T]
	WhereUserIdIN(vals ...int) orderReadQuery[T]
	orderReadQuery[T]
}

// Find methods

// Select returns rows from `profile` table with index awared query
func Select() iQuery[Profile] {
	return new(idxQuery[Profile])
}

// SelectFields returns rows with selected fields from `profile` table with index awared query
func SelectFields[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereUserIdEQ(val int) orderReadQuery[T] {
	q.whereSql += " where `user_id` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...int) orderReadQuery[T] {
	q.whereSql = " where `user_id` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
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
