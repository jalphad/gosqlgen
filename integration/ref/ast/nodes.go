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

func (n *ColumnNode) toSQL(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	// Track column reference in context
	if ctx != nil {
		ctx.ColumnReferences = append(ctx.ColumnReferences, ColumnReference{
			Table:  n.Column.Table,
			Column: n.Column.Column,
		})
	}
	builder.WriteString(n.Column.Table + "." + n.Column.Column)
}

func NewLiteralExpression(val any) *LiteralNode {
	return &LiteralNode{
		Literal: val,
	}
}

// LiteralNode represents literal values passed as arguments to the SQL prepared statement
type LiteralNode ExpressionNode

func (n *LiteralNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	*params = append(*params, n.Literal)
	builder.WriteString(fmt.Sprintf("$%d", len(*params)))
}

type UnaryNode ExpressionNode

func (n *UnaryNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	builder.WriteString(n.Op + " ")
	n.Args[0].toSQL(builder, params, ctx)
}

type BinaryNode ExpressionNode

func (n *BinaryNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	left := n.Args[0]
	right := n.Args[1]
	if _, ok := right.(*BinaryNode); ok {
		right = NewGroupedExpression(right)
	}

	left.toSQL(builder, params, ctx)
	builder.WriteString(" " + n.Op + " ")
	right.toSQL(builder, params, ctx)
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

func (n *FunctionNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if n.renderFn == nil {
		n.renderFn = FunctionDefaultRender
	}
	n.renderFn(n.node, builder, params, ctx)
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

type RenderFunc func(ExpressionNode, *strings.Builder, *[]any, *QueryContext)

func FunctionDefaultRender(node ExpressionNode, builder *strings.Builder, params *[]any, ctx *QueryContext) {
	builder.WriteString(node.Op + "(")
	for i := 0; i < len(node.Args)-1; i++ {
		node.Args[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	node.Args[len(node.Args)-1].toSQL(builder, params, ctx)
	builder.WriteString(")")
}
