package builder

import (
	"context"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type SelectQuery[O any] interface {
	Select(columns ...ast.NamedExpression) SelectFromQuery[O]
}

type KnownTableSelectQuery[O any] interface {
	Select(columns ...ast.NamedExpression) SelectJoinQuery[O]
}

type SelectFromQuery[O any] interface {
	From(table *ast.TableSource) SelectWhereQuery[O]
	SelectWhereQuery[O]
}

type SelectJoinQuery[O any] interface {
	Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) SelectJoinQuery[O]
	SelectWhereQuery[O]
}

type SelectWhereQuery[O any] interface {
	Where(expr ast.OfType[bool]) SelectGroupByQuery[O]
	SelectGroupByQuery[O]
}

type SelectGroupByQuery[O any] interface {
	GroupBy(columns ...ast.Expression) SelectHavingQuery[O]
	SelectHavingQuery[O]
}

type SelectHavingQuery[O any] interface {
	Having(expr ast.OfType[bool]) SelectOrderByQuery[O]
	SelectOrderByQuery[O]
}

type SelectOrderByQuery[O any] interface {
	OrderBy(orderBy ...*ast.OrderByItem) SelectPagingQuery[O]
	SelectPagingQuery[O]
}

type SelectPagingQuery[O any] interface {
	Limit(limit int) SelectPagingQuery[O]
	Offset(offset int) SelectPagingQuery[O]
	SelectFinalizeQuery[O]
}

type SelectFinalizeQuery[O any] interface {
	Find(ctx context.Context) ([]O, error)
	FindOne(ctx context.Context) (O, error)
	ToSql() string
}

type UpdateFinalizeQuery[O any] interface {
	Exec(context.Context, O) (int64, error)
	ToSql() string
}
