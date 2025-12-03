package query

import "github.com/jalphad/gosqlgen/integration/ref/ast"

func Coalesce[T ast.MappedTypes](expressions ...ast.OfType[T]) ast.Expression {
	args := make([]ast.Expression, 0, len(expressions))
	for _, expression := range expressions {
		args = append(args, expression)
	}
	return &ast.FunctionNode{
		Op:   "COALESCE",
		Args: args,
	}
}

func JsonAgg(expression ast.Expression) *AggregationFunction {
	return &AggregationFunction{
		function: ast.FunctionNode{
			Op:   "json_agg",
			Args: []ast.Expression{expression},
		},
	}
}

func JsonbBuildObject(expressions ...ast.NamedExpression) ast.Expression {
	args := make([]ast.Expression, 0, len(expressions))
	for _, expression := range expressions {
		args = append(args, expression)
	}
	return &ast.FunctionNode{
		Op:   "jsonb_build_object",
		Args: args,
	}
}

func Pair(name string, expression ast.Expression) ast.NamedExpression {
	return ast.NewNamedExpression(name, expression)
}

type function = ast.FunctionNode
type AggregationFunction struct {
	function
}

func (f *AggregationFunction) Filter(filter ast.OfType[bool]) ast.Expression {
	return &ast.BinaryNode{
		Op: "FILTER",
		Args: []ast.Expression{
			f,
			&ast.GroupedExpression{
				Args: []ast.Expression{
					&ast.UnaryNode{
						Op: "WHERE",
						Args: []ast.Expression{
							filter,
						},
					},
				},
			},
		},
	}
}
