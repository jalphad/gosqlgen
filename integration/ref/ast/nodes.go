package ast

import (
	"fmt"
	"strings"
)

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	toSQL(*[]any) string
}

func NewColumnExpresion(ref *ColumnRef) *ColumnExpression {
	return &ColumnExpression{
		expr: &ExpressionNode{
			Field: ref,
		},
	}
}

type ColumnExpression struct {
	expr *ExpressionNode
}

func (e *ColumnExpression) toSQL(_ *[]any) string {
	return e.expr.Field.Table + "." + e.expr.Field.Column
}
func (e *ColumnExpression) selectExprMarker()  {}
func (e *ColumnExpression) whereExprMarker()   {}
func (e *ColumnExpression) groupByExprMarker() {}
func (e *ColumnExpression) havingExprMarker()  {}
func (e *ColumnExpression) orderByExprMarker() {}

func NewLiteralExpression(val any) *LiteralExpression {
	return &LiteralExpression{
		expr: &ExpressionNode{Literal: val},
	}
}

type LiteralExpression struct {
	expr *ExpressionNode
}

func (e *LiteralExpression) toSQL(params *[]any) string {
	*params = append(*params, e.expr.Literal)
	return fmt.Sprintf("$%d", len(*params))
}
func (e *LiteralExpression) selectExprMarker()  {}
func (e *LiteralExpression) whereExprMarker()   {}
func (e *LiteralExpression) groupByExprMarker() {}
func (e *LiteralExpression) havingExprMarker()  {}
func (e *LiteralExpression) orderByExprMarker() {}

func NewUnaryExpression(expr *ExpressionNode) *UnaryExpression {
	return &UnaryExpression{
		expr: expr,
	}
}

type UnaryExpression struct {
	expr *ExpressionNode
}

func (e *UnaryExpression) toSQL(params *[]any) string {
	return e.expr.Op + " " + e.expr.Args[0].toSQL(params)
}

func NewBinaryExpression(expr *ExpressionNode) *BinaryExpression {
	return &BinaryExpression{
		expr: expr,
	}
}

type BinaryExpression struct {
	expr *ExpressionNode
}

func (e *BinaryExpression) toSQL(params *[]any) string {
	left := e.expr.Args[0]
	right := e.expr.Args[1]
	if _, ok := right.(*BinaryExpression); ok {
		right = NewGroupedExpression(&ExpressionNode{Args: []Expression{right}})
	}

	return left.toSQL(params) + " " + e.expr.Op + " " + right.toSQL(params)
}

func NewFunctionExpression(expr *ExpressionNode) *FunctionExpression {
	return &FunctionExpression{
		expr: expr,
	}
}

type FunctionExpression struct {
	expr *ExpressionNode
}

func (e *FunctionExpression) toSQL(params *[]any) string {
	parts := []string{}
	for _, arg := range e.expr.Args {
		parts = append(parts, arg.toSQL(params))
	}
	return e.expr.Op + "(" + strings.Join(parts, ", ") + ")"
}

func NewGroupedExpression(expr *ExpressionNode) *GroupedExpression {
	return &GroupedExpression{
		expr: expr,
	}
}

type GroupedExpression struct {
	expr *ExpressionNode
}

func (e *GroupedExpression) toSQL(params *[]any) string {
	return "(" + e.expr.Args[0].toSQL(params) + ")"
}

// ExpressionNode represents a node in the SQL Expression AST
type ExpressionNode struct {
	Op      string       // Operator or function name
	Args    []Expression // Arguments (for functions or nested ops)
	Literal any          // Literal value (if leaf)
	Field   *ColumnRef   // Optional field reference
}

// ColumnRef represents a type-safe field reference.
// Can be used in SELECT, WHERE, ORDER BY, etc.
type ColumnRef struct {
	Table  string // Optional: table name
	Column string // Column name or SQL ExpressionNode
}
