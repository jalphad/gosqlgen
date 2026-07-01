package comments

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoCommentsDto{}
var Column = columns{}

type intoCommentsDto struct{}
type columns struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("comments")
}

func Id() *ast.IntColumnProjection[models.CommentsDto, **int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"id",
		func(c *models.CommentsDto) **int64 {
			return &c.Id
		},
	)
}

func Content() *ast.StringColumnProjection[models.CommentsDto, *string] {
	return ast.NewStringColumnProjection(
		"comments",
		"content",
		func(c *models.CommentsDto) *string {
			return &c.Content
		},
	)
}

func UserId() *ast.UUIDColumnProjection[models.CommentsDto, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"comments",
		"user_id",
		func(c *models.CommentsDto) *uuid.UUID {
			return &c.UserId
		},
	)
}

func PostId() *ast.IntColumnProjection[models.CommentsDto, *int64] {
	return ast.NewIntColumnProjection(
		"comments",
		"post_id",
		func(c *models.CommentsDto) *int64 {
			return &c.PostId
		},
	)
}

func IsApproved() *ast.BoolColumnProjection[models.CommentsDto, **bool] {
	return ast.NewBoolColumnProjection(
		"comments",
		"is_approved",
		func(c *models.CommentsDto) **bool {
			return &c.IsApproved
		},
	)
}

func CreatedAt() *ast.TimestampColumnProjection[models.CommentsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"comments",
		"created_at",
		func(c *models.CommentsDto) **time.Time {
			return &c.CreatedAt
		},
	)
}

func (columns) Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "id")
}

func (columns) Content() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("comments", "content")
}

func (columns) UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("comments", "user_id")
}

func (columns) PostId() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("comments", "post_id")
}

func (columns) IsApproved() *ast.BoolColumnExpression {
	return ast.NewBoolColumnExpression("comments", "is_approved")
}

func (columns) CreatedAt() *ast.TimestampColumnExpression {
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
	return Id()
}

func (intoCommentsDto) Content() ast.Projection[models.CommentsDto] {
	return Content()
}

func (intoCommentsDto) UserId() ast.Projection[models.CommentsDto] {
	return UserId()
}

func (intoCommentsDto) PostId() ast.Projection[models.CommentsDto] {
	return PostId()
}

func (intoCommentsDto) IsApproved() ast.Projection[models.CommentsDto] {
	return IsApproved()
}

func (intoCommentsDto) CreatedAt() ast.Projection[models.CommentsDto] {
	return CreatedAt()
}
