package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
)

// UsersDto represents the users table
type UsersDto struct {
	Id        *uuid.UUID `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	FullName  *string    `db:"full_name" json:"full_name"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	IsActive  *bool      `db:"is_active" json:"is_active"`

	// One-to-many reverse relationships (populated via LoadXxx methods)
	Comments []CommentsDto `reverse:"comments" fk:"user_id"`
	Posts    []PostsDto    `reverse:"posts" fk:"user_id"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]any `json:"from_expressions,omitempty"`
}

// TableName returns the table name for UsersDto
func (u *UsersDto) TableName() string {
	return "users"
}

func (u *UsersDto) UnmarshalJSON(b []byte) error {
	type UsersDto_ UsersDto
	type DtoWrapper struct {
		UsersDto_
		CreatedAt NoTimezoneWrapper `db:"created_at" json:"created_at"`
		UpdatedAt NoTimezoneWrapper `db:"updated_at" json:"updated_at"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*u = UsersDto(wrapper.UsersDto_)
	u.CreatedAt = &wrapper.CreatedAt.Time
	u.UpdatedAt = &wrapper.UpdatedAt.Time

	return nil
}

// scanInto scans a row into a struct
func (u *UsersDto) ScanInto(row pgx.Row, stmt ast.SqlStatement) error {
	var scanDest []any
	var jsonUnmarshalFuncs []func() error

	columns := stmt.Returns()
	if len(columns) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range columns {
			destPtr, unmarshalFunc := u.getScanDestForField(field)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		// Default behavior - scan all table columns + joined tables
		scanDest = append(scanDest, &u.Id)
		scanDest = append(scanDest, &u.Username)
		scanDest = append(scanDest, &u.Email)
		scanDest = append(scanDest, &u.FullName)
		scanDest = append(scanDest, &u.CreatedAt)
		scanDest = append(scanDest, &u.UpdatedAt)
		scanDest = append(scanDest, &u.IsActive)
		// Add joined table columns
		ctx := stmt.GetQueryContext()
		if ctx != nil && ctx.JoinedTables != nil {
			if _, ok := ctx.JoinedTables["comments"]; ok {
				joinedComments := CommentsDto{}
				scanDest = append(scanDest, &joinedComments.Id)
				scanDest = append(scanDest, &joinedComments.PostId)
				scanDest = append(scanDest, &joinedComments.UserId)
				scanDest = append(scanDest, &joinedComments.Content)
				scanDest = append(scanDest, &joinedComments.IsApproved)
				scanDest = append(scanDest, &joinedComments.CreatedAt)
				scanDest = append(scanDest, &joinedComments.TestDate)
				u.Comments = []CommentsDto{joinedComments}
			}
			if _, ok := ctx.JoinedTables["posts"]; ok {
				joinedPosts := PostsDto{}
				scanDest = append(scanDest, &joinedPosts.Id)
				scanDest = append(scanDest, &joinedPosts.UserId)
				scanDest = append(scanDest, &joinedPosts.Title)
				scanDest = append(scanDest, &joinedPosts.Content)
				scanDest = append(scanDest, &joinedPosts.Status)
				scanDest = append(scanDest, &joinedPosts.PublishedAt)
				scanDest = append(scanDest, &joinedPosts.ViewCount)
				scanDest = append(scanDest, &joinedPosts.CreatedAt)
				scanDest = append(scanDest, &joinedPosts.UpdatedAt)
				u.Posts = []PostsDto{joinedPosts}
			}
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
func (u *UsersDto) getScanDestForField(ref ast.NamedExpression) (any, func() error) {
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
				u.Comments = items
			} else {
				u.Comments = []CommentsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}
	if aliasLower == "posts" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []PostsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Posts: %w", err)
				}
				u.Posts = items
			} else {
				u.Posts = []PostsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}

	// Check if alias matches a regular table column
	if refTable == "users" {

		if aliasLower == "id" {
			return &u.Id, nil
		}

		if aliasLower == "username" {
			return &u.Username, nil
		}

		if aliasLower == "email" {
			return &u.Email, nil
		}

		if aliasLower == "full_name" || aliasLower == "fullname" {
			return &u.FullName, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return &u.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return &u.UpdatedAt, nil
		}

		if aliasLower == "is_active" || aliasLower == "isactive" {
			return &u.IsActive, nil
		}
	}

	// No match - store in FromExpressions as any
	if u.FromExpressions == nil {
		u.FromExpressions = make(map[string]any)
	}
	var value any
	// Store a pointer to value that we'll populate after scan
	unmarshalFunc := func() error {
		u.FromExpressions[ref.Name()] = value
		return nil
	}
	return &value, unmarshalFunc
}

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (u *UsersDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "id" {
		if u.Id == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.Id), nil
	}

	if columnLower == "username" {
		return ast.NewSQLType(u.Username), nil
	}

	if columnLower == "email" {
		return ast.NewSQLType(u.Email), nil
	}

	if columnLower == "full_name" || columnLower == "fullname" {
		if u.FullName == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.FullName), nil
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		if u.CreatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.CreatedAt), nil
	}

	if columnLower == "updated_at" || columnLower == "updatedat" {
		if u.UpdatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.UpdatedAt), nil
	}

	if columnLower == "is_active" || columnLower == "isactive" {
		if u.IsActive == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*u.IsActive), nil
	}

	return nil, errors.New("unknown column")
}

