package ast

import (
	"time"

	"github.com/google/uuid"
)

type StringFieldProjection[R any, D ScanDest[string]] struct {
	*StringColumnExpression
	dest func(*R) D
}

func NewStringFieldProjection[R any, D ScanDest[string]](
	table string,
	column string,
	dest func(*R) D,
) *StringFieldProjection[R, D] {
	return &StringFieldProjection[R, D]{
		StringColumnExpression: NewStringColumnExpression(table, column),
		dest:                   dest,
	}
}

func (p *StringFieldProjection[R, D]) BindScan(r *R) ScanBinding {
	return scalarScanBinding[string, D]{dest: p.dest(r)}
}

type IntFieldProjection[R any, D ScanDest[int64]] struct {
	*IntColumnExpression
	dest func(*R) D
}

func NewIntFieldProjection[R any, D ScanDest[int64]](
	table string,
	column string,
	dest func(*R) D,
) *IntFieldProjection[R, D] {
	return &IntFieldProjection[R, D]{
		IntColumnExpression: NewIntColumnExpression(table, column),
		dest:                dest,
	}
}

func (p *IntFieldProjection[R, D]) BindScan(r *R) ScanBinding {
	return scalarScanBinding[int64, D]{dest: p.dest(r)}
}

type UUIDFieldProjection[R any, D ScanDest[uuid.UUID]] struct {
	*UUIDColumnExpression
	dest func(*R) D
}

func NewUUIDFieldProjection[R any, D ScanDest[uuid.UUID]](
	table string,
	column string,
	dest func(*R) D,
) *UUIDFieldProjection[R, D] {
	return &UUIDFieldProjection[R, D]{
		UUIDColumnExpression: NewUUIDColumnExpression(table, column),
		dest:                 dest,
	}
}

func (p *UUIDFieldProjection[R, D]) BindScan(r *R) ScanBinding {
	return scalarScanBinding[uuid.UUID, D]{dest: p.dest(r)}
}

type BoolFieldProjection[R any, D ScanDest[bool]] struct {
	*BoolColumnExpression
	dest func(*R) D
}

func NewBoolFieldProjection[R any, D ScanDest[bool]](
	table string,
	column string,
	dest func(*R) D,
) *BoolFieldProjection[R, D] {
	return &BoolFieldProjection[R, D]{
		BoolColumnExpression: NewBoolColumnExpression(table, column),
		dest:                 dest,
	}
}

func (p *BoolFieldProjection[R, D]) BindScan(r *R) ScanBinding {
	return scalarScanBinding[bool, D]{dest: p.dest(r)}
}

type TimestampFieldProjection[R any, D ScanDest[time.Time]] struct {
	*TimestampColumnExpression
	dest func(*R) D
}

func NewTimestampFieldProjection[R any, D ScanDest[time.Time]](
	table string,
	column string,
	dest func(*R) D,
) *TimestampFieldProjection[R, D] {
	return &TimestampFieldProjection[R, D]{
		TimestampColumnExpression: NewTimestampColumnExpression(table, column),
		dest:                      dest,
	}
}

func (p *TimestampFieldProjection[R, D]) BindScan(r *R) ScanBinding {
	return scalarScanBinding[time.Time, D]{dest: p.dest(r)}
}
