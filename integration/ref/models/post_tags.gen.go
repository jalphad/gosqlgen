package models

import (
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

// PostTagsDto represents the post_tags table
type PostTagsDto struct {
	PostId int64 `db:"post_id" json:"post_id"`
	TagId  int64 `db:"tag_id" json:"tag_id"`

	// Joined relationships (populated when corresponding Join method is called)
	PostIdRef *PostsDto `joined:"posts" fk:"post_id"`
	TagIdRef  *TagsDto  `joined:"tags" fk:"tag_id"`
}
