package ast

import (
	"fmt"
	"strconv"
	"strings"
)

type TableValuesProvider interface {
	RowCount() int
	ColumnCount() int
	Value(row int, column int) (any, error)
}

type TypedTableValuesProvider interface {
	TableValuesProvider
	ColumnSQLType(column int) (string, error)
}

type ValuesTable struct {
	provider TableValuesProvider
}

func NewValuesTable(provider TableValuesProvider) *ValuesTable {
	return &ValuesTable{provider: provider}
}

func (v *ValuesTable) As(alias *TableAlias) *NamedValuesTable {
	return &NamedValuesTable{
		table: v,
		alias: alias,
	}
}

func (v *ValuesTable) isTableExpression() {}

func (v *ValuesTable) Join(join *JoinExpr) TableExpression {
	return joinTableExpression(v, join)
}

func (v *ValuesTable) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	v.render(builder, params, ctx, true)
}

func (v *ValuesTable) render(builder *strings.Builder, params *[]any, ctx *QueryContext, parenthesized bool) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if v == nil || v.provider == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("VALUES table requires a provider")
		}
		return
	}

	rowCount := v.provider.RowCount()
	columnCount := v.provider.ColumnCount()
	if rowCount == 0 || columnCount == 0 {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("VALUES table requires at least one row and one column")
		}
		return
	}

	valueCount := rowCount * columnCount
	typedProvider, hasColumnTypes := v.provider.(TypedTableValuesProvider)
	// Pre-grow for the common placeholder shape: about 8 bytes per value
	// for "$N" plus separators, about 4 bytes per row for parentheses and
	// row separators, and the "(VALUES " prefix. This is a heuristic; the
	// params slice remains the source of truth for exact placeholder numbers.
	builder.Grow(valueCount*8 + rowCount*4 + 8)
	if cap(*params) < len(*params)+valueCount {
		next := make([]any, len(*params), len(*params)+valueCount)
		copy(next, *params)
		*params = next
	}
	currentParams := *params

	if parenthesized {
		builder.WriteString("(VALUES ")
	} else {
		builder.WriteString("VALUES ")
	}
	var placeholder [20]byte
	for row := 0; row < rowCount; row++ {
		if row > 0 {
			builder.WriteString(", ")
		}
		builder.WriteByte('(')
		for column := 0; column < columnCount; column++ {
			if column > 0 {
				builder.WriteString(", ")
			}
			value, err := v.provider.Value(row, column)
			if err != nil {
				*params = currentParams
				if ctx != nil && ctx.Error == nil {
					ctx.Error = err
				}
				return
			}
			currentParams = append(currentParams, value)
			sqlType := ""
			if hasColumnTypes && row == 0 {
				sqlType, err = typedProvider.ColumnSQLType(column)
				if err != nil {
					*params = currentParams
					if ctx != nil && ctx.Error == nil {
						ctx.Error = err
					}
					return
				}
			}
			if sqlType != "" {
				builder.WriteString("CAST($")
			} else {
				builder.WriteByte('$')
			}
			placeholderBytes := strconv.AppendInt(placeholder[:0], int64(len(currentParams)), 10)
			_, _ = builder.Write(placeholderBytes)
			if sqlType != "" {
				builder.WriteString(" AS ")
				builder.WriteString(sqlType)
				builder.WriteByte(')')
			}
		}
		builder.WriteByte(')')
	}
	*params = currentParams
	if parenthesized {
		builder.WriteByte(')')
	}
}

type NamedValuesTable struct {
	table *ValuesTable
	alias *TableAlias
}

func (v *NamedValuesTable) GetName() string {
	return v.alias.GetName()
}

func (v *NamedValuesTable) isTableExpression() {}

func (v *NamedValuesTable) Join(join *JoinExpr) TableExpression {
	return joinTableExpression(v, join)
}

func (v *NamedValuesTable) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	v.table.toSQL(builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(" AS ")
	v.alias.renderDefinition(builder)
	if ctx != nil {
		ctx.registerRangeVariable(v.alias)
	}
}
