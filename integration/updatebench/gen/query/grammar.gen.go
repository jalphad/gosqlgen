package query

import (
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/expr"
)

func Distinct(expressions ...ast.Expression) *ast.KeywordExpression {
	return expr.Distinct(expressions...)
}

func IsNotNull(expression ast.Expression) *ast.BoolType {
	return expr.IsNotNull(expression)
}
