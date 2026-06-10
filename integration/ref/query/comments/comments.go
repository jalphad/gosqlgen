package comments

import (
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func Table() *ast.TableSource {
	return &ast.TableSource{Table: "comments"}
}

func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "id")
}

func Content() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("comments", "content")
}

func UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("comments", "user_id")
}

func PostId() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "post_id")
}

func IsApproved() *ast.BoolColumnExpression {
	return ast.NewBoolColumnExpression("comments", "is_approved")
}

func CreatedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("comments", "created_at")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Content(),
		UserId(),
		PostId(),
		IsApproved(),
		CreatedAt(),
	}
}
