package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

type UsersDtos []*UsersDto

func (u UsersDtos) Id() ast.OfType[[]uuid.UUID] {
	out := make([]*uuid.UUID, len(u))
	for idx, dto := range u {
		out[idx] = dto.Id
	}
	return ast.SetType[[]uuid.UUID](ast.NewLiteralExpression(out))
}

func (u UsersDtos) Username() ast.OfType[[]string] {
	out := make([]string, len(u))
	for idx, dto := range u {
		out[idx] = dto.Username
	}
	return ast.NewSQLType(out)
}

func (u UsersDtos) Email() ast.OfType[[]string] {
	out := make([]string, len(u))
	for idx, dto := range u {
		out[idx] = dto.Email
	}
	return ast.NewSQLType(out)
}

func (u UsersDtos) FullName() ast.OfType[[]string] {
	out := make([]*string, len(u))
	for idx, dto := range u {
		out[idx] = dto.FullName
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

func (u UsersDtos) CreatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(u))
	for idx, dto := range u {
		out[idx] = dto.CreatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (u UsersDtos) UpdatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(u))
	for idx, dto := range u {
		out[idx] = dto.UpdatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (u UsersDtos) IsActive() ast.OfType[[]bool] {
	out := make([]*bool, len(u))
	for idx, dto := range u {
		out[idx] = dto.IsActive
	}
	return ast.SetType[[]bool](ast.NewLiteralExpression(out))
}

func (u UsersDtos) Attributes() ast.OfType[[]json.RawMessage] {
	out := make([]*json.RawMessage, len(u))
	for idx, dto := range u {
		out[idx] = dto.Attributes
	}
	return ast.SetType[[]json.RawMessage](ast.NewLiteralExpression(out))
}

// UsersDto represents the "public"."users" table
type UsersDto struct {
	Id         *uuid.UUID       `db:"id" json:"id"`
	Username   string           `db:"username" json:"username"`
	Email      string           `db:"email" json:"email"`
	FullName   *string          `db:"full_name" json:"full_name"`
	CreatedAt  *time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time       `db:"updated_at" json:"updated_at"`
	IsActive   *bool            `db:"is_active" json:"is_active"`
	Attributes *json.RawMessage `db:"attributes" json:"attributes"`

	// One-to-many reverse relationships (populated via LoadXxx methods)
	CommentsByUserId []CommentsDto `reverse:"public.comments" fk:"user_id"`
	PostsByUserId    []PostsDto    `reverse:"public.posts" fk:"user_id"`
}

func (u *UsersDto) UnmarshalJSON(bts []byte) error {
	type UsersDto_ UsersDto
	type DtoWrapper struct {
		UsersDto_
		CreatedAt NoTimezoneWrapper `db:"created_at" json:"created_at"`
		UpdatedAt NoTimezoneWrapper `db:"updated_at" json:"updated_at"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(bts, &wrapper)
	if err != nil {
		return err
	}

	*u = UsersDto(wrapper.UsersDto_)
	u.CreatedAt = &wrapper.CreatedAt.Time
	u.UpdatedAt = &wrapper.UpdatedAt.Time

	return nil
}
