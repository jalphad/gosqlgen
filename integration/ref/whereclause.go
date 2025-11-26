package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Expression represents a SQL expression
type Expression[C any] interface {
	ToSQL() (string, []interface{})
	getBuilder() *ConditionBuilder[C]
	copy() Expression[C]
}

type GroupedCondition[C any] struct {
	inner   Expression[C]
	builder *ConditionBuilder[C]
}

// GroupClause creates a grouped expression with parentheses
func GroupClause[C any](expr Expression[C]) *GroupedCondition[C] {
	builder := expr.getBuilder()
	return &GroupedCondition[C]{inner: expr, builder: builder}
}

func (g *GroupedCondition[C]) ToSQL() (string, []interface{}) {
	sql, args := g.inner.ToSQL()
	return fmt.Sprintf("(%s)", sql), args
}

func (g *GroupedCondition[C]) getBuilder() *ConditionBuilder[C] {
	return g.builder
}

func (g *GroupedCondition[C]) copy() Expression[C] {
	cp := *g
	return &cp
}

func (g *GroupedCondition[C]) And() C {
	condition := *g
	condition.inner = g.inner.copy()

	builder := ConditionBuilder[C]{
		clauses: g.builder.clauses,
		left:    &condition,
		op:      AND,
	}
	*g.builder = builder

	return builder.clauses
}

func (g *GroupedCondition[C]) Or() C {
	var condition Expression[C]
	condition = g
	if g.builder.left != nil {
		g.builder.right = g
		cb := *g.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: g.builder.clauses,
		left:    condition,
		op:      OR,
	}
	*g.builder = builder

	return g.builder.clauses
}

// UUIDColumnClause represents a UUID column
type UUIDColumnClause[C any] struct {
	name    string
	builder *ConditionBuilder[C]
}

func (c *UUIDColumnClause[C]) Eq(val uuid.UUID) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s = $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *UUIDColumnClause[C]) In(vals ...uuid.UUID) *ConditionBuilder[C] {
	placeholders := make([]string, len(vals))
	args := make([]interface{}, len(vals))
	for i, v := range vals {
		placeholders[i] = "$?"
		args[i] = v
	}
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s IN (%s)", c.name, strings.Join(placeholders, ", ")),
		args: args,
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *UUIDColumnClause[C]) IsNull() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s IS NULL", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *UUIDColumnClause[C]) IsNotNull() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s IS NOT NULL", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

// StringColumnClause represents a VARCHAR column
type StringColumnClause[C any] struct {
	ref     *FieldRef
	name    string
	builder *ConditionBuilder[C]
}

func (c *StringColumnClause[C]) Eq(val string) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s = $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *StringColumnClause[C]) Like(pattern string) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s LIKE $?", c.name),
		args: []interface{}{pattern},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *StringColumnClause[C]) In(vals ...string) *ConditionBuilder[C] {
	placeholders := make([]string, len(vals))
	args := make([]interface{}, len(vals))
	for i, v := range vals {
		placeholders[i] = "$?"
		args[i] = v
	}
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s IN (%s)", c.name, strings.Join(placeholders, ", ")),
		args: args,
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *StringColumnClause[C]) IsNull() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s IS NULL", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *StringColumnClause[C]) IsNotNull() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s IS NOT NULL", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

// StringFunction represents SQL string functions
type StringFunction[C any] struct {
	column *StringColumnClause[C]
}

func (c *StringColumnClause[C]) Fn() *StringFunction[C] {
	return &StringFunction[C]{column: c}
}

func (sf *StringFunction[C]) Length() *IntExpression[C] {
	return &IntExpression[C]{
		expr:    fmt.Sprintf("LENGTH(%s)", sf.column.name),
		builder: sf.column.builder,
	}
}

// IntExpression represents an integer expression
type IntExpression[C any] struct {
	expr    string
	builder *ConditionBuilder[C]
}

func (c *IntExpression[C]) Gt(val int) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s > $?", c.expr),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *IntExpression[C]) Lt(val int) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s < $?", c.expr),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *IntExpression[C]) Gte(val int) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s >= $?", c.expr),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *IntExpression[C]) Lte(val int) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s <= $?", c.expr),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *IntExpression[C]) Eq(val int) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s = $?", c.expr),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

// BoolColumnClause represents a BOOLEAN column
type BoolColumnClause[C any] struct {
	name    string
	builder *ConditionBuilder[C]
}

func (c *BoolColumnClause[C]) Eq(val bool) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s = $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *BoolColumnClause[C]) IsTrue() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s = true", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *BoolColumnClause[C]) IsFalse() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s = false", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *BoolColumnClause[C]) IsNull() *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{expr: fmt.Sprintf("%s IS NULL", c.name), args: nil}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

// TimeColumnClause represents a TIMESTAMP column
type TimeColumnClause[C any] struct {
	name    string
	builder *ConditionBuilder[C]
}

func (c *TimeColumnClause[C]) Eq(val time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s = $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *TimeColumnClause[C]) Gt(val time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s > $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *TimeColumnClause[C]) Lt(val time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s < $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *TimeColumnClause[C]) Gte(val time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s >= $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *TimeColumnClause[C]) Lte(val time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s <= $?", c.name),
		args: []interface{}{val},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}

func (c *TimeColumnClause[C]) Between(start, end time.Time) *ConditionBuilder[C] {
	var condition Expression[C]
	condition = &Condition[C]{
		expr: fmt.Sprintf("%s BETWEEN $? AND $?", c.name),
		args: []interface{}{start, end},
	}
	if c.builder.left != nil {
		c.builder.right = condition
		cb := *c.builder
		condition = &cb
	}

	builder := ConditionBuilder[C]{
		clauses: c.builder.clauses,
		left:    condition,
	}
	*c.builder = builder

	return c.builder
}
