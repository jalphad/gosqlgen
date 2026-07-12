package ast

import (
	"fmt"
	"strings"
)

type AliasedFunction[T MappedTypes] interface {
	ofType[T]
	NamedExpression
	isFunction()
}

type Function[T MappedTypes] interface {
	ofType[T]
	As(r *Alias) AliasedFunction[T]
	Over() WindowFunction[T]
	isFunction()
}
type function[T MappedTypes] struct {
	sqlType[T]
}

func NewFunction[T MappedTypes](node *FunctionNode) Function[T] {
	return &function[T]{sqlType: sqlType[T]{node}}
}

func (f function[T]) As(alias *Alias) AliasedFunction[T] {
	return &aliasedFunction[T]{function: f, alias: alias}
}

func (f function[T]) Over() WindowFunction[T] {
	return &windowFunction[T]{function: f}
}

func (f function[T]) isFunction() {}

type aliasedFunction[T MappedTypes] struct {
	function[T]
	alias *Alias
}

func (n *aliasedFunction[T]) Name() string {
	return n.alias.name
}

func (n *aliasedFunction[T]) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	n.function.toSQL(builder, params, ctx)
	builder.WriteString(" AS ")
	n.alias.toSQL(builder, params, ctx)
}

type AggregationFunction[T MappedTypes] function[T]

func NewAggregationFunction[T MappedTypes](node *FunctionNode) *AggregationFunction[T] {
	return &AggregationFunction[T]{sqlType: sqlType[T]{node}}
}

func (f *AggregationFunction[T]) As(alias *Alias) AliasedFunction[T] {
	return (*function[T])(f).As(alias)
}

func (f *AggregationFunction[T]) Over() WindowFunction[T] {
	return (*function[T])(f).Over()
}

func (f *AggregationFunction[T]) Filter(filter OfType[bool]) Expression {
	return &BinaryNode{
		Op: "FILTER",
		Args: []Expression{
			f,
			NewGroupedExpression(
				&UnaryNode{
					Op: "WHERE",
					Args: []Expression{
						filter,
					},
				},
			),
		},
	}
}

type AliasedWindowFunction[T MappedTypes] interface {
	ofType[T]
	NamedExpression
	isWindowFunction()
}

type WindowFunction[T MappedTypes] interface {
	ofType[T]
	PartitionBy(partitions ...Expression) WindowFunction[T]
	OrderBy(orderBy ...*OrderByItem) WindowFunction[T]
	As(alias *Alias) AliasedWindowFunction[T]
	isWindowFunction()
}

type windowFunction[T MappedTypes] struct {
	function   function[T]
	partition []Expression
	orderBy   []*OrderByItem
}

func (f *windowFunction[T]) isOfType(_ T) {}

func (f *windowFunction[T]) isWindowFunction() {}

func (f *windowFunction[T]) PartitionBy(partitions ...Expression) WindowFunction[T] {
	f.partition = append(f.partition, partitions...)
	return f
}

func (f *windowFunction[T]) OrderBy(orderBy ...*OrderByItem) WindowFunction[T] {
	f.orderBy = append(f.orderBy, orderBy...)
	return f
}

func (f *windowFunction[T]) As(alias *Alias) AliasedWindowFunction[T] {
	return &aliasedWindowFunction[T]{windowFunction: f, alias: alias}
}

func (f *windowFunction[T]) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if f == nil {
		if ctx != nil && ctx.Error == nil {
			ctx.Error = fmt.Errorf("nil WindowFunction")
		}
		return
	}
	f.function.toSQL(builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(" OVER (")
	if len(f.partition) > 0 {
		builder.WriteString("PARTITION BY ")
		for i := 0; i < len(f.partition)-1; i++ {
			f.partition[i].toSQL(builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(", ")
		}
		f.partition[len(f.partition)-1].toSQL(builder, params, ctx)
		if ctx != nil && ctx.Error != nil {
			return
		}
	}
	if len(f.orderBy) > 0 {
		if len(f.partition) > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString("ORDER BY ")
		for i := 0; i < len(f.orderBy)-1; i++ {
			renderOrderByItem(f.orderBy[i], builder, params, ctx)
			if ctx != nil && ctx.Error != nil {
				return
			}
			builder.WriteString(", ")
		}
		renderOrderByItem(f.orderBy[len(f.orderBy)-1], builder, params, ctx)
	}
	builder.WriteString(")")
}

type aliasedWindowFunction[T MappedTypes] struct {
	*windowFunction[T]
	alias *Alias
}

func (n *aliasedWindowFunction[T]) Name() string {
	return n.alias.name
}

func (n *aliasedWindowFunction[T]) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	n.windowFunction.toSQL(builder, params, ctx)
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(" AS ")
	n.alias.toSQL(builder, params, ctx)
}

type NamedSetReturningFunction struct {
	aliasedFunction[[]any]
}

func (f *NamedSetReturningFunction) isTableExpression() {}

func (f *NamedSetReturningFunction) Join(jointype JoinType, table NamedTableExpression, on OfType[bool]) TableExpression {
	return joinTableExpression(f, jointype, table, on)
}

type SetReturningFunction struct {
	function[[]any]
	columns []NamedExpression
}

func NewSetReturningFunction(node *FunctionNode) *SetReturningFunction {
	return &SetReturningFunction{function: function[[]any]{sqlType[[]any]{node}}}
}

func (f *SetReturningFunction) As(alias *Alias) *NamedSetReturningFunction {
	if f == nil {
		return nil
	}

	return &NamedSetReturningFunction{
		aliasedFunction: aliasedFunction[[]any]{
			function: f.function,
			alias:    alias,
		},
	}
}

func (f *SetReturningFunction) isTableExpression() {}

func (f *SetReturningFunction) Join(jointype JoinType, table NamedTableExpression, on OfType[bool]) TableExpression {
	return joinTableExpression(f, jointype, table, on)
}

type Alias struct {
	name    string
	columns []NamedExpression
	source  *TableSource
}

func (r *Alias) isTableExpression() {}

func (r *Alias) Join(jointype JoinType, table NamedTableExpression, on OfType[bool]) TableExpression {
	return joinTableExpression(r, jointype, table, on)
}

func (r *Alias) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	if r.source != nil {
		r.source.toSQL(builder, params, ctx)
		if ctx != nil && ctx.Error != nil {
			return
		}
		builder.WriteString(" AS ")
		builder.WriteString(r.name)
		return
	}
	builder.WriteString(r.name)
	if len(r.columns) == 0 {
		return
	}
	builder.WriteString("(")
	for i := 0; i < len(r.columns)-1; i++ {
		builder.WriteString(r.columns[i].Name())
		builder.WriteString(", ")
	}
	builder.WriteString(r.columns[len(r.columns)-1].Name())
	builder.WriteString(")")
}

func (r *Alias) Name() string {
	return r.name
}

func (r *Alias) Columns() []NamedExpression {
	if r.columns == nil {
		return []NamedExpression{}
	}

	return r.columns
}

func (r *Alias) Column(name string) NamedExpression {
	for _, e := range r.columns {
		if e.Name() == name {
			return NewColumnNode(r.name, e.Name())
		}
	}

	return NewNamedExpression(r.name,
		NewErrorExpression(fmt.Errorf("column does not exist for alias %s", r.name)))
}

func NewAlias(alias string, columns ...NamedExpression) *Alias {
	return &Alias{name: alias,
		columns: columns}
}

type NamedAndTyped[T MappedTypes] interface {
	NamedExpression
	OfType[T]
}
