package ast

import (
	"fmt"
	"strings"
)

type UpdateStatement struct {
	Table   string
	SetList []*SetKV
	Where   OfType[bool]
}

func (s *UpdateStatement) getNode() ExpressionNode {
	panic("not implemented")
}

func (s *UpdateStatement) toSQL(params *[]any) string {
	var query strings.Builder

	query.WriteString("UPDATE users SET ")

	setClauses := make([]string, 0, len(s.SetList))
	for _, set := range s.SetList {
		*params = append(*params, set.Value)
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", set.Key.Name(), len(*params)))
	}
	query.WriteString(strings.Join(setClauses, ", "))

	// Add WHERE conditions
	if s.Where != nil {
		query.WriteString(" WHERE ")
		query.WriteString(s.Where.toSQL(params))
	}

	return query.String()
}

type SetKV struct {
	Key   NamedExpression
	Value any
}
