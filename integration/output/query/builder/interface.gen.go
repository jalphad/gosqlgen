package builder

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
)

type ResultSelectQuery[O any] interface {
	With(ctes ...*ast.CTE) ResultSelectQuery[O]
	Select(projections ...ast.Projection[O]) ResultSelectJoinQuery[O]
}

type ResultSelectFromQuery[O any] interface {
	From(table *ast.TableSource) ResultSelectWhereQuery[O]
	ResultSelectWhereQuery[O]
}

type ResultSelectJoinQuery[O any] interface {
	With(ctes ...*ast.CTE) ResultSelectJoinQuery[O]
	Join(joinType ast.JoinType, table ast.NamedTableExpression, expr ast.OfType[bool]) ResultSelectJoinQuery[O]
	ResultSelectWhereQuery[O]
}

type ResultSelectWhereQuery[O any] interface {
	Where(expr ast.OfType[bool]) ResultSelectGroupByQuery[O]
	ResultSelectGroupByQuery[O]
}

type ResultSelectGroupByQuery[O any] interface {
	GroupBy(columns ...ast.Expression) ResultSelectHavingQuery[O]
	ResultSelectHavingQuery[O]
}

type ResultSelectHavingQuery[O any] interface {
	Having(expr ast.OfType[bool]) ResultSelectOrderByQuery[O]
	ResultSelectOrderByQuery[O]
}

type ResultSelectOrderByQuery[O any] interface {
	OrderBy(orderBy ...*ast.OrderByItem) ResultSelectPagingQuery[O]
	ResultSelectPagingQuery[O]
}

type ResultSelectPagingQuery[O any] interface {
	Limit(limit int) ResultSelectPagingQuery[O]
	Offset(offset int) ResultSelectPagingQuery[O]
	ResultSelectFinalizeQuery[O]
}

type ResultSelectFinalizeQuery[O any] interface {
	Find(ctx context.Context) ([]O, error)
	FindOne(ctx context.Context) (O, error)
	ToSql() (string, error)
	ToSqlArgs() (string, []any, error)
	Statement() ast.SqlStatement
}

type InsertQuery[T any, O DTO[T]] interface {
	Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T, O]
}

type InsertOnConflictQuery[T any, O DTO[T]] interface {
	OnConflict(columns ...ast.NamedExpression) InsertOnConflictDoQuery[T, O]
	InsertReturningQuery[T, O]
}

type InsertOnConflictDoQuery[T any, O DTO[T]] interface {
	Do(expr ast.OnConflictDoExpression) InsertReturningQuery[T, O]
	InsertReturningQuery[T, O]
}

type InsertReturningQuery[T any, O DTO[T]] interface {
	Returning(projections ...ast.Projection[T]) InsertValuesQuery[T, O]
	InsertValuesQuery[T, O]
}

type InsertValuesQuery[T any, O DTO[T]] interface {
	Values(values ...O) InsertFinalizeQuery[T, O]
}

type InsertFinalizeQuery[T any, O DTO[T]] interface {
	Exec(context.Context) error
	ToSql() (string, error)
	ToSqlArgs() (string, []any, error)
	Statement() ast.SqlStatement
}

type UpdateQuery[T any, O DTO[T]] interface {
	Update(toSet ...ast.UpdateSetExpr) UpdateFromQuery[T, O]
}

type UpdateFromQuery[T any, O DTO[T]] interface {
	From(table ast.NamedTableExpression) UpdateWhereQuery[T, O]
	UpdateWhereQuery[T, O]
}

type UpdateJoinQuery[T any, O DTO[T]] interface {
	Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) UpdateJoinQuery[T, O]
	UpdateWhereQuery[T, O]
}

type UpdateWhereQuery[T any, O DTO[T]] interface {
	Where(expr ast.OfType[bool]) UpdateReturningQuery[T, O]
	UpdateReturningQuery[T, O]
}

type UpdateReturningQuery[T any, O DTO[T]] interface {
	Returning(projections ...ast.Projection[T]) UpdateFinalizeQuery[T, O]
	UpdateFinalizeQuery[T, O]
}

type UpdateFinalizeQuery[T any, O DTO[T]] interface {
	Exec(context.Context) (int64, []T, error)
	ToSql() (string, error)
	ToSqlArgs() (string, []any, error)
	Statement() ast.SqlStatement
}

type DeleteQuery[T any, O DTO[T]] interface {
	Delete(from string) DeleteUsingQuery[T, O]
}

type KnownTableDeleteQuery[T any, O DTO[T]] interface {
	Delete() DeleteUsingQuery[T, O]
}

type DeleteUsingQuery[T any, O DTO[T]] interface {
	Using(tables ...ast.NamedExpression) DeleteWhereQuery[T, O]
	DeleteWhereQuery[T, O]
}

type DeleteWhereQuery[T any, O DTO[T]] interface {
	Where(expr ast.OfType[bool]) DeleteReturningQuery[T, O]
	DeleteReturningQuery[T, O]
}

type DeleteReturningQuery[T any, O DTO[T]] interface {
	Returning(projections ...ast.Projection[T]) DeleteFinalizeQuery[T, O]
	DeleteFinalizeQuery[T, O]
}

type DeleteFinalizeQuery[T any, O DTO[T]] interface {
	Exec(ctx context.Context) (int64, []T, error)
	ToSql() (string, error)
	ToSqlArgs() (string, []any, error)
	Statement() ast.SqlStatement
}

type KnownTableStartQuery[T any, O DTO[T]] interface {
	With(ctes ...*ast.CTE) KnownTableStartQuery[T, O]
	WithTx(tx pgx.Tx) KnownTableStartQuery[T, O]
	KnownTableDeleteQuery[T, O]
	InsertQuery[T, O]
	UpdateQuery[T, O]
}
