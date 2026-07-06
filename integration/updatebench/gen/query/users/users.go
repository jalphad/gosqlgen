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

// NewQuery returns a query builder for users
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return builder.NewKnownTableBuilder[models.UsersDto](pool, Table())
}

var Into = intoUsersDto{}

type intoUsersDto struct{}

type forUsersDto[E models.ExportsUsersDto[T], T any] struct{}

func For[E models.ExportsUsersDto[T], T any]() forUsersDto[E, T] {
	return forUsersDto[E, T]{}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("users")
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"users",
		"id",
		func(u *models.UsersDto) **uuid.UUID {
			return &u.Id
		},
	)
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

func (forUsersDto[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"users",
		"id",
		func(t *T) **uuid.UUID {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Id
		},
	)
}

func (f forUsersDto[E, T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
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

func (a *Alias) Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	column := Id()
	alias := ast.NewUUIDColumnProjection[models.UsersDto, **uuid.UUID](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **uuid.UUID {
			return &u.Id
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	alias.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
