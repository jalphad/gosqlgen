package query

import (
	"context"

	"github.com/jalphad/gosqlgen/integration/models"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type Tables interface {
	*models.UsersDto
}

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[O any] struct {
	builder[O]
}

func (qb *KnownTableBuilder[O]) Select(columns ...ast.NamedExpression) JoinQuery[O] {
	qb.builder.stmt.SelectList = append(qb.builder.stmt.SelectList, columns...)
	return qb
}

func (qb *KnownTableBuilder[O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) JoinQuery[O] {
	_ = qb.builder.stmt.From.Join(joinType, table, expr)
	return qb
}

func (qb *KnownTableBuilder[O]) Where(expr ast.OfType[bool]) GroupByQuery[O] {
	qb.builder.stmt.Where = expr
	return &qb.builder
}

type Builder[O any] struct {
	stmt   *ast.SelectStatement
	params []interface{}
}

func (qb *Builder[O]) Select(columns ...ast.NamedExpression) FromQuery[O] {
	qb.stmt.SelectList = append(qb.stmt.SelectList, columns...)
	return qb
}

func (qb *Builder[O]) From(table *ast.TableSource) WhereQuery[O] {
	qb.stmt.From = table
	return qb
}

func (qb *Builder[O]) Where(expr ast.OfType[bool]) GroupByQuery[O] {
	qb.stmt.Where = expr
	return qb
}

func (qb *Builder[O]) GroupBy(columns ...ast.Expression) HavingQuery[O] {
	qb.stmt.GroupBy = columns
	return qb
}

func (qb *Builder[O]) Having(expr ast.OfType[bool]) OrderByQuery[O] {
	qb.stmt.Having = expr
	return qb
}

func (qb *Builder[O]) OrderBy(orderBy ...*ast.OrderByItem) PagingQuery[O] {
	qb.stmt.OrderBy = orderBy
	return qb
}

func (qb *Builder[O]) Limit(limit int) PagingQuery[O] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Limit = limit
	return qb
}

func (qb *Builder[O]) Offset(offset int) PagingQuery[O] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Offset = offset
	return qb
}

func (qb *Builder[O]) ToSql() (string, []interface{}) {
	sql := ast.Render(qb.stmt, &qb.params)
	return sql, qb.params
}

func (qb *Builder[O]) Find(ctx context.Context) ([]O, error) {
	//TODO implement me
	panic("implement me")
}

func (qb *Builder[O]) FindOne(ctx context.Context) (O, error) {
	//TODO implement me
	panic("implement me")
}

func NewUsersQuery() *KnownTableBuilder[*models.UsersDto] {
	return &KnownTableBuilder[*models.UsersDto]{
		builder: Builder[*models.UsersDto]{
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: "users"},
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

// expression helpers
func Column(table, col string) ast.Expression {
	return &ast.ColumnNode{
		Column: &ast.ColumnRef{
			Table:  table,
			Column: col,
		},
	}
}

type builder[O any] = Builder[O]
