package tags

import (
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

func (intoTagsDto) Posts(columns ...ast.NamedExpression) ast.Projection[models.TagsDto] {
	if len(columns) == 0 {
		columns = postsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(t *models.TagsDto) *[]models.PostsDto {
			return &t.Posts
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
