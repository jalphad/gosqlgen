package query

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func NVal[T ast.MappedTypes, N ast.NullableMappedTypes[T]](n N) ast.OfType[T] {
	if n == nil {
		return ast.SetType[T](ast.NewLiteralExpression(nil))
	}

	return ast.SetType[T](ast.NewLiteralExpression(n))
}

func Val[T ast.MappedTypes](t T) ast.OfType[T] {
	return ast.NewSQLType(t)
}

func Rel[T ast.MappedTypes](r *ast.Relation, c ast.NamedAndTyped[T]) ast.OfType[T] {
	return ast.SetType[T](ast.NewColumnNode(r.Name(), c.Name()))
}

func Set[T ast.MappedTypes](c ast.NamedAndTyped[T]) *SetPart[T] {
	return &SetPart[T]{c}
}

type SetPart[T ast.MappedTypes] struct {
	c ast.NamedAndTyped[T]
}

func (s *SetPart[T]) To(val ast.OfType[T]) ast.UpdateSet {
	return ast.UpdateSet{
		Key:   s.c,
		Value: val,
	}
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
