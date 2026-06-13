package query

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

func Null() ast.Expression {
	return expr.Null()
}

func InlineStringLiteral(value string) ast.Expression {
	return expr.InlineStringLiteral(value)
}

func Distinct(expressions ...ast.Expression) *ast.KeywordExpression {
	return expr.Distinct(expressions...)
}

func IsNull(expression ast.Expression) *ast.BoolType {
	return expr.IsNull(expression)
}

func IsNotNull(expression ast.Expression) *ast.BoolType {
	return expr.IsNotNull(expression)
}
