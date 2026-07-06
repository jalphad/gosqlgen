package models

import (
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

// TagsDto represents the tags table
type TagsDto struct {
	Id   *int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Posts []PostsDto `manytomany:"post_tags"`
}

type ExportsTagsDto[T any] interface {
	*T
	GetTagsDto() *TagsDto
}
