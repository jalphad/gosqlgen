package query

import "github.com/jalphad/gosqlgen/integration/ref/ast"

func Coalesce(expressions ...ast.Expression) ast.Expression {
	return ast.NewFunctionExpression(&ast.ExpressionNode{
		Op:      "COALESCE",
		Args:    expressions,
		Literal: nil,
		Field:   nil,
	})
}

func JsonbBuildObject(expressions ...ast.Expression) ast.Expression {
	return ast.NewFunctionExpression(&ast.ExpressionNode{
		Op: "jsonb_build_object",
	})
}
