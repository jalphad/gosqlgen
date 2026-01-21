package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
)

// TagsDto represents the tags table
type TagsDto struct {
	Id   *int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Posts []PostsDto `manytomany:"post_tags"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]any `json:"from_expressions,omitempty"`
}

// TableName returns the table name for TagsDto
func (t *TagsDto) TableName() string {
	return "tags"
}

// scanInto scans a row into a struct
func (t *TagsDto) ScanInto(row pgx.Row, stmt ast.SqlStatement) error {
	var scanDest []any
	var jsonUnmarshalFuncs []func() error

	columns := stmt.Returns()
	if len(columns) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range columns {
			destPtr, unmarshalFunc := t.getScanDestForField(field)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		// Default behavior - scan all table columns + joined tables
		scanDest = append(scanDest, &t.Id)
		scanDest = append(scanDest, &t.Name)
		scanDest = append(scanDest, &t.Slug)

		// Add joined table columns if active
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
func (t *TagsDto) getScanDestForField(ref ast.NamedExpression) (any, func() error) {
	var refTable string
	if split := strings.Split(ast.Render(ref, &[]any{}), "."); len(split) == 2 {
		refTable = split[0]
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(ref.Name())
	// Check if alias matches a many-to-many collection field
	if aliasLower == "posts" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []PostsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Posts: %w", err)
				}
				t.Posts = items
			} else {
				t.Posts = []PostsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}

	// Check if alias matches a regular table column
	if refTable == "tags" {

		if aliasLower == "id" {
			return &t.Id, nil
		}

		if aliasLower == "name" {
			return &t.Name, nil
		}

		if aliasLower == "slug" {
			return &t.Slug, nil
		}
	}

	// No match - store in FromExpressions as any
	if t.FromExpressions == nil {
		t.FromExpressions = make(map[string]any)
	}
	var value any
	// Store a pointer to value that we'll populate after scan
	unmarshalFunc := func() error {
		t.FromExpressions[ref.Name()] = value
		return nil
	}
	return &value, unmarshalFunc
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

// LoadPosts loads associated posts through post_tags
func (t *TagsDto) LoadPosts(ctx context.Context, db *DB) error {
	if t.Id == nil {
		return nil
	}

	// Query junction table to get referenced IDs
	qry := "SELECT post_id FROM post_tags WHERE tag_id = $1"

	rows, err := db.pool.Query(ctx, qry, *t.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		t.Posts = []PostsDto{}
		return nil
	}

	results, err := db.Posts().Where(posts.Id().In(ast.NewSQLType(ids))).Find(ctx)
	if err != nil {
		return err
	}
	t.Posts = results
	return nil
}
