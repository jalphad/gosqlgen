package ast

import (
	"strings"
)

type UpdateStatement struct {
	Table     string
	SetList   []*SetKV
	From      *TableSource
	Where     OfType[bool]
	Returning []NamedExpression
}

func (s *UpdateStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *UpdateStatement) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString("UPDATE " + s.Table + " SET ")

	for i := 0; i < len(s.SetList)-1; i++ {
		builder.WriteString(s.SetList[i].Key.Name() + " = ")
		s.SetList[i].Value.toSQL(builder, params)
		builder.WriteString(", ")
	}
	builder.WriteString(s.SetList[len(s.SetList)-1].Key.Name() + " = ")
	s.SetList[len(s.SetList)-1].Value.toSQL(builder, params)

	if s.From != nil {
		builder.WriteString(" FROM " + s.From.Table)
		for _, join := range s.From.joins {
			join.toSQL(builder, params)
		}
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

type SetKV struct {
	Key   NamedExpression
	Value Expression
}
