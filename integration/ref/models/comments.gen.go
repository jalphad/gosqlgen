package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
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
	PostIdRef *PostsDto `joined:"posts" fk:"post_id"`
	UserIdRef *UsersDto `joined:"users" fk:"user_id"`
}

func (c *CommentsDto) UnmarshalJSON(bts []byte) error {
	type CommentsDto_ CommentsDto
	type DtoWrapper struct {
		CommentsDto_
		TestDate DateWrapper `db:"test_date" json:"test_date"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(bts, &wrapper)
	if err != nil {
		return err
	}

	*c = CommentsDto(wrapper.CommentsDto_)
	c.TestDate = &wrapper.TestDate.Time

	return nil
}

func (c *CommentsDto) GetCommentsDto() *CommentsDto {
	return c
}

type ExportsCommentsDto[T any] interface {
	*T
	GetCommentsDto() *CommentsDto
}
