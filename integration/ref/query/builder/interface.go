package builder

import (
	"context"

	"github.com/jackc/pgx/v5"
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
	Returning(columns ...ast.NamedExpression) InsertFinalizeQuery[T, O]
	InsertFinalizeQuery[T, O]
}

type InsertFinalizeQuery[T any, O DTO[T]] interface {
	Exec(context.Context, O) error
	ToSql() string
}

type UpdateQuery[T any, O DTO[T]] interface {
	Update(columns ...ast.NamedExpression) UpdateFromQuery[T, O]
}

type UpdateFromQuery[T any, O DTO[T]] interface {
	From(table *ast.TableSource) UpdateWhereQuery[T, O]
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
	Returning(columns ...ast.NamedExpression) UpdateFinalizeQuery[T, O]
	UpdateFinalizeQuery[T, O]
}

type UpdateFinalizeQuery[T any, O DTO[T]] interface {
	Exec(context.Context, O) (int64, []T, error)
	ToSql() string
}

type KnownTableStartQuery[T any, O DTO[T]] interface {
	WithTx(tx pgx.Tx) KnownTableStartQuery[T, O]
	KnownTableSelectQuery[T]
	InsertQuery[T, O]
	UpdateQuery[T, O]
}
