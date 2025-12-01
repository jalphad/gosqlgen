package ast

func Render(expression Expression, params *[]any) string {
	return expression.toSQL(params)
}

// SelectExpr can be used in SELECT clause (almost any expression)
type SelectExpr interface {
	Expression
	selectExprMarker()
}

// WhereExpr can be used in WHERE clause (must return Bool)
type WhereExpr interface {
	Expression
	whereExprMarker()
}

// GroupByExpr can be used in GROUP BY clause (non-aggregate expressions)
type GroupByExpr interface {
	Expression
	groupByExprMarker()
}

// HavingExpr can be used in HAVING clause (must return Bool, can use aggregates)
type HavingExpr interface {
	Expression
	havingExprMarker()
}

// OrderByExpr can be used in ORDER BY clause (any expression)
type OrderByExpr interface {
	Expression
	orderByExprMarker()
}

type LogicalExpr struct {
	BinaryExpression
}

func (e LogicalExpr) selectExprMarker()  {}
func (e LogicalExpr) whereExprMarker()   {}
func (e LogicalExpr) havingExprMarker()  {}
func (e LogicalExpr) orderByExprMarker() {}

func NewComparisonExpression(e *BinaryExpression) *ComparisonExpr {
	return &ComparisonExpr{e}
}

type ComparisonExpr struct {
	*BinaryExpression
}

func (e *ComparisonExpr) selectExprMarker()  {}
func (e *ComparisonExpr) whereExprMarker()   {}
func (e *ComparisonExpr) havingExprMarker()  {}
func (e *ComparisonExpr) orderByExprMarker() {}

func NewArithmeticExpr(e *BinaryExpression) *ArithmeticExpr {
	return &ArithmeticExpr{e}
}

type ArithmeticExpr struct {
	*BinaryExpression
}

func (e *ArithmeticExpr) selectExprMarker()  {}
func (e *ArithmeticExpr) whereExprMarker()   {}
func (e *ArithmeticExpr) groupByExprMarker() {}
func (e *ArithmeticExpr) havingExprMarker()  {}
func (e *ArithmeticExpr) orderByExprMarker() {}

func NewAggregateExpr(e *FunctionExpression) *AggregateExpr {
	return &AggregateExpr{e}
}

type AggregateExpr struct {
	*FunctionExpression
}

func (e *AggregateExpr) selectExprMarker()  {}
func (e *AggregateExpr) havingExprMarker()  {}
func (e *AggregateExpr) orderByExprMarker() {}
