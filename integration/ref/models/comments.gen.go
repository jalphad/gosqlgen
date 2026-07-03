package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
)

type CommentsDtos []*CommentsDto

func (c CommentsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(c))
	for idx, dto := range c {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) PostId() ast.OfType[[]int64] {
	out := make([]int64, len(c))
	for idx, dto := range c {
		out[idx] = dto.PostId
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) UserId() ast.OfType[[]uuid.UUID] {
	out := make([]uuid.UUID, len(c))
	for idx, dto := range c {
		out[idx] = dto.UserId
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) Content() ast.OfType[[]string] {
	out := make([]string, len(c))
	for idx, dto := range c {
		out[idx] = dto.Content
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) IsApproved() ast.OfType[[]bool] {
	out := make([]*bool, len(c))
	for idx, dto := range c {
		out[idx] = dto.IsApproved
	}
	return ast.SetType[[]bool](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) CreatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(c))
	for idx, dto := range c {
		out[idx] = dto.CreatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) TestDate() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(c))
	for idx, dto := range c {
		out[idx] = dto.TestDate
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(c)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "post_id" || columnLower == "postid" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.PostId))
			}
		}

		if columnLower == "user_id" || columnLower == "userid" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.UserId))
			}
		}

		if columnLower == "content" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.Content))
			}
		}

		if columnLower == "is_approved" || columnLower == "isapproved" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.IsApproved))
			}
		}

		if columnLower == "created_at" || columnLower == "createdat" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "test_date" || columnLower == "testdate" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.TestDate))
			}
		}
	}

	values := make([]ast.Expression, 0, len(c))
	for i := 0; i < len(c); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(c)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

// CommentsDto represents the comments table
type CommentsDto struct {
	Id         *int64     `db:"id" json:"id"`
	PostId     int64      `db:"post_id" json:"post_id"`
	UserId     uuid.UUID  `db:"user_id" json:"user_id"`
	Content    string     `db:"content" json:"content"`
	IsApproved *bool      `db:"is_approved" json:"is_approved"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	TestDate   *time.Time `db:"test_date" json:"test_date"`

	// Joined relationships (populated when corresponding Join method is called)
	Post *PostsDto `joined:"posts" fk:"post_id"`
	User *UsersDto `joined:"users" fk:"user_id"`
}

// TableName returns the table name for CommentsDto
func (c *CommentsDto) TableName() string {
	return "comments"
}

func (c *CommentsDto) UnmarshalJSON(b []byte) error {
	type CommentsDto_ CommentsDto
	type DtoWrapper struct {
		CommentsDto_
		CreatedAt NoTimezoneWrapper `db:"created_at" json:"created_at"`
		TestDate  DateWrapper       `db:"test_date" json:"test_date"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*c = CommentsDto(wrapper.CommentsDto_)
	c.CreatedAt = &wrapper.CreatedAt.Time
	c.TestDate = &wrapper.TestDate.Time

	return nil
}
