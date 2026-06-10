package comments

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoCommentsDto{}

type intoCommentsDto struct{}

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

func (intoCommentsDto) Id() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(Id(), func(c *models.CommentsDto) **int64 {
		return &c.Id
	})
}

func (intoCommentsDto) Content() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(Content(), func(c *models.CommentsDto) *string {
		return &c.Content
	})
}

func (intoCommentsDto) UserId() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(UserId(), func(c *models.CommentsDto) *uuid.UUID {
		return &c.UserId
	})
}

func (intoCommentsDto) PostId() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(PostId(), func(c *models.CommentsDto) *int64 {
		return &c.PostId
	})
}

func (intoCommentsDto) IsApproved() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(IsApproved(), func(c *models.CommentsDto) **bool {
		return &c.IsApproved
	})
}

func (intoCommentsDto) CreatedAt() ast.Projection[models.CommentsDto] {
	return ast.NewProjection(CreatedAt(), func(c *models.CommentsDto) **time.Time {
		return &c.CreatedAt
	})
}
