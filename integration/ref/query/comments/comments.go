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

type forCommentsDto[T any] struct {
	dest func(*T) *models.CommentsDto
}

func For[T any](dest func(*T) *models.CommentsDto) forCommentsDto[T] {
	return forCommentsDto[T]{dest: dest}
}

func (f forCommentsDto[T]) Id() *ast.IntColumnProjection[T, **int64] {
	return newId("comments", func(t *T) **int64 {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forCommentsDto[T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return newPostId("comments", func(t *T) *int64 {
		dto := f.dest(t)
		return &dto.PostId
	})
}

func (f forCommentsDto[T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	return newUserId("comments", func(t *T) *uuid.UUID {
		dto := f.dest(t)
		return &dto.UserId
	})
}

func (f forCommentsDto[T]) Content() *ast.StringColumnProjection[T, *string] {
	return newContent("comments", func(t *T) *string {
		dto := f.dest(t)
		return &dto.Content
	})
}

func (f forCommentsDto[T]) IsApproved() *ast.BoolColumnProjection[T, **bool] {
	return newIsApproved("comments", func(t *T) **bool {
		dto := f.dest(t)
		return &dto.IsApproved
	})
}

func (f forCommentsDto[T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newCreatedAt("comments", func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.CreatedAt
	})
}

func (f forCommentsDto[T]) TestDate() *ast.DateColumnProjection[T, **time.Time] {
	return newTestDate("comments", func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.TestDate
	})
}

func (f forCommentsDto[T]) AllColumns() []ast.Projection[T] {
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

type Alias[T any] struct {
	*ast.Alias
	forCommentsDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.CommentsDto] {
	return Alias[models.CommentsDto]{
		Alias: ast.NewAlias(name, columns...),
		forCommentsDto: forCommentsDto[models.CommentsDto]{
			dest: func(c *models.CommentsDto) *models.CommentsDto {
				return c
			},
		},
	}
}

func (a Alias[T]) Id() *ast.IntColumnProjection[T, **int64] {
	projection := newId(a.Alias.Name(), func(t *T) **int64 {
		dto := a.dest(t)
		return &dto.Id
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) PostId() *ast.IntColumnProjection[T, *int64] {
	projection := newPostId(a.Alias.Name(), func(t *T) *int64 {
		dto := a.dest(t)
		return &dto.PostId
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'post_id' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	projection := newUserId(a.Alias.Name(), func(t *T) *uuid.UUID {
		dto := a.dest(t)
		return &dto.UserId
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'user_id' in alias %s", a.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) Content() *ast.StringColumnProjection[T, *string] {
	projection := newContent(a.Alias.Name(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Content
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'content' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) IsApproved() *ast.BoolColumnProjection[T, **bool] {
	projection := newIsApproved(a.Alias.Name(), func(t *T) **bool {
		dto := a.dest(t)
		return &dto.IsApproved
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_approved' in alias %s", a.Alias.Name()))
	projection.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newCreatedAt(a.Alias.Name(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.CreatedAt
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) TestDate() *ast.DateColumnProjection[T, **time.Time] {
	projection := newTestDate(a.Alias.Name(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.TestDate
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'test_date' in alias %s", a.Alias.Name()))
	projection.DateColumnExpression = ast.NewDateColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
		a.PostId(),
		a.UserId(),
		a.Content(),
		a.IsApproved(),
		a.CreatedAt(),
		a.TestDate(),
	}
}

func AliasFor[T any](alias *ast.Alias, dest func(*T) *models.CommentsDto) Alias[T] {
	return Alias[T]{
		Alias:          alias,
		forCommentsDto: forCommentsDto[T]{dest: dest},
	}
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
