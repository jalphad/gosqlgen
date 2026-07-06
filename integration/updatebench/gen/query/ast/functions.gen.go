package ast

import "strings"

type AliasedFunction[T MappedTypes] interface {
	ofType[T]
	NamedExpression
	isFunction()
}

type Function[T MappedTypes] interface {
	ofType[T]
	As(r *Alias) AliasedFunction[T]
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

type NamedSetReturningFunction struct {
	aliasedFunction[[]any]
}

func (f *NamedSetReturningFunction) isTableExpression() {}

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

type Alias struct {
	name    string
	columns []NamedExpression
}

func (r Alias) toSQL(builder *strings.Builder, _ *[]any, ctx *QueryContext) {
	if ctx != nil && ctx.Error != nil {
		return
	}
	builder.WriteString(r.name)
	if len(r.columns) == 0 {
		return
	}
	builder.WriteString("(")
	for i := 0; i < len(r.columns)-1; i++ {
		builder.WriteString(r.columns[i].Name() + ", ")
	}
	builder.WriteString(r.columns[len(r.columns)-1].Name() + ")")
}

func (r Alias) Name() string {
	return r.name
}

func (r Alias) Columns() []NamedExpression {
	if r.columns == nil {
		return []NamedExpression{}
	}

	return r.columns
}

func NewAlias(alias string, columns ...NamedExpression) *Alias {
	return &Alias{name: alias,
		columns: columns}
}

type NamedAndTyped[T MappedTypes] interface {
	NamedExpression
	OfType[T]
}
