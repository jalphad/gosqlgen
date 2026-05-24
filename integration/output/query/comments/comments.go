package comments

import (
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
)


func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "id")
}

func PostId() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "post_id")
}

func UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("comments", "user_id")
}

func Content() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("comments", "content")
}

func IsApproved() *ast.BoolColumnExpression {
	return ast.NewBoolColumnExpression("comments", "is_approved")
}

func CreatedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("comments", "created_at")
}

func TestDate() *ast.DateColumnExpression {
	return ast.NewDateColumnExpression("comments", "test_date")
}


func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		PostId(),
		UserId(),
		Content(),
		IsApproved(),
		CreatedAt(),
		TestDate(),
	}
}
