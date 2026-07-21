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
	ofType[T]
	NamedExpression
}

type AsExpression[T MappedTypes, C AsExprConstraint[T]] interface {
	Expression
	As(alias *ColumnAlias) C
}

type NamedExpression interface {
	Expression
	GetName() string
}

type NamedTableExpression interface {
	TableExpression
	NamedExpression
}

type TableExpression interface {
	Expression
	isTableExpression()
	Join(join *JoinExpr) TableExpression
}

type namedExpression struct {
	expression
	name string
}

func NewNamedExpression(name string, expression Expression) NamedExpression {
	return &namedExpression{
		expression: expression,
		name:       name,
	}
}

func (e namedExpression) GetName() string {
	return e.name
}

func NewNamedExpressionOfType[T MappedTypes](name string, expression OfType[T]) *NamedExpressionWrapper[T] {
	return &NamedExpressionWrapper[T]{
		ofType: expression,
		name:   name,
	}
}

type NamedExpressionWrapper[T MappedTypes] struct {
	ofType[T]
	name string
}

func (e *NamedExpressionWrapper[T]) GetName() string {
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

func (e *StringColumnExpression) GetName() string {
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

func (e *IntColumnExpression) GetName() string {
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

func (e *FloatColumnExpression) GetName() string {
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

func (e *NumericColumnExpression) GetName() string {
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

func (e *BoolColumnExpression) GetName() string {
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

func (e *UUIDColumnExpression) GetName() string {
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

func (e *TimestampColumnExpression) GetName() string {
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

func (e *DateColumnExpression) GetName() string {
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

func (e *TimeColumnExpression) GetName() string {
	return e.name
}

type JsonColumnExpression struct {
	*jsonType
	name string
}

func NewJsonColumnExpression(table, column string) *JsonColumnExpression {
	return &JsonColumnExpression{
		jsonType: Json(NewColumnNode(table, column)),
		name:     column,
	}
}

func NewJsonColumnExpressionFromExpr(name string, expression Expression) *JsonColumnExpression {
	return &JsonColumnExpression{
		jsonType: Json(expression),
		name:     name,
	}
}

func (e *JsonColumnExpression) GetName() string {
	return e.name
}

type BytesColumnExpression struct {
	*bytesType
	name string
}

func NewBytesColumnExpression(table, column string) *BytesColumnExpression {
	return &BytesColumnExpression{
		bytesType: Bytes(NewColumnNode(table, column)),
		name:      column,
	}
}

func NewBytesColumnExpressionFromExpr(name string, expression Expression) *BytesColumnExpression {
	return &BytesColumnExpression{
		bytesType: Bytes(expression),
		name:      name,
	}
}

func (e *BytesColumnExpression) GetName() string {
	return e.name
}

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	table string      // Base table Name
}

func NewTableSource(name string) *TableSource {
	return &TableSource{
		table: name,
	}
}

func (s *TableSource) As(alias string) *TableAlias {
	return &TableAlias{
		name:   alias,
		source: s,
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
			ctx.PrimaryTable = s.GetName()
		}
		// Add to AllTables if not already there
		if ctx.AllTables == nil {
			ctx.AllTables = []string{s.GetName()}
		} else if !slices.Contains(ctx.AllTables, s.GetName()) {
			ctx.AllTables = append(ctx.AllTables, s.GetName())
		}
	}

	builder.WriteString(s.table)
}

func (s *TableSource) GetName() string {
	return s.table
}

func (s *TableSource) Join(join *JoinExpr) TableExpression {
	return joinTableExpression(s, join)
}

func renderTableExpression(table TableExpression, builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if alias, ok := table.(*TableAlias); ok && alias.source == nil {
		builder.WriteString(alias.GetName())
		return
	}
	if alias, ok := table.(interface {
		GetName() string
		Columns() []NamedExpression
	}); ok && len(alias.Columns()) > 0 {
		builder.WriteString(alias.GetName())
		return
	}
	table.toSQL(builder, params, ctx)
}

func TableExpressionName(table TableExpression) string {
	switch table := table.(type) {
	case NamedTableExpression:
		return table.GetName()
	case *JoinedTableExpression:
		return TableExpressionName(table.Left)
	default:
		return ""
	}
}

type JoinedTableExpression struct {
	Left  TableExpression
	Joins []*JoinExpr
}

func joinTableExpression(left TableExpression, join *JoinExpr) TableExpression {
	if joined, ok := left.(*JoinedTableExpression); ok {
		joined.Joins = append(joined.Joins, join)
		return joined
	}
	return &JoinedTableExpression{
		Left:  left,
		Joins: []*JoinExpr{join},
	}
}

func (j *JoinedTableExpression) isTableExpression() {}

func (j *JoinedTableExpression) Join(join *JoinExpr) TableExpression {
	return joinTableExpression(j, join)
}

func (j *JoinedTableExpression) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if j == nil || j.Left == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("joined table expression requires a left table expression")
		}
		return
	}
	renderTableExpression(j.Left, builder, params, ctx)
	for _, join := range j.Joins {
		join.toSQL(builder, params, ctx)
		if ctx != nil && ctx.Error != nil {
			return
		}
	}
}

func joinedTableNames(table TableExpression) []string {
	joined, ok := table.(*JoinedTableExpression)
	if !ok {
		return nil
	}
	ret := make([]string, 0, len(joined.Joins))
	for _, join := range joined.Joins {
		if join != nil && join.right != nil {
			ret = append(ret, join.right.GetName())
		}
	}
	return ret
}

// JoinExpr represents a complete JOIN operation.
type JoinExpr struct {
	joinType  JoinType
	right     NamedTableExpression
	qualifier JoinQualifier
}

// JoinQualifier represents the qualification of a non-cross join.
// Values are created with On, Using, or Natural.
type JoinQualifier interface {
	isJoinQualifier()
	beforeJoin(builder *strings.Builder)
	afterJoin(builder *strings.Builder, params *[]any, ctx *QueryContext)
}

type onJoinQualifier struct {
	condition OfType[bool]
}

func (*onJoinQualifier) isJoinQualifier() {}

func (*onJoinQualifier) beforeJoin(_ *strings.Builder) {}

func (q *onJoinQualifier) afterJoin(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if q == nil || q.condition == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("ON join qualifier requires a condition")
		}
		return
	}
	builder.WriteString(" ON ")
	q.condition.toSQL(builder, params, ctx)
}

// On qualifies a join with a boolean expression.
func On(condition OfType[bool]) JoinQualifier {
	return &onJoinQualifier{condition: condition}
}

type usingJoinQualifier struct {
	columns []NamedExpression
}

func (*usingJoinQualifier) isJoinQualifier() {}

func (*usingJoinQualifier) beforeJoin(_ *strings.Builder) {}

func (q *usingJoinQualifier) afterJoin(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	if q == nil || len(q.columns) == 0 {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("USING join qualifier requires at least one column")
		}
		return
	}
	builder.WriteString(" USING (")
	for i, column := range q.columns {
		if column == nil {
			if ctx != nil && ctx.Error == nil {
				ctx.Error = fmt.Errorf("USING join qualifier contains a nil column")
			}
			return
		}
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(column.GetName())
	}
	builder.WriteString(")")
}

// Using qualifies a join with one or more shared column names.
func Using(columns ...NamedExpression) JoinQualifier {
	return &usingJoinQualifier{columns: columns}
}

type naturalJoinQualifier struct{}

func (*naturalJoinQualifier) isJoinQualifier() {}

func (*naturalJoinQualifier) beforeJoin(builder *strings.Builder) {
	builder.WriteString("NATURAL ")
}

func (*naturalJoinQualifier) afterJoin(_ *strings.Builder, _ *[]any, _ *QueryContext) {}

// Natural qualifies a join using all identically named columns.
func Natural() JoinQualifier {
	return &naturalJoinQualifier{}
}

func qualifiedJoin(joinType JoinType, right NamedTableExpression, qualifier JoinQualifier) *JoinExpr {
	return &JoinExpr{joinType: joinType, right: right, qualifier: qualifier}
}

func InnerJoin(right NamedTableExpression, qualifier JoinQualifier) *JoinExpr {
	return qualifiedJoin(JoinInner, right, qualifier)
}

func LeftJoin(right NamedTableExpression, qualifier JoinQualifier) *JoinExpr {
	return qualifiedJoin(JoinLeft, right, qualifier)
}

func RightJoin(right NamedTableExpression, qualifier JoinQualifier) *JoinExpr {
	return qualifiedJoin(JoinRight, right, qualifier)
}

func FullJoin(right NamedTableExpression, qualifier JoinQualifier) *JoinExpr {
	return qualifiedJoin(JoinFull, right, qualifier)
}

func CrossJoin(right NamedTableExpression) *JoinExpr {
	return &JoinExpr{joinType: JoinCross, right: right}
}

type lateralTableExpression struct {
	table NamedTableExpression
}

func (l *lateralTableExpression) isTableExpression() {}

func (l *lateralTableExpression) Join(join *JoinExpr) TableExpression {
	return joinTableExpression(l, join)
}

func (l *lateralTableExpression) GetName() string {
	if l == nil || l.table == nil {
		return ""
	}
	return l.table.GetName()
}

func (l *lateralTableExpression) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if l == nil || l.table == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("LATERAL requires a table expression")
		}
		return
	}
	builder.WriteString("LATERAL ")
	renderTableExpression(l.table, builder, params, ctx)
}

