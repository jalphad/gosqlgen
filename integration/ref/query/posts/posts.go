package posts

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoPostsDto{}
var Column = columns{}

type intoPostsDto struct{}
type columns struct{}

func Table() *ast.TableSource {
	return &ast.TableSource{Table: "posts"}
}

func Id() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return ast.NewIntColumnProjection(
		"posts",
		"id",
		func(p *models.PostsDto) **int64 {
			return &p.Id
		},
	)
}

func Title() *ast.StringColumnProjection[models.PostsDto, *string] {
	return ast.NewStringColumnProjection(
		"posts",
		"title",
		func(p *models.PostsDto) *string {
			return &p.Title
		},
	)
}

func UserId() *ast.UUIDColumnProjection[models.PostsDto, *uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"posts",
		"user_id",
		func(p *models.PostsDto) *uuid.UUID {
			return &p.UserId
		},
	)
}

func Content() *ast.StringColumnProjection[models.PostsDto, **string] {
	return ast.NewStringColumnProjection(
		"posts",
		"content",
		func(p *models.PostsDto) **string {
			return &p.Content
		},
	)
}

func Status() *ast.StringColumnProjection[models.PostsDto, **string] {
	return ast.NewStringColumnProjection(
		"posts",
		"status",
		func(p *models.PostsDto) **string {
			return &p.Status
		},
	)
}

func PublishedAt() *ast.TimestampColumnProjection[models.PostsDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"posts",
		"published_at",
		func(p *models.PostsDto) **time.Time {
			return &p.PublishedAt
		},
	)
}

func ViewCount() *ast.IntColumnProjection[models.PostsDto, **int64] {
	return ast.NewIntColumnProjection(
		"posts",
		"view_count",
		func(p *models.PostsDto) **int64 {
			return &p.ViewCount
		},
	)
}

func (columns) Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("posts", "id")
}

func (columns) Title() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "title")
}

func (columns) UserId() *ast.UUIDColumnExpression {
	return ast.NewUUIDColumnExpression("posts", "user_id")
}

func (columns) Content() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "content")
}

func (columns) Status() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("posts", "status")
}

func (columns) PublishedAt() *ast.TimestampColumnExpression {
	return ast.NewTimestampColumnExpression("posts", "published_at")
}

func (columns) ViewCount() *ast.IntColumnExpression {
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

func (intoPostsDto) Id() ast.Projection[models.PostsDto] {
	return Id()
}

func (intoPostsDto) Title() ast.Projection[models.PostsDto] {
	return Title()
}

func (intoPostsDto) UserId() ast.Projection[models.PostsDto] {
	return UserId()
}

func (intoPostsDto) Content() ast.Projection[models.PostsDto] {
	return Content()
}

func (intoPostsDto) Status() ast.Projection[models.PostsDto] {
	return Status()
}

func (intoPostsDto) PublishedAt() ast.Projection[models.PostsDto] {
	return PublishedAt()
}

func (intoPostsDto) ViewCount() ast.Projection[models.PostsDto] {
	return ViewCount()
}
