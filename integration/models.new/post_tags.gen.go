package models

import (
	"errors"
	"strings"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

// PostTagsDto represents the post_tags table
type PostTagsDto struct {
	PostId int64 `db:"post_id" json:"post_id"`
	TagId  int64 `db:"tag_id" json:"tag_id"`

	// Joined relationships (populated when corresponding Join method is called)
	Post *PostsDto `joined:"posts" fk:"post_id"`
	Tag  *TagsDto  `joined:"tags" fk:"tag_id"`
}

// TableName returns the table name for PostTagsDto
func (p *PostTagsDto) TableName() string {
	return "post_tags"
}

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (p *PostTagsDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "post_id" || columnLower == "postid" {
		return ast.NewSQLType(p.PostId), nil
	}

	if columnLower == "tag_id" || columnLower == "tagid" {
		return ast.NewSQLType(p.TagId), nil
	}

	return nil, errors.New("unknown column")
}
