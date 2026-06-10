package posts

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Table() *ast.TableSource {
	return &ast.TableSource{Table: "posts"}
}

func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("posts", "id")
}

func Title() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "title")
}

func UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("posts", "user_id")
}

func Content() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "content")
}

func Status() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "status")
}

func PublishedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("posts", "published_at")
}

func ViewCount() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("posts", "view_count")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Title(),
		UserId(),
		Content(),
		Status(),
		PublishedAt(),
		ViewCount(),
	}
}
