package users

import (
	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Id() *ast.ColumnExpression[uuid.UUID] {
	return &ast.ColumnExpression[uuid.UUID]{
		ColumnNode: ast.ColumnNode{
			Table:  "users",
			Column: "id",
		},
	}
}

func Username() *ast.ColumnExpression[string] {
	return &ast.ColumnExpression[string]{
		ColumnNode: ast.ColumnNode{
			Table:  "users",
			Column: "username",
		},
	}
}

func Email() *ast.ColumnExpression[string] {
	return &ast.ColumnExpression[string]{
		ColumnNode: ast.ColumnNode{
			Table:  "users",
			Column: "email",
		},
	}
}

func FullName() *ast.ColumnExpression[string] {
	return &ast.ColumnExpression[string]{
		ColumnNode: ast.ColumnNode{
			Table:  "users",
			Column: "full_name",
		},
	}
}

func IsActive() *ast.ColumnExpression[bool] {
	return &ast.ColumnExpression[bool]{
		ColumnNode: ast.ColumnNode{
			Table:  "users",
			Column: "is_active",
		},
	}
}

func AllColumns() []ast.Expression {
	return []ast.Expression{
		Id(),
		Username(),
		Email(),
		FullName(),
		IsActive(),
	}
}
