package models

import (
	"strings"

	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
)

type PostTagsDtos []*PostTagsDto

func (p PostTagsDtos) PostId() ast.OfType[[]int64] {
	out := make([]int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.PostId
	}
	return ast.NewSQLType(out)
}

func (p PostTagsDtos) TagId() ast.OfType[[]int64] {
	out := make([]int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.TagId
	}
	return ast.NewSQLType(out)
}

func (p PostTagsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(p)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "post_id" || columnLower == "postid" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.PostId))
			}
		}

		if columnLower == "tag_id" || columnLower == "tagid" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.TagId))
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

// PostTagsDto represents the post_tags table
type PostTagsDto struct {
	PostId int64 `db:"post_id" json:"post_id"`
	TagId  int64 `db:"tag_id" json:"tag_id"`

	// Joined relationships (populated when corresponding Join method is called)
	Post *PostsDto `joined:"posts" fk:"post_id"`
	Tag  *TagsDto  `joined:"tags" fk:"tag_id"`
}

type ExportsPostTagsDto[T any] interface {
	*T
	GetPostTagsDto() *PostTagsDto
}
