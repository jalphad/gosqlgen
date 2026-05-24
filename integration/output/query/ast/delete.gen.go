package ast

import "strings"

type DeleteStatement struct {
	Table     string
	Using     []NamedExpression
	Where     OfType[bool]
	Returning []NamedExpression
	context   *QueryContext
}

func (s *DeleteStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *DeleteStatement) GetQueryContext() *QueryContext {
	return s.context
}

func (s *DeleteStatement) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	s.context = ctx
	if ctx != nil {
		ctx.Type = QueryTypeDelete
	}

	builder.WriteString("DELETE FROM " + s.Table)

	if len(s.Using) > 0 {
		if ctx != nil {
			ctx.CurrentPart = QueryPartUsing
		}
		builder.WriteString(" USING ")
		for i := 0; i < len(s.Using)-1; i++ {
			s.Using[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.Using[len(s.Using)-1].toSQL(builder, params, ctx)
	}

	// Add WHERE conditions
	if s.Where != nil {
		if ctx != nil {
			ctx.CurrentPart = QueryPartWhere
		}
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params, ctx)
	}

	if len(s.Returning) > 0 {
		if ctx != nil {
			ctx.CurrentPart = QueryPartReturning
		}
		returning := make([]string, 0, len(s.Returning))
		for _, val := range s.Returning {
			returning = append(returning, val.Name())
		}
		builder.WriteString(" RETURNING " + strings.Join(returning, ", "))
	}
}
