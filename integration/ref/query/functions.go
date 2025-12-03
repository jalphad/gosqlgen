package query

import (
	"encoding/json"
	"strings"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Coalesce[T ast.MappedTypes](expressions ...ast.OfType[T]) *ast.Function[T] {
	args := make([]ast.Expression, 0, len(expressions))
	for _, expression := range expressions {
		args = append(args, expression)
	}
	return ast.NewFunction[T](ast.NewFunctionNode("COALESCE", args, nil))
}

func JsonAgg(expression ast.Expression) *ast.AggregationFunction[json.RawMessage] {
	return ast.NewAggregationFunction[json.RawMessage](
		ast.NewFunctionNode("json_agg", []ast.Expression{expression}, nil))
}

func JsonbBuildObject(expressions ...ast.NamedExpression) ast.Expression {
	return ast.NewFunctionNode(
		"jsonb_build_object",
		nil,
		func(node ast.ExpressionNode, params *[]any) string {
			parts := make([]string, 0, len(expressions))
			for _, expression := range expressions {
				parts = append(parts, "'"+expression.Name()+"', "+ast.Render(expression, params))
			}
			return node.Op + "(" + strings.Join(parts, ", ") + ")"
		},
	)
}
