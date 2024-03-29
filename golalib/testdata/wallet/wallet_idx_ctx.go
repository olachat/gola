// Code generated by gola 0.1.1; DO NOT EDIT.

package wallet

import (
	"context"
	"fmt"

	"github.com/olachat/gola/v2/coredb"
)

func (q *idxQuery[T]) AllCtx(ctx context.Context) ([]*T, error) {
	return coredb.FindCtx[T](ctx, DBName, TableName, q)
}

func (q *idxQuery[T]) LimitCtx(ctx context.Context, offset, limit int) ([]*T, error) {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, limit)
	return coredb.FindCtx[T](ctx, DBName, TableName, q)
}

func (q *idxQuery[T]) AllFromMasterCtx(ctx context.Context) ([]*T, error) {
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, q)
}

func (q *idxQuery[T]) LimitFromMasterCtx(ctx context.Context, offset, limit int) ([]*T, error) {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, limit)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, q)
}
