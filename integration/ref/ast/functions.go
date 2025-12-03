package ast

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
