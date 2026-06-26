package ast

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ScanDest[T MappedTypes] interface {
	*T | **T
}

type ScanBinding interface {
	Destination() any
	Assign() error
}

type Projection[R any] interface {
	NamedExpression
	BindScan(*R) ScanBinding
}

type BaseProjection[R any, T MappedTypes, D ScanDest[T]] struct {
	OfType[T]
	name string
	dest func(*R) D
}

func NewProjection[R any, T MappedTypes, D ScanDest[T]](expr OfType[T], dest func(*R) D) *BaseProjection[R, T, D] {
	projection := &BaseProjection[R, T, D]{
		OfType: expr,
		dest:   dest,
	}
	if named, ok := expr.(NamedExpression); ok {
		projection.name = named.Name()
	}

	return projection
}

func (p *BaseProjection[R, T, D]) Name() string {
	return p.name
}

func (p *BaseProjection[R, T, D]) BindScan(t *R) ScanBinding {
	return scalarScanBinding[T, D]{dest: p.dest(t)}
}

type scalarScanBinding[T MappedTypes, D ScanDest[T]] struct {
	dest D
}

func (b scalarScanBinding[T, D]) Destination() any {
	return b.dest
}

func (b scalarScanBinding[T, D]) Assign() error {
	return nil
}

type CustomProjection[R any, T MappedTypes, O any] struct {
	OfType[T]
	name    string
	dest    func(*R) *O
	convert func(T) (O, error)
}

func NewCustomProjection[R any, T MappedTypes, O any](expr OfType[T], dest func(*R) *O, convert func(T) (O, error)) *CustomProjection[R, T, O] {
	projection := &CustomProjection[R, T, O]{
		OfType:  expr,
		dest:    dest,
		convert: convert,
	}
	if named, ok := expr.(NamedExpression); ok {
		projection.name = named.Name()
	}

	return projection
}

func NewJSONProjection[R any, O any](expr OfType[json.RawMessage], dest func(*R) *O) Projection[R] {
	return NewCustomProjection(expr, dest, func(raw json.RawMessage) (O, error) {
		var out O
		if err := json.Unmarshal(raw, &out); err != nil {
			return out, err
		}
		return out, nil
	})
}

func (p *CustomProjection[R, T, O]) Name() string {
	return p.name
}

func (p *CustomProjection[R, T, O]) BindScan(t *R) ScanBinding {
	return &customScanBinding[T, O]{
		dest:    p.dest(t),
		convert: p.convert,
	}
}

type customScanBinding[T MappedTypes, O any] struct {
	input   T
	dest    *O
	convert func(T) (O, error)
}

func (b *customScanBinding[T, O]) Destination() any {
	return &b.input
}

func (b *customScanBinding[T, O]) Assign() error {
	out, err := b.convert(b.input)
	if err != nil {
		return err
	}
	*b.dest = out
	return nil
}

type StringColumnProjection[R any, D ScanDest[string]] struct {
	*StringColumnExpression
	dest func(*R) D
}

func NewStringColumnProjection[R any, D ScanDest[string]](
	table string,
	column string,
	ref func(*R) D,
) *StringColumnProjection[R, D] {
	return &StringColumnProjection[R, D]{
		StringColumnExpression: NewStringColumnExpression(table, column),
		dest:                   ref,
	}
}

func (p *StringColumnProjection[R, D]) BindScan(r *R) ScanBinding {
	return &scalarScanBinding[string, D]{
		dest: p.dest(r),
	}
}

type IntColumnProjection[R any, D ScanDest[int64]] struct {
	*IntColumnExpression
	dest func(*R) D
}

func NewIntColumnProjection[R any, D ScanDest[int64]](
	table string,
	column string,
	ref func(*R) D,
) *IntColumnProjection[R, D] {
	return &IntColumnProjection[R, D]{
		IntColumnExpression: NewIntColumnExpression(table, column),
		dest:                ref,
	}
}

func (p *IntColumnProjection[R, D]) BindScan(r *R) ScanBinding {
	return &scalarScanBinding[int64, D]{
		dest: p.dest(r),
	}
}

type UUIDColumnProjection[R any, D ScanDest[uuid.UUID]] struct {
	*UUIDColumnExpression
	dest func(*R) D
}

func NewUUIDColumnProjection[R any, D ScanDest[uuid.UUID]](
	table string,
	column string,
	ref func(*R) D,
) *UUIDColumnProjection[R, D] {
	return &UUIDColumnProjection[R, D]{
		UUIDColumnExpression: NewUUIDColumnExpression(table, column),
		dest:                 ref,
	}
}

func (p *UUIDColumnProjection[R, D]) BindScan(r *R) ScanBinding {
	return &scalarScanBinding[uuid.UUID, D]{
		dest: p.dest(r),
	}
}

type BoolColumnProjection[R any, D ScanDest[bool]] struct {
	*BoolColumnExpression
	dest func(*R) D
}

func NewBoolColumnProjection[R any, D ScanDest[bool]](
	table string,
	column string,
	ref func(*R) D,
) *BoolColumnProjection[R, D] {
	return &BoolColumnProjection[R, D]{
		BoolColumnExpression: NewBoolColumnExpression(table, column),
		dest:                 ref,
	}
}

func (p *BoolColumnProjection[R, D]) BindScan(r *R) ScanBinding {
	return &scalarScanBinding[bool, D]{
		dest: p.dest(r),
	}
}

type TimestampColumnProjection[R any, D ScanDest[time.Time]] struct {
	*TimestampColumnExpression
	dest func(*R) D
}

func NewTimestampColumnProjection[R any, D ScanDest[time.Time]](
	table string,
	column string,
	ref func(*R) D,
) *TimestampColumnProjection[R, D] {
	return &TimestampColumnProjection[R, D]{
		TimestampColumnExpression: NewTimestampColumnExpression(table, column),
		dest:                      ref,
	}
}

func (p *TimestampColumnProjection[R, D]) BindScan(r *R) ScanBinding {
	return &scalarScanBinding[time.Time, D]{
		dest: p.dest(r),
	}
}
