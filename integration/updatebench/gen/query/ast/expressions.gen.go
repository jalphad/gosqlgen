package ast

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
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
	name string
}

func NewStringColumnExpression(table, column string) *StringColumnExpression {
	return &StringColumnExpression{
		stringType: String(NewColumnNode(table, column)),
		name:       column,
	}
}

func NewStringColumnExpressionFromExpr(name string, expression Expression) *StringColumnExpression {
	return &StringColumnExpression{
		stringType: String(expression),
		name:       name,
	}
}

func (e *StringColumnExpression) Name() string {
	return e.name
}

type IntColumnExpression struct {
	*intType
	name string
}

func NewIntColumnExpression(table, column string) *IntColumnExpression {
	return &IntColumnExpression{
		intType: Int(NewColumnNode(table, column)),
		name:    column,
	}
}

func NewIntColumnExpressionFromExpr(name string, expression Expression) *IntColumnExpression {
	return &IntColumnExpression{
		intType: Int(expression),
		name:    name,
	}
}

func (e *IntColumnExpression) Name() string {
	return e.name
}

type FloatColumnExpression struct {
	*floatType
	name string
}

func NewFloatColumnExpression(table, column string) *FloatColumnExpression {
	return &FloatColumnExpression{
		floatType: Float(NewColumnNode(table, column)),
		name:      column,
	}
}

func NewFloatColumnExpressionFromExpr(name string, expression Expression) *FloatColumnExpression {
	return &FloatColumnExpression{
		floatType: Float(expression),
		name:      name,
	}
}

func (e *FloatColumnExpression) Name() string {
	return e.name
}

type NumericColumnExpression struct {
	*numericType
	name string
}

func NewNumericColumnExpression(table, column string) *NumericColumnExpression {
	return &NumericColumnExpression{
		numericType: &NumericType{sqlType[pgtype.Numeric]{NewColumnNode(table, column)}},
		name:        column,
	}
}

func NewNumericColumnExpressionFromExpr(name string, expression Expression) *NumericColumnExpression {
	return &NumericColumnExpression{
		numericType: &NumericType{sqlType[pgtype.Numeric]{expression}},
		name:        name,
	}
}

func (e *NumericColumnExpression) Name() string {
	return e.name
}

type BoolColumnExpression struct {
	*boolType
	name string
}

func NewBoolColumnExpression(table, column string) *BoolColumnExpression {
	return &BoolColumnExpression{
		boolType: Bool(NewColumnNode(table, column)),
		name:     column,
	}
}

func NewBoolColumnExpressionFromExpr(name string, expression Expression) *BoolColumnExpression {
	return &BoolColumnExpression{
		boolType: Bool(expression),
		name:     name,
	}
}

func (e *BoolColumnExpression) Name() string {
	return e.name
}

type UUIDColumnExpression struct {
	*uuidType
	name string
}

func NewUUIDColumnExpression(table, column string) *UUIDColumnExpression {
	return &UUIDColumnExpression{
		uuidType: UUID(NewColumnNode(table, column)),
		name:     column,
	}
}

func NewUUIDColumnExpressionFromExpr(name string, expression Expression) *UUIDColumnExpression {
	return &UUIDColumnExpression{
		uuidType: UUID(expression),
		name:     name,
	}
}

func (e *UUIDColumnExpression) Name() string {
	return e.name
}

type TimestampColumnExpression struct {
	*timestampType
	name string
}

func NewTimestampColumnExpression(table, column string) *TimestampColumnExpression {
	return &TimestampColumnExpression{
		timestampType: Timestamp(NewColumnNode(table, column)),
		name:          column,
	}
}

func NewTimestampColumnExpressionFromExpr(name string, expression Expression) *TimestampColumnExpression {
	return &TimestampColumnExpression{
		timestampType: Timestamp(expression),
		name:          name,
	}
}

func (e *TimestampColumnExpression) Name() string {
	return e.name
}

type DateColumnExpression struct {
	*dateType
	name string
}

func NewDateColumnExpression(table, column string) *DateColumnExpression {
	return &DateColumnExpression{
		dateType: Date(NewColumnNode(table, column)),
		name:     column,
	}
}

func NewDateColumnExpressionFromExpr(name string, expression Expression) *DateColumnExpression {
	return &DateColumnExpression{
		dateType: Date(expression),
		name:     name,
	}
}

func (e *DateColumnExpression) Name() string {
	return e.name
}

type TimeColumnExpression struct {
	*timeType
	name string
}

func NewTimeColumnExpression(table, column string) *TimeColumnExpression {
	return &TimeColumnExpression{
		timeType: Time(NewColumnNode(table, column)),
		name:     column,
	}
}

func NewTimeColumnExpressionFromExpr(name string, expression Expression) *TimeColumnExpression {
	return &TimeColumnExpression{
		timeType: Time(expression),
		name:     name,
	}
}

func (e *TimeColumnExpression) Name() string {
	return e.name
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
