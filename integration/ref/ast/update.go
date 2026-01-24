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
}

func (s *UpdateStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *UpdateStatement) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString("UPDATE " + s.Table + " SET ")

	for i := 0; i < len(s.SetList)-1; i++ {
		s.SetList[i].toSQL(builder, params)
		builder.WriteString(", ")
	}
	s.SetList[len(s.SetList)-1].toSQL(builder, params)

	if s.From != nil {
		builder.WriteString(" FROM ")
		s.From.toSQL(builder, params)
	}

	// Add WHERE conditions
	if s.Where != nil {
		builder.WriteString(" WHERE ")
		s.Where.toSQL(builder, params)
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

func (s UpdateSet) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString(s.Key.Name() + " = ")
	s.Value.toSQL(builder, params)
}

func (s UpdateSet) forUpdate() {}

type UpdateSetExpr interface {
	expression
	forUpdate()
}
