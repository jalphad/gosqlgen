package query

import (
	"encoding/json"

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

func Into[R any, T ast.MappedTypes, D ast.ScanDest[T]](expr ast.OfType[T], dest func(*R) D) ast.Projection[R] {
	return ast.NewProjection(expr, dest)
}

func IntoCustom[R any, I ast.MappedTypes, O any](expr ast.OfType[I], dest func(*R) *O, convert func(I) (O, error)) ast.Projection[R] {
	return ast.NewCustomProjection(expr, dest, convert)
}

func IntoJSON[R any, O any](expr ast.OfType[json.RawMessage], dest func(*R) *O) ast.Projection[R] {
	return ast.NewJSONProjection(expr, dest)
}

func Rel[T ast.MappedTypes](r *ast.Alias, c ast.NamedAndTyped[T]) ast.OfType[T] {
	return ast.SetType[T](ast.NewColumnNode(r.Name(), c.Name()))
}

func Set[T ast.MappedTypes](c ast.NamedAndTyped[T]) *SetPart[T] {
	return &SetPart[T]{c}
}

func As[T ast.MappedTypes, C ast.AsExprConstraint[T]](alias string, expression ast.AsExpression[T, C]) C {
	return expression.As(ast.NewAlias(alias))
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
