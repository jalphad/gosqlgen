package ast

import (
	"encoding/json"
	"strings"
)

type ScanDest[T MappedTypes] interface {
	*T | **T
}

type ScanBinding interface {
	Destination() any
	Assign() error
}

type Projection[T any] interface {
	NamedExpression
	BindScan(*T) ScanBinding
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

func (p *projection[T, V, D]) BindScan(t *T) ScanBinding {
	return scalarScanBinding[V, D]{dest: p.dest(t)}
}

type scalarScanBinding[V MappedTypes, D ScanDest[V]] struct {
	dest D
}

func (b scalarScanBinding[V, D]) Destination() any {
	return b.dest
}

func (b scalarScanBinding[V, D]) Assign() error {
	return nil
}

type customProjection[T any, I MappedTypes, O any] struct {
	expression OfType[I]
	name       string
	dest       func(*T) *O
	convert    func(I) (O, error)
}

func NewCustomProjection[T any, I MappedTypes, O any](expr OfType[I], dest func(*T) *O, convert func(I) (O, error)) Projection[T] {
	name := ""
	if named, ok := expr.(NamedExpression); ok {
		name = named.Name()
	}
	return &customProjection[T, I, O]{
		expression: expr,
		name:       name,
		dest:       dest,
		convert:    convert,
	}
}

func NewJSONProjection[T any, O any](expr OfType[json.RawMessage], dest func(*T) *O) Projection[T] {
	return NewCustomProjection(expr, dest, func(raw json.RawMessage) (O, error) {
		var out O
		if err := json.Unmarshal(raw, &out); err != nil {
			return out, err
		}
		return out, nil
	})
}

func (p *customProjection[T, I, O]) toSQL(builder *strings.Builder, params *[]any, ctx *QueryContext) {
	p.expression.toSQL(builder, params, ctx)
}

func (p *customProjection[T, I, O]) Name() string {
	return p.name
}

func (p *customProjection[T, I, O]) BindScan(t *T) ScanBinding {
	return &customScanBinding[I, O]{
		dest:    p.dest(t),
		convert: p.convert,
	}
}

type customScanBinding[I MappedTypes, O any] struct {
	input   I
	dest    *O
	convert func(I) (O, error)
}

func (b *customScanBinding[I, O]) Destination() any {
	return &b.input
}

func (b *customScanBinding[I, O]) Assign() error {
	out, err := b.convert(b.input)
	if err != nil {
		return err
	}
	*b.dest = out
	return nil
}
