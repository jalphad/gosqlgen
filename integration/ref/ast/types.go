package ast

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	expression = Expression
)

type MappedTypes interface {
	float64 | []float64 |
		int | []int |
		string | []string |
		bool |
		time.Time | []time.Time |
		time.Duration |
		uuid.UUID | []uuid.UUID |
		pgtype.Numeric |
		[]byte
}

type OfType[T MappedTypes] interface {
	expression
	isOfType(_ T)
}

func NewSQLType[T MappedTypes](t T) *SQLType[T] {
	return &SQLType[T]{NewLiteralExpression(t)}
}

// SQLType is the underlying type for interacting with SQL types through queries
//
// 2025-12-1: [T any] necessary because Go cannot infer ArrayType[[]T]
type SQLType[T any] struct {
	expression
}

func (t *SQLType[T]) isOfType(_ T) {}

type sqlType[T any] = SQLType[T]

func String(e Expression) *StringType {
	return &StringType{sqlType[string]{e}}
}

// StringType represents a VARCHAR column
type StringType struct {
	sqlType[string]
}

func (t *StringType) Eq(expr OfType[string]) *BoolType {

	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, expr}}))
}

func (t *StringType) Like(pattern string) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "LIKE", Args: []Expression{t.expression, NewLiteralExpression(pattern)}}))
}

func (t *StringType) In(expr OfType[[]string]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "IN", Args: []Expression{t.expression, expr}}))
}

func (t *StringType) Between(start, end OfType[string]) *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op: "BETWEEN",
		Args: []Expression{
			NewBinaryExpression(&ExpressionNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			}),
		},
	}))
}

func (t *StringType) IsNull() *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{Op: "IS NULL", Args: []Expression{t.expression}}))
}

func (t *StringType) IsNotNull() *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{Op: "IS NOT NULL", Args: []Expression{t.expression}}))
}

// IntType represents an integer expression
type IntType struct {
	sqlType[int]
}

func (t *IntType) Eq(expr OfType[int]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, expr}}))
}

func (t *IntType) Gt(expr OfType[int]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: ">", Args: []Expression{t.expression, expr}}))
}

func (t *IntType) Lt(expr OfType[int]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "<", Args: []Expression{t.expression, expr}}))
}

func (t *IntType) Gte(expr OfType[int]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: ">=", Args: []Expression{t.expression, expr}}))
}

func (t *IntType) Lte(expr OfType[int]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "<=", Args: []Expression{t.expression, expr}}))
}

func (t *IntType) Between(start, end OfType[int]) *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op: "BETWEEN",
		Args: []Expression{
			NewBinaryExpression(&ExpressionNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			}),
		},
	}))
}

func Bool(e Expression) *BoolType {
	return &BoolType{sqlType[bool]{e}}
}

// BoolType represents a BOOLEAN column
type BoolType struct {
	sqlType[bool]
}

func (t *BoolType) And(expr *BoolType) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "AND", Args: []Expression{t.expression, expr}}))
}

func (t *BoolType) Or(expr *BoolType) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "OR", Args: []Expression{t.expression, expr}}))
}

func (t *BoolType) Eq(expr *BoolType) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, expr}}))
}

func (t *BoolType) IsTrue() *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, NewLiteralExpression(true)}}))
}

func (t *BoolType) IsFalse() *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, NewLiteralExpression(false)}}))
}

func (t *BoolType) IsNull() *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op:   "IS NULL",
		Args: []Expression{t.expression},
	}))
}

func Timestamp(e sqlType[time.Time]) *TimestampType {
	return &TimestampType{e}
}

type TimestampType struct {
	sqlType[time.Time]
}

func (t *TimestampType) isOfType(_ time.Time) {}

func Date(e sqlType[time.Time]) *DateType {
	return &DateType{e}
}

type DateType struct {
	sqlType[time.Time]
}

func (d *DateType) isOfType(_ time.Time) {}

func Time(e sqlType[time.Time]) *TimeType {
	return &TimeType{e}
}

// TimeType represents a TIMESTAMP column
type TimeType struct {
	sqlType[time.Time]
}

func (t *TimeType) Eq(expr OfType[time.Time]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, expr}}))
}

func (t *TimeType) Gt(expr OfType[time.Time]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: ">", Args: []Expression{t.expression, expr}}))
}

func (t *TimeType) Lt(expr OfType[time.Time]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "<", Args: []Expression{t.expression, expr}}))
}

func (t *TimeType) Gte(expr OfType[time.Time]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: ">=", Args: []Expression{t.expression, expr}}))
}

func (t *TimeType) Lte(expr OfType[time.Time]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "<=", Args: []Expression{t.expression, expr}}))
}

func (t *TimeType) Between(start, end OfType[time.Time]) *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op: "BETWEEN",
		Args: []Expression{
			NewBinaryExpression(&ExpressionNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			}),
		},
	}))
}

func UUID(e Expression) *UUIDType {
	return &UUIDType{sqlType[uuid.UUID]{e}}
}

// UUIDType represents a UUID column
type UUIDType struct {
	sqlType[uuid.UUID]
}

func (t *UUIDType) Eq(expr OfType[uuid.UUID]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "=", Args: []Expression{t.expression, expr}}))
}

func (t *UUIDType) In(expr OfType[[]uuid.UUID]) *BoolType {
	return Bool(NewBinaryExpression(&ExpressionNode{Op: "IN", Args: []Expression{t.expression, expr}}))
}

func (t *UUIDType) IsNull() *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op:   "IS NULL",
		Args: []Expression{t.expression},
	}))
}

func (t *UUIDType) IsNotNull() *BoolType {
	return Bool(NewUnaryExpression(&ExpressionNode{
		Op:   "IS NOT NULL",
		Args: []Expression{t.expression},
	}))
}

type ArrayType[T MappedTypes] struct {
	sqlType[[]T]
}

type Bytes struct {
	sqlType[[]byte]
}
