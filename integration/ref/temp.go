package models

import (
	"fmt"
)

// UsersClauses represents the users table columns
type UsersClauses struct {
	Id        *StringColumnClause[*UsersClauses]
	Username  *StringColumnClause[*UsersClauses]
	Email     *StringColumnClause[*UsersClauses]
	FullName  *StringColumnClause[*UsersClauses]
	CreatedAt *TimeColumnClause[*UsersClauses]
	UpdatedAt *TimeColumnClause[*UsersClauses]
	IsActive  *BoolColumnClause[*UsersClauses]
}

func (c *UsersClauses) GroupClause(expr Expression[*UsersClauses]) *ConditionBuilder[*UsersClauses] {
	var condition Expression[*UsersClauses]
	condition = &GroupedCondition[*UsersClauses]{
		inner: expr,
	}
	current := c.Id.builder
	if current.left != nil {
		current.right = condition
		cb := *current
		condition = &cb
	}

	builder := ConditionBuilder[*UsersClauses]{
		clauses: current.clauses,
		left:    condition,
	}
	*c.Id.builder = builder

	return c.Id.builder
}

func UsersClause() *UsersClauses {
	builder := &ConditionBuilder[*UsersClauses]{}
	clauses := &UsersClauses{
		Id:        &StringColumnClause[*UsersClauses]{name: "id", builder: builder},
		Username:  &StringColumnClause[*UsersClauses]{name: "username", builder: builder},
		Email:     &StringColumnClause[*UsersClauses]{name: "email", builder: builder},
		FullName:  &StringColumnClause[*UsersClauses]{name: "full_name", builder: builder},
		CreatedAt: &TimeColumnClause[*UsersClauses]{name: "created_at", builder: builder},
		UpdatedAt: &TimeColumnClause[*UsersClauses]{name: "updated_at", builder: builder},
		IsActive:  &BoolColumnClause[*UsersClauses]{name: "is_active", builder: builder},
	}
	builder.clauses = clauses

	return clauses
}

// Condition represents a WHERE condition
type Condition[C any] struct {
	expr string
	args []interface{}
}

func (c *Condition[C]) ToSQL() (string, []interface{}) {
	return c.expr, c.args
}

func (c *Condition[C]) getBuilder() *ConditionBuilder[C] {
	return nil
}

func (c *Condition[C]) copy() Expression[C] {
	cp := *c
	return &cp
}

type ConditionBuilder[C any] struct {
	clauses C
	left    Expression[C]
	right   Expression[C]
	op      LogicalOp
}

func (b *ConditionBuilder[C]) ToSQL() (string, []any) {
	if b.right != nil {
		leftSQL, leftArgs := b.left.ToSQL()
		rightSQL, rightArgs := b.right.ToSQL()
		expr := fmt.Sprintf("%s %s %s", leftSQL, b.op, rightSQL)
		args := append(leftArgs, rightArgs...)
		return expr, args
	}

	return b.left.ToSQL()
}

func (b *ConditionBuilder[C]) And() C {
	b.op = AND
	return b.clauses
}

func (b *ConditionBuilder[C]) Or() C {
	b.op = OR
	return b.clauses
}

func (b *ConditionBuilder[C]) getBuilder() *ConditionBuilder[C] {
	return b
}

func (b *ConditionBuilder[C]) copy() Expression[C] {
	cp := *b
	return &cp
}
