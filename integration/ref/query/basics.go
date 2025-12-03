package query

import "github.com/jalphad/gosqlgen/integration/ref/ast"

func Distinct(expressions ...ast.Expression) ast.Expression {
	return &ast.KeywordExpression{
		Op:   "DISTINCT",
		Args: expressions,
	}
}
