package ast

import (
	"fmt"
	"strings"
)

type InsertStatement struct {
	With       []*CTE
	Table      string
	Into       []NamedExpression
	Values     []Expression
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
	for _, column := range s.Into {
		columns = append(columns, column.Name())
	}
	builder.WriteString("(" + strings.Join(columns, ", ") + ")")

	switch {
	case len(s.Values) > 0 && s.Select != nil:
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("INSERT requires exactly one VALUES or SELECT source")
		}
		return
	case len(s.Values) > 0:
		if ctx != nil {
			ctx.CurrentPart = QueryPartValues
		}
		builder.WriteString(" VALUES ")
		for i := 0; i < len(s.Values)-1; i++ {
			s.Values[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.Values[len(s.Values)-1].toSQL(builder, params, ctx)
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
		columns = append(columns, column.Name())
	}
	builder.WriteString(" ON CONFLICT (" + strings.Join(columns, ", ") + ") DO" + Render(c.Action, params))
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

func Update(set ...NamedExpression) OnConflictDoExpression {
	return &ConflictAction{
		Do:  " UPDATE",
		Set: set,
	}
}

type ConflictAction struct {
	Do        string
	Set       []NamedExpression
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
	if len(a.Set) > 0 {
		var elems []string
		for _, expr := range a.Set {
			elems = append(elems, fmt.Sprintf("%s = EXCLUDED.%s", expr.Name(), expr.Name()))
		}
		builder.WriteString(" SET " + strings.Join(elems, ", "))
	}
	if a.WhereExpr != nil {
		builder.WriteString(" WHERE " + Render(a.WhereExpr, params))
	}
}
