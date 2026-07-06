package models

import (
	"strings"

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

func (u UsersDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(u)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range u {
				im = append(im, ast.NewLiteralExpression(dto.Id))
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

// UsersDto represents the users table
type UsersDto struct {
	Id *uuid.UUID `db:"id" json:"id"`
}

type ExportsUsersDto[T any] interface {
	*T
	GetUsersDto() *UsersDto
}
