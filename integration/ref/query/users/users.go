package users

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

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
