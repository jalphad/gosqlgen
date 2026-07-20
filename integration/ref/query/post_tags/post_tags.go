package post_tags

import (
	"fmt"
	"slices"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
)

// NewQuery returns a query builder for "public"."post_tags"
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.PostTagsDto] {
	return builder.NewKnownTableBuilder[models.PostTagsDto](pool, Table())
}

var Into = intoPostTagsDto{}

type intoPostTagsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource(`"public"."post_tags"`)
}

func PostId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return newPostId(`"public"."post_tags"`, func(p *models.PostTagsDto) *int64 {
		return &p.PostId
	})
}

func TagId() *ast.IntColumnProjection[models.PostTagsDto, *int64] {
	return newTagId(`"public"."post_tags"`, func(p *models.PostTagsDto) *int64 {
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

type forPostTagsDto[T any] struct {
	dest func(*T) *models.PostTagsDto
}

func For[T any](dest func(*T) *models.PostTagsDto) forPostTagsDto[T] {
	return forPostTagsDto[T]{dest: dest}
}

func (f forPostTagsDto[T]) PostId() *ast.IntColumnProjection[T, *int64] {
	return newPostId(`"public"."post_tags"`, func(t *T) *int64 {
		dto := f.dest(t)
		return &dto.PostId
	})
}

func (f forPostTagsDto[T]) TagId() *ast.IntColumnProjection[T, *int64] {
	return newTagId(`"public"."post_tags"`, func(t *T) *int64 {
		dto := f.dest(t)
		return &dto.TagId
	})
}

func (f forPostTagsDto[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.PostId(),
		f.TagId(),
	}
}

type Alias[T any] struct {
	*ast.TableAlias
	forPostTagsDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.PostTagsDto] {
	return Alias[models.PostTagsDto]{
		TableAlias: ast.NewTableAlias(name, columns...),
		forPostTagsDto: forPostTagsDto[models.PostTagsDto]{
			dest: func(p *models.PostTagsDto) *models.PostTagsDto {
				return p
			},
		},
	}
}

func (a Alias[T]) PostId() *ast.IntColumnProjection[T, *int64] {
	projection := newPostId(a.TableAlias.Name(), func(t *T) *int64 {
		dto := a.dest(t)
		return &dto.PostId
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'post_id' in alias %s", a.TableAlias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) TagId() *ast.IntColumnProjection[T, *int64] {
	projection := newTagId(a.TableAlias.Name(), func(t *T) *int64 {
		dto := a.dest(t)
		return &dto.TagId
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'tag_id' in alias %s", a.TableAlias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.PostId(),
		a.TagId(),
	}
}

func AliasFor[T any](alias *ast.TableAlias, dest func(*T) *models.PostTagsDto) Alias[T] {
	return Alias[T]{
		TableAlias:     alias,
		forPostTagsDto: forPostTagsDto[T]{dest: dest},
	}
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
