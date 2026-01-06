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

func (s *UpdateStatement) toSQL(params *[]any) string {
	var query strings.Builder

	query.WriteString("UPDATE " + s.Table + " SET ")

	setClauses := make([]string, 0, len(s.SetList))
	for _, set := range s.SetList {
		setClauses = append(setClauses, set.Key.Name()+" = "+Render(set.Value, params))
	}
	query.WriteString(strings.Join(setClauses, ", "))

	if s.From != nil {
		query.WriteString(" FROM " + s.From.Name)
		for _, join := range s.From.joins {
			query.WriteString(join.toSQL(params))
		}
	}

	// Add WHERE conditions
	if s.Where != nil {
		query.WriteString(" WHERE ")
		query.WriteString(s.Where.toSQL(params))
	}

	if len(s.Returning) > 0 {
		returning := make([]string, 0, len(s.Returning))
		for _, val := range s.Returning {
			returning = append(returning, val.Name())
		}
		query.WriteString(" RETURNING " + strings.Join(returning, ", "))
	}

	return query.String()
}

type SetKV struct {
	Key   NamedExpression
	Value Expression
}
