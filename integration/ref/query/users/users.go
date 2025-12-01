package users

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Id() *ast.UUIDType {
	return ast.UUID(ast.NewColumnExpresion(&ast.ColumnRef{
		Table:  "users",
		Column: "id",
	}))
}

func Username() *ast.StringType {
	return ast.String(ast.NewColumnExpresion(&ast.ColumnRef{
		Table:  "users",
		Column: "username",
	}))
}

func Email() *ast.StringType {
	return ast.String(ast.NewColumnExpresion(&ast.ColumnRef{
		Table:  "users",
		Column: "email",
	}))
}

func FullName() *ast.StringType {
	return ast.String(ast.NewColumnExpresion(&ast.ColumnRef{
		Table:  "users",
		Column: "full_name",
	}))
}

func IsActive() *ast.BoolType {
	return ast.Bool(ast.NewColumnExpresion(&ast.ColumnRef{
		Table:  "users",
		Column: "is_active",
	}))
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
