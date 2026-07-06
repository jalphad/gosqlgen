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
	if ctx != nil && ctx.Error != nil {
		return
	}
	// Track column reference in context
	if ctx != nil {
		currentColumn := ColumnReference{
			Table:  n.Column.Table,
			Column: n.Column.Column,
		}
		ctx.ColumnReferences = append(ctx.ColumnReferences, currentColumn)
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
	if ctx != nil && ctx.Error != nil {
		return
	}
	*params = append(*params, n.Literal)
	builder.WriteString(fmt.Sprintf("$%d", len(*params)))
}

type KeywordNode struct {
	keyword string
}

func NewKeywordNode(keyword string) *KeywordNode {
	return &KeywordNode{keyword: keyword}
}

func (n *KeywordNode) toSQL(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if !isKeywordToken(n.keyword) {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("invalid SQL keyword token %q", n.keyword)
		}
		return
	}
	builder.WriteString(n.keyword)
}

type InlineStringLiteralNode struct {
	value string
}

// NewInlineStringLiteralNode returns an expression rendered directly into SQL
// as a quoted string literal instead of as a bind parameter. Use this only for
// trusted, static values controlled by the program. Never pass user input,
// request data, database values, or partially sanitized strings.
func NewInlineStringLiteralNode(value string) *InlineStringLiteralNode {
	return &InlineStringLiteralNode{value: value}
}

func (n *InlineStringLiteralNode) toSQL(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString("'")
	builder.WriteString(strings.ReplaceAll(n.value, "'", "''"))
	builder.WriteString("'")
}

type UnaryNode ExpressionNode

func (n *UnaryNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(n.Op + " ")
	n.Args[0].toSQL(builder, params, ctx)
}

type BinaryNode ExpressionNode

func (n *BinaryNode) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
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
	if ctx != nil && ctx.Error != nil {
		return
	}
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
	if ctx != nil && ctx.Error != nil {
		return
	}
	if len(node.Args) == 0 {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("function %q has no arguments", node.Op)
		}
		return
	}
	builder.WriteString(node.Op + "(")
	for i := 0; i < len(node.Args)-1; i++ {
		node.Args[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	node.Args[len(node.Args)-1].toSQL(builder, params, ctx)
	builder.WriteString(")")
}

func isKeywordToken(keyword string) bool {
	if keyword == "" {
		return false
	}
	for i, r := range keyword {
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r == '_' {
			continue
		}
		if i > 0 && r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}
