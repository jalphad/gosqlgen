package users

import (
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

func (intoUsersDto) Posts(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = postsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(u *models.UsersDto) *[]models.PostsDto {
			return &u.Posts
		},
	)
}

func (intoUsersDto) Comments(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = commentsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("comments", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(u *models.UsersDto) *[]models.CommentsDto {
			return &u.Comments
		},
	)
}

func postsColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("posts", "id"),
		ast.NewUUIDColumnExpression("posts", "user_id"),
		ast.NewStringColumnExpression("posts", "title"),
		ast.NewStringColumnExpression("posts", "content"),
		ast.NewStringColumnExpression("posts", "status"),
		ast.NewTimestampColumnExpression("posts", "published_at"),
		ast.NewIntColumnExpression("posts", "view_count"),
		ast.NewTimestampColumnExpression("posts", "created_at"),
		ast.NewTimestampColumnExpression("posts", "updated_at"),
	}
}

func commentsColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("comments", "id"),
		ast.NewStringColumnExpression("comments", "content"),
		ast.NewUUIDColumnExpression("comments", "user_id"),
		ast.NewIntColumnExpression("comments", "post_id"),
		ast.NewBoolColumnExpression("comments", "is_approved"),
		ast.NewTimestampColumnExpression("comments", "created_at"),
	}
}
