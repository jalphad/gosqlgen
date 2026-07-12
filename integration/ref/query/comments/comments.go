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

func Table() *ast.TableSource {
	return ast.NewTableSource("comments")
}

func Id() *ast.IntColumnProjection[models.CommentsDto, **int64] {
	return newId("comments", func(c *models.CommentsDto) **int64 {
		return &c.Id
	})
}

func PostId() *ast.IntColumnProjection[models.CommentsDto, *int64] {
	return newPostId("comments", func(c *models.CommentsDto) *int64 {
		return &c.PostId
	})
}

func UserId() *ast.UUIDColumnProjection[models.CommentsDto, *uuid.UUID] {
	return newUserId("comments", func(c *models.CommentsDto) *uuid.UUID {
		return &c.UserId
	})
}

func Content() *ast.StringColumnProjection[models.CommentsDto, *string] {
	return newContent("comments", func(c *models.CommentsDto) *string {
		return &c.Content
	})
}

func IsApproved() *ast.BoolColumnProjection[models.CommentsDto, **bool] {
	return newIsApproved("comments", func(c *models.CommentsDto) **bool {
		return &c.IsApproved
	})
}

func CreatedAt() *ast.TimestampColumnProjection[models.CommentsDto, **time.Time] {
	return newCreatedAt("comments", func(c *models.CommentsDto) **time.Time {
		return &c.CreatedAt
	})
}

func TestDate() *ast.DateColumnProjection[models.CommentsDto, **time.Time] {
	return newTestDate("comments", func(c *models.CommentsDto) **time.Time {
		return &c.TestDate
	})
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

type forCommentsDto[E models.ExportsCommentsDto[T], T any] struct{}

func For[E models.ExportsCommentsDto[T], T any]() forCommentsDto[E, T] {
	return forCommentsDto[E, T]{}
}

func (f forCommentsDto[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	return newId("comments", func(t *T) **int64 {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.Id
	})
}

func (f forCommentsDto[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return newPostId("comments", func(t *T) *int64 {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.PostId
	})
}

func (f forCommentsDto[E, T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	return newUserId("comments", func(t *T) *uuid.UUID {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.UserId
	})
}

func (f forCommentsDto[E, T]) Content() *ast.StringColumnProjection[T, *string] {
	return newContent("comments", func(t *T) *string {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.Content
	})
}

func (f forCommentsDto[E, T]) IsApproved() *ast.BoolColumnProjection[T, **bool] {
	return newIsApproved("comments", func(t *T) **bool {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.IsApproved
	})
}

func (f forCommentsDto[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newCreatedAt("comments", func(t *T) **time.Time {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.CreatedAt
	})
}

func (f forCommentsDto[E, T]) TestDate() *ast.DateColumnProjection[T, **time.Time] {
	return newTestDate("comments", func(t *T) **time.Time {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.TestDate
	})
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

type Alias[E models.ExportsCommentsDto[T], T any] struct {
	*ast.Alias
	forCommentsDto[E, T]
}

func As[E models.ExportsCommentsDto[T], T any](name string, columns ...ast.NamedExpression) Alias[E, T] {
	return Alias[E, T]{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	projection := newId(a.Alias.Name(), func(t *T) **int64 {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.Id
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	projection := newPostId(a.Alias.Name(), func(t *T) *int64 {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.PostId
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'post_id' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	projection := newUserId(a.Alias.Name(), func(t *T) *uuid.UUID {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.UserId
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'user_id' in alias %s", a.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) Content() *ast.StringColumnProjection[T, *string] {
	projection := newContent(a.Alias.Name(), func(t *T) *string {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.Content
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'content' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) IsApproved() *ast.BoolColumnProjection[T, **bool] {
	projection := newIsApproved(a.Alias.Name(), func(t *T) **bool {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.IsApproved
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_approved' in alias %s", a.Alias.Name()))
	projection.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newCreatedAt(a.Alias.Name(), func(t *T) **time.Time {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.CreatedAt
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) TestDate() *ast.DateColumnProjection[T, **time.Time] {
	projection := newTestDate(a.Alias.Name(), func(t *T) **time.Time {
		e := E(t)
		dto := e.GetCommentsDto()
		return &dto.TestDate
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'test_date' in alias %s", a.Alias.Name()))
	projection.DateColumnExpression = ast.NewDateColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func newId[T any](table string, ref func(*T) **int64) *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		table,
		"id",
		ref,
	)
}

func newPostId[T any](table string, ref func(*T) *int64) *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		table,
		"post_id",
		ref,
	)
}

func newUserId[T any](table string, ref func(*T) *uuid.UUID) *ast.UUIDColumnProjection[T, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		table,
		"user_id",
		ref,
	)
}

func newContent[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"content",
		ref,
	)
}

func newIsApproved[T any](table string, ref func(*T) **bool) *ast.BoolColumnProjection[T, **bool] {
	return ast.NewBoolColumnProjection(
		table,
		"is_approved",
		ref,
	)
}

func newCreatedAt[T any](table string, ref func(*T) **time.Time) *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		table,
		"created_at",
		ref,
	)
}

func newTestDate[T any](table string, ref func(*T) **time.Time) *ast.DateColumnProjection[T, **time.Time] {
	return ast.NewDateColumnProjection(
		table,
		"test_date",
		ref,
	)
}
