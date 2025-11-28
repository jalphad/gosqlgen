package types

import "github.com/jalphad/gosqlgen/integration/ref/ast"

// UsersExpr represents the users table columns
type UsersExpr struct {
	Id       *UUIDType[*UsersExpr]
	Username *StringType[*UsersExpr]
	FullName *StringType[*UsersExpr]
	Email    *StringType[*UsersExpr]
	IsActive *BoolType[*UsersExpr]
	expression
}

func (c *UsersExpr) And(next ast.Expression) ast.BooleanExpression {
	if cl, ok := next.(*UsersExpr); ok {
		next = cl.expression
	}
	c.expression = ast.NewLogicalExpression(&ast.ExpressionNode{Op: "AND", Args: []ast.Expression{c.expression, next}})
	return c
}

func (c *UsersExpr) Or(next ast.Expression) ast.BooleanExpression {
	if cl, ok := next.(*UsersExpr); ok {
		next = cl.expression
	}
	c.expression = ast.NewLogicalExpression(&ast.ExpressionNode{Op: "OR", Args: []ast.Expression{c.expression, next}})
	return c
}

func (c *UsersExpr) setExpression(expr ast.Expression) {
	c.expression = expr
}
