package comments

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
)

// NewQuery returns a query builder for comments
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.CommentsDto] {
	return builder.NewKnownTableBuilder[models.CommentsDto](pool, Table())
}

var Into = intoCommentsDto{}

type intoCommentsDto struct{}

type forCommentsDto[E models.ExportsCommentsDto[T], T any] struct{}

func For[E models.ExportsCommentsDto[T], T any]() forCommentsDto[E, T] {
	return forCommentsDto[E, T]{}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("comments")
}

func Id() *ast.IntColumnProjection[models.CommentsDto, **int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"id",
		func(c *models.CommentsDto) **int64 {
			return &c.Id
		},
	)
}

func PostId() *ast.IntColumnProjection[models.CommentsDto, *int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"post_id",
		func(c *models.CommentsDto) *int64 {
			return &c.PostId
		},
	)
}

func UserId() *ast.UUIDColumnProjection[models.CommentsDto, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"comments",
		"user_id",
		func(c *models.CommentsDto) *uuid.UUID {
			return &c.UserId
		},
	)
}

func Content() *ast.StringColumnProjection[models.CommentsDto, *string] {
	return ast.NewStringColumnProjection(
		"comments",
		"content",
		func(c *models.CommentsDto) *string {
			return &c.Content
		},
	)
}

func IsApproved() *ast.BoolColumnProjection[models.CommentsDto, **bool] {
	return ast.NewBoolColumnProjection(
		"comments",
		"is_approved",
		func(c *models.CommentsDto) **bool {
			return &c.IsApproved
		},
	)
}

func CreatedAt() *ast.TimestampColumnProjection[models.CommentsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"comments",
		"created_at",
		func(c *models.CommentsDto) **time.Time {
			return &c.CreatedAt
		},
	)
}

func TestDate() *ast.DateColumnProjection[models.CommentsDto, **time.Time] {
	return ast.NewDateColumnProjection(
		"comments",
		"test_date",
		func(c *models.CommentsDto) **time.Time {
			return &c.TestDate
		},
	)
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		PostId(),
		UserId(),
		Content(),
		IsApproved(),
		CreatedAt(),
		TestDate(),
	}
}

func (intoCommentsDto) Id() ast.Projection[models.CommentsDto] {
	return Id()
}

func (intoCommentsDto) PostId() ast.Projection[models.CommentsDto] {
	return PostId()
}

func (intoCommentsDto) UserId() ast.Projection[models.CommentsDto] {
	return UserId()
}

func (intoCommentsDto) Content() ast.Projection[models.CommentsDto] {
	return Content()
}

func (intoCommentsDto) IsApproved() ast.Projection[models.CommentsDto] {
	return IsApproved()
}

func (intoCommentsDto) CreatedAt() ast.Projection[models.CommentsDto] {
	return CreatedAt()
}

func (intoCommentsDto) TestDate() ast.Projection[models.CommentsDto] {
	return TestDate()
}

func (intoCommentsDto) AllColumns() []ast.Projection[models.CommentsDto] {
	return []ast.Projection[models.CommentsDto]{
		Into.Id(),
		Into.PostId(),
		Into.UserId(),
		Into.Content(),
		Into.IsApproved(),
		Into.CreatedAt(),
		Into.TestDate(),
	}
}

func (forCommentsDto[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"id",
		func(t *T) **int64 {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.Id
		},
	)
}

func (forCommentsDto[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"post_id",
		func(t *T) *int64 {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.PostId
		},
	)
}

func (forCommentsDto[E, T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"comments",
		"user_id",
		func(t *T) *uuid.UUID {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.UserId
		},
	)
}

func (forCommentsDto[E, T]) Content() *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		"comments",
		"content",
		func(t *T) *string {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.Content
		},
	)
}

func (forCommentsDto[E, T]) IsApproved() *ast.BoolColumnProjection[T, **bool] {
	return ast.NewBoolColumnProjection(
		"comments",
		"is_approved",
		func(t *T) **bool {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.IsApproved
		},
	)
}

func (forCommentsDto[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"comments",
		"created_at",
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.CreatedAt
		},
	)
}

func (forCommentsDto[E, T]) TestDate() *ast.DateColumnProjection[T, **time.Time] {
	return ast.NewDateColumnProjection(
		"comments",
		"test_date",
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetCommentsDto()
			return &dto.TestDate
		},
	)
}

func (f forCommentsDto[E, T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
		f.PostId(),
		f.UserId(),
		f.Content(),
		f.IsApproved(),
		f.CreatedAt(),
		f.TestDate(),
	}
}

type Alias struct {
	*ast.Alias
}

func As(name string, columns ...ast.NamedExpression) *Alias {
	return &Alias{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias) Id() *ast.IntColumnProjection[models.CommentsDto, **int64] {
	column := Id()
	alias := ast.NewIntColumnProjection[models.CommentsDto, **int64](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) **int64 {
			return &c.Id
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	alias.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) PostId() *ast.IntColumnProjection[models.CommentsDto, *int64] {
	column := PostId()
	alias := ast.NewIntColumnProjection[models.CommentsDto, *int64](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) *int64 {
			return &c.PostId
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'post_id' in alias %s", a.Alias.Name()))
	alias.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) UserId() *ast.UUIDColumnProjection[models.CommentsDto, *uuid.UUID] {
	column := UserId()
	alias := ast.NewUUIDColumnProjection[models.CommentsDto, *uuid.UUID](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) *uuid.UUID {
			return &c.UserId
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'user_id' in alias %s", a.Alias.Name()))
	alias.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Content() *ast.StringColumnProjection[models.CommentsDto, *string] {
	column := Content()
	alias := ast.NewStringColumnProjection[models.CommentsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) *string {
			return &c.Content
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'content' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) IsApproved() *ast.BoolColumnProjection[models.CommentsDto, **bool] {
	column := IsApproved()
	alias := ast.NewBoolColumnProjection[models.CommentsDto, **bool](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) **bool {
			return &c.IsApproved
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_approved' in alias %s", a.Alias.Name()))
	alias.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) CreatedAt() *ast.TimestampColumnProjection[models.CommentsDto, **time.Time] {
	column := CreatedAt()
	alias := ast.NewTimestampColumnProjection[models.CommentsDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) **time.Time {
			return &c.CreatedAt
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.Alias.Name()))
	alias.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) TestDate() *ast.DateColumnProjection[models.CommentsDto, **time.Time] {
	column := TestDate()
	alias := ast.NewDateColumnProjection[models.CommentsDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(c *models.CommentsDto) **time.Time {
			return &c.TestDate
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'test_date' in alias %s", a.Alias.Name()))
	alias.DateColumnExpression = ast.NewDateColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
