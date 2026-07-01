package query

import (
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
	"github.com/jalphad/gosqlgen/integration/output/query/expr"
)

func Distinct(expressions ...ast.Expression) *ast.KeywordExpression {
	return expr.Distinct(expressions...)
}

func IsNotNull(expression ast.Expression) *ast.BoolType {
	return expr.IsNotNull(expression)
}
