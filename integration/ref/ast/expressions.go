package ast

import (
	"fmt"
	"strings"
)

func Render(expression Expression, params *[]any) string {
	return expression.toSQL(params)
}

type ValueExpression interface {
	Expression
	isValueExpression()
}

type BooleanExpression interface {
	Expression
	And(next Expression) BooleanExpression
	Or(next Expression) BooleanExpression
}

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	toSQL(*[]any) string
	getAst() *ExpressionNode
}

func NewFieldExpresion(ref *FieldRef) *FieldExpression {
	return &FieldExpression{
		expr: &ExpressionNode{
			Field: ref,
		},
	}
}

type FieldExpression struct {
	expr *ExpressionNode
}

func (e *FieldExpression) toSQL(_ *[]any) string {
	return e.expr.Field.Table + "." + e.expr.Field.Column
}

func (e *FieldExpression) getAst() *ExpressionNode {
	return e.expr
}

func NewLiteralExpression(expr *ExpressionNode) *LiteralExpression {
	return &LiteralExpression{
		expr: expr,
	}
}

type LiteralExpression struct {
	expr *ExpressionNode
}

func (e *LiteralExpression) toSQL(params *[]any) string {
	*params = append(*params, e.expr.Literal)
	return fmt.Sprintf("$%d", len(*params))
}

func (e *LiteralExpression) getAst() *ExpressionNode {
	return e.expr
}

func NewLogicalExpression(expr *ExpressionNode) *LogicalExpression {
	return &LogicalExpression{
		expr: expr,
	}
}

type LogicalExpression struct {
	expr *ExpressionNode
}

func (e *LogicalExpression) toSQL(params *[]any) string {
	left := e.expr.Args[0]
	right := e.expr.Args[1]
	if _, ok := right.(*LogicalExpression); ok {
		right = NewGroupedExpression(&ExpressionNode{Args: []Expression{right}})
	}
	return fmt.Sprintf("%s %s %s", left.toSQL(params), e.expr.Op, right.toSQL(params))
}

func (e *LogicalExpression) getAst() *ExpressionNode {
	return e.expr
}

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

func (e *UnaryExpression) getAst() *ExpressionNode {
	return e.expr
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
	return e.expr.Args[0].toSQL(params) + " " + e.expr.Op + " " + e.expr.Args[1].toSQL(params)
}

func (e *BinaryExpression) getAst() *ExpressionNode {
	return e.expr
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

func (e *FunctionExpression) getAst() *ExpressionNode {
	return e.expr
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

func (e *GroupedExpression) getAst() *ExpressionNode {
	return e.expr
}

// ExpressionNode represents a node in the SQL Expression AST
type ExpressionNode struct {
	Op      string       // Operator or function name
	Args    []Expression // Arguments (for functions or nested ops)
	Literal any          // Literal value (if leaf)
	Field   *FieldRef    // Optional field reference
}

func NewValueExpression(expr Expression) ValueExpression {
	return &valueExpressionWrapper{
		Expression: expr,
	}
}

type valueExpressionWrapper struct {
	Expression
}

func (*valueExpressionWrapper) isValueExpression() {}
