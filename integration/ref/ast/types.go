package ast

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ArrayTypes interface {
	[]float64 | []int | []int64 | []string | []bool |
		[]time.Time | []uuid.UUID | []pgtype.Numeric |
		[]json.RawMessage | []any
}

type NullableMappedTypes[T MappedTypes] interface {
	*T
}

type MappedTypes interface {
	float64 | int | int64 | string | bool |
		time.Time | time.Duration |
		uuid.UUID |
		pgtype.Numeric |
		[]byte | json.RawMessage |
		ArrayTypes
}

type ofType[T MappedTypes] = OfType[T]
type OfType[T MappedTypes] interface {
	expression
	isOfType(_ T)
}

func SetType[T MappedTypes](e Expression) OfType[T] {
	return &SQLType[T]{e}
}

func NewSQLType[T MappedTypes](t T) OfType[T] {
	return &SQLType[T]{NewLiteralExpression(t)}
}

// SQLType is the underlying type for interacting with SQL types through queries
type SQLType[T MappedTypes] struct {
	expression
}

func (t *SQLType[T]) isOfType(_ T) {}

type sqlType[T MappedTypes] = SQLType[T]

func String(e Expression) *StringType {
	return &StringType{SQLType[string]{e}}
}

// StringType represents a VARCHAR/TEXT type
type StringType struct {
	sqlType[string]
}

func (t *StringType) Eq(expr OfType[string]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *StringType) Like(pattern string) *BoolType {
	return Bool(&BinaryNode{Op: "LIKE", Args: []Expression{t, NewLiteralExpression(pattern)}})
}

func (t *StringType) In(expr OfType[[]string]) *BoolType {
	return Bool(&BinaryNode{Op: "IN", Args: []Expression{t, expr}})
}

func (t *StringType) Between(start, end OfType[string]) *BoolType {
	return Bool(&UnaryNode{
		Op: "BETWEEN",
		Args: []Expression{
			&BinaryNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			},
		},
	})
}

func (t *StringType) IsNull() *BoolType {
	return Bool(&UnaryNode{Op: "IS NULL", Args: []Expression{t}})
}

func (t *StringType) IsNotNull() *BoolType {
	return Bool(&UnaryNode{Op: "IS NOT NULL", Args: []Expression{t}})
}

func Int(e Expression) *IntType {
	return &IntType{SQLType[int64]{e}}
}

// IntType represents an integer expression
type IntType struct {
	sqlType[int64]
}

func (t *IntType) Eq(expr OfType[int64]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *IntType) Gt(expr OfType[int64]) *BoolType {
	return Bool(&BinaryNode{Op: ">", Args: []Expression{t, expr}})
}

func (t *IntType) Lt(expr OfType[int64]) *BoolType {
	return Bool(&BinaryNode{Op: "<", Args: []Expression{t, expr}})
}

func (t *IntType) Gte(expr OfType[int64]) *BoolType {
	return Bool(&BinaryNode{Op: ">=", Args: []Expression{t, expr}})
}

func (t *IntType) Lte(expr OfType[int64]) *BoolType {
	return Bool(&BinaryNode{Op: "<=", Args: []Expression{t, expr}})
}

func (t *IntType) In(expr OfType[[]int64]) *BoolType {
	return Bool(&BinaryNode{Op: "IN", Args: []Expression{t, expr}})
}

func (t *IntType) Between(start, end OfType[int64]) *BoolType {
	return Bool(&UnaryNode{
		Op: "BETWEEN",
		Args: []Expression{
			&BinaryNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			},
		},
	})
}

func Float(e Expression) *FloatType {
	return &FloatType{SQLType[float64]{e}}
}

// FloatType represents an floating point expression
type FloatType struct {
	sqlType[float64]
}

func (t *FloatType) Eq(expr OfType[float64]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *FloatType) Gt(expr OfType[float64]) *BoolType {
	return Bool(&BinaryNode{Op: ">", Args: []Expression{t, expr}})
}

func (t *FloatType) Lt(expr OfType[float64]) *BoolType {
	return Bool(&BinaryNode{Op: "<", Args: []Expression{t, expr}})
}

func (t *FloatType) Gte(expr OfType[float64]) *BoolType {
	return Bool(&BinaryNode{Op: ">=", Args: []Expression{t, expr}})
}

func (t *FloatType) Lte(expr OfType[float64]) *BoolType {
	return Bool(&BinaryNode{Op: "<=", Args: []Expression{t, expr}})
}

func (t *FloatType) In(expr OfType[[]float64]) *BoolType {
	return Bool(&BinaryNode{Op: "IN", Args: []Expression{t, expr}})
}

func (t *FloatType) Between(start, end OfType[float64]) *BoolType {
	return Bool(&UnaryNode{
		Op: "BETWEEN",
		Args: []Expression{
			&BinaryNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			},
		},
	})
}

// NumericType represents an integer expression
type NumericType struct {
	sqlType[pgtype.Numeric]
}

func (t *NumericType) Eq(expr OfType[pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *NumericType) Gt(expr OfType[pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: ">", Args: []Expression{t, expr}})
}

