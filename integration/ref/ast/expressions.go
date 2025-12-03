package ast

import "strings"

func Render(e Expression, params *[]any) string {
	return e.toSQL(params)
}

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	toSQL(*[]any) string
	getNode() ExpressionNode
}

type AliasedExpression interface {
	NamedExpression
	hasAlias()
}

type NamedExpression interface {
	Expression
	Name() string
}

func NewAsExpression[T MappedTypes](expression OfType[T], alias string) *AsExpression[T] {
	return &AsExpression[T]{
		ofType: expression,
		alias:  alias,
	}
}

type AsExpression[T MappedTypes] struct {
	ofType[T]
	alias string
}

func (a *AsExpression[T]) Name() string {
	return a.alias
}

func (a *AsExpression[T]) hasAlias() {}

func NewNamedExpression[T MappedTypes](name string, expression OfType[T]) *NamedExpressionWrapper[T] {
	return &NamedExpressionWrapper[T]{
		ofType: expression,
		name:   name,
	}
}

type NamedExpressionWrapper[T MappedTypes] struct {
	ofType[T]
	name string
}

func (e *NamedExpressionWrapper[T]) Name() string {
	return e.name
}

// GroupedExpression represents an expression surrounded by parentheses
type GroupedExpression struct {
	expression
}

func NewGroupedExpression(e Expression) *GroupedExpression {
	return &GroupedExpression{e}
}

func (e *GroupedExpression) toSQL(params *[]any) string {
	return "(" + e.expression.toSQL(params) + ")"
}

// KeywordExpression represents known SQL keywords in the SQL language (like DISTINCT)
type KeywordExpression struct {
	node ExpressionNode
}

func NewKeywordExpression(op string, expressions ...Expression) *KeywordExpression {
	return &KeywordExpression{ExpressionNode{
		Op:   op,
		Args: expressions,
	}}
}

func (e *KeywordExpression) toSQL(params *[]any) string {
	parts := []string{}
	for _, arg := range e.node.Args {
		parts = append(parts, arg.toSQL(params))
	}
	return e.node.Op + strings.Join(parts, ", ")
}

type ColumnExpression struct {
	columnNode
}

func NewColumnExpresion(table, column string) *ColumnExpression {
	return &ColumnExpression{
		columnNode: columnNode{
			Column: &ColumnRef{
				Table:  table,
				Column: column,
			},
		},
	}
}

type StringColumnExpression struct {
	*stringType
}

func NewStringColumnExpression(table, column string) *StringColumnExpression {
	return &StringColumnExpression{
		String(NewColumnNode(table, column)),
	}
}

func (e *StringColumnExpression) Name() string {
	return e.getNode().Column.Column
}

type IntColumnExpression struct {
	*intType
}

func (e *IntColumnExpression) Name() string {
	return e.getNode().Column.Column
}

type BoolColumnExpression struct {
	*boolType
}

func NewBoolColumnExpression(table, column string) *BoolColumnExpression {
	return &BoolColumnExpression{
		Bool(NewColumnNode(table, column)),
	}
}

func (e *BoolColumnExpression) Name() string {
	return e.getNode().Column.Column
}

type UUIDColumnExpression struct {
	*uuidType
}

func NewUUIDColumnExpression(table, column string) *UUIDColumnExpression {
	return &UUIDColumnExpression{
		UUID(NewColumnNode(table, column)),
	}
}

func (e *UUIDColumnExpression) Name() string {
	return e.getNode().Column.Column
}

type (
	expression = Expression
)
