package expr

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

func Null() ast.Expression {
	return ast.NewKeywordNode("NULL")
}

// InlineStringLiteral returns an expression rendered directly into SQL as a
// quoted string literal instead of as a bind parameter. Use this only for
// trusted, static values controlled by the program. Never pass user input,
// request data, database values, or partially sanitized strings.
func InlineStringLiteral(value string) ast.Expression {
	return ast.NewInlineStringLiteralNode(value)
}

func Coalesce[T ast.MappedTypes](expressions ...ast.OfType[T]) ast.Function[T] {
	args := make([]ast.Expression, 0, len(expressions))
	for _, expression := range expressions {
		args = append(args, expression)
	}
	return ast.NewFunction[T](ast.NewFunctionNode("COALESCE", args, nil))
}

func JsonAgg(expression ast.Expression) *ast.AggregationFunction[json.RawMessage] {
	return ast.NewAggregationFunction[json.RawMessage](
		ast.NewFunctionNode("json_agg", []ast.Expression{expression}, nil))
}

func Count(expression ast.Expression) ast.Function[int64] {
	return ast.NewFunction[int64](ast.NewFunctionNode("count", []ast.Expression{expression}, nil))
}

func Sum[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("sum", []ast.Expression{expression}, nil))
}

func Avg[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("avg", []ast.Expression{expression}, nil))
}

func Min[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("min", []ast.Expression{expression}, nil))
}

func Max[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("max", []ast.Expression{expression}, nil))
}

func RowNumber() ast.Function[int64] {
	return ast.NewFunction[int64](newEmptyFunctionNode("row_number"))
}

func Rank() ast.Function[int64] {
	return ast.NewFunction[int64](newEmptyFunctionNode("rank"))
}

func DenseRank() ast.Function[int64] {
	return ast.NewFunction[int64](newEmptyFunctionNode("dense_rank"))
}

func Lag[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("lag", []ast.Expression{expression}, nil))
}

func Lead[T ast.MappedTypes](expression ast.OfType[T]) ast.Function[T] {
	return ast.NewFunction[T](ast.NewFunctionNode("lead", []ast.Expression{expression}, nil))
}

