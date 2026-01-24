package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

// PostTagsDto represents the post_tags table
type PostTagsDto struct {
	PostId int64 `db:"post_id" json:"post_id"`
	TagId  int64 `db:"tag_id" json:"tag_id"`

	// Joined relationships (populated when corresponding Join method is called)
	Post *PostsDto `joined:"posts" fk:"post_id"`
	Tag  *TagsDto  `joined:"tags" fk:"tag_id"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]any `json:"from_expressions,omitempty"`
}

// TableName returns the table name for PostTagsDto
func (p *PostTagsDto) TableName() string {
	return "post_tags"
}

// scanInto scans a row into a struct
func (p *PostTagsDto) ScanInto(row pgx.Row, stmt ast.SqlStatement) error {
	var scanDest []any
	var jsonUnmarshalFuncs []func() error

	columns := stmt.Returns()
	if len(columns) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range columns {
			destPtr, unmarshalFunc := p.getScanDestForField(field)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		//// Default behavior - scan all table columns + joined tables
		//scanDest = append(scanDest, &p.PostId)
		//scanDest = append(scanDest, &p.TagId)
		//
		//// Add joined table columns if active
		//if slices.Contains(stmt.GetJoinedTables(), "posts") {
		//	joinedPost := &PostsDto{}
		//	scanDest = append(scanDest, &joinedPost.Id)
		//	scanDest = append(scanDest, &joinedPost.UserId)
		//	scanDest = append(scanDest, &joinedPost.Title)
		//	scanDest = append(scanDest, &joinedPost.Content)
		//	scanDest = append(scanDest, &joinedPost.Status)
		//	scanDest = append(scanDest, &joinedPost.PublishedAt)
		//	scanDest = append(scanDest, &joinedPost.ViewCount)
		//	scanDest = append(scanDest, &joinedPost.CreatedAt)
		//	scanDest = append(scanDest, &joinedPost.UpdatedAt)
		//	p.Post = joinedPost
		//}
		//if slices.Contains(stmt.GetJoinedTables(), "tags") {
		//	joinedTag := &TagsDto{}
		//	scanDest = append(scanDest, &joinedTag.Id)
		//	scanDest = append(scanDest, &joinedTag.Name)
		//	scanDest = append(scanDest, &joinedTag.Slug)
		//	p.Tag = joinedTag
		//}
	}

	// Perform scan
	if err := row.Scan(scanDest...); err != nil {
		return err
	}

	// Execute JSON unmarshal functions and collect errors
	var unmarshalErrors []error
	for _, fn := range jsonUnmarshalFuncs {
		if err := fn(); err != nil {
			unmarshalErrors = append(unmarshalErrors, err)
		}
	}

	if len(unmarshalErrors) > 0 {
		return fmt.Errorf("JSON unmarshal errors: %v", unmarshalErrors)
	}

	return nil
}

// getScanDestForField returns the appropriate scan destination for a field
// and optionally a function to unmarshal JSON data after scanning
func (p *PostTagsDto) getScanDestForField(ref ast.NamedExpression) (any, func() error) {
	var refTable string
	if split := strings.Split(ast.Render(ref, &[]any{}), "."); len(split) == 2 {
		refTable = split[0]
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(ref.Name())

	// Check if alias matches a regular table column
	if refTable == "post_tags" {

		if aliasLower == "post_id" || aliasLower == "postid" {
			return &p.PostId, nil
		}

		if aliasLower == "tag_id" || aliasLower == "tagid" {
			return &p.TagId, nil
		}
	}
	//// Check if it matches a joined table column
	//if refTable == "posts" && slices.Contains(stmt.GetJoinedTables(), "posts") {
	//	if p.Post == nil {
	//		p.Post = &PostsDto{}
	//	}
	//
	//	if aliasLower == "id" {
	//		return p.Post.Id, nil
	//	}
	//
	//	if aliasLower == "user_id" || aliasLower == "userid" {
	//		return p.Post.UserId, nil
	//	}
	//
	//	if aliasLower == "title" {
	//		return p.Post.Title, nil
	//	}
	//
	//	if aliasLower == "content" {
	//		return p.Post.Content, nil
	//	}
	//
	//	if aliasLower == "status" {
	//		return p.Post.Status, nil
	//	}
	//
	//	if aliasLower == "published_at" || aliasLower == "publishedat" {
	//		return p.Post.PublishedAt, nil
	//	}
	//
	//	if aliasLower == "view_count" || aliasLower == "viewcount" {
	//		return p.Post.ViewCount, nil
	//	}
	//
	//	if aliasLower == "created_at" || aliasLower == "createdat" {
	//		return p.Post.CreatedAt, nil
	//	}
	//
	//	if aliasLower == "updated_at" || aliasLower == "updatedat" {
	//		return p.Post.UpdatedAt, nil
	//	}
	//}
	//if refTable == "tags" && slices.Contains(stmt.GetJoinedTables(), "tags") {
	//	if p.Tag == nil {
	//		p.Tag = &TagsDto{}
	//	}
	//
	//	if aliasLower == "id" {
	//		return p.Tag.Id, nil
	//	}
	//
	//	if aliasLower == "name" {
	//		return p.Tag.Name, nil
	//	}
	//
	//	if aliasLower == "slug" {
	//		return p.Tag.Slug, nil
	//	}
	//}

	// No match - store in FromExpressions as any
	if p.FromExpressions == nil {
		p.FromExpressions = make(map[string]any)
	}
	var value any
	// Store a pointer to value that we'll populate after scan
	unmarshalFunc := func() error {
		p.FromExpressions[ref.Name()] = value
		return nil
	}
	return &value, unmarshalFunc
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
