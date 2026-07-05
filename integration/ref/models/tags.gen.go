package models

import (
	"strings"

	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

type TagsDtos []*TagsDto

func (t TagsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(t))
	for idx, dto := range t {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (t TagsDtos) Name() ast.OfType[[]string] {
	out := make([]string, len(t))
	for idx, dto := range t {
		out[idx] = dto.Name
	}
	return ast.NewSQLType(out)
}

func (t TagsDtos) Slug() ast.OfType[[]string] {
	out := make([]string, len(t))
	for idx, dto := range t {
		out[idx] = dto.Slug
	}
	return ast.NewSQLType(out)
}

func (t TagsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(t)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range t {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "name" {
			for _, dto := range t {
				im = append(im, ast.NewLiteralExpression(dto.Name))
			}
		}

		if columnLower == "slug" {
			for _, dto := range t {
				im = append(im, ast.NewLiteralExpression(dto.Slug))
			}
		}
	}

	values := make([]ast.Expression, 0, len(t))
	for i := 0; i < len(t); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(t)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

// TagsDto represents the tags table
type TagsDto struct {
	Id   *int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Posts []PostsDto `manytomany:"post_tags"`
}
