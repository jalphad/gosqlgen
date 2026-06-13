package posts

import (
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

func (intoPostsDto) Comments(columns ...ast.NamedExpression) ast.Projection[models.PostsDto] {
	if len(columns) == 0 {
		columns = commentsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("comments", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(p *models.PostsDto) *[]models.CommentsDto {
			return &p.Comments
		},
	)
}

func (intoPostsDto) Tags(columns ...ast.NamedExpression) ast.Projection[models.PostsDto] {
	if len(columns) == 0 {
		columns = tagsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("tags", expr.IsNotNull(ast.NewIntColumnExpression("tags", "id")), columns...),
		func(p *models.PostsDto) *[]models.TagsDto {
			return &p.Tags
		},
	)
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

func tagsColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("tags", "id"),
		ast.NewStringColumnExpression("tags", "name"),
		ast.NewStringColumnExpression("tags", "slug"),
	}
}
