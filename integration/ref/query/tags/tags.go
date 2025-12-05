package tags

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("tags", "id")
}

func Name() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("tags", "name")
}

func Slug() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("tags", "slug")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Name(),
		Slug(),
	}
}
