package ast

func Render(e Expression, params *[]any) string {
	return e.toSQL(params)
}

// Expression represents any SQL Expression (binary, unary, function, literal).
type Expression interface {
	toSQL(*[]any) string
}

type AliasedExpression interface {
	NamedExpression
	hasAlias()
}

type NamedExpression interface {
	Name() string
}

func NewAliasedExpression[T MappedTypes](expression OfType[T], alias string) *AsExpression[T] {
	return &AsExpression[T]{
		OfType: expression,
		alias:  alias,
	}
}

type AsExpression[T MappedTypes] struct {
	OfType[T]
	alias string
}

func (a *AsExpression[T]) Name() string {
	return a.alias
}

func (a *AsExpression[T]) hasAlias() {}

func NewNamedExpression(name string, expression Expression) NamedExpression {
	return &namedExpression{
		Expression: expression,
		name:       name,
	}
}

type namedExpression struct {
	Expression
	name string
}

func (e *namedExpression) Name() string {
	return e.name
}

type ColumnExpression[T MappedTypes] struct {
	OfType[T]
	ColumnNode
}
