package posts

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

// NewQuery returns a query builder for posts
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.PostsDto] {
	return builder.NewKnownTableBuilder[models.PostsDto](pool, Table())
}

var Into = intoPostsDto{}

type intoPostsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("posts")
}

func Id() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return newId("posts", func(p *models.PostsDto) **int64 {
		return &p.Id
	})
}

func UserId() *ast.UUIDColumnProjection[models.PostsDto, *uuid.UUID] {
	return newUserId("posts", func(p *models.PostsDto) *uuid.UUID {
		return &p.UserId
	})
}

func Title() *ast.StringColumnProjection[models.PostsDto, *string] {
	return newTitle("posts", func(p *models.PostsDto) *string {
		return &p.Title
	})
}

func Content() *ast.StringColumnProjection[models.PostsDto, **string] {
	return newContent("posts", func(p *models.PostsDto) **string {
		return &p.Content
	})
}

func Status() *ast.StringColumnProjection[models.PostsDto, **string] {
	return newStatus("posts", func(p *models.PostsDto) **string {
		return &p.Status
	})
}

func PublishedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return newPublishedAt("posts", func(p *models.PostsDto) **time.Time {
		return &p.PublishedAt
	})
}

func ViewCount() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return newViewCount("posts", func(p *models.PostsDto) **int64 {
		return &p.ViewCount
	})
}

func CreatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return newCreatedAt("posts", func(p *models.PostsDto) **time.Time {
		return &p.CreatedAt
	})
}

func UpdatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return newUpdatedAt("posts", func(p *models.PostsDto) **time.Time {
		return &p.UpdatedAt
	})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		UserId(),
		Title(),
		Content(),
		Status(),
		PublishedAt(),
		ViewCount(),
		CreatedAt(),
		UpdatedAt(),
	}
}

func (intoPostsDto) Id() ast.Projection[models.PostsDto] {
	return Id()
}

func (intoPostsDto) UserId() ast.Projection[models.PostsDto] {
	return UserId()
}

func (intoPostsDto) Title() ast.Projection[models.PostsDto] {
	return Title()
}

func (intoPostsDto) Content() ast.Projection[models.PostsDto] {
	return Content()
}

func (intoPostsDto) Status() ast.Projection[models.PostsDto] {
	return Status()
}

func (intoPostsDto) PublishedAt() ast.Projection[models.PostsDto] {
	return PublishedAt()
}

func (intoPostsDto) ViewCount() ast.Projection[models.PostsDto] {
	return ViewCount()
}

func (intoPostsDto) CreatedAt() ast.Projection[models.PostsDto] {
	return CreatedAt()
}

func (intoPostsDto) UpdatedAt() ast.Projection[models.PostsDto] {
	return UpdatedAt()
}

func (intoPostsDto) AllColumns() []ast.Projection[models.PostsDto] {
	return []ast.Projection[models.PostsDto]{
		Into.Id(),
		Into.UserId(),
		Into.Title(),
		Into.Content(),
		Into.Status(),
		Into.PublishedAt(),
		Into.ViewCount(),
		Into.CreatedAt(),
		Into.UpdatedAt(),
	}
}

type forPostsDto[T any] struct {
	dest func(*T) *models.PostsDto
}

func For[T any](dest func(*T) *models.PostsDto) forPostsDto[T] {
	return forPostsDto[T]{dest: dest}
}

func (f forPostsDto[T]) Id() *ast.IntColumnProjection[T, **int64] {
	return newId("posts", func(t *T) **int64 {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forPostsDto[T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	return newUserId("posts", func(t *T) *uuid.UUID {
		dto := f.dest(t)
		return &dto.UserId
	})
}

func (f forPostsDto[T]) Title() *ast.StringColumnProjection[T, *string] {
	return newTitle("posts", func(t *T) *string {
		dto := f.dest(t)
		return &dto.Title
	})
}

func (f forPostsDto[T]) Content() *ast.StringColumnProjection[T, **string] {
	return newContent("posts", func(t *T) **string {
		dto := f.dest(t)
		return &dto.Content
	})
}

func (f forPostsDto[T]) Status() *ast.StringColumnProjection[T, **string] {
	return newStatus("posts", func(t *T) **string {
		dto := f.dest(t)
		return &dto.Status
	})
}

func (f forPostsDto[T]) PublishedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newPublishedAt("posts", func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.PublishedAt
	})
}

func (f forPostsDto[T]) ViewCount() *ast.IntColumnProjection[T, **int64] {
	return newViewCount("posts", func(t *T) **int64 {
		dto := f.dest(t)
		return &dto.ViewCount
	})
}

func (f forPostsDto[T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newCreatedAt("posts", func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.CreatedAt
	})
}

func (f forPostsDto[T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newUpdatedAt("posts", func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.UpdatedAt
	})
}

func (f forPostsDto[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
		f.UserId(),
		f.Title(),
		f.Content(),
		f.Status(),
		f.PublishedAt(),
		f.ViewCount(),
		f.CreatedAt(),
		f.UpdatedAt(),
	}
}

func (intoPostsDto) CommentsByPostId(columns ...ast.NamedExpression) ast.Projection[models.PostsDto] {
	if len(columns) == 0 {
		columns = defaultCommentsByPostIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbypostid", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(p *models.PostsDto) *[]models.CommentsDto {
			return &p.CommentsByPostId
		},
	)
}

func defaultCommentsByPostIdColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("comments", "id"),
		ast.NewIntColumnExpression("comments", "post_id"),
		ast.NewUUIDColumnExpression("comments", "user_id"),
		ast.NewStringColumnExpression("comments", "content"),
		ast.NewBoolColumnExpression("comments", "is_approved"),
		ast.NewTimestampColumnExpression("comments", "created_at"),
		ast.NewDateColumnExpression("comments", "test_date"),
	}
}

func (f forPostsDto[T]) CommentsByPostId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultCommentsByPostIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbypostid", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(t *T) *[]models.CommentsDto {
			dto := f.dest(t)
			return &dto.CommentsByPostId
		},
	)
}

func (intoPostsDto) Tags(columns ...ast.NamedExpression) ast.Projection[models.PostsDto] {
	if len(columns) == 0 {
		columns = defaultTagsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("tags", expr.IsNotNull(ast.NewIntColumnExpression("tags", "id")), columns...),
		func(p *models.PostsDto) *[]models.TagsDto {
			return &p.Tags
		},
	)
}

func defaultTagsColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("tags", "id"),
		ast.NewStringColumnExpression("tags", "name"),
		ast.NewStringColumnExpression("tags", "slug"),
	}
}

func (f forPostsDto[T]) Tags(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultTagsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("tags", expr.IsNotNull(ast.NewIntColumnExpression("tags", "id")), columns...),
		func(t *T) *[]models.TagsDto {
			dto := f.dest(t)
			return &dto.Tags
		},
	)
}

type Alias[T any] struct {
	*ast.Alias
	forPostsDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.PostsDto] {
	return Alias[models.PostsDto]{
		Alias: ast.NewAlias(name, columns...),
		forPostsDto: forPostsDto[models.PostsDto]{
			dest: func(p *models.PostsDto) *models.PostsDto {
				return p
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

func (a Alias[T]) Title() *ast.StringColumnProjection[T, *string] {
	projection := newTitle(a.Alias.Name(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Title
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'title' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) Content() *ast.StringColumnProjection[T, **string] {
	projection := newContent(a.Alias.Name(), func(t *T) **string {
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

func (a Alias[T]) Status() *ast.StringColumnProjection[T, **string] {
	projection := newStatus(a.Alias.Name(), func(t *T) **string {
		dto := a.dest(t)
		return &dto.Status
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'status' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) PublishedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newPublishedAt(a.Alias.Name(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.PublishedAt
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'published_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) ViewCount() *ast.IntColumnProjection[T, **int64] {
	projection := newViewCount(a.Alias.Name(), func(t *T) **int64 {
		dto := a.dest(t)
		return &dto.ViewCount
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'view_count' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
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

func (a Alias[T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newUpdatedAt(a.Alias.Name(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.UpdatedAt
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
		a.UserId(),
		a.Title(),
		a.Content(),
		a.Status(),
		a.PublishedAt(),
		a.ViewCount(),
		a.CreatedAt(),
		a.UpdatedAt(),
	}
}

func AliasFor[T any](alias *ast.Alias, dest func(*T) *models.PostsDto) Alias[T] {
	return Alias[T]{
		Alias:       alias,
		forPostsDto: forPostsDto[T]{dest: dest},
	}
}

func newId[T any](table string, ref func(*T) **int64) *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		table,
		"id",
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

func newTitle[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"title",
		ref,
	)
}

func newContent[T any](table string, ref func(*T) **string) *ast.StringColumnProjection[T, **string] {
	return ast.NewStringColumnProjection(
		table,
		"content",
		ref,
	)
}

func newStatus[T any](table string, ref func(*T) **string) *ast.StringColumnProjection[T, **string] {
	return ast.NewStringColumnProjection(
		table,
		"status",
		ref,
	)
}

func newPublishedAt[T any](table string, ref func(*T) **time.Time) *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		table,
		"published_at",
		ref,
	)
}

func newViewCount[T any](table string, ref func(*T) **int64) *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		table,
		"view_count",
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

func newUpdatedAt[T any](table string, ref func(*T) **time.Time) *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		table,
		"updated_at",
		ref,
	)
}
