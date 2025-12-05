package posts

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("posts", "id")
}

func Title() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "title")
}

func UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("posts", "user_id")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Title(),
		UserId(),
	}
}
