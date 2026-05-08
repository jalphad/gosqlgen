package query

import "github.com/jalphad/gosqlgen/integration/ref/ast"

type SelectOpt struct {
	selectExpr ast.NamedExpression
	joinExpr   JoinExpr
}

type JoinExpr struct {
	Type      ast.JoinType
	Table     string
	Condition ast.OfType[bool]
}
