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

type forPostTagsDto[E models.ExportsPostTagsDto[T], T any] struct {
	alias *Alias
	err   error
}

// TODO(go1.27): if type parameters on methods are available, consider
// supporting alias.For[*resultRow]() as the primary alias API.
func For[E models.ExportsPostTagsDto[T], T any](alias ...*Alias) forPostTagsDto[E, T] {
	switch len(alias) {
	case 0:
		return forPostTagsDto[E, T]{}
	case 1:
		if alias[0] == nil || alias[0].Alias == nil {
			return forPostTagsDto[E, T]{err: fmt.Errorf("post_tags.For requires a non-nil alias")}
		}
		return forPostTagsDto[E, T]{alias: alias[0]}
	default:
		return forPostTagsDto[E, T]{err: fmt.Errorf("post_tags.For accepts at most one alias")}
	}
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

func (f forPostTagsDto[E, T]) PostId() *ast.IntColumnProjection[T, *int64] {
	column := PostId()
	projection := ast.NewIntColumnProjection[T, *int64](
		f.tableName(),
		column.Name(),
		func(t *T) *int64 {
			e := E(t)
			dto := e.GetPostTagsDto()
			return &dto.PostId
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
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'post_id' in alias %s", f.alias.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostTagsDto[E, T]) TagId() *ast.IntColumnProjection[T, *int64] {
	column := TagId()
	projection := ast.NewIntColumnProjection[T, *int64](
		f.tableName(),
		column.Name(),
		func(t *T) *int64 {
			e := E(t)
			dto := e.GetPostTagsDto()
			return &dto.TagId
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
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'tag_id' in alias %s", f.alias.Alias.Name()))
	projection.IntColumnExpression = ast.NewIntColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forPostTagsDto[E, T]) tableName() string {
	if f.alias == nil || f.alias.Alias == nil {
		return "post_tags"
	}
	return f.alias.Alias.Name()
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
