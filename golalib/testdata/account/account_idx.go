// Code generated by gola 0.1.1; DO NOT EDIT.

package account

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
	TypeAsc
	TypeDesc
	CountryCodeAsc
	CountryCodeDesc
	MoneyAsc
	MoneyDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case UserIdAsc:
			q.orders[i] = "`user_id` asc"
		case UserIdDesc:
			q.orders[i] = "`user_id` desc"
		case TypeAsc:
			q.orders[i] = "`type` asc"
		case TypeDesc:
			q.orders[i] = "`type` desc"
		case CountryCodeAsc:
			q.orders[i] = "`country_code` asc"
		case CountryCodeDesc:
			q.orders[i] = "`country_code` desc"
		case MoneyAsc:
			q.orders[i] = "`money` asc"
		case MoneyDesc:
			q.orders[i] = "`money` desc"
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
	WhereUserIdEQ(val int) iQuery1[T]
	WhereUserIdIN(vals ...int) iQuery1[T]
	orderReadQuery[T]
}
type iQuery1[T any] interface {
	AndCountryCodeEQ(val uint) orderReadQuery[T]
	AndCountryCodeIN(vals ...uint) orderReadQuery[T]
	orderReadQuery[T]
}

type idxQuery1[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery1[T]) AndCountryCodeEQ(val uint) orderReadQuery[T] {
	q.whereSql += " and `country_code` = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery1[T]) AndCountryCodeIN(vals ...uint) orderReadQuery[T] {
	q.whereSql += " and `country_code` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

// Find methods

// Select returns rows from `account` table with index awared query
func Select() iQuery[Account] {
	return new(idxQuery[Account])
}

// SelectFields returns rows with selected fields from `account` table with index awared query
func SelectFields[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereUserIdEQ(val int) iQuery1[T] {
	q.whereSql += " where `user_id` = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereUserIdIN(vals ...int) iQuery1[T] {
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
