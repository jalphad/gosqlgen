package ast

import (
	"fmt"
	"strings"
)

type InsertStatement struct {
	Table      string
	Into       []NamedExpression
	Values     []Expression
	OnConflict *Conflict
	Returning  []NamedExpression
}

func (s *InsertStatement) Returns() []NamedExpression {
	return s.Returning
}

func (s *InsertStatement) toSQL(params *[]any) string {
	var query strings.Builder
	s.toSQLBuilder(&query, params)

	return query.String()
}

func (s *InsertStatement) toSQLBuilder(query *strings.Builder, params *[]any) {
	query.WriteString("INSERT INTO " + s.Table)

	columns := make([]string, 0, len(s.Into))
	for _, column := range s.Into {
		columns = append(columns, column.Name())
	}
	query.WriteString("(" + strings.Join(columns, ", ") + ")")
	query.WriteString(" VALUES (")
	for i := 0; i < len(s.Values)-1; i++ {
		query.WriteString(Render(s.Values[i], params) + ", ")
	}
	query.WriteString(Render(s.Values[len(s.Values)-1], params) + ")")

	if s.OnConflict != nil {
		query.WriteString(Render(s.OnConflict, params))
	}

	if len(s.Returning) > 0 {
		returning := make([]string, 0, len(s.Returning))
		for _, val := range s.Returning {
			returning = append(returning, val.Name())
		}
		query.WriteString(" RETURNING " + strings.Join(returning, ", "))
	}
}

type Conflict struct {
	Columns []NamedExpression
	Action  OnConflictDoExpression
}

func (c *Conflict) toSQL(params *[]any) string {
	var query strings.Builder
	c.toSQLBuilder(&query, params)

	return query.String()
}

func (c *Conflict) toSQLBuilder(query *strings.Builder, params *[]any) {
	columns := make([]string, 0, len(c.Columns))
	for _, column := range c.Columns {
		columns = append(columns, column.Name())
	}
	query.WriteString(" ON CONFLICT (" + strings.Join(columns, ", ") + ") DO" + Render(c.Action, params))
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

func (a *ConflictAction) toSQL(params *[]any) string {
	var query strings.Builder
	a.toSQLBuilder(&query, params)
	return query.String()
}

func (a *ConflictAction) toSQLBuilder(query *strings.Builder, params *[]any) {
	query.WriteString(a.Do)
	if len(a.Set) > 0 {
		var elems []string
		for _, expr := range a.Set {
			elems = append(elems, fmt.Sprintf("%s = EXCLUDED.%s", expr.Name(), expr.Name()))
		}
		query.WriteString(" SET " + strings.Join(elems, ", "))
	}
	if a.WhereExpr != nil {
		query.WriteString(" WHERE " + Render(a.WhereExpr, params))
	}
}