// LoadPosts loads associated posts for this users
func (u *UsersDto) LoadPosts(ctx context.Context, db *DB) error {
	if u.Id == nil {
		return nil
	}
	results, err := db.Posts().Select().Where(posts.UserId().Eq(ast.NewSQLType(*u.Id))).Find(ctx)
	if err != nil {
		return err
	}
	u.Posts = results
	return nil
}

// LoadComments loads associated comments for this users
func (u *UsersDto) LoadComments(ctx context.Context, db *DB) error {
	if u.Id == nil {
		return nil
	}
	results, err := db.Comments().Select().Where(comments.UserId().Eq(ast.NewSQLType(*u.Id))).Find(ctx)
	if err != nil {
		return err
	}
	u.Comments = results
	return nil
}

// UsersDtos is a collection of UsersDto DTOs used for bulk operations
type UsersDtos []*UsersDto

// Id returns a slice of all Id values from the collection
func (d UsersDtos) Id() ast.OfType[[]uuid.UUID] {
	out := make([]*uuid.UUID, len(d))
	for idx, dto := range d {
		out[idx] = dto.Id
	}
	return ast.SetType[[]uuid.UUID](ast.NewLiteralExpression(out))
}

// Username returns a slice of all Username values from the collection
func (d UsersDtos) Username() ast.OfType[[]string] {
	out := make([]string, len(d))
	for idx, dto := range d {
		out[idx] = dto.Username
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

// Email returns a slice of all Email values from the collection
func (d UsersDtos) Email() ast.OfType[[]string] {
	out := make([]string, len(d))
	for idx, dto := range d {
		out[idx] = dto.Email
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

// FullName returns a slice of all FullName values from the collection
func (d UsersDtos) FullName() ast.OfType[[]string] {
	out := make([]*string, len(d))
	for idx, dto := range d {
		out[idx] = dto.FullName
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

// CreatedAt returns a slice of all CreatedAt values from the collection
func (d UsersDtos) CreatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(d))
	for idx, dto := range d {
		out[idx] = dto.CreatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

// UpdatedAt returns a slice of all UpdatedAt values from the collection
func (d UsersDtos) UpdatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(d))
	for idx, dto := range d {
		out[idx] = dto.UpdatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

// IsActive returns a slice of all IsActive values from the collection
func (d UsersDtos) IsActive() ast.OfType[[]bool] {
	out := make([]*bool, len(d))
	for idx, dto := range d {
		out[idx] = dto.IsActive
	}
	return ast.SetType[[]bool](ast.NewLiteralExpression(out))
}

func (d UsersDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(d)*len(columns))
	for _, column := range columns {
		// Normalize column name for matching (lowercase)
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "username" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.Username))
			}
		}

		if columnLower == "email" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.Email))
			}
		}

		if columnLower == "full_name" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.FullName))
			}
		}

		if columnLower == "created_at" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "updated_at" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.UpdatedAt))
			}
		}

		if columnLower == "is_active" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.IsActive))
			}
		}

	}

	values := make([]ast.Expression, 0, len(d))
	for i := 0; i < len(d); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(d)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

func (d UsersDtos) GetArgs(column ast.NamedExpression) ast.Expression {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(column.Name())

	if columnLower == "id" {
		out := d.Id()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "username" {
		out := d.Username()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "email" {
		out := d.Email()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "full_name" {
		out := d.FullName()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "created_at" {
		out := d.CreatedAt()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "updated_at" {
		out := d.UpdatedAt()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "is_active" {
		out := d.IsActive()
		return ast.NewLiteralExpression(out)
	}

	return ast.NewErrorExpression(errors.New("unknown column"))
}
