package models

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
)

// PostsDto represents the posts table
type PostsDto struct {
	Id          *int64     `db:"id" json:"id"`
	UserId      uuid.UUID  `db:"user_id" json:"user_id"`
	Title       string     `db:"title" json:"title"`
	Content     *string    `db:"content" json:"content"`
	Status      *string    `db:"status" json:"status"`
	PublishedAt *time.Time `db:"published_at" json:"published_at"`
	ViewCount   *int64     `db:"view_count" json:"view_count"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`

	// Joined relationships (populated when corresponding Join method is called)
	User *UsersDto `joined:"users" fk:"user_id"`

	// One-to-many reverse relationships (populated via LoadXxx methods)
	Comments []CommentsDto `reverse:"comments" fk:"post_id"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Tags []TagsDto `manytomany:"post_tags"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]any `json:"from_expressions,omitempty"`
}

func (p *PostsDto) GetArgs(expressions []ast.NamedExpression) []any {
	//TODO implement me
	panic("implement me")
}

// TableName returns the table name for PostsDto
func (p *PostsDto) TableName() string {
	return "posts"
}

func (p *PostsDto) UnmarshalJSON(b []byte) error {
	type PostsDto_ PostsDto
	type DtoWrapper struct {
		PostsDto_
		PublishedAt NoTimezoneWrapper `db:"published_at" json:"published_at"`
		CreatedAt   NoTimezoneWrapper `db:"created_at" json:"created_at"`
		UpdatedAt   NoTimezoneWrapper `db:"updated_at" json:"updated_at"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*p = PostsDto(wrapper.PostsDto_)
	p.PublishedAt = &wrapper.PublishedAt.Time
	p.CreatedAt = &wrapper.CreatedAt.Time
	p.UpdatedAt = &wrapper.UpdatedAt.Time

	return nil
}

// scanInto scans a row into a struct
func (p *PostsDto) ScanInto(row pgx.Row, stmt ast.SqlStatement) error {
	var scanDest []any
	var jsonUnmarshalFuncs []func() error

	columns := stmt.Returns()
	if len(columns) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range columns {
			destPtr, unmarshalFunc := p.getScanDestForField(field, stmt)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		// Default behavior - scan all table columns + joined tables
		scanDest = append(scanDest, &p.Id)
		scanDest = append(scanDest, &p.UserId)
		scanDest = append(scanDest, &p.Title)
		scanDest = append(scanDest, &p.Content)
		scanDest = append(scanDest, &p.Status)
		scanDest = append(scanDest, &p.PublishedAt)
		scanDest = append(scanDest, &p.ViewCount)
		scanDest = append(scanDest, &p.CreatedAt)
		scanDest = append(scanDest, &p.UpdatedAt)

		// Add joined table columns if active
		if selectStmt, ok := stmt.(*ast.SelectStatement); ok && slices.Contains(selectStmt.GetJoinedTables(), "users") {
			joinedUser := &UsersDto{}
			scanDest = append(scanDest, &joinedUser.Id)
			scanDest = append(scanDest, &joinedUser.Username)
			scanDest = append(scanDest, &joinedUser.Email)
			scanDest = append(scanDest, &joinedUser.FullName)
			scanDest = append(scanDest, &joinedUser.CreatedAt)
			scanDest = append(scanDest, &joinedUser.UpdatedAt)
			scanDest = append(scanDest, &joinedUser.IsActive)
			p.User = joinedUser
		}
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
func (p *PostsDto) getScanDestForField(ref ast.NamedExpression, stmt ast.SqlStatement) (any, func() error) {
	var refTable string
	if split := strings.Split(ast.Render(ref, &[]any{}), "."); len(split) == 2 {
		refTable = split[0]
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(ref.Name())
	// Check if alias matches a reverse relationship collection field
	if aliasLower == "comments" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []CommentsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Comments: %w", err)
				}
				p.Comments = items
			} else {
				p.Comments = []CommentsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}
	// Check if alias matches a many-to-many collection field
	if aliasLower == "tags" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []TagsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Tags: %w", err)
				}
				p.Tags = items
			} else {
				p.Tags = []TagsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}

	// Check if alias matches a regular table column
	if refTable == "posts" {

		if aliasLower == "id" {
			return &p.Id, nil
		}

		if aliasLower == "user_id" || aliasLower == "userid" {
			return &p.UserId, nil
		}

		if aliasLower == "title" {
			return &p.Title, nil
		}

		if aliasLower == "content" {
			return &p.Content, nil
		}

		if aliasLower == "status" {
			return &p.Status, nil
		}

		if aliasLower == "published_at" || aliasLower == "publishedat" {
			return &p.PublishedAt, nil
		}

		if aliasLower == "view_count" || aliasLower == "viewcount" {
			return &p.ViewCount, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return &p.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return &p.UpdatedAt, nil
		}
	}
	// Check if it matches a joined table column
	if selectStmt, ok := stmt.(*ast.SelectStatement); ok && refTable == "users" && slices.Contains(selectStmt.GetJoinedTables(), "users") {
		if p.User == nil {
			p.User = &UsersDto{}
		}

		if aliasLower == "id" {
			return p.User.Id, nil
		}

		if aliasLower == "username" {
			return p.User.Username, nil
		}

		if aliasLower == "email" {
			return p.User.Email, nil
		}

		if aliasLower == "full_name" || aliasLower == "fullname" {
			return p.User.FullName, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return p.User.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return p.User.UpdatedAt, nil
		}

		if aliasLower == "is_active" || aliasLower == "isactive" {
			return p.User.IsActive, nil
		}
	}

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

// PrepareInsert returns the data about the columns to be inserted
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (p *PostsDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	var columns []string
	var placeholders []string
	var args []any
	argIdx := 1

	// Dynamically build column list based on non-nil values
	// Required field (not nullable, no default)
	columns = append(columns, "user_id")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, p.UserId)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "title")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, p.Title)
	argIdx++

	if p.Content != nil {
		columns = append(columns, "content")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.Content)
		argIdx++
	}

	if p.Status != nil {
		columns = append(columns, "status")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.Status)
		argIdx++
	}

	if p.PublishedAt != nil {
		columns = append(columns, "published_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.PublishedAt)
		argIdx++
	}

	if p.ViewCount != nil {
		columns = append(columns, "view_count")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.ViewCount)
		argIdx++
	}

	if p.CreatedAt != nil {
		columns = append(columns, "created_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.CreatedAt)
		argIdx++
	}

	if p.UpdatedAt != nil {
		columns = append(columns, "updated_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, p.UpdatedAt)
		argIdx++
	}

	_ = fmt.Sprintf("INSERT INTO posts (%s) VALUES (%s) RETURNING id",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	return nil, nil
}

// LoadComments loads associated comments for this posts
func (p *PostsDto) LoadComments(ctx context.Context, db *DB) error {
	if p.Id == nil {
		return nil
	}
	results, err := db.Comments().Select().Where(comments.PostId().Eq(ast.NewSQLType(*p.Id))).Find(ctx)
	if err != nil {
		return err
	}
	p.Comments = results
	return nil
}

// LoadTags loads associated tags through post_tags
func (p *PostsDto) LoadTags(ctx context.Context, db *DB) error {
	if p.Id == nil {
		return nil
	}

	// Query junction table to get referenced IDs
	qry := "SELECT tag_id FROM post_tags WHERE post_id = $1"

	rows, err := db.pool.Query(ctx, qry, *p.Id)
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
		p.Tags = []TagsDto{}
		return nil
	}

	results, err := db.Tags().Select().Where(tags.Id().In(ast.NewSQLType(ids))).Find(ctx)
	if err != nil {
		return err
	}
	p.Tags = results
	return nil
}
