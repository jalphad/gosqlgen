package post_tags

import (
	"fmt"
	"slices"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
)

// NewQuery returns a query builder for post_tags
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.PostTagsDto] {
	return builder.NewKnownTableBuilder[models.PostTagsDto](pool, Table())
}

var Into = intoPostTagsDto{}

type intoPostTagsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("post_tags")
}

func PostId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return newPostId("post_tags", func(p *models.PostTagsDto) *int64 {
		return &p.PostId
	})
}

func TagId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return newTagId("post_tags", func(p *models.PostTagsDto) *int64 {
		return &p.TagId
	})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		PostId(),
		TagId(),
	}
}

func (intoPostTagsDto) PostId() ast.Projection[models.PostTagsDto] {
	return PostId()
}

func (intoPostTagsDto) TagId() ast.Projection[models.PostTagsDto] {
	return TagId()
}

func (intoPostTagsDto) AllColumns() []ast.Projection[models.PostTagsDto] {
	return []ast.Projection[models.PostTagsDto]{
		Into.PostId(),
		Into.TagId(),
	}
}

type forPostTagsDto[E models.ExportsPostTagsDto[T], T any] struct{}

func For[E models.ExportsPostTagsDto[T], T any]() forPostTagsDto[E, T] {
	return forPostTagsDto[E, T]{}
}

func (f forPostTagsDto[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return newPostId("post_tags", func(t *T) *int64 {
		e := E(t)
		dto := e.GetPostTagsDto()
		return &dto.PostId
	})
}

func (f forPostTagsDto[E, T]) TagId() *ast.IntColumnProjection[T, *int64] {
	return newTagId("post_tags", func(t *T) *int64 {
		e := E(t)
		dto := e.GetPostTagsDto()
		return &dto.TagId
	})
}

func (f forPostTagsDto[E, T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.PostId(),
		f.TagId(),
	}
}

type Alias[E models.ExportsPostTagsDto[T], T any] struct {
	*ast.Alias
	forPostTagsDto[E, T]
}

func As[E models.ExportsPostTagsDto[T], T any](name string, columns ...ast.NamedExpression) Alias[E, T] {
	return Alias[E, T]{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	projection := newPostId(a.Alias.Name(), func(t *T) *int64 {
		e := E(t)
		dto := e.GetPostTagsDto()
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

func (a *Alias[E, T]) TagId() *ast.IntColumnProjection[T, *int64] {
	projection := newTagId(a.Alias.Name(), func(t *T) *int64 {
		e := E(t)
		dto := e.GetPostTagsDto()
		return &dto.TagId
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'tag_id' in alias %s", a.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func newPostId[T any](table string, ref func(*T) *int64) *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		table,
		"post_id",
		ref,
	)
}

func newTagId[T any](table string, ref func(*T) *int64) *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		table,
		"tag_id",
		ref,
	)
}
