package users

import (
	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoUsersDto{}
var Column = columns{}

type intoUsersDto struct{}
type columns struct{}

func Table() *ast.TableSource {
	return &ast.TableSource{Table: "users"}
}

func Id() *ast.UUIDFieldProjection[models.UsersDto, **uuid.UUID] {
	return ast.NewUUIDFieldProjection("users", "id", func(u *models.UsersDto) **uuid.UUID {
		return &u.Id
	})
}

func Username() *ast.StringFieldProjection[models.UsersDto, *string] {
	return ast.NewStringFieldProjection("users", "username", func(u *models.UsersDto) *string {
		return &u.Username
	})
}

func Email() *ast.StringFieldProjection[models.UsersDto, *string] {
	return ast.NewStringFieldProjection("users", "email", func(u *models.UsersDto) *string {
		return &u.Email
	})
}

func FullName() *ast.StringFieldProjection[models.UsersDto, **string] {
	return ast.NewStringFieldProjection("users", "full_name", func(u *models.UsersDto) **string {
		return &u.FullName
	})
}

func IsActive() *ast.BoolFieldProjection[models.UsersDto, **bool] {
	return ast.NewBoolFieldProjection("users", "is_active", func(u *models.UsersDto) **bool {
		return &u.IsActive
	})
}

func (columns) Id() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("users", "id")
}

func (columns) Username() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "username")
}

func (columns) Email() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "email")
}

func (columns) FullName() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("users", "full_name")
}

func (columns) IsActive() *ast.BoolColumnExpression {
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