func JsonbBuildObject(expressions ...ast.NamedExpression) ast.Function[json.RawMessage] {
	return ast.NewFunction[json.RawMessage](ast.NewFunctionNode(
		"jsonb_build_object",
		nil,
		func(node ast.ExpressionNode, builder *strings.Builder, params *[]any, ctx *ast.QueryContext) {
			builder.WriteString(node.Op + "(")
			for i := 0; i < len(expressions)-1; i++ {
				builder.WriteString("'" + expressions[i].Name() + "', ")
				ast.BuildQueryWithContext(expressions[i], builder, params, ctx)
				if ctx != nil && ctx.Error != nil {
					return
				}
				builder.WriteString(", ")
			}
			builder.WriteString("'" + expressions[len(expressions)-1].Name() + "', ")
			ast.BuildQueryWithContext(expressions[len(expressions)-1], builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(")")
		},
	))
}

func Unnest(arrays ...ast.Expression) *ast.SetReturningFunction {
	return ast.NewSetReturningFunction(
		ast.NewFunctionNode("unnest", arrays, nil))
}

func Values(provider ast.TableValuesProvider) *ast.ValuesTable {
	return ast.NewValuesTable(provider)
}

func Distinct(expressions ...ast.Expression) *ast.KeywordExpression {
	return ast.NewKeywordExpression("DISTINCT", expressions...)
}

func newEmptyFunctionNode(op string) *ast.FunctionNode {
	return ast.NewFunctionNode(op, nil,
		func(node ast.ExpressionNode, builder *strings.Builder, _ *[]any, ctx *ast.QueryContext) {
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(node.Op)
			builder.WriteString("()")
		})
}

func IsNull(expression ast.Expression) *ast.BoolType {
	return ast.Bool(&ast.BinaryNode{
		Op: "IS",
		Args: []ast.Expression{
			expression,
			Null(),
		},
	})
}

func IsNotNull(expression ast.Expression) *ast.BoolType {
	return ast.Bool(&ast.BinaryNode{
		Op: "IS",
		Args: []ast.Expression{
			expression,
			&ast.UnaryNode{
				Op: "NOT",
				Args: []ast.Expression{
					Null(),
				},
			},
		},
	})
}

func Cast[T ast.MappedTypes](expression ast.OfType[T]) *PartialCast[T] {
	return &PartialCast[T]{expression: expression}
}

type PartialCast[T ast.MappedTypes] struct {
	expression ast.Expression
}

func (c *PartialCast[T]) AsText() ast.OfType[string] {
	return toCastExpression[string](c.expression, "text")
}

func (c *PartialCast[T]) AsTextArray() ast.OfType[[]string] {
	return toCastExpression[[]string](c.expression, "text[]")
}

func (c *PartialCast[T]) AsInteger() ast.OfType[int] {
	return toCastExpression[int](c.expression, "integer")
}

func (c *PartialCast[T]) AsIntegerArray() ast.OfType[[]int] {
	return toCastExpression[[]int](c.expression, "integer[]")
}

func (c *PartialCast[T]) AsBigInt() ast.OfType[int64] {
	return toCastExpression[int64](c.expression, "bigint")
}

func (c *PartialCast[T]) AsBigIntArray() ast.OfType[[]int64] {
	return toCastExpression[[]int64](c.expression, "bigint[]")
}

func (c *PartialCast[T]) AsFloat() ast.OfType[float64] {
	return toCastExpression[float64](c.expression, "float")
}

func (c *PartialCast[T]) AsFloatArray() ast.OfType[[]float64] {
	return toCastExpression[[]float64](c.expression, "float[]")
}

func (c *PartialCast[T]) AsDouble() ast.OfType[float64] {
	return toCastExpression[float64](c.expression, "double")
}

func (c *PartialCast[T]) AsDoubleArray() ast.OfType[[]float64] {
	return toCastExpression[[]float64](c.expression, "double[]")
}

func (c *PartialCast[T]) AsNumeric() ast.OfType[pgtype.Numeric] {
	return toCastExpression[pgtype.Numeric](c.expression, "numeric")
}

func (c *PartialCast[T]) AsNumericArray() ast.OfType[[]pgtype.Numeric] {
	return toCastExpression[[]pgtype.Numeric](c.expression, "numeric[]")
}

func (c *PartialCast[T]) AsBoolean() ast.OfType[bool] {
	return toCastExpression[bool](c.expression, "boolean")
}

func (c *PartialCast[T]) AsBooleanArray() ast.OfType[[]bool] {
	return toCastExpression[[]bool](c.expression, "boolean[]")
}

func (c *PartialCast[T]) AsTimestamp() ast.OfType[time.Time] {
	return toCastExpression[time.Time](c.expression, "timestamp")
}

func (c *PartialCast[T]) AsTimestampArray() ast.OfType[[]time.Time] {
	return toCastExpression[[]time.Time](c.expression, "timestamp[]")
}

func (c *PartialCast[T]) AsDate() ast.OfType[time.Time] {
	return toCastExpression[time.Time](c.expression, "date")
}

func (c *PartialCast[T]) AsDateArray() ast.OfType[[]time.Time] {
	return toCastExpression[[]time.Time](c.expression, "date[]")
}

func (c *PartialCast[T]) AsTime() ast.OfType[time.Time] {
	return toCastExpression[time.Time](c.expression, "time")
}

func (c *PartialCast[T]) AsTimeArray() ast.OfType[[]time.Time] {
	return toCastExpression[[]time.Time](c.expression, "time[]")
}

func (c *PartialCast[T]) AsJson() ast.OfType[json.RawMessage] {
	return toCastExpression[json.RawMessage](c.expression, "json")
}

func (c *PartialCast[T]) AsJsonArray() ast.OfType[[]json.RawMessage] {
	return toCastExpression[[]json.RawMessage](c.expression, "json[]")
}

func (c *PartialCast[T]) AsJsonb() ast.OfType[json.RawMessage] {
	return toCastExpression[json.RawMessage](c.expression, "jsonb")
}

func (c *PartialCast[T]) AsJsonbArray() ast.OfType[[]json.RawMessage] {
	return toCastExpression[[]json.RawMessage](c.expression, "json[]")
}

func (c *PartialCast[T]) AsBytea() ast.OfType[[]byte] {
	return toCastExpression[[]byte](c.expression, "bytea")
}

func (c *PartialCast[T]) AsByteaArray() ast.OfType[[][]byte] {
	return toCastExpression[[][]byte](c.expression, "bytea[]")
}

func (c *PartialCast[T]) AsUUID() ast.OfType[uuid.UUID] {
	return toCastExpression[uuid.UUID](c.expression, "uuid")
}

func (c *PartialCast[T]) AsUUIDArray() ast.OfType[[]uuid.UUID] {
	return toCastExpression[[]uuid.UUID](c.expression, "uuid[]")
}

func toCastExpression[T ast.MappedTypes](expression ast.Expression, sqlType string) ast.OfType[T] {
	return ast.SetType[T](ast.NewFunctionNode("CAST", []ast.Expression{expression},
		func(node ast.ExpressionNode, builder *strings.Builder, params *[]any, ctx *ast.QueryContext) {
			if ctx != nil && ctx.Error != nil {
				return
			}
			if sqlType == "" {
				if ctx != nil && ctx.Error == nil {
					ctx.Error = errors.New("CAST expects a type")
				}
				return
			}
			if len(node.Args) != 1 {
				if ctx != nil && ctx.Error == nil {
					ctx.Error = errors.New("CAST expects exactly one expression")
				}
				return
			}
			builder.WriteString(node.Op + "(")
			ast.BuildQueryWithContext(node.Args[0], builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(" AS ")
			builder.WriteString(sqlType)
			builder.WriteString(")")
		}))
}
