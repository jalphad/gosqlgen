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

type forTagsDto[E models.ExportsTagsDto[T], T any] struct{}

func For[E models.ExportsTagsDto[T], T any]() forTagsDto[E, T] {
	return forTagsDto[E, T]{}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("tags")
}

func Id() *ast.IntColumnProjection[models.TagsDto, **int64] {
	return ast.NewIntColumnProjection(
		"tags",
		"id",
		func(t *models.TagsDto) **int64 {
			return &t.Id
		},
	)
}

func Name() *ast.StringColumnProjection[models.TagsDto, *string] {
	return ast.NewStringColumnProjection(
		"tags",
		"name",
		func(t *models.TagsDto) *string {
			return &t.Name
		},
	)
}

func Slug() *ast.StringColumnProjection[models.TagsDto, *string] {
	return ast.NewStringColumnProjection(
		"tags",
		"slug",
		func(t *models.TagsDto) *string {
			return &t.Slug
		},
	)
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

func (forTagsDto[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	return ast.NewIntColumnProjection(
		"tags",
		"id",
		func(t *T) **int64 {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Id
		},
	)
}

func (forTagsDto[E, T]) Name() *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		"tags",
		"name",
		func(t *T) *string {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Name
		},
	)
}

func (forTagsDto[E, T]) Slug() *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		"tags",
		"slug",
		func(t *T) *string {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Slug
		},
	)
}

func (f forTagsDto[E, T]) AllColumns() []ast.Projection[T] {
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

func (forTagsDto[E, T]) Posts(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultPostsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(t *T) *[]models.PostsDto {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Posts
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

func (a *Alias) Id() *ast.IntColumnProjection[models.TagsDto, **int64] {
	column := Id()
	alias := ast.NewIntColumnProjection[models.TagsDto, **int64](
		a.Alias.Name(),
		column.Name(),
		func(t *models.TagsDto) **int64 {
			return &t.Id
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

func (a *Alias) Name() *ast.StringColumnProjection[models.TagsDto, *string] {
	column := Name()
	alias := ast.NewStringColumnProjection[models.TagsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(t *models.TagsDto) *string {
			return &t.Name
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'name' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Slug() *ast.StringColumnProjection[models.TagsDto, *string] {
	column := Slug()
	alias := ast.NewStringColumnProjection[models.TagsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(t *models.TagsDto) *string {
			return &t.Slug
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'slug' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
