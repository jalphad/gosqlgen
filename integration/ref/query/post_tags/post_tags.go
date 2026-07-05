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

type forPostTagsDto[E models.ExportsPostTagsDto[T], T any] struct{}

func For[E models.ExportsPostTagsDto[T], T any]() forPostTagsDto[E, T] {
	return forPostTagsDto[E, T]{}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("post_tags")
}

func PostId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return ast.NewIntColumnProjection(
		"post_tags",
		"post_id",
		func(p *models.PostTagsDto) *int64 {
			return &p.PostId
		},
	)
}

func TagId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return ast.NewIntColumnProjection(
		"post_tags",
		"tag_id",
		func(p *models.PostTagsDto) *int64 {
			return &p.TagId
		},
	)
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

func (forPostTagsDto[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		"post_tags",
		"post_id",
		func(t *T) *int64 {
			e := E(t)
			dto := e.GetPostTagsDto()
			return &dto.PostId
		},
	)
}

func (forPostTagsDto[E, T]) TagId() *ast.IntColumnProjection[T, *int64] {
	return ast.NewIntColumnProjection(
		"post_tags",
		"tag_id",
		func(t *T) *int64 {
			e := E(t)
			dto := e.GetPostTagsDto()
			return &dto.TagId
		},
	)
}

func (f forPostTagsDto[E, T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.PostId(),
		f.TagId(),
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

func (a *Alias) PostId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	column := PostId()
	alias := ast.NewIntColumnProjection[models.PostTagsDto, *int64](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostTagsDto) *int64 {
			return &p.PostId
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

func (a *Alias) TagId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	column := TagId()
	alias := ast.NewIntColumnProjection[models.PostTagsDto, *int64](
		a.Alias.Name(),
		column.Name(),
		func(p *models.PostTagsDto) *int64 {
			return &p.TagId
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'tag_id' in alias %s", a.Alias.Name()))
	alias.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
