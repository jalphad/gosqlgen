package posts

import (
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
)


func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("posts", "id")
}

func UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("posts", "user_id")
}

func Title() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "title")
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

func CreatedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("posts", "created_at")
}

func UpdatedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("posts", "updated_at")
}


func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		UserId(),
		Title(),
		Content(),
		Status(),
		PublishedAt(),
		ViewCount(),
		CreatedAt(),
		UpdatedAt(),
	}
}
