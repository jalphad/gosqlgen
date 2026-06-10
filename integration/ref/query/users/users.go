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

func Id() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("users", "id")
}

func Username() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "username")
}

func Email() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "email")
}

func FullName() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "full_name")
}

func IsActive() *ast.BoolColumnExpression {
	return ast.NewBoolColumnExpression("users", "is_active")
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
	return ast.NewProjection(Id(), func(u *models.UsersDto) **uuid.UUID {
		return &u.Id
	})
}

func (intoUsersDto) Username() ast.Projection[models.UsersDto] {
	return ast.NewProjection(Username(), func(u *models.UsersDto) *string {
		return &u.Username
	})
}

func (intoUsersDto) Email() ast.Projection[models.UsersDto] {
	return ast.NewProjection(Email(), func(u *models.UsersDto) *string {
		return &u.Email
	})
}

func (intoUsersDto) FullName() ast.Projection[models.UsersDto] {
	return ast.NewProjection(FullName(), func(u *models.UsersDto) **string {
		return &u.FullName
	})
}

func (intoUsersDto) IsActive() ast.Projection[models.UsersDto] {
	return ast.NewProjection(IsActive(), func(u *models.UsersDto) **bool {
		return &u.IsActive
	})
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
