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

type forTagsDto[E models.ExportsTagsDto[T], T any] struct {
	alias *Alias
	err   error
}

// TODO(go1.27): if type parameters on methods are available, consider
// supporting alias.For[*resultRow]() as the primary alias API.
func For[E models.ExportsTagsDto[T], T any](alias ...*Alias) forTagsDto[E, T] {
	switch len(alias) {
	case 0:
		return forTagsDto[E, T]{}
	case 1:
		if alias[0] == nil || alias[0].Alias == nil {
			return forTagsDto[E, T]{err: fmt.Errorf("tags.For requires a non-nil alias")}
		}
		return forTagsDto[E, T]{alias: alias[0]}
	default:
		return forTagsDto[E, T]{err: fmt.Errorf("tags.For accepts at most one alias")}
	}
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

func (f forTagsDto[E, T]) Id() *ast.IntColumnProjection[T, **int64] {
	column := Id()
	projection := ast.NewIntColumnProjection[T, **int64](
		f.tableName(),
		column.Name(),
		func(t *T) **int64 {
			e := E(t)
			dto := e.GetTagsDto()
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

func (f forTagsDto[E, T]) Name() *ast.StringColumnProjection[T, *string] {
	column := Name()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Name
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
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'name' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forTagsDto[E, T]) Slug() *ast.StringColumnProjection[T, *string] {
	column := Slug()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetTagsDto()
			return &dto.Slug
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
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'slug' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forTagsDto[E, T]) tableName() string {
	if f.alias == nil || f.alias.Alias == nil {
		return "tags"
	}
	return f.alias.Alias.Name()
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
