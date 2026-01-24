package ast

import (
	"strings"
)

func Render(e Expression, params *[]any) string {
	var query strings.Builder
	e.toSQL(&query, params)

	return query.String()
}

func BuildQuery(e Expression, builder *strings.Builder, params *[]any) {
	e.toSQL(builder, params)
}

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	// TODO: refactor to allow returning an error
	toSQL(builder *strings.Builder, params *[]any)
}

type AliasedExpression interface {
	NamedExpression
	hasAlias()
}

type NamedExpression interface {
	Expression
	Name() string
}

type NamedTableExpression interface {
	TableExpression
	NamedExpression
}

type TableExpression interface {
	Expression
	isTableExpression()
}

type AsExpression[T MappedTypes] struct {
	ofType[T]
	alias string
}

func NewAsExpression[T MappedTypes](expression OfType[T], alias string) *AsExpression[T] {
	return &AsExpression[T]{
		ofType: expression,
		alias:  alias,
	}
}

func (a *AsExpression[T]) Name() string {
	return a.alias
}

func (a *AsExpression[T]) hasAlias() {}

func (a *AsExpression[T]) toSQL(builder *strings.Builder, params *[]any) {
	a.ofType.toSQL(builder, params)
	builder.WriteString(" AS " + a.alias)
}

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
	expressions []expression
}

func NewGroupedExpression(e ...Expression) *GroupedExpression {
	return &GroupedExpression{e}
}

func (e *GroupedExpression) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString("(")
	for i := 0; i < len(e.expressions)-1; i++ {
		e.expressions[i].toSQL(builder, params)
		builder.WriteString(", ")
	}
	e.expressions[len(e.expressions)-1].toSQL(builder, params)
	builder.WriteString(")")
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

func (e *KeywordExpression) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString(e.node.Op + " ")
	for i := 0; i < len(e.node.Args)-1; i++ {
		e.node.Args[i].toSQL(builder, params)
		builder.WriteString(", ")
	}
	e.node.Args[len(e.node.Args)-1].toSQL(builder, params)
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
	return e.stringType.sqlType.expression.(*ColumnNode).Column.Column
}

type IntColumnExpression struct {
	*intType
}

func NewIntColumnExpression(table, column string) *IntColumnExpression {
	return &IntColumnExpression{
		Int(NewColumnNode(table, column)),
	}
}

func (e *IntColumnExpression) Name() string {
	return e.intType.sqlType.expression.(*ColumnNode).Column.Column
}

type FloatColumnExpression struct {
	*floatType
}

func NewFloatColumnExpression(table, column string) *FloatColumnExpression {
	return &FloatColumnExpression{
		Float(NewColumnNode(table, column)),
	}
}

func (e *FloatColumnExpression) Name() string {
	return e.floatType.sqlType.expression.(*ColumnNode).Column.Column
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
	return e.boolType.sqlType.expression.(*ColumnNode).Column.Column
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
	return e.uuidType.sqlType.expression.(*ColumnNode).Column.Column
}

type (
	expression = Expression
)

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	Table string      // Base table Name
	joins []*JoinExpr // Optional joins
}

func (s *TableSource) isTableExpression() {}

func (s *TableSource) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString(" " + s.Table)
	for _, join := range s.joins {
		join.toSQL(builder, params)
	}
}

func (s *TableSource) Name() string {
	return s.Table
}

func (s *TableSource) Join(jointype JoinType, table string, on OfType[bool]) *TableSource {
	s.joins = append(s.joins, &JoinExpr{
		Type:      jointype,
		Right:     &TableSource{Table: table},
		Condition: on,
	})
	return s
}

// JoinExpr represents a JOIN operation.
type JoinExpr struct {
	Type      JoinType     // INNER, LEFT, RIGHT, FULL
	Right     *TableSource // The table being joined
	Condition OfType[bool] // ON condition
}

func (j *JoinExpr) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString(string(j.Type) + " JOIN " + j.Right.Table + " ON ")
	j.Condition.toSQL(builder, params)
}

// JoinType enumerates join types.
type JoinType string

const (
	JoinInner JoinType = " INNER"
	JoinLeft  JoinType = " LEFT"
	JoinRight JoinType = " RIGHT"
	JoinFull  JoinType = " FULL"
)
