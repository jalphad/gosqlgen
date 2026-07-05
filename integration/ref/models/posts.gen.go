package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

type PostsDtos []*PostsDto

func (p PostsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (p PostsDtos) UserId() ast.OfType[[]uuid.UUID] {
	out := make([]uuid.UUID, len(p))
	for idx, dto := range p {
		out[idx] = dto.UserId
	}
	return ast.NewSQLType(out)
}

func (p PostsDtos) Title() ast.OfType[[]string] {
	out := make([]string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Title
	}
	return ast.NewSQLType(out)
}

func (p PostsDtos) Content() ast.OfType[[]string] {
	out := make([]*string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Content
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

func (p PostsDtos) Status() ast.OfType[[]string] {
	out := make([]*string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Status
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

func (p PostsDtos) PublishedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(p))
	for idx, dto := range p {
		out[idx] = dto.PublishedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (p PostsDtos) ViewCount() ast.OfType[[]int64] {
	out := make([]*int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.ViewCount
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (p PostsDtos) CreatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(p))
	for idx, dto := range p {
		out[idx] = dto.CreatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (p PostsDtos) UpdatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(p))
	for idx, dto := range p {
		out[idx] = dto.UpdatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (p PostsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(p)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "user_id" || columnLower == "userid" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.UserId))
			}
		}

		if columnLower == "title" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Title))
			}
		}

		if columnLower == "content" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Content))
			}
		}

		if columnLower == "status" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Status))
			}
		}

		if columnLower == "published_at" || columnLower == "publishedat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.PublishedAt))
			}
		}

		if columnLower == "view_count" || columnLower == "viewcount" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.ViewCount))
			}
		}

		if columnLower == "created_at" || columnLower == "createdat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "updated_at" || columnLower == "updatedat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.UpdatedAt))
			}
		}
	}

	values := make([]ast.Expression, 0, len(p))
	for i := 0; i < len(p); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(p)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

// PostsDto represents the posts table
type PostsDto struct {
	Id          *int64     `db:"id" json:"id"`
	UserId      uuid.UUID  `db:"user_id" json:"user_id"`
	Title       string     `db:"title" json:"title"`
	Content     *string    `db:"content" json:"content"`
	Status      *string    `db:"status" json:"status"`
	PublishedAt *time.Time `db:"published_at" json:"published_at"`
	ViewCount   *int64     `db:"view_count" json:"view_count"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`

	// Joined relationships (populated when corresponding Join method is called)
	User *UsersDto `joined:"users" fk:"user_id"`

	// One-to-many reverse relationships (populated via LoadXxx methods)
	Comments []CommentsDto `reverse:"comments" fk:"post_id"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Tags []TagsDto `manytomany:"post_tags"`
}

func (p *PostsDto) UnmarshalJSON(b []byte) error {
	type PostsDto_ PostsDto
	type DtoWrapper struct {
		PostsDto_
		PublishedAt NoTimezoneWrapper `db:"published_at" json:"published_at"`
		CreatedAt   NoTimezoneWrapper `db:"created_at" json:"created_at"`
		UpdatedAt   NoTimezoneWrapper `db:"updated_at" json:"updated_at"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*p = PostsDto(wrapper.PostsDto_)
	p.PublishedAt = &wrapper.PublishedAt.Time
	p.CreatedAt = &wrapper.CreatedAt.Time
	p.UpdatedAt = &wrapper.UpdatedAt.Time

	return nil
}

type ExportsPostsDto[T any] interface {
	*T
	GetPostsDto() *PostsDto
}
