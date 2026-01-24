package ast

import (
	"fmt"
	"strings"
)

type SqlStatement interface {
	Returns() []NamedExpression
}

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

func (s *SelectStatement) Returns() []NamedExpression {
	return s.SelectList
}

func (s *SelectStatement) GetJoinedTables() []string {
	ret := make([]string, 0, len(s.From.joins))
	for _, join := range s.From.joins {
		ret = append(ret, join.Right.Table)
	}
	return ret
}

func (s *SelectStatement) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString("SELECT ")
	if len(s.SelectList) > 0 {
		for i := 0; i < len(s.SelectList)-1; i++ {
			s.SelectList[i].toSQL(builder, params)
			builder.WriteString(", ")
		}
		s.SelectList[len(s.SelectList)-1].toSQL(builder, params)
	} else {
		builder.WriteString("*")
	}

	builder.WriteString(" FROM")
	s.From.toSQL(builder, params)

	if s.Where != nil {
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params)
	}

	if len(s.GroupBy) > 0 {
		builder.WriteString(" GROUP BY ")
		for i := 0; i < len(s.GroupBy)-1; i++ {
			s.GroupBy[i].toSQL(builder, params)
			builder.WriteString(", ")
		}
		s.GroupBy[len(s.GroupBy)-1].toSQL(builder, params)
	}

	if s.Having != nil {
		builder.WriteString(" HAVING ")
		s.Having.toSQL(builder, params)
	}

	if len(s.OrderBy) > 0 {
		builder.WriteString(" ORDER BY ")
		for i := 0; i < len(s.OrderBy)-1; i++ {
			s.OrderBy[i].Field.toSQL(builder, params)
			builder.WriteString(" " + string(s.OrderBy[i].Direction))
			builder.WriteString(", ")
		}
		s.OrderBy[len(s.OrderBy)-1].Field.toSQL(builder, params)
		builder.WriteString(" " + string(s.OrderBy[len(s.OrderBy)-1].Direction))
	}

	if s.Limit != nil {
		if s.Limit.Limit > 0 {
			builder.WriteString(fmt.Sprintf(" LIMIT %d", s.Limit.Limit))
		}
		if s.Limit.Offset > 0 {
			builder.WriteString(fmt.Sprintf(" OFFSET %d", s.Limit.Offset))
		}
	}
}

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
