package ast

// SelectStatement represents a full SELECT query AST.
type SelectStatement struct {
	With       []*CTE            // Common Table Expressions
	SelectList []NamedExpression // Columns or expressions in SELECT
	From       []*TableSource    // Tables and joins
	Where      *BoolType         // WHERE clause
	GroupBy    []Expression      // GROUP BY fields
	Having     *BoolType         // HAVING clause
	OrderBy    []*OrderByItem    // ORDER BY items
	Limit      *LimitClause      // LIMIT/OFFSET
}

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	TableName string    // Base table Name
	Alias     string    // Optional alias
	Join      *JoinExpr // Optional join ExpressionNode
}

// JoinExpr represents a JOIN operation.
type JoinExpr struct {
	Type      JoinType     // INNER, LEFT, RIGHT, FULL
	Right     *TableSource // The table being joined
	Condition *BoolType    // ON condition
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
	Field     Expression
	Direction SortDirection
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
