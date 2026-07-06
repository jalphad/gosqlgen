package models

import (
	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
)

type UsersDtos []*UsersDto

func (u UsersDtos) Id() ast.OfType[[]uuid.UUID] {
	out := make([]*uuid.UUID, len(u))
	for idx, dto := range u {
		out[idx] = dto.Id
	}
	return ast.SetType[[]uuid.UUID](ast.NewLiteralExpression(out))
}

// UsersDto represents the users table
type UsersDto struct {
	Id *uuid.UUID `db:"id" json:"id"`
}

type ExportsUsersDto[T any] interface {
	*T
	GetUsersDto() *UsersDto
}
