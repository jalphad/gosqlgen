package builder

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

type KnownTableSelectQuery[O any] interface {
	Select(projections ...ast.Projection[O]) SelectJoinQuery[O]
}

type SelectFromQuery[O any] interface {
	From(table *ast.TableSource) SelectWhereQuery[O]
	SelectWhereQuery[O]
}

type SelectJoinQuery[O any] interface {
	With(ctes ...*ast.CTE) SelectJoinQuery[O]
	Join(joinType ast.JoinType, table ast.NamedTableExpression, expr ast.OfType[bool]) SelectJoinQuery[O]
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
	ToSql() (string, []any, error)
	Statement() ast.SqlStatement
}

type InsertQuery[T any] interface {
	Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T]
}

type InsertOnConflictQuery[T any] interface {
	OnConflict(columns ...ast.NamedExpression) InsertOnConflictDoQuery[T]
	InsertReturningQuery[T]
}

type InsertOnConflictDoQuery[T any] interface {
	Do(expr ast.OnConflictDoExpression) InsertReturningQuery[T]
	InsertReturningQuery[T]
}

type InsertReturningQuery[T any] interface {
	Returning(projections ...ast.Projection[T]) InsertValuesQuery[T]
	InsertValuesQuery[T]
}

type InsertValuesQuery[T any] interface {
	Values(values ...*T) InsertFinalizeQuery[T]
}

type InsertFinalizeQuery[T any] interface {
	Exec(context.Context) error
	ToSql() (string, []any, error)
	Statement() ast.SqlStatement
}

type UpdateQuery[T any] interface {
	Update(toSet ...ast.UpdateSetExpr) UpdateFromQuery[T]
}

type UpdateFromQuery[T any] interface {
	From(table ast.NamedTableExpression) UpdateWhereQuery[T]
	UpdateWhereQuery[T]
}

type UpdateJoinQuery[T any] interface {
	Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) UpdateJoinQuery[T]
	UpdateWhereQuery[T]
}

type UpdateWhereQuery[T any] interface {
	Where(expr ast.OfType[bool]) UpdateReturningQuery[T]
	UpdateReturningQuery[T]
}

type UpdateReturningQuery[T any] interface {
	Returning(projections ...ast.Projection[T]) UpdateFinalizeQuery[T]
	UpdateFinalizeQuery[T]
}

type UpdateFinalizeQuery[T any] interface {
	Exec(context.Context) (int64, []T, error)
	ToSql() (string, []any, error)
	Statement() ast.SqlStatement
}

type DeleteQuery[T any] interface {
	Delete(from string) DeleteUsingQuery[T]
}

type KnownTableDeleteQuery[T any] interface {
	Delete() DeleteUsingQuery[T]
}

type DeleteUsingQuery[T any] interface {
	Using(tables ...ast.NamedExpression) DeleteWhereQuery[T]
	DeleteWhereQuery[T]
}

type DeleteWhereQuery[T any] interface {
	Where(expr ast.OfType[bool]) DeleteReturningQuery[T]
	DeleteReturningQuery[T]
}

type DeleteReturningQuery[T any] interface {
	Returning(projections ...ast.Projection[T]) DeleteFinalizeQuery[T]
	DeleteFinalizeQuery[T]
}

type DeleteFinalizeQuery[T any] interface {
	Exec(ctx context.Context) (int64, []T, error)
	ToSql() (string, []any, error)
	Statement() ast.SqlStatement
}

type KnownTableStartQuery[T any] interface {
	With(ctes ...*ast.CTE) KnownTableStartQuery[T]
	WithTx(tx pgx.Tx) KnownTableStartQuery[T]
	KnownTableSelectQuery[T]
	KnownTableDeleteQuery[T]
	InsertQuery[T]
	UpdateQuery[T]
}
