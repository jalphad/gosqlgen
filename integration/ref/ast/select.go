package ast

import (
	"fmt"
	"strings"
)

type SqlStatement interface {
	Returns() []NamedExpression
	GetQueryContext() *QueryContext
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
	context    *QueryContext
}

func (s *SelectStatement) Returns() []NamedExpression {
	return s.SelectList
}

func (s *SelectStatement) GetQueryContext() *QueryContext {
	return s.context
}

func (s *SelectStatement) GetJoinedTables() []string {
	ret := make([]string, 0, len(s.From.joins))
	for _, join := range s.From.joins {
		ret = append(ret, join.Right.Table)
	}
	return ret
}

func (s *SelectStatement) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	s.context = ctx
	builder.WriteString("SELECT ")
	if len(s.SelectList) > 0 {
		for i := 0; i < len(s.SelectList)-1; i++ {
			s.SelectList[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.SelectList[len(s.SelectList)-1].toSQL(builder, params, ctx)
	} else {
		builder.WriteString("*")
	}

	builder.WriteString(" FROM")
	s.From.toSQL(builder, params, ctx)

	if s.Where != nil {
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params, ctx)
	}

	if len(s.GroupBy) > 0 {
		builder.WriteString(" GROUP BY ")
		for i := 0; i < len(s.GroupBy)-1; i++ {
			s.GroupBy[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.GroupBy[len(s.GroupBy)-1].toSQL(builder, params, ctx)
	}

	if s.Having != nil {
		builder.WriteString(" HAVING ")
		s.Having.toSQL(builder, params, ctx)
	}

	if len(s.OrderBy) > 0 {
		builder.WriteString(" ORDER BY ")
		for i := 0; i < len(s.OrderBy)-1; i++ {
			s.OrderBy[i].Field.toSQL(builder, params, ctx)
			builder.WriteString(" " + string(s.OrderBy[i].Direction))
			builder.WriteString(", ")
		}
		s.OrderBy[len(s.OrderBy)-1].Field.toSQL(builder, params, ctx)
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
