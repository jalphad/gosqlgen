package ast

import "strings"

type ScanDest[T MappedTypes] interface {
	*T | **T
}

type Projection[T any] interface {
	NamedExpression
	ScanDestination(*T) any
}

type projection[T any, V MappedTypes, D ScanDest[V]] struct {
	expression OfType[V]
	name       string
	dest       func(*T) D
}

func NewProjection[T any, V MappedTypes, D ScanDest[V]](expr OfType[V], dest func(*T) D) Projection[T] {
	name := ""
	if named, ok := expr.(NamedExpression); ok {
		name = named.Name()
	}
	return &projection[T, V, D]{
		expression: expr,
		name:       name,
		dest:       dest,
	}
}

func (p *projection[T, V, D]) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	p.expression.toSQL(builder, params, ctx)
}

func (p *projection[T, V, D]) Name() string {
	return p.name
}

func (p *projection[T, V, D]) ScanDestination(t *T) any {
	return p.dest(t)
}
