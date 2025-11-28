package ast

// SelectStatement represents a full SELECT query AST.
type SelectStatement struct {
	With       []*CTE            // Common Table Expressions
	SelectList []ValueExpression // Columns or expressions in SELECT
	From       []*TableSource    // Tables and joins
	Where      Expression        // WHERE clause
	GroupBy    []ValueExpression // GROUP BY fields
	Having     Expression        // HAVING clause
	OrderBy    []*OrderByItem    // ORDER BY items
	Limit      *LimitClause      // LIMIT/OFFSET
}

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	TableName string    // Base table name
	Alias     string    // Optional alias
	Join      *JoinExpr // Optional join ExpressionNode
}

// JoinExpr represents a JOIN operation.
type JoinExpr struct {
	Type      JoinType          // INNER, LEFT, RIGHT, FULL
	Right     *TableSource      // The table being joined
	Condition BooleanExpression // ON condition
}

// JoinType enumerates join types.
type JoinType string

const (
	JoinInner JoinType = "INNER"
	JoinLeft  JoinType = "LEFT"
	JoinRight JoinType = "RIGHT"
	JoinFull  JoinType = "FULL"
)

// OrderByItem represents an ORDER BY element.
type OrderByItem struct {
	Field     ValueExpression
	Direction SortDirection
}

// FieldRef represents a type-safe field reference or SQL ExpressionNode.
// Can be used in SELECT, WHERE, ORDER BY, etc.
type FieldRef struct {
	Table  string // Optional: table name or alias
	Column string // Column name or SQL ExpressionNode
}

// SortDirection enumerates sort directions.
type SortDirection string

const (
	Asc  SortDirection = "ASC"
	Desc SortDirection = "DESC"
)

// LimitClause represents LIMIT and OFFSET.
type LimitClause struct {
	Limit  int
	Offset int
}

// CTE represents a Common Table ExpressionNode (WITH clause).
type CTE struct {
	Name  string
	Query *SelectStatement
}
