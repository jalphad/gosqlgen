package ast

import (
	"fmt"
	"strings"
)

func NewColumnExpresion(table, column string) *ColumnNode {
	return &ColumnNode{
		Table:  table,
		Column: column,
	}
}

type ColumnNode struct {
	Table  string // Optional: table Name
	Column string // Column Name
}

func (e *ColumnNode) toSQL(_ *[]any) string {
	return e.Table + "." + e.Column
}

func (e *ColumnNode) Name() string {
	return e.Column
}

func NewLiteralExpression(val any) *LiteralNode {
	return &LiteralNode{
		value: val,
	}
}

// LiteralNode represents literal values passed as arguments to the SQL prepared statement
type LiteralNode struct {
	value any // Literal value
}

func (e *LiteralNode) toSQL(params *[]any) string {
	*params = append(*params, e.value)
	return fmt.Sprintf("$%d", len(*params))
}

type UnaryNode ExpressionNode

func (e *UnaryNode) toSQL(params *[]any) string {
	return e.Op + " " + e.Args[0].toSQL(params)
}

type BinaryNode ExpressionNode

func (e *BinaryNode) toSQL(params *[]any) string {
	left := e.Args[0]
	right := e.Args[1]
	if _, ok := right.(*BinaryNode); ok {
		right = &GroupedExpression{Args: []Expression{right}}
	}

	return left.toSQL(params) + " " + e.Op + " " + right.toSQL(params)
}

type FunctionNode ExpressionNode

func (e *FunctionNode) toSQL(params *[]any) string {
	parts := make([]string, 0, len(e.Args))
	for _, arg := range e.Args {
		parts = append(parts, arg.toSQL(params))
	}
	return e.Op + "(" + strings.Join(parts, ", ") + ")"
}

// GroupedExpression represents an expression surrounded by parentheses
type GroupedExpression ExpressionNode

func (e *GroupedExpression) toSQL(params *[]any) string {
	return "(" + e.Args[0].toSQL(params) + ")"
}

// KeywordExpression represents known SQL keywords in the SQL language (like DISTINCT)
type KeywordExpression ExpressionNode

func (e *KeywordExpression) toSQL(params *[]any) string {
	parts := []string{}
	for _, arg := range e.Args {
		parts = append(parts, arg.toSQL(params))
	}
	return e.Op + strings.Join(parts, ", ")
}

// ExpressionNode represents a non-leaf node in the SQL Expression AST
type ExpressionNode struct {
	Op   string       // Operator or function Name
	Args []Expression // Arguments (for functions or nested ops)
}
