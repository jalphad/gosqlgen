package ast

import (
	"fmt"
	"strings"
)

// SelectStatement represents a full SELECT query AST.
type SelectStatement struct {
	With       []*CTE            // Common Table Expressions
	SelectList []NamedExpression // Columns or expressions in SELECT
	From       *TableSource      // Tables and joins
	Where      OfType[bool]      // WHERE clause
	GroupBy    []Expression      // GROUP BY fields
	Having     OfType[bool]      // HAVING clause
	OrderBy    []*OrderByItem    // ORDER BY items
	Limit      *LimitClause      // LIMIT/OFFSET
}

func (s *SelectStatement) getNode() ExpressionNode {
	panic("not implemented")
}

func (s *SelectStatement) toSQL(params *[]any) string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	for i, expr := range s.SelectList {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(Render(expr, params))
	}
	sb.WriteString(" FROM " + s.From.Name)
	if s.Where != nil {
		sb.WriteString(" WHERE " + Render(s.Where, params))
	}
	if s.GroupBy != nil {

	}
	if s.OrderBy != nil {
		sb.WriteString(" ORDER BY ")
		for i, expr := range s.OrderBy {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(Render(expr.Field, params) + " " + string(expr.Direction))
		}
	}
	if s.Having != nil {

	}
	if s.Limit != nil {
		if s.Limit.Limit > 0 {
			sb.WriteString(fmt.Sprintf(" LIMIT %d", s.Limit.Limit))
		}
		if s.Limit.Offset > 0 {
			sb.WriteString(fmt.Sprintf(" OFFSET %d", s.Limit.Offset))
		}
	}
	return sb.String()
}

// TableSource represents a table or a join in the FROM clause.
type TableSource struct {
	Name string      // Base table Name
	join []*JoinExpr // Optional join ExpressionNode
}

func (s *TableSource) Join(jointype JoinType, table string, on OfType[bool]) *TableSource {
	s.join = append(s.join, &JoinExpr{
		Type:      jointype,
		Right:     &TableSource{Name: table},
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
