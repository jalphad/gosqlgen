package ast

import (
	"fmt"
	"strings"
)

type InsertStatement struct {
	With       []*CTE
	Table      string
	Into       []NamedExpression
	Values     *ValuesTable
	Select     *SelectStatement
	OnConflict *Conflict
	Returning  []NamedExpression
	context    *QueryContext
}

func (s *InsertStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *InsertStatement) GetQueryContext() *QueryContext {
	return s.context
}

func (s *InsertStatement) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	s.context = ctx
	if ctx != nil {
		ctx.Type = QueryTypeInsert
	}

	renderWithClause(s.With, builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	if len(s.With) > 0 {
		builder.WriteString(" ")
	}

	builder.WriteString("INSERT INTO " + s.Table)

	if ctx != nil {
		ctx.CurrentPart = QueryPartInto
	}
	columns := make([]string, 0, len(s.Into))
	for i, column := range s.Into {
		if column == nil {
			if ctx != nil && ctx.Error == nil {
				ctx.Error = fmt.Errorf("INSERT column %d is nil", i)
			}
			return
		}
		columns = append(columns, column.GetName())
	}
	builder.WriteString("(" + strings.Join(columns, ", ") + ")")

	switch {
	case s.Values != nil && s.Select != nil:
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("INSERT requires exactly one VALUES or SELECT source")
		}
		return
	case s.Values != nil:
		if ctx != nil {
			ctx.CurrentPart = QueryPartValues
		}
		builder.WriteString(" ")
		s.Values.render(builder, params, ctx, false)
	case s.Select != nil:
		if len(s.Select.SelectList) > 0 && len(s.Into) != len(s.Select.SelectList) {
			if ctx != nil && ctx.Error == nil {
				ctx.Error = fmt.Errorf("INSERT has %d target columns but SELECT returns %d columns", len(s.Into), len(s.Select.SelectList))
			}
			return
		}
		builder.WriteString(" ")
		savedType := QueryType("")
		savedPart := QueryPart("")
		if ctx != nil {
			savedType = ctx.Type
			savedPart = ctx.CurrentPart
		}
		s.Select.toSQL(builder, params, ctx)
		if ctx != nil {
			ctx.Type = savedType
			ctx.CurrentPart = savedPart
		}
	default:
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("INSERT requires a VALUES or SELECT source")
		}
		return
	}

	if s.OnConflict != nil {
		s.OnConflict.toSQL(builder, params, ctx)
	}

	if len(s.Returning) > 0 {
		if ctx != nil {
			ctx.CurrentPart = QueryPartReturning
		}
		builder.WriteString(" RETURNING ")
		for i := 0; i < len(s.Returning)-1; i++ {
			s.Returning[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.Returning[len(s.Returning)-1].toSQL(builder, params, ctx)
	}
}

type Conflict struct {
	Columns []NamedExpression
	Action  OnConflictDoExpression
}

func (c *Conflict) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	columns := make([]string, 0, len(c.Columns))
	for _, column := range c.Columns {
		columns = append(columns, column.GetName())
	}
	builder.WriteString(" ON CONFLICT (" + strings.Join(columns, ", ") + ") DO")
	if c.Action == nil {
		if ctx != nil {
			ctx.Error = fmt.Errorf("ON CONFLICT requires an action")
		}
		return
	}
	c.Action.toSQL(builder, params, ctx)
}

type OnConflictDoExpression interface {
	expression
	Where(expr OfType[bool]) OnConflictDoExpression
}

func Nothing() OnConflictDoExpression {
	return &ConflictAction{
		Do: " NOTHING",
	}
}

func Update(set ...UpdateSetExpr) OnConflictDoExpression {
	return &ConflictAction{
		Do:  " UPDATE",
		Set: set,
	}
}

type ConflictAction struct {
	Do        string
	Set       []UpdateSetExpr
	WhereExpr OfType[bool]
}

func (a *ConflictAction) Where(expr OfType[bool]) OnConflictDoExpression {
	a.WhereExpr = expr
	return a
}

func (a *ConflictAction) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(a.Do)
	if a.Do == " UPDATE" {
		if len(a.Set) == 0 {
			if ctx != nil {
				ctx.Error = fmt.Errorf("ON CONFLICT DO UPDATE requires at least one SET assignment")
			}
			return
		}
		builder.WriteString(" SET ")
		for i, set := range a.Set {
			if set == nil {
				if ctx != nil {
					ctx.Error = fmt.Errorf("ON CONFLICT DO UPDATE assignment %d is nil", i)
				}
				return
			}
			if i > 0 {
				builder.WriteString(", ")
			}
			set.toSQL(builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
		}
	}
	if a.WhereExpr != nil {
		builder.WriteString(" WHERE ")
		a.WhereExpr.toSQL(builder, params, ctx)
	}
}
