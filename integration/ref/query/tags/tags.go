package tags

import (
	"fmt"
	"slices"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

// NewQuery returns a query builder for tags
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.TagsDto] {
	return builder.NewKnownTableBuilder[models.TagsDto](pool, Table())
}

var Into = intoTagsDto{}

type intoTagsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("tags")
}

func Id() *ast.IntColumnProjection[models.TagsDto, **int64] {
	return newId("tags", func(t *models.TagsDto) **int64 {
		return &t.Id
	})
}

func Name() *ast.StringColumnProjection[models.TagsDto, *string] {
	return newName("tags", func(t *models.TagsDto) *string {
		return &t.Name
	})
}

func Slug() *ast.StringColumnProjection[models.TagsDto, *string] {
	return newSlug("tags", func(t *models.TagsDto) *string {
		return &t.Slug
	})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Name(),
		Slug(),
	}
}

func (intoTagsDto) Id() ast.Projection[models.TagsDto] {
	return Id()
}

func (intoTagsDto) Name() ast.Projection[models.TagsDto] {
	return Name()
}

func (intoTagsDto) Slug() ast.Projection[models.TagsDto] {
	return Slug()
}

func (intoTagsDto) AllColumns() []ast.Projection[models.TagsDto] {
	return []ast.Projection[models.TagsDto]{
		Into.Id(),
		Into.Name(),
		Into.Slug(),
	}
}

type forTagsDto[T any] struct {
	dest func(*T) *models.TagsDto
}

func For[T any](dest func(*T) *models.TagsDto) forTagsDto[T] {
	return forTagsDto[T]{dest: dest}
}

func (f forTagsDto[T]) Id() *ast.IntColumnProjection[T, **int64] {
	return newId("tags", func(t *T) **int64 {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forTagsDto[T]) Name() *ast.StringColumnProjection[T, *string] {
	return newName("tags", func(t *T) *string {
		dto := f.dest(t)
		return &dto.Name
	})
}

func (f forTagsDto[T]) Slug() *ast.StringColumnProjection[T, *string] {
	return newSlug("tags", func(t *T) *string {
		dto := f.dest(t)
		return &dto.Slug
	})
}

func (f forTagsDto[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
		f.Name(),
		f.Slug(),
	}
}

func (intoTagsDto) Posts(columns ...ast.NamedExpression) ast.Projection[models.TagsDto] {
	if len(columns) == 0 {
		columns = defaultPostsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(t *models.TagsDto) *[]models.PostsDto {
			return &t.Posts
		},
	)
}

func defaultPostsColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("posts", "id"),
		ast.NewUUIDColumnExpression("posts", "user_id"),
		ast.NewStringColumnExpression("posts", "title"),
		ast.NewStringColumnExpression("posts", "content"),
		ast.NewStringColumnExpression("posts", "status"),
		ast.NewTimestampColumnExpression("posts", "published_at"),
		ast.NewIntColumnExpression("posts", "view_count"),
		ast.NewTimestampColumnExpression("posts", "created_at"),
		ast.NewTimestampColumnExpression("posts", "updated_at"),
	}
}

func (f forTagsDto[T]) Posts(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultPostsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(t *T) *[]models.PostsDto {
			dto := f.dest(t)
			return &dto.Posts
		},
	)
}

type Alias[T any] struct {
	*ast.Alias
	forTagsDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.TagsDto] {
	return Alias[models.TagsDto]{
		Alias: ast.NewAlias(name, columns...),
		forTagsDto: forTagsDto[models.TagsDto]{
			dest: func(t *models.TagsDto) *models.TagsDto {
				return t
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

func (a Alias[T]) Name() *ast.StringColumnProjection[T, *string] {
	projection := newName(a.Alias.Name(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Name
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'name' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) Slug() *ast.StringColumnProjection[T, *string] {
	projection := newSlug(a.Alias.Name(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Slug
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'slug' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
		a.Name(),
		a.Slug(),
	}
}

func AliasFor[T any](alias *ast.Alias, dest func(*T) *models.TagsDto) Alias[T] {
	return Alias[T]{
		Alias:      alias,
		forTagsDto: forTagsDto[T]{dest: dest},
	}
}

func newId[T any](table string, ref func(*T) **int64) *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		table,
		"id",
		ref,
	)
}

func newName[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"name",
		ref,
	)
}

func newSlug[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"slug",
		ref,
	)
}
