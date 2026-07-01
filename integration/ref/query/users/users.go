package users

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoUsersDto{}

type intoUsersDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("users")
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"users",
		"id",
		func(r *models.UsersDto) **uuid.UUID {
			return &r.Id
		})
}

func Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"username",
		func(r *models.UsersDto) *string {
			return &r.Username
		})
}

func Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"email",
		func(r *models.UsersDto) *string {
			return &r.Email
		})
}

func FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	return ast.NewStringColumnProjection(
		"users",
		"full_name",
		func(r *models.UsersDto) **string {
			return &r.FullName
		})
}

func IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	return ast.NewBoolColumnProjection(
		"users",
		"is_active",
		func(r *models.UsersDto) **bool {
			return &r.IsActive
		})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Username(),
		Email(),
		FullName(),
		IsActive(),
	}
}

func (intoUsersDto) Id() ast.Projection[models.UsersDto] {
	return Id()
}

func (intoUsersDto) Username() ast.Projection[models.UsersDto] {
	return Username()
}

func (intoUsersDto) Email() ast.Projection[models.UsersDto] {
	return Email()
}

func (intoUsersDto) FullName() ast.Projection[models.UsersDto] {
	return FullName()
}

func (intoUsersDto) IsActive() ast.Projection[models.UsersDto] {
	return IsActive()
}

func (intoUsersDto) AllColumns() []ast.Projection[models.UsersDto] {
	return []ast.Projection[models.UsersDto]{
		Into.Id(),
		Into.Username(),
		Into.Email(),
		Into.FullName(),
		Into.IsActive(),
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
		a.Name(),
		column.Name(),
		func(r *models.UsersDto) **uuid.UUID {
			return &r.Id
		})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Name()))
	alias.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	column := Username()
	alias := ast.NewStringColumnProjection[models.UsersDto, *string](
		a.Name(),
		column.Name(),
		func(r *models.UsersDto) *string {
			return &r.Username
		})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'username' in alias %s", a.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	column := Email()
	alias := ast.NewStringColumnProjection[models.UsersDto, *string](
		a.Name(),
		column.Name(),
		func(r *models.UsersDto) *string {
			return &r.Email
		})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'email' in alias %s", a.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	column := FullName()
	alias := ast.NewStringColumnProjection[models.UsersDto, **string](
		a.Name(),
		column.Name(),
		func(r *models.UsersDto) **string {
			return &r.FullName
		})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'full_name' in alias %s", a.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	column := IsActive()
	alias := ast.NewBoolColumnProjection[models.UsersDto, **bool](
		a.Name(),
		column.Name(),
		func(r *models.UsersDto) **bool {
			return &r.IsActive
		})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_active' in alias %s", a.Name()))
	alias.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
