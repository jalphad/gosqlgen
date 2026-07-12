package ast

import (
	"fmt"
	"strings"
)

type SqlStatement interface {
	Expression
	Returns() []NamedExpression
	GetQueryContext() *QueryContext
}

// SelectStatement represents a full SELECT query AST.
type SelectStatement struct {
	With       []*CTE            // Common Table Expressions
	SelectList []NamedExpression // Columns or expressions in SELECT
	From       TableExpression   // Tables and joins
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
	return joinedTableNames(s.From)
}

func (s *SelectStatement) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	s.context = ctx
	if ctx != nil {
		ctx.Type = QueryTypeSelect
	}

	renderWithClause(s.With, builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	if len(s.With) > 0 {
		builder.WriteString(" ")
	}

	builder.WriteString("SELECT ")
	if ctx != nil {
		ctx.CurrentPart = QueryPartSelectList
	}
	if len(s.SelectList) > 0 {
		for i := 0; i < len(s.SelectList)-1; i++ {
			if s.SelectList[i] == nil {
				continue
			}
			s.SelectList[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.SelectList[len(s.SelectList)-1].toSQL(builder, params, ctx)
	} else {
		builder.WriteString("*")
	}

	if ctx != nil {
		ctx.CurrentPart = QueryPartFrom
	}
	builder.WriteString(" FROM ")
	renderTableExpression(s.From, builder, params, ctx)

	if s.Where != nil {
		if ctx != nil {
			ctx.CurrentPart = QueryPartWhere
		}
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params, ctx)
	}

	if len(s.GroupBy) > 0 {
		if ctx != nil {
			ctx.CurrentPart = QueryPartGroupBy
		}
		builder.WriteString(" GROUP BY ")
		for i := 0; i < len(s.GroupBy)-1; i++ {
			s.GroupBy[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.GroupBy[len(s.GroupBy)-1].toSQL(builder, params, ctx)
	}

	if s.Having != nil {
		if ctx != nil {
			ctx.CurrentPart = QueryPartHaving
		}
		builder.WriteString(" HAVING ")
		s.Having.toSQL(builder, params, ctx)
	}

	if len(s.OrderBy) > 0 {
		if ctx != nil {
			ctx.CurrentPart = QueryPartOrderBy
		}
		builder.WriteString(" ORDER BY ")
		for i := 0; i < len(s.OrderBy)-1; i++ {
			renderOrderByItem(s.OrderBy[i], builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(", ")
		}
		renderOrderByItem(s.OrderBy[len(s.OrderBy)-1], builder, params, ctx)
	}

	if s.Limit != nil {
		if ctx != nil {
			ctx.CurrentPart = QueryPartLimit
		}
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

func renderOrderByItem(item *OrderByItem, builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if item == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("nil ORDER BY item")
		}
		return
	}
	if item.Field == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("ORDER BY item requires a field")
		}
		return
	}
	item.Field.toSQL(builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(" " + string(item.Direction))
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
	Alias *Alias
	Query SqlStatement
}

func NewCTE(alias *Alias, query SqlStatement) *CTE {
	return &CTE{
		Alias: alias,
		Query: query,
	}
}

func renderWithClause(ctes []*CTE, builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if len(ctes) == 0 {
		return
	}
	builder.WriteString("WITH ")
	for i := 0; i < len(ctes); i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		ctes[i].toSQL(builder, params, ctx)
		if ctx != nil && ctx.Error != nil {
			return
		}
	}
}

func (c *CTE) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if c == nil {
		if ctx != nil {
			ctx.Error = fmt.Errorf("nil CTE")
		}
		return
	}
	if c.Alias == nil {
		if ctx != nil {
			ctx.Error = fmt.Errorf("CTE requires an alias")
		}
		return
	}
	if c.Query == nil {
		if ctx != nil {
			ctx.Error = fmt.Errorf("CTE %q requires a query", c.Alias.Name())
		}
		return
	}

	c.Alias.toSQL(builder, params, ctx)
	builder.WriteString(" AS (")
	savedType := QueryType("")
	savedPart := QueryPart("")
	if ctx != nil {
		savedType = ctx.Type
		savedPart = ctx.CurrentPart
	}
	c.Query.toSQL(builder, params, ctx)
	if ctx != nil {
		ctx.Type = savedType
		ctx.CurrentPart = savedPart
	}
	builder.WriteString(")")
}
