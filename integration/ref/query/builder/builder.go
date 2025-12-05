package query

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type DTO[T any] interface {
	*T
	ScanInto(rows pgx.Rows, stmt *ast.SelectStatement) error
}

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[T any, O DTO[T]] struct {
	builder[T, O]
}

func (qb *KnownTableBuilder[T, O]) Select(columns ...ast.NamedExpression) JoinQuery[T] {
	qb.builder.stmt.SelectList = append(qb.builder.stmt.SelectList, columns...)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) JoinQuery[T] {
	_ = qb.builder.stmt.From.Join(joinType, table, expr)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Where(expr ast.OfType[bool]) GroupByQuery[T] {
	qb.builder.stmt.Where = expr
	return &qb.builder
}

type Builder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.SelectStatement
	params []interface{}
}

func (qb *Builder[T, O]) Select(columns ...ast.NamedExpression) FromQuery[T] {
	qb.stmt.SelectList = append(qb.stmt.SelectList, columns...)
	return qb
}

func (qb *Builder[T, O]) From(table *ast.TableSource) WhereQuery[T] {
	qb.stmt.From = table
	return qb
}

func (qb *Builder[T, O]) Where(expr ast.OfType[bool]) GroupByQuery[T] {
	qb.stmt.Where = expr
	return qb
}

func (qb *Builder[T, O]) GroupBy(columns ...ast.Expression) HavingQuery[T] {
	qb.stmt.GroupBy = columns
	return qb
}

func (qb *Builder[T, O]) Having(expr ast.OfType[bool]) OrderByQuery[T] {
	qb.stmt.Having = expr
	return qb
}

func (qb *Builder[T, O]) OrderBy(orderBy ...*ast.OrderByItem) PagingQuery[T] {
	qb.stmt.OrderBy = orderBy
	return qb
}

func (qb *Builder[T, O]) Limit(limit int) PagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Limit = limit
	return qb
}

func (qb *Builder[T, O]) Offset(offset int) PagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Offset = offset
	return qb
}

func (qb *Builder[T, O]) ToSql() (string, []any) {
	sql := ast.Render(qb.stmt, &qb.params)
	return sql, qb.params
}

func (qb *Builder[T, O]) Find(ctx context.Context) ([]T, error) {
	query, args := qb.ToSql()

	var rows pgx.Rows
	var err error

	if qb.tx != nil {
		rows, err = qb.tx.Query(ctx, query, args...)
	} else {
		rows, err = qb.pool.Query(ctx, query, args...)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var t T
		result := O(&t)
		if err := result.ScanInto(rows, qb.stmt); err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}

func (qb *Builder[T, O]) FindOne(ctx context.Context) (T, error) {
	qb.Limit(1)
	results, err := qb.Find(ctx)
	if err != nil {
		var t T
		return t, err
	}
	if len(results) == 0 {
		var t T
		return t, pgx.ErrNoRows
	}
	return results[0], nil
}

func NewUsersQuery(pool *pgxpool.Pool) *KnownTableBuilder[models.UsersDto, *models.UsersDto] {
	return &KnownTableBuilder[models.UsersDto, *models.UsersDto]{
		builder: Builder[models.UsersDto, *models.UsersDto]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "users"},
			},
		},
	}
}

func NewPostsQuery(pool *pgxpool.Pool) *KnownTableBuilder[models.PostsDto, *models.PostsDto] {
	return &KnownTableBuilder[models.PostsDto, *models.PostsDto]{
		builder: Builder[models.PostsDto, *models.PostsDto]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "posts"},
			},
		},
	}
}

func NewCommentsQuery(pool *pgxpool.Pool) *KnownTableBuilder[models.CommentsDto, *models.CommentsDto] {
	return &KnownTableBuilder[models.CommentsDto, *models.CommentsDto]{
		builder: Builder[models.CommentsDto, *models.CommentsDto]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "comments"},
			},
		},
	}
}

func NewTagsQuery(pool *pgxpool.Pool) *KnownTableBuilder[models.TagsDto, *models.TagsDto] {
	return &KnownTableBuilder[models.TagsDto, *models.TagsDto]{
		builder: Builder[models.TagsDto, *models.TagsDto]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "tags"},
			},
		},
	}
}

func NewPostTagsQuery(pool *pgxpool.Pool) *KnownTableBuilder[models.PostTagsDto, *models.PostTagsDto] {
	return &KnownTableBuilder[models.PostTagsDto, *models.PostTagsDto]{
		builder: Builder[models.PostTagsDto, *models.PostTagsDto]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "post_tags"},
			},
		},
	}
}

func Asc(column ast.Expression) *ast.OrderByItem {
	return &ast.OrderByItem{
		Field:     column,
		Direction: ast.Asc,
	}
}

func Desc(column ast.Expression) *ast.OrderByItem {
	return &ast.OrderByItem{
		Field:     column,
		Direction: ast.Desc,
	}
}

type builder[T any, O DTO[T]] = Builder[T, O]
