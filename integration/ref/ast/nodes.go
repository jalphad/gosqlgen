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

func (n *ColumnNode) toSQL(builder *strings.Builder, _ *[]any) {
	builder.WriteString(n.Column.Table + "." + n.Column.Column)
}

func NewLiteralExpression(val any) *LiteralNode {
	return &LiteralNode{
		Literal: val,
	}
}

// LiteralNode represents literal values passed as arguments to the SQL prepared statement
type LiteralNode ExpressionNode

func (n *LiteralNode) toSQL(builder *strings.Builder, params *[]any) {
	*params = append(*params, n.Literal)
	builder.WriteString(fmt.Sprintf("$%d", len(*params)))
}

type UnaryNode ExpressionNode

func (n *UnaryNode) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString(n.Op + " ")
	n.Args[0].toSQL(builder, params)
}

type BinaryNode ExpressionNode

func (n *BinaryNode) toSQL(builder *strings.Builder, params *[]any) {
	left := n.Args[0]
	right := n.Args[1]
	if _, ok := right.(*BinaryNode); ok {
		right = NewGroupedExpression(right)
	}

	left.toSQL(builder, params)
	builder.WriteString(" " + n.Op + " ")
	right.toSQL(builder, params)
}

type FunctionNode struct {
	node     ExpressionNode
	renderFn RenderFunc
}

func NewFunctionNode(op string, args []Expression, fn RenderFunc) *FunctionNode {
	return &FunctionNode{
		node: ExpressionNode{
			Op:   op,
			Args: args,
		},
		renderFn: fn,
	}
}

func (n *FunctionNode) toSQL(builder *strings.Builder, params *[]any) {
	if n.renderFn == nil {
		n.renderFn = FunctionDefaultRender
	}
	n.renderFn(n.node, builder, params)
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

type RenderFunc func(ExpressionNode, *strings.Builder, *[]any)

func FunctionDefaultRender(node ExpressionNode, builder *strings.Builder, params *[]any) {
	builder.WriteString(node.Op + "(")
	for i := 0; i < len(node.Args)-1; i++ {
		node.Args[i].toSQL(builder, params)
		builder.WriteString(", ")
	}
	node.Args[len(node.Args)-1].toSQL(builder, params)
	builder.WriteString(")")
}
