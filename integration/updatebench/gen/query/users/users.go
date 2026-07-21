package users

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/builder"
)

// NewQuery returns a query builder for "public"."users"
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return builder.NewKnownTableBuilder[models.UsersDto](pool, Table())
}

var Into = intoUsersDto{}

type intoUsersDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource(`"public"."users"`)
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return newId(`"public"."users"`, func(u *models.UsersDto) **uuid.UUID {
		return &u.Id
	})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
	}
}

func (intoUsersDto) Id() ast.Projection[models.UsersDto] {
	return Id()
}

func (intoUsersDto) AllColumns() []ast.Projection[models.UsersDto] {
	return []ast.Projection[models.UsersDto]{
		Into.Id(),
	}
}

type forUsersDto[T any] struct {
	dest func(*T) *models.UsersDto
}

func For[T any](dest func(*T) *models.UsersDto) forUsersDto[T] {
	return forUsersDto[T]{dest: dest}
}

func (f forUsersDto[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return newId(`"public"."users"`, func(t *T) **uuid.UUID {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forUsersDto[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
	}
}

type Alias[T any] struct {
	*ast.TableAlias
	forUsersDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.UsersDto] {
	return Alias[models.UsersDto]{
		TableAlias: ast.NewTableAlias(name, columns...),
		forUsersDto: forUsersDto[models.UsersDto]{
			dest: func(u *models.UsersDto) *models.UsersDto {
				return u
			},
		},
	}
}

func (a Alias[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	projection := newId(a.TableAlias.GetName(), func(t *T) **uuid.UUID {
		dto := a.dest(t)
		return &dto.Id
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.TableAlias.GetName()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
	}
}

func AliasFor[T any](alias *ast.TableAlias, dest func(*T) *models.UsersDto) Alias[T] {
	return Alias[T]{
		TableAlias:  alias,
		forUsersDto: forUsersDto[T]{dest: dest},
	}
}

func newId[T any](table string, ref func(*T) **uuid.UUID) *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		table,
		"id",
		ref,
	)
}
