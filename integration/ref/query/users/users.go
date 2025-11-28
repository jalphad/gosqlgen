package users

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/ast/types"
)

func Id() *types.UUIDType[*types.UsersExpr] {
	expr := newExpr()
	return expr.Id
}

func Username() *types.StringType[*types.UsersExpr] {
	expr := newExpr()
	return expr.Username
}

func Email() *types.StringType[*types.UsersExpr] {
	expr := newExpr()
	return expr.Email
}

func FullName() *types.StringType[*types.UsersExpr] {
	expr := newExpr()
	return expr.FullName
}

func IsActive() *types.BoolType[*types.UsersExpr] {
	expr := newExpr()
	return expr.IsActive
}

func AllColumns() []ast.ValueExpression {
	expr := newExpr()
	return []ast.ValueExpression{
		expr.Id,
		expr.Username,
		expr.Email,
		expr.FullName,
		expr.IsActive,
	}
}

func newExpr() *types.UsersExpr {
	expr := &types.UsersExpr{}
	expr.Id = types.NewUUIDType(&ast.FieldRef{
		Table:  "users",
		Column: "id",
	}, expr)
	expr.Email = types.NewStringType(&ast.FieldRef{
		Table:  "users",
		Column: "email",
	}, expr)
	expr.Username = types.NewStringType(&ast.FieldRef{
		Table:  "users",
		Column: "username",
	}, expr)
	expr.FullName = types.NewStringType(&ast.FieldRef{
		Table:  "users",
		Column: "full_name",
	}, expr)
	expr.IsActive = types.NewBoolType(&ast.FieldRef{
		Table:  "users",
		Column: "is_active",
	}, expr)
	return expr
}
