package post_tags

import "github.com/jalphad/gosqlgen/integration/ref/ast"

func PostId() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("post_tags", "post_id")
}

func TagId() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("post_tags", "tag_id")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		PostId(),
		TagId(),
	}
}
