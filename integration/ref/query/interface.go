package query

import (
	"context"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type SelectQuery[O any] interface {
	Select(columns ...ast.ValueExpression) JoinQuery[O]
}

type JoinQuery[O any] interface {
	Join(joinType ast.JoinType, table string, expr ast.BooleanExpression) WhereQuery[O]
	WhereQuery[O]
}

type WhereQuery[O any] interface {
	Where(expr ast.BooleanExpression) GroupByQuery[O]
	GroupByQuery[O]
}

type GroupByQuery[O any] interface {
	GroupBy(columns ...ast.ValueExpression) HavingQuery[O]
	HavingQuery[O]
}

type HavingQuery[O any] interface {
	Having(expr ast.BooleanExpression) OrderByQuery[O]
	OrderByQuery[O]
}

type OrderByQuery[O any] interface {
	OrderBy(orderBy ...*ast.OrderByItem) PagingQuery[O]
	PagingQuery[O]
}

type PagingQuery[O any] interface {
	Limit(limit int) PagingQuery[O]
	Offset(offset int) PagingQuery[O]
	CompleteQuery[O]
}

type CompleteQuery[O any] interface {
	Find(ctx context.Context) ([]O, error)
	FindOne(ctx context.Context) (O, error)
	ToSql() (string, []interface{})
}
