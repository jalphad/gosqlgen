package query

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Lit[T ast.MappedTypes](t T) ast.OfType[T] {
	return ast.NewSQLType(t)
}

func As[T ast.MappedTypes](alias string, expr ast.OfType[T]) *ast.AsExpression[T] {
	return ast.NewAsExpression(expr, alias)
}

func Pair[T ast.MappedTypes](name string, expression ast.OfType[T]) *ast.NamedExpressionWrapper[T] {
	return ast.NewNamedExpression(name, expression)
}

// expression helpers
func Column[T ast.MappedTypes](table, col string) *ast.NamedExpressionWrapper[T] {
	//return ast.NewNamedExpression(col, ast.NewNewColumnNode(table, col))
	panic("not implemented")
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
