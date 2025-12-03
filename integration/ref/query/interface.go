package query

import (
	"context"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type SelectQuery[O any] interface {
	Select(columns ...ast.NamedExpression) FromQuery[O]
}

type KnownTableSelectQuery[O any] interface {
	Select(columns ...ast.NamedExpression) JoinQuery[O]
}

type FromQuery[O any] interface {
	From(table *ast.TableSource) WhereQuery[O]
	WhereQuery[O]
}

type JoinQuery[O any] interface {
	Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) JoinQuery[O]
	WhereQuery[O]
}

type WhereQuery[O any] interface {
	Where(expr ast.OfType[bool]) GroupByQuery[O]
	GroupByQuery[O]
}

type GroupByQuery[O any] interface {
	GroupBy(columns ...ast.Expression) HavingQuery[O]
	HavingQuery[O]
}

type HavingQuery[O any] interface {
	Having(expr ast.OfType[bool]) OrderByQuery[O]
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
