package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type (
	expression      ast.Expression
	valueExpression ast.ValueExpression
)

func NewStringType[B ExpressionBuilder](ref *ast.FieldRef, builder B) *StringType[B] {
	return &StringType[B]{
		valueExpression: ast.NewValueExpression(ast.NewFieldExpresion(ref)),
		b:               builder,
	}
}

// StringType represents a VARCHAR column
type StringType[B ExpressionBuilder] struct {
	valueExpression
	b B
}

func (c *StringType[B]) Eq(val string) ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewBinaryExpression(&ast.ExpressionNode{Op: "=", Args: []ast.Expression{left, Lit(val)}}))
	return c.b
}

func (c *StringType[B]) Like(pattern string) ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewBinaryExpression(&ast.ExpressionNode{Op: "LIKE", Args: []ast.Expression{left, Lit(pattern)}}))
	return c.b
}

func (c *StringType[B]) In(vals ...string) ast.BooleanExpression {

	return c.b
}

func (c *StringType[B]) IsNull() ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewUnaryExpression(&ast.ExpressionNode{Op: "IS NULL", Args: []ast.Expression{left}}))
	return c.b
}

func (c *StringType[B]) IsNotNull() ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewUnaryExpression(&ast.ExpressionNode{Op: "IS NOT NULL", Args: []ast.Expression{left}}))
	return c.b
}

func NewIntType[B ExpressionBuilder](ref *ast.FieldRef, builder B) *IntType[B] {
	return &IntType[B]{
		valueExpression: ast.NewValueExpression(ast.NewFieldExpresion(ref)),
		b:               builder,
	}
}

// IntType represents an integer expression
type IntType[B ExpressionBuilder] struct {
	valueExpression
	b B
}

func (c *IntType[B]) Gt(val int) ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewBinaryExpression(&ast.ExpressionNode{Op: ">", Args: []ast.Expression{left, Lit(val)}}))
	return c.b
}

func (c *IntType[B]) Lt(val int) ast.BooleanExpression {

	return c.b
}

func (c *IntType[B]) Gte(val int) ast.BooleanExpression {

	return c.b
}

func (c *IntType[B]) Lte(val int) ast.BooleanExpression {

	return c.b
}

func (c *IntType[B]) Eq(val int) ast.BooleanExpression {

	return c.b
}

func NewBoolType[B ExpressionBuilder](ref *ast.FieldRef, builder B) *BoolType[B] {
	return &BoolType[B]{
		valueExpression: ast.NewValueExpression(ast.NewFieldExpresion(ref)),
		b:               builder,
	}
}

// BoolType represents a BOOLEAN column
type BoolType[B ExpressionBuilder] struct {
	valueExpression
	b B
}

func (c *BoolType[B]) Eq(val bool) ast.BooleanExpression {

	return c.b
}

func (c *BoolType[B]) IsTrue() ast.BooleanExpression {
	left := c.valueExpression
	c.b.setExpression(ast.NewBinaryExpression(&ast.ExpressionNode{Op: "=", Args: []ast.Expression{left, Lit(true)}}))
	return c.b
}

func (c *BoolType[B]) IsFalse() ast.BooleanExpression {

	return c.b
}

func (c *BoolType[B]) IsNull() ast.BooleanExpression {

	return c.b
}

func NewTimeType[B ExpressionBuilder](ref *ast.FieldRef, builder B) *TimeType[B] {
	return &TimeType[B]{
		valueExpression: ast.NewValueExpression(ast.NewFieldExpresion(ref)),
		b:               builder,
	}
}

// TimeType represents a TIMESTAMP column
type TimeType[B ExpressionBuilder] struct {
	valueExpression
	b B
}

func (c *TimeType[B]) Eq(val time.Time) ast.BooleanExpression {

	return c.b
}

func (c *TimeType[B]) Gt(val time.Time) ast.BooleanExpression {

	return c.b
}

func (c *TimeType[B]) Lt(val time.Time) ast.BooleanExpression {

	return c.b
}

func (c *TimeType[B]) Gte(val time.Time) ast.BooleanExpression {

	return c.b
}

func (c *TimeType[B]) Lte(val time.Time) ast.BooleanExpression {

	return c.b
}

func (c *TimeType[B]) Between(start, end time.Time) ast.BooleanExpression {

	return c.b
}

func Lit(val interface{}) ast.Expression {
	return ast.NewLiteralExpression(&ast.ExpressionNode{Literal: val})
}

func NewUUIDType[B ExpressionBuilder](ref *ast.FieldRef, builder B) *UUIDType[B] {
	return &UUIDType[B]{
		valueExpression: ast.NewValueExpression(ast.NewFieldExpresion(ref)),
		b:               builder,
	}
}

// UUIDType represents a UUID column
type UUIDType[B ExpressionBuilder] struct {
	valueExpression
	b B
}

func (c *UUIDType[B]) Eq(val uuid.UUID) ast.BooleanExpression {

	return c.b
}

func (c *UUIDType[B]) In(vals ...uuid.UUID) ast.BooleanExpression {

	return c.b
}

func (c *UUIDType[B]) IsNull() ast.BooleanExpression {

	return c.b
}

func (c *UUIDType[B]) IsNotNull() ast.BooleanExpression {

	return c.b
}

type ExpressionBuilder interface {
	ast.BooleanExpression
	setExpression(ast.Expression)
}
