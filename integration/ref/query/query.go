package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/jalphad/gosqlgen/integration/models"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type Tables interface {
	*models.UsersDto
}

// Builder
type Builder[O Tables] struct {
	o      O
	stmt   *ast.SelectStatement
	params []interface{}
}

func (qb *Builder[O]) Select(columns ...ast.Expression) JoinQuery[O] {
	qb.stmt.SelectList = append(qb.stmt.SelectList, columns...)
	return qb
}

func (qb *Builder[O]) Join(joinType ast.JoinType, table string, expr *ast.BoolType) WhereQuery[O] {
	qb.stmt.From = append(qb.stmt.From, &ast.TableSource{
		Join: &ast.JoinExpr{
			Type:      joinType,
			Right:     &ast.TableSource{TableName: table},
			Condition: expr,
		},
	})
	return qb
}

func (qb *Builder[O]) Where(expr *ast.BoolType) GroupByQuery[O] {
	qb.stmt.Where = expr
	return qb
}

func (qb *Builder[O]) GroupBy(columns ...ast.Expression) HavingQuery[O] {
	qb.stmt.GroupBy = columns
	return qb
}

func (qb *Builder[O]) Having(expr *ast.BoolType) OrderByQuery[O] {
	//TODO implement me
	panic("implement me")
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
	sql := renderSelect(qb.stmt, &qb.params)
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

func NewUsersQuery() *Builder[*models.UsersDto] {
	return &Builder[*models.UsersDto]{
		stmt: &ast.SelectStatement{
			From: []*ast.TableSource{{TableName: "users", Alias: "u"}},
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
func Field(table, col string) ast.Expression {
	return ast.NewColumnExpresion(&ast.ColumnRef{Table: table, Column: col})
}

// Rendering
func renderSelect(stmt *ast.SelectStatement, params *[]interface{}) string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	for i, expr := range stmt.SelectList {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(ast.Render(expr, params))
	}
	sb.WriteString(" FROM " + stmt.From[0].TableName)
	if stmt.Where != nil {
		sb.WriteString(" WHERE " + ast.Render(stmt.Where, params))
	}
	if stmt.GroupBy != nil {

	}
	if stmt.OrderBy != nil {
		sb.WriteString(" ORDER BY ")
		for i, expr := range stmt.OrderBy {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(ast.Render(expr.Field, params) + " " + string(expr.Direction))
		}
	}
	if stmt.Having != nil {

	}
	if stmt.Limit != nil {
		if stmt.Limit.Limit > 0 {
			sb.WriteString(fmt.Sprintf(" LIMIT %d", stmt.Limit.Limit))
		}
		if stmt.Limit.Offset > 0 {
			sb.WriteString(fmt.Sprintf(" OFFSET %d", stmt.Limit.Offset))
		}
	}
	return sb.String()
}
