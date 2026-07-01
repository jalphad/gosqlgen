package ast

import (
	"fmt"
	"slices"
	"strings"
)

func Render(e Expression, params *[]any) string {
	query, _ := RenderWithContext(e, params, &QueryContext{})
	return query
}

func RenderWithContext(e Expression, params *[]any, ctx *QueryContext) (string, error) {
	var query strings.Builder
	e.toSQL(&query, params, ctx)

	if ctx != nil && ctx.Error != nil {
		return query.String(), ctx.Error
	}
	return query.String(), nil
}

func BuildQuery(e Expression, builder *strings.Builder, params *[]any) {
	e.toSQL(builder, params, &QueryContext{})
}

func BuildQueryWithContext(e Expression, builder *strings.Builder, params *[]any, ctx *QueryContext) {
	e.toSQL(builder, params, ctx)
}

// ErrorExpression represents an expression that sets an error in the context during rendering
type ErrorExpression struct {
	err error
}

func NewErrorExpression(err error) *ErrorExpression {
	return &ErrorExpression{err: err}
}

func (e *ErrorExpression) toSQL(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error == nil {
		ctx.Error = e.err
	}
}

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext)
}
type expression = Expression

type AsExprConstraint[T MappedTypes] interface {
	AliasedFunction[T]
}

type AsExpression[T MappedTypes, C AsExprConstraint[T]] interface {
	Expression
	As(alias *Alias) C
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

func (e *GroupedExpression) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if len(e.expressions) == 0 {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("GroupedExpression has no expressions")
		}
		return
	}
	builder.WriteString("(")
	for i := 0; i < len(e.expressions)-1; i++ {
		e.expressions[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	e.expressions[len(e.expressions)-1].toSQL(builder, params, ctx)
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

func (e *KeywordExpression) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if len(e.node.Args) == 0 {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("KeywordExpression %q has no arguments", e.node.Op)
		}
		return
	}
	builder.WriteString(e.node.Op + " ")
	for i := 0; i < len(e.node.Args)-1; i++ {
		e.node.Args[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	e.node.Args[len(e.node.Args)-1].toSQL(builder, params, ctx)
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
	if colNode, ok := e.stringType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
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
	if colNode, ok := e.intType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
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
	if colNode, ok := e.floatType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
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
	if colNode, ok := e.boolType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
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
	if colNode, ok := e.uuidType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
}

type TimestampColumnExpression struct {
	*TimestampType
}

func NewTimestampColumnExpression(table, column string) *TimestampColumnExpression {
	return &TimestampColumnExpression{
		Timestamp(NewColumnNode(table, column)),
	}
}

func (e *TimestampColumnExpression) Name() string {
	if colNode, ok := e.TimestampType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
}

type DateColumnExpression struct {
	*dateType
}

func NewDateColumnExpression(table, column string) *DateColumnExpression {
	return &DateColumnExpression{
		Date(NewColumnNode(table, column)),
	}
}

func (e *DateColumnExpression) Name() string {
	if colNode, ok := e.dateType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
}

type TimeColumnExpression struct {
	*timeType
}

func NewTimeColumnExpression(table, column string) *TimeColumnExpression {
	return &TimeColumnExpression{
		Time(NewColumnNode(table, column)),
	}
}

func (e *TimeColumnExpression) Name() string {
	if colNode, ok := e.timeType.sqlType.expression.(*ColumnNode); ok {
		return colNode.Column.Column
	}
	return ""
}

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	table string      // Base table Name
	joins []*JoinExpr // Optional joins
}

func NewTableSource(name string) *TableSource {
	return &TableSource{
		table: name,
	}
}

func (s *TableSource) isTableExpression() {}

func (s *TableSource) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	// Track primary table
	if ctx != nil {
		if ctx.PrimaryTable == "" {
			ctx.PrimaryTable = s.table
		}
		// Add to AllTables if not already there
		if ctx.AllTables == nil {
			ctx.AllTables = []string{s.table}
		} else if !slices.Contains(ctx.AllTables, s.table) {
			ctx.AllTables = append(ctx.AllTables, s.table)
		}
	}

	builder.WriteString(s.table)
	for _, join := range s.joins {
		join.toSQL(builder, params, ctx)
	}
}

func (s *TableSource) Name() string {
	return s.table
}

func (s *TableSource) Join(jointype JoinType, table NamedTableExpression, on OfType[bool]) *TableSource {
	s.joins = append(s.joins, &JoinExpr{
		Type:      jointype,
		Right:     table,
		Condition: on,
	})
	return s
}

// JoinExpr represents a JOIN operation.
type JoinExpr struct {
	Type      JoinType             // INNER, LEFT, RIGHT, FULL
	Right     NamedTableExpression // The table being joined
	Condition OfType[bool]         // ON condition
}

func (j *JoinExpr) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	// Track joined table in context
	if ctx != nil {
		if ctx.JoinedTables == nil {
			ctx.JoinedTables = make(map[string]JoinType)
		}
		ctx.JoinedTables[j.Right.Name()] = j.Type
		if !slices.Contains(ctx.AllTables, j.Right.Name()) {
			ctx.AllTables = append(ctx.AllTables, j.Right.Name())
		}
		ctx.CurrentPart = QueryPartJoin
	}

	builder.WriteString(string(j.Type) + " JOIN ")
	j.Right.toSQL(builder, params, ctx)
	builder.WriteString(" ON ")
	j.Condition.toSQL(builder, params, ctx)
}

// JoinType enumerates join types.
type JoinType string

const (
	JoinInner   JoinType = " INNER"
	JoinLeft    JoinType = " LEFT"
	JoinRight   JoinType = " RIGHT"
	JoinFull    JoinType = " FULL"
	JoinLateral JoinType = " LATERAL"
	JoinCross   JoinType = " CROSS"
)
