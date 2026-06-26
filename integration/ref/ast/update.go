package ast

import (
	"fmt"
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
	if ctx != nil {
		ctx.Type = QueryTypeUpdate
	}

	builder.WriteString("UPDATE " + s.Table + " SET ")

	if ctx != nil {
		ctx.CurrentPart = QueryPartSet
	}
	for i := 0; i < len(s.SetList)-1; i++ {
		s.SetList[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	s.SetList[len(s.SetList)-1].toSQL(builder, params, ctx)

	if s.From != nil {
		if ctx != nil {
			ctx.CurrentPart = QueryPartFrom
		}
		builder.WriteString(" FROM ")
		s.From.toSQL(builder, params, ctx)
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
		builder.WriteString(" RETURNING ")
		for i := 0; i < len(s.Returning)-1; i++ {
			s.Returning[i].toSQL(builder, params, ctx)
			builder.WriteString(", ")
		}
		s.Returning[len(s.Returning)-1].toSQL(builder, params, ctx)
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

type UpdateSetList struct {
	Sets []UpdateSet
	Err  error
}

func NewUpdateSetError(err error) UpdateSetList {
	return UpdateSetList{Err: err}
}

func (s UpdateSetList) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if s.Err != nil {
		if ctx != nil {
			ctx.Error = s.Err
		}
		return
	}
	if len(s.Sets) == 0 {
		if ctx != nil {
			ctx.Error = fmt.Errorf("UpdateSetList has no sets")
		}
		return
	}
	for i := 0; i < len(s.Sets)-1; i++ {
		s.Sets[i].toSQL(builder, params, ctx)
		builder.WriteString(", ")
	}
	s.Sets[len(s.Sets)-1].toSQL(builder, params, ctx)
}

func (s UpdateSetList) forUpdate() {}

type UpdateSetExpr interface {
	expression
	forUpdate()
}
