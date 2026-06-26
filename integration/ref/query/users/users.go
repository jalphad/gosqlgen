package users

import (
	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoUsersDto{}

type intoUsersDto struct{}

func Table() *ast.TableSource {
	return &ast.TableSource{Table: "users"}
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
		"username",
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
