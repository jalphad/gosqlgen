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

type forUsersDto[E models.ExportsUsersDto[T], T any] struct {
	alias *Alias
	err   error
}

// TODO(go1.27): if type parameters on methods are available, consider
// supporting alias.For[*resultRow]() as the primary alias API.
func For[E models.ExportsUsersDto[T], T any](alias ...*Alias) forUsersDto[E, T] {
	switch len(alias) {
	case 0:
		return forUsersDto[E, T]{}
	case 1:
		if alias[0] == nil || alias[0].Alias == nil {
			return forUsersDto[E, T]{err: fmt.Errorf("users.For requires a non-nil alias")}
		}
		return forUsersDto[E, T]{alias: alias[0]}
	default:
		return forUsersDto[E, T]{err: fmt.Errorf("users.For accepts at most one alias")}
	}
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

func (f forUsersDto[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	column := Id()
	projection := ast.NewUUIDColumnProjection[T, **uuid.UUID](
		f.tableName(),
		column.Name(),
		func(t *T) **uuid.UUID {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Id
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
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", f.alias.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forUsersDto[E, T]) tableName() string {
	if f.alias == nil || f.alias.Alias == nil {
		return "users"
	}
	return f.alias.Alias.Name()
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
