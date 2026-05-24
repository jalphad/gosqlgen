package query

import "github.com/jalphad/gosqlgen/integration/output/query/ast"

func Distinct(expressions ...ast.Expression) *ast.KeywordExpression {
	return ast.NewKeywordExpression("DISTINCT", expressions...)
}