// Lateral marks a named FROM item as LATERAL.
func Lateral(table NamedTableExpression) NamedTableExpression {
	return &lateralTableExpression{table: table}
}

func (j *JoinExpr) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if j == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("join expression is nil")
		}
		return
	}
	if j.right == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("join expression requires a right table expression")
		}
		return
	}
	if j.joinType != JoinCross && j.qualifier == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("%s JOIN requires a qualifier", j.joinType)
		}
		return
	}
	// Track joined table in context
	if ctx != nil {
		if ctx.JoinedTables == nil {
			ctx.JoinedTables = make(map[string]JoinType)
		}
		ctx.JoinedTables[j.right.GetName()] = j.joinType
		if !slices.Contains(ctx.AllTables, j.right.GetName()) {
			ctx.AllTables = append(ctx.AllTables, j.right.GetName())
		}
		ctx.CurrentPart = QueryPartJoin
	}

	builder.WriteString(" ")
	if j.qualifier != nil {
		j.qualifier.beforeJoin(builder)
	}
	builder.WriteString(string(j.joinType) + " JOIN ")
	renderTableExpression(j.right, builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	if j.qualifier != nil {
		j.qualifier.afterJoin(builder, params, ctx)
	}
}

// JoinType enumerates join types.
type JoinType string

const (
	JoinInner JoinType = "INNER"
	JoinLeft  JoinType = "LEFT"
	JoinRight JoinType = "RIGHT"
	JoinFull  JoinType = "FULL"
	JoinCross JoinType = "CROSS"
)
