package posts

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoPostsDto{}

type intoPostsDto struct{}

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

func (intoPostsDto) Id() ast.Projection[models.PostsDto] {
	return ast.NewProjection(Id(), func(p *models.PostsDto) **int64 {
		return &p.Id
	})
}

func (intoPostsDto) Title() ast.Projection[models.PostsDto] {
	return ast.NewProjection(Title(), func(p *models.PostsDto) *string {
		return &p.Title
	})
}

func (intoPostsDto) UserId() ast.Projection[models.PostsDto] {
	return ast.NewProjection(UserId(), func(p *models.PostsDto) *uuid.UUID {
		return &p.UserId
	})
}

func (intoPostsDto) Content() ast.Projection[models.PostsDto] {
	return ast.NewProjection(Content(), func(p *models.PostsDto) **string {
		return &p.Content
	})
}

func (intoPostsDto) Status() ast.Projection[models.PostsDto] {
	return ast.NewProjection(Status(), func(p *models.PostsDto) **string {
		return &p.Status
	})
}

func (intoPostsDto) PublishedAt() ast.Projection[models.PostsDto] {
	return ast.NewProjection(PublishedAt(), func(p *models.PostsDto) **time.Time {
		return &p.PublishedAt
	})
}

func (intoPostsDto) ViewCount() ast.Projection[models.PostsDto] {
	return ast.NewProjection(ViewCount(), func(p *models.PostsDto) **int64 {
		return &p.ViewCount
	})
}
