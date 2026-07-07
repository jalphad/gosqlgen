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

type forPostsDto[E models.ExportsPostsDto[T], T any] struct {
	alias *Alias
	err   error
}

// TODO(go1.27): if type parameters on methods are available, consider
// supporting alias.For[*resultRow]() as the primary alias API.
func For[E models.ExportsPostsDto[T], T any](alias ...*Alias) forPostsDto[E, T] {
	switch len(alias) {
	case 0:
		return forPostsDto[E, T]{}
	case 1:
		if alias[0] == nil || alias[0].Alias == nil {
			return forPostsDto[E, T]{err: fmt.Errorf("posts.For requires a non-nil alias")}
		}
		return forPostsDto[E, T]{alias: alias[0]}
	default:
		return forPostsDto[E, T]{err: fmt.Errorf("posts.For accepts at most one alias")}
	}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("posts")
}

func Id() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return ast.NewIntColumnProjection(
		"posts",
		"id",
		func(p *models.PostsDto) **int64 {
			return &p.Id
		},
	)
}

func UserId() *ast.UUIDColumnProjection[models.PostsDto, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"posts",
		"user_id",
		func(p *models.PostsDto) *uuid.UUID {
			return &p.UserId
		},
	)
}

func Title() *ast.StringColumnProjection[models.PostsDto, *string] {
	return ast.NewStringColumnProjection(
		"posts",
		"title",
		func(p *models.PostsDto) *string {
			return &p.Title
		},
	)
}

func Content() *ast.StringColumnProjection[models.PostsDto, **string] {
	return ast.NewStringColumnProjection(
		"posts",
		"content",
		func(p *models.PostsDto) **string {
			return &p.Content
		},
	)
}

func Status() *ast.StringColumnProjection[models.PostsDto, **string] {
	return ast.NewStringColumnProjection(
		"posts",
		"status",
		func(p *models.PostsDto) **string {
			return &p.Status
		},
	)
}

func PublishedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"posts",
		"published_at",
		func(p *models.PostsDto) **time.Time {
			return &p.PublishedAt
		},
	)
}

func ViewCount() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return ast.NewIntColumnProjection(
		"posts",
		"view_count",
		func(p *models.PostsDto) **int64 {
			return &p.ViewCount
		},
	)
}

func CreatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"posts",
		"created_at",
		func(p *models.PostsDto) **time.Time {
			return &p.CreatedAt
		},
	)
}

func UpdatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"posts",
		"updated_at",
		func(p *models.PostsDto) **time.Time {
			return &p.UpdatedAt
		},
	)
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

func (f forPostsDto[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	column := Id()
	projection := ast.NewIntColumnProjection[T, **int64](
		f.tableName(),
		column.Name(),
		func(t *T) **int64 {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.Id
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", f.alias.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) UserId() *ast.UUIDColumnProjection[T, *uuid.UUID] {
	column := UserId()
	projection := ast.NewUUIDColumnProjection[T, *uuid.UUID](
		f.tableName(),
		column.Name(),
		func(t *T) *uuid.UUID {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.UserId
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'user_id' in alias %s", f.alias.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) Title() *ast.StringColumnProjection[T, *string] {
	column := Title()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.Title
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'title' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) Content() *ast.StringColumnProjection[T, **string] {
	column := Content()
	projection := ast.NewStringColumnProjection[T, **string](
		f.tableName(),
		column.Name(),
		func(t *T) **string {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.Content
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'content' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) Status() *ast.StringColumnProjection[T, **string] {
	column := Status()
	projection := ast.NewStringColumnProjection[T, **string](
		f.tableName(),
		column.Name(),
		func(t *T) **string {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.Status
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'status' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) PublishedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	column := PublishedAt()
	projection := ast.NewTimestampColumnProjection[T, **time.Time](
		f.tableName(),
		column.Name(),
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.PublishedAt
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'published_at' in alias %s", f.alias.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) ViewCount() *ast.IntColumnProjection[T, **int64] {
	column := ViewCount()
	projection := ast.NewIntColumnProjection[T, **int64](
		f.tableName(),
		column.Name(),
		func(t *T) **int64 {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.ViewCount
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'view_count' in alias %s", f.alias.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	column := CreatedAt()
	projection := ast.NewTimestampColumnProjection[T, **time.Time](
		f.tableName(),
		column.Name(),
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.CreatedAt
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", f.alias.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	column := UpdatedAt()
	projection := ast.NewTimestampColumnProjection[T, **time.Time](
		f.tableName(),
		column.Name(),
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.UpdatedAt
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", f.alias.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostsDto[E, T]) tableName() string {
	if f.alias == nil || f.alias.Alias == nil {
		return "posts"
	}
	return f.alias.Alias.Name()
}

func (f forPostsDto[E, T]) AllColumns() []ast.Projection[T] {
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

func (forPostsDto[E, T]) CommentsByPostId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultCommentsByPostIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbypostid", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(t *T) *[]models.CommentsDto {
			e := E(t)
			dto := e.GetPostsDto()
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

func (forPostsDto[E, T]) Tags(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultTagsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("tags", expr.IsNotNull(ast.NewIntColumnExpression("tags", "id")), columns...),
		func(t *T) *[]models.TagsDto {
			e := E(t)
			dto := e.GetPostsDto()
			return &dto.Tags
		},
	)
}

type Alias struct {
	*ast.Alias
}

func As(name string, columns ...ast.NamedExpression) *Alias {
	return &Alias{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias) Id() *ast.IntColumnProjection[models.PostsDto, **int64] {
	column := Id()
	alias := ast.NewIntColumnProjection[models.PostsDto, **int64](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **int64 {
			return &p.Id
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

func (a *Alias) UserId() *ast.UUIDColumnProjection[models.PostsDto, *uuid.UUID] {
	column := UserId()
	alias := ast.NewUUIDColumnProjection[models.PostsDto, *uuid.UUID](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) *uuid.UUID {
			return &p.UserId
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

func (a *Alias) Title() *ast.StringColumnProjection[models.PostsDto, *string] {
	column := Title()
	alias := ast.NewStringColumnProjection[models.PostsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) *string {
			return &p.Title
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'title' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Content() *ast.StringColumnProjection[models.PostsDto, **string] {
	column := Content()
	alias := ast.NewStringColumnProjection[models.PostsDto, **string](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **string {
			return &p.Content
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

func (a *Alias) Status() *ast.StringColumnProjection[models.PostsDto, **string] {
	column := Status()
	alias := ast.NewStringColumnProjection[models.PostsDto, **string](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **string {
			return &p.Status
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'status' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) PublishedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	column := PublishedAt()
	alias := ast.NewTimestampColumnProjection[models.PostsDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **time.Time {
			return &p.PublishedAt
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'published_at' in alias %s", a.Alias.Name()))
	alias.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) ViewCount() *ast.IntColumnProjection[models.PostsDto, **int64] {
	column := ViewCount()
	alias := ast.NewIntColumnProjection[models.PostsDto, **int64](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **int64 {
			return &p.ViewCount
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'view_count' in alias %s", a.Alias.Name()))
	alias.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) CreatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	column := CreatedAt()
	alias := ast.NewTimestampColumnProjection[models.PostsDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **time.Time {
			return &p.CreatedAt
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

func (a *Alias) UpdatedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	column := UpdatedAt()
	alias := ast.NewTimestampColumnProjection[models.PostsDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostsDto) **time.Time {
			return &p.UpdatedAt
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", a.Alias.Name()))
	alias.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
