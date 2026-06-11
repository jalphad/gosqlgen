package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
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

func (u UsersDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(u)*len(columns))
	for _, column := range columns {
		// Normalize column name for matching (lowercase)
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "username" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.Username))
			}
		}

		if columnLower == "email" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.Email))
			}
		}

		if columnLower == "full_name" || columnLower == "fullname" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.FullName))
			}
		}

		if columnLower == "created_at" || columnLower == "createdat" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "updated_at" || columnLower == "updatedat" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.UpdatedAt))
			}
		}

		if columnLower == "is_active" || columnLower == "isactive" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.IsActive))
			}
		}
	}

	values := make([]ast.Expression, 0, len(u))
	for i := 0; i < len(u); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(u)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

// TODO: return an ast.ErrorExpression when implemented
func (u UsersDtos) GetArgs(column ast.NamedExpression) ast.Expression {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(column.Name())

	if columnLower == "id" {
		out := u.Id()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "username" {
		out := make([]string, len(u))
		for idx, dto := range u {
			out[idx] = dto.Username
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "email" {
		out := make([]string, len(u))
		for idx, dto := range u {
			out[idx] = dto.Email
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "full_name" || columnLower == "fullname" {
		out := make([]*string, len(u))
		for idx, dto := range u {
			out[idx] = dto.FullName
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		out := make([]*time.Time, len(u))
		for idx, dto := range u {
			out[idx] = dto.CreatedAt
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "updated_at" || columnLower == "updatedat" {
		out := make([]*time.Time, len(u))
		for idx, dto := range u {
			out[idx] = dto.UpdatedAt
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "is_active" || columnLower == "isactive" {
		out := make([]*bool, len(u))
		for idx, dto := range u {
			out[idx] = dto.IsActive
		}
		return ast.NewLiteralExpression(out)
	}

	return nil //, errors.New("unknown column")
}

// UsersDto represents the users table
type UsersDto struct {
	Id        *uuid.UUID `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	FullName  *string    `db:"full_name" json:"full_name"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	IsActive  *bool      `db:"is_active" json:"is_active"`

	// One-to-many reverse relationships (populated via LoadXxx methods)
	Comments []CommentsDto `reverse:"comments" fk:"user_id"`
	Posts    []PostsDto    `reverse:"posts" fk:"user_id"`
}

// TableName returns the table name for UsersDto
func (u *UsersDto) TableName() string {
	return "users"
}

func (u *UsersDto) UnmarshalJSON(b []byte) error {
	type UsersDto_ UsersDto
	type DtoWrapper struct {
		UsersDto_
		CreatedAt NoTimezoneWrapper `db:"created_at" json:"created_at"`
		UpdatedAt NoTimezoneWrapper `db:"updated_at" json:"updated_at"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*u = UsersDto(wrapper.UsersDto_)
	u.CreatedAt = &wrapper.CreatedAt.Time
	u.UpdatedAt = &wrapper.UpdatedAt.Time

	return nil
}

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (u *UsersDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "id" {
		if u.Id == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.Id), nil
	}

	if columnLower == "username" {
		return ast.NewSQLType(u.Username), nil
	}

	if columnLower == "email" {
		return ast.NewSQLType(u.Email), nil
	}

	if columnLower == "full_name" || columnLower == "fullname" {
		if u.FullName == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.FullName), nil
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		if u.CreatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.CreatedAt), nil
	}

	if columnLower == "updated_at" || columnLower == "updatedat" {
		if u.UpdatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.UpdatedAt), nil
	}

	if columnLower == "is_active" || columnLower == "isactive" {
		if u.IsActive == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.IsActive), nil
	}

	return nil, errors.New("unknown column")
}
