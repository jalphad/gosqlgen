package ast

import "strings"

type DeleteStatement struct {
	Table     string
	Using     []NamedExpression
	Where     OfType[bool]
	Returning []NamedExpression
}

func (s *DeleteStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *DeleteStatement) toSQL(builder *strings.Builder, params *[]any) {
	builder.WriteString("DELETE FROM " + s.Table)

	if len(s.Using) > 0 {
		builder.WriteString(" USING ")
		for i := 0; i < len(s.Using)-1; i++ {
			s.Using[i].toSQL(builder, params)
			builder.WriteString(", ")
		}
		s.Using[len(s.Using)-1].toSQL(builder, params)
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
