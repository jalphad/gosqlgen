package models

import (
	"fmt"
)

// OrderDirection represents sort order
type OrderDirection string

const (
	ASC  OrderDirection = "ASC"
	DESC OrderDirection = "DESC"
)

// ComparisonOp represents comparison operators
type ComparisonOp string

const (
	EQ   ComparisonOp = "="
	NEQ  ComparisonOp = "!="
	GT   ComparisonOp = ">"
	GTE  ComparisonOp = ">="
	LT   ComparisonOp = "<"
	LTE  ComparisonOp = "<="
	LIKE ComparisonOp = "LIKE"
	IN   ComparisonOp = "IN"
)

// LogicalOp represents logical operators
type LogicalOp string

const (
	AND LogicalOp = "AND"
	OR  LogicalOp = "OR"
)

// FieldRef represents a type-safe field reference
type FieldRef struct {
	Table  string
	Column string
	GoType string
	Alias  string
}

// As creates an aliased field reference
func (f *FieldRef) As(alias string) *FieldRef {
	return &FieldRef{
		Table:  f.Table,
		Column: f.Column,
		GoType: f.GoType,
		Alias:  alias,
	}
}

// String returns the SQL representation
func (f *FieldRef) String() string {
	if f.Alias != "" {
		if f.Table == "" {
			return fmt.Sprintf("%s AS %s", f.Column, f.Alias)
		}
		return fmt.Sprintf("%s.%s AS %s", f.Table, f.Column, f.Alias)
	}
	if f.Table == "" {
		return f.Column
	}
	return fmt.Sprintf("%s.%s", f.Table, f.Column)
}

// Condition represents a WHERE condition
type Condition struct {
	Field    *FieldRef
	Op       ComparisonOp
	Value    interface{}
	Logical  LogicalOp
	SubConds []Condition
}

// JoinClause represents a JOIN operation
type JoinClause struct {
	Type       string
	Table      string
	LeftField  *FieldRef
	RightField *FieldRef
}

// OrderClause represents an ORDER BY clause
type OrderClause struct {
	Field     *FieldRef
	Direction OrderDirection
}
