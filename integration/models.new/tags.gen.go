package models

import (
	"errors"
	"strings"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

// TagsDto represents the tags table
type TagsDto struct {
	Id   *int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Posts []PostsDto `manytomany:"post_tags"`
}

// TableName returns the table name for TagsDto
func (t *TagsDto) TableName() string {
	return "tags"
}

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (t *TagsDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "id" {
		if t.Id == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*t.Id), nil
	}

	if columnLower == "name" {
		return ast.NewSQLType(t.Name), nil
	}

	if columnLower == "slug" {
		return ast.NewSQLType(t.Slug), nil
	}

	return nil, errors.New("unknown column")
}

//// LoadPosts loads associated posts through post_tags
//func (t *TagsDto) LoadPosts(ctx context.Context, db *DB) error {
//	if t.Id == nil {
//		return nil
//	}
//
//	// Query junction table to get referenced IDs
//	qry := "SELECT post_id FROM post_tags WHERE tag_id = $1"
//
//	rows, err := db.pool.Query(ctx, qry, *t.Id)
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//
//	var ids []int64
//	for rows.Next() {
//		var id int64
//		if err := rows.Scan(&id); err != nil {
//			return err
//		}
//		ids = append(ids, id)
//	}
//
//	if len(ids) == 0 {
//		t.Posts = []PostsDto{}
//		return nil
//	}
//
//	results, err := db.Posts().Where(posts.Id().In(ast.NewSQLType(ids))).Find(ctx)
//	if err != nil {
//		return err
//	}
//	t.Posts = results
//	return nil
//}
