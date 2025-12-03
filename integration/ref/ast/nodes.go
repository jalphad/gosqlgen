package ast

import (
	"fmt"
	"strings"
)

type ColumnNode ExpressionNode

func NewColumnNode(table, column string) *ColumnNode {
	return &ColumnNode{
		Column: &ColumnRef{
			Table:  table,
			Column: column,
		},
	}
}

func (n *ColumnNode) getNode() ExpressionNode {
	return ExpressionNode(*n)
}

func (n *ColumnNode) toSQL(_ *[]any) string {
	return n.Column.Table + "." + n.Column.Column
}

func NewLiteralExpression(val any) *LiteralNode {
	return &LiteralNode{
		Literal: val,
	}
}

// LiteralNode represents literal values passed as arguments to the SQL prepared statement
type LiteralNode ExpressionNode

func (n *LiteralNode) getNode() ExpressionNode {
	return ExpressionNode(*n)
}

func (n *LiteralNode) toSQL(params *[]any) string {
	*params = append(*params, n.Literal)
	return fmt.Sprintf("$%d", len(*params))
}

type UnaryNode ExpressionNode

func (n *UnaryNode) getNode() ExpressionNode {
	return ExpressionNode(*n)
}

func (n *UnaryNode) toSQL(params *[]any) string {
	return n.Op + " " + n.Args[0].toSQL(params)
}

type BinaryNode ExpressionNode

func (n *BinaryNode) getNode() ExpressionNode {
	return ExpressionNode(*n)
}

func (n *BinaryNode) toSQL(params *[]any) string {
	left := n.Args[0]
	right := n.Args[1]
	if _, ok := right.(*BinaryNode); ok {
		right = NewGroupedExpression(right)
	}

	return left.toSQL(params) + " " + n.Op + " " + right.toSQL(params)
}

type FunctionNode struct {
	node     ExpressionNode
	renderFn renderFunc
}

func NewFunctionNode(op string, args []Expression, fn renderFunc) *FunctionNode {
	return &FunctionNode{
		node: ExpressionNode{
			Op:   op,
			Args: args,
		},
		renderFn: fn,
	}
}

func (n *FunctionNode) getNode() ExpressionNode {
	return n.node
}

func (n *FunctionNode) toSQL(params *[]any) string {
	if n.renderFn == nil {
		n.renderFn = FunctionDefaultRender
	}
	return n.renderFn(n.node, params)
}

// ExpressionNode represents a non-leaf node in the SQL Expression AST
type ExpressionNode struct {
	Op      string       // Operator or function Name
	Args    []Expression // Arguments (for functions or nested ops)
	Literal any          // Literal value as passed to prepared statement
	Column  *ColumnRef   // Reference to column
}

type ColumnRef struct {
	Table  string // Optional: table Name
	Column string // Column Name
}

type (
	columnNode   = ColumnNode
	literalNode  = LiteralNode
	unaryNode    = UnaryNode
	binaryNode   = BinaryNode
	functionNode = FunctionNode
)

type renderFunc func(ExpressionNode, *[]any) string

func FunctionDefaultRender(node ExpressionNode, params *[]any) string {
	parts := make([]string, 0, len(node.Args))
	for _, arg := range node.Args {
		parts = append(parts, arg.toSQL(params))
	}
	return node.Op + "(" + strings.Join(parts, ", ") + ")"
}