func (t *NumericType) Lt(expr OfType[pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: "<", Args: []Expression{t, expr}})
}

func (t *NumericType) Gte(expr OfType[pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: ">=", Args: []Expression{t, expr}})
}

func (t *NumericType) Lte(expr OfType[pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: "<=", Args: []Expression{t, expr}})
}

func (t *NumericType) In(expr OfType[[]pgtype.Numeric]) *BoolType {
	return Bool(&BinaryNode{Op: "IN", Args: []Expression{t, expr}})
}

func (t *NumericType) Between(start, end OfType[pgtype.Numeric]) *BoolType {
	return Bool(&UnaryNode{
		Op: "BETWEEN",
		Args: []Expression{
			&BinaryNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			},
		},
	})
}

func Bool(e Expression) *BoolType {
	return &BoolType{sqlType[bool]{e}}
}

// BoolType represents a BOOLEAN column
type BoolType struct {
	sqlType[bool]
}

func (t *BoolType) And(expr OfType[bool]) *BoolType {
	return Bool(&BinaryNode{Op: "AND", Args: []Expression{t, expr}})
}

func (t *BoolType) Or(expr OfType[bool]) *BoolType {
	return Bool(&BinaryNode{Op: "OR", Args: []Expression{t, expr}})
}

func (t *BoolType) Eq(expr OfType[bool]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *BoolType) IsTrue() *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, NewLiteralExpression(true)}})
}

func (t *BoolType) IsFalse() *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, NewLiteralExpression(false)}})
}

func (t *BoolType) IsNull() *BoolType {
	return Bool(&UnaryNode{
		Op:   "IS NULL",
		Args: []Expression{t},
	})
}

type TimestampType struct {
	sqlType[time.Time]
}

func Timestamp(e sqlType[time.Time]) *TimestampType {
	return &TimestampType{e}
}

func (t *TimestampType) isOfType(_ time.Time) {}

type DateType struct {
	sqlType[time.Time]
}

func Date(e sqlType[time.Time]) *DateType {
	return &DateType{e}
}

func (d *DateType) isOfType(_ time.Time) {}

// TimeType represents a TIMESTAMP column
type TimeType struct {
	sqlType[time.Time]
}

func Time(e sqlType[time.Time]) *TimeType {
	return &TimeType{e}
}

func (t *TimeType) Eq(expr OfType[time.Time]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *TimeType) Gt(expr OfType[time.Time]) *BoolType {
	return Bool(&BinaryNode{Op: ">", Args: []Expression{t, expr}})
}

func (t *TimeType) Lt(expr OfType[time.Time]) *BoolType {
	return Bool(&BinaryNode{Op: "<", Args: []Expression{t, expr}})
}

func (t *TimeType) Gte(expr OfType[time.Time]) *BoolType {
	return Bool(&BinaryNode{Op: ">=", Args: []Expression{t, expr}})
}

func (t *TimeType) Lte(expr OfType[time.Time]) *BoolType {
	return Bool(&BinaryNode{Op: "<=", Args: []Expression{t, expr}})
}

func (t *TimeType) Between(start, end OfType[time.Time]) *BoolType {
	return Bool(&UnaryNode{
		Op: "BETWEEN",
		Args: []Expression{
			&BinaryNode{
				Op: "AND",
				Args: []Expression{
					start,
					end,
				},
			},
		},
	})
}

func UUID(e Expression) *UUIDType {
	return &UUIDType{sqlType[uuid.UUID]{e}}
}

// UUIDType represents a UUID column
type UUIDType struct {
	sqlType[uuid.UUID]
}

func (t *UUIDType) Eq(expr OfType[uuid.UUID]) *BoolType {
	return Bool(&BinaryNode{Op: "=", Args: []Expression{t, expr}})
}

func (t *UUIDType) In(expr OfType[[]uuid.UUID]) *BoolType {
	return Bool(&BinaryNode{Op: "IN", Args: []Expression{t, expr}})
}

func (t *UUIDType) IsNull() *BoolType {
	return Bool(&UnaryNode{
		Op:   "IS NULL",
		Args: []Expression{t},
	})
}

func (t *UUIDType) IsNotNull() *BoolType {
	return Bool(&UnaryNode{
		Op:   "IS NOT NULL",
		Args: []Expression{t},
	})
}

func Array[T ArrayTypes](t T) *ArrayType[T] {
	return &ArrayType[T]{NewSQLType(t)}
}

type ArrayType[T ArrayTypes] struct {
	ofType[T]
}

type BytesType struct {
	sqlType[[]byte]
}

func Json(e Expression) *JsonType {
	return &JsonType{sqlType[json.RawMessage]{e}}
}

type JsonType struct {
	sqlType[json.RawMessage]
}

type (
	stringType  = StringType
	intType     = IntType
	floatType   = FloatType
	numericType = NumericType
	boolType    = BoolType
	timeType    = TimeType
	dateType    = DateType
	uuidType    = UUIDType
	bytesType   = BytesType
	jsonType    = JsonType
)
