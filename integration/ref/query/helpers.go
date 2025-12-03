package query

import "github.com/jalphad/gosqlgen/integration/ref/ast"

func Lit[T ast.MappedTypes](t T) ast.OfType[T] {
	return ast.NewSQLType(t)
}

func As[T ast.MappedTypes](expr ast.OfType[T], alias string) *ast.AsExpression[T] {
	return ast.NewAsExpression(expr, alias)
}

func Pair[T ast.MappedTypes](name string, expression ast.OfType[T]) ast.NamedExpression {
	return ast.NewNamedExpression(name, expression)
}
