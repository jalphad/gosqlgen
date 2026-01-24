package ast

import (
	"strings"
)

type function[T MappedTypes] = Function[T]
type Function[T MappedTypes] struct {
	sqlType[T]
}

func NewFunction[T MappedTypes](node *FunctionNode) *Function[T] {
	return &Function[T]{sqlType[T]{node}}
}

type AggregationFunction[T MappedTypes] Function[T]

func NewAggregationFunction[T MappedTypes](node *FunctionNode) *AggregationFunction[T] {
	return &AggregationFunction[T]{sqlType[T]{node}}
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
	srf
	relation *Relation
}

func (n NamedSetReturningFunction) Name() string {
	return n.relation.alias
}

type SetReturningFunction struct {
	function[[]any]
}
type srf = SetReturningFunction

func NewSetReturningFunction(node *FunctionNode) *SetReturningFunction {
	return &SetReturningFunction{function: function[[]any]{sqlType[[]any]{node}}}
}

func (f *SetReturningFunction) As(r *Relation, columns ...NamedExpression) *NamedSetReturningFunction {
	if f == nil {
		return nil
	}
	r.columns = columns

	return &NamedSetReturningFunction{srf: *f, relation: r}
}

func (f *SetReturningFunction) isTableExpression() {}

type Relation struct {
	alias   string
	columns []NamedExpression
}

func (r Relation) toSQL(builder *strings.Builder, _ *[]any) {
	builder.WriteString(r.alias + "(")
	for i := 0; i < len(r.columns)-1; i++ {
		builder.WriteString(r.columns[i].Name() + ", ")
	}
	builder.WriteString(r.columns[len(r.columns)-1].Name() + ")")
}

func (r Relation) Name() string {
	return r.alias
}

func NewRelation(alias string) *Relation {
	return &Relation{alias: alias}
}

type NamedAndTyped[T MappedTypes] interface {
	NamedExpression
	OfType[T]
}
