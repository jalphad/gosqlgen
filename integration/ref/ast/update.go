package ast

import (
	"strings"
)

type UpdateStatement struct {
	Table     string
	SetList   []UpdateSetExpr
	From      TableExpression
	Where     OfType[bool]
	Returning []NamedExpression
	context   *QueryContext
}

func (s *UpdateStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *UpdateStatement) GetQueryContext() *QueryContext {
	return s.context
}

func (s *UpdateStatement) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	s.context = ctx
	builder.WriteString("UPDATE " + s.Table + " SET ")

	for i := 0; i < len(s.SetList)-1; i++ {
		s.SetList[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	s.SetList[len(s.SetList)-1].toSQL(builder, params, ctx)

	if s.From != nil {
		builder.WriteString(" FROM ")
		s.From.toSQL(builder, params, ctx)
	}

	// Add WHERE conditions
	if s.Where != nil {
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params, ctx)
	}

	if len(s.Returning) > 0 {
		returning := make([]string, 0, len(s.Returning))
		for _, val := range s.Returning {
			returning = append(returning, val.Name())
		}
		builder.WriteString(" RETURNING " + strings.Join(returning, ", "))
	}
}

type UpdateSet struct {
	Key   NamedExpression
	Value Expression
}

func (s UpdateSet) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(s.Key.Name() + " = ")
	s.Value.toSQL(builder, params, ctx)
}

func (s UpdateSet) forUpdate() {}

type UpdateSetExpr interface {
	expression
	forUpdate()
}
