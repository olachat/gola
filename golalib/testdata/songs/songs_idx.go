// Code generated by gola 0.0.6; DO NOT EDIT.

package songs

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
	TitleAsc
	TitleDesc
	RankAsc
	RankDesc
	TypeAsc
	TypeDesc
	HashAsc
	HashDesc
	ManifestAsc
	ManifestDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case IdAsc:
			q.orders[i] = "`id` asc"
		case IdDesc:
			q.orders[i] = "`id` desc"
		case TitleAsc:
			q.orders[i] = "`title` asc"
		case TitleDesc:
			q.orders[i] = "`title` desc"
		case RankAsc:
			q.orders[i] = "`rank` asc"
		case RankDesc:
			q.orders[i] = "`rank` desc"
		case TypeAsc:
			q.orders[i] = "`type` asc"
		case TypeDesc:
			q.orders[i] = "`type` desc"
		case HashAsc:
			q.orders[i] = "`hash` asc"
		case HashDesc:
			q.orders[i] = "`hash` desc"
		case ManifestAsc:
			q.orders[i] = "`manifest` asc"
		case ManifestDesc:
			q.orders[i] = "`manifest` desc"
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
	WhereHashEQ(val string) orderReadQuery[T]
	WhereHashIN(vals ...string) orderReadQuery[T]
	orderReadQuery[T]
}

// Find methods

// Select returns rows from `songs` table with index awared query
func Select() iQuery[Song] {
	return new(idxQuery[Song])
}

// SelectFields returns rows with selected fields from `songs` table with index awared query
func SelectFields[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereHashEQ(val string) orderReadQuery[T] {
	q.whereSql += " where `hash` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereHashIN(vals ...string) orderReadQuery[T] {
	q.whereSql = " where `hash` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
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
