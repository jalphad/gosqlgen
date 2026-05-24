package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
)

// CommentsDto represents the comments table
type CommentsDto struct {
	Id         *int64     `db:"id" json:"id"`
	PostId     int64      `db:"post_id" json:"post_id"`
	UserId     uuid.UUID  `db:"user_id" json:"user_id"`
	Content    string     `db:"content" json:"content"`
	IsApproved *bool      `db:"is_approved" json:"is_approved"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	TestDate   *time.Time `db:"test_date" json:"test_date"`

	// Joined relationships (populated when corresponding Join method is called)
	Post *PostsDto `joined:"posts" fk:"post_id"`
	User *UsersDto `joined:"users" fk:"user_id"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]any `json:"from_expressions,omitempty"`
}

// TableName returns the table name for CommentsDto
func (c *CommentsDto) TableName() string {
	return "comments"
}

func (c *CommentsDto) UnmarshalJSON(b []byte) error {
	type CommentsDto_ CommentsDto
	type DtoWrapper struct {
		CommentsDto_
		TestDate DateWrapper `db:"test_date" json:"test_date"`
	}

	var wrapper DtoWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}

	*c = CommentsDto(wrapper.CommentsDto_)
	c.TestDate = &wrapper.TestDate.Time

	return nil
}

// scanInto scans a row into a struct
func (c *CommentsDto) ScanInto(row pgx.Row, stmt ast.SqlStatement) error {
	var scanDest []any
	var jsonUnmarshalFuncs []func() error

	columns := stmt.Returns()
	if len(columns) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range columns {
			destPtr, unmarshalFunc := c.getScanDestForField(field)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		// Default behavior - scan all table columns + joined tables
		scanDest = append(scanDest, &c.Id)
		scanDest = append(scanDest, &c.PostId)
		scanDest = append(scanDest, &c.UserId)
		scanDest = append(scanDest, &c.Content)
		scanDest = append(scanDest, &c.IsApproved)
		scanDest = append(scanDest, &c.CreatedAt)
		scanDest = append(scanDest, &c.TestDate)
		// Add joined table columns
		ctx := stmt.GetQueryContext()
		if ctx != nil && ctx.JoinedTables != nil {
			if _, ok := ctx.JoinedTables["posts"]; ok {
				joinedPost := &PostsDto{}
				scanDest = append(scanDest, &joinedPost.Id)
				scanDest = append(scanDest, &joinedPost.UserId)
				scanDest = append(scanDest, &joinedPost.Title)
				scanDest = append(scanDest, &joinedPost.Content)
				scanDest = append(scanDest, &joinedPost.Status)
				scanDest = append(scanDest, &joinedPost.PublishedAt)
				scanDest = append(scanDest, &joinedPost.ViewCount)
				scanDest = append(scanDest, &joinedPost.CreatedAt)
				scanDest = append(scanDest, &joinedPost.UpdatedAt)
				c.Post = joinedPost
			}
			if _, ok := ctx.JoinedTables["users"]; ok {
				joinedUser := &UsersDto{}
				scanDest = append(scanDest, &joinedUser.Id)
				scanDest = append(scanDest, &joinedUser.Username)
				scanDest = append(scanDest, &joinedUser.Email)
				scanDest = append(scanDest, &joinedUser.FullName)
				scanDest = append(scanDest, &joinedUser.CreatedAt)
				scanDest = append(scanDest, &joinedUser.UpdatedAt)
				scanDest = append(scanDest, &joinedUser.IsActive)
				c.User = joinedUser
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
func (c *CommentsDto) getScanDestForField(ref ast.NamedExpression) (any, func() error) {
	var refTable string
	if split := strings.Split(ast.Render(ref, &[]any{}), "."); len(split) == 2 {
		refTable = split[0]
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(ref.Name())

	// Check if alias matches a regular table column
	if refTable == "comments" {

		if aliasLower == "id" {
			return &c.Id, nil
		}

		if aliasLower == "post_id" || aliasLower == "postid" {
			return &c.PostId, nil
		}

		if aliasLower == "user_id" || aliasLower == "userid" {
			return &c.UserId, nil
		}

		if aliasLower == "content" {
			return &c.Content, nil
		}

		if aliasLower == "is_approved" || aliasLower == "isapproved" {
			return &c.IsApproved, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return &c.CreatedAt, nil
		}

		if aliasLower == "test_date" || aliasLower == "testdate" {
			return &c.TestDate, nil
		}
	}
	// Check if it matches a joined table column
	if refTable == "posts" {
		if c.Post == nil {
			c.Post = &PostsDto{}
		}

		if aliasLower == "id" {
			return c.Post.Id, nil
		}

		if aliasLower == "user_id" || aliasLower == "userid" {
			return c.Post.UserId, nil
		}

		if aliasLower == "title" {
			return c.Post.Title, nil
		}

		if aliasLower == "content" {
			return c.Post.Content, nil
		}

		if aliasLower == "status" {
			return c.Post.Status, nil
		}

		if aliasLower == "published_at" || aliasLower == "publishedat" {
			return c.Post.PublishedAt, nil
		}

		if aliasLower == "view_count" || aliasLower == "viewcount" {
			return c.Post.ViewCount, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return c.Post.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return c.Post.UpdatedAt, nil
		}
	}
	if refTable == "users" {
		if c.User == nil {
			c.User = &UsersDto{}
		}

		if aliasLower == "id" {
			return c.User.Id, nil
		}

		if aliasLower == "username" {
			return c.User.Username, nil
		}

		if aliasLower == "email" {
			return c.User.Email, nil
		}

		if aliasLower == "full_name" || aliasLower == "fullname" {
			return c.User.FullName, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return c.User.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return c.User.UpdatedAt, nil
		}

		if aliasLower == "is_active" || aliasLower == "isactive" {
			return c.User.IsActive, nil
		}
	}

	// No match - store in FromExpressions as any
	if c.FromExpressions == nil {
		c.FromExpressions = make(map[string]any)
	}
	var value any
	// Store a pointer to value that we'll populate after scan
	unmarshalFunc := func() error {
		c.FromExpressions[ref.Name()] = value
		return nil
	}
	return &value, unmarshalFunc
}

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (c *CommentsDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "id" {
		if c.Id == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*c.Id), nil
	}

	if columnLower == "post_id" || columnLower == "postid" {
		return ast.NewSQLType(c.PostId), nil
	}

	if columnLower == "user_id" || columnLower == "userid" {
		return ast.NewSQLType(c.UserId), nil
	}

	if columnLower == "content" {
		return ast.NewSQLType(c.Content), nil
	}

	if columnLower == "is_approved" || columnLower == "isapproved" {
		if c.IsApproved == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*c.IsApproved), nil
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		if c.CreatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*c.CreatedAt), nil
	}

	if columnLower == "test_date" || columnLower == "testdate" {
		if c.TestDate == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*c.TestDate), nil
	}

	return nil, errors.New("unknown column")
}

// CommentsDtos is a collection of CommentsDto DTOs used for bulk operations
type CommentsDtos []*CommentsDto

// Id returns a slice of all Id values from the collection
func (d CommentsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(d))
	for idx, dto := range d {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

// PostId returns a slice of all PostId values from the collection
func (d CommentsDtos) PostId() ast.OfType[[]int64] {
	out := make([]int64, len(d))
	for idx, dto := range d {
		out[idx] = dto.PostId
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

// UserId returns a slice of all UserId values from the collection
func (d CommentsDtos) UserId() ast.OfType[[]uuid.UUID] {
	out := make([]uuid.UUID, len(d))
	for idx, dto := range d {
		out[idx] = dto.UserId
	}
	return ast.SetType[[]uuid.UUID](ast.NewLiteralExpression(out))
}

// Content returns a slice of all Content values from the collection
func (d CommentsDtos) Content() ast.OfType[[]string] {
	out := make([]string, len(d))
	for idx, dto := range d {
		out[idx] = dto.Content
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

// IsApproved returns a slice of all IsApproved values from the collection
func (d CommentsDtos) IsApproved() ast.OfType[[]bool] {
	out := make([]*bool, len(d))
	for idx, dto := range d {
		out[idx] = dto.IsApproved
	}
	return ast.SetType[[]bool](ast.NewLiteralExpression(out))
}

// CreatedAt returns a slice of all CreatedAt values from the collection
func (d CommentsDtos) CreatedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(d))
	for idx, dto := range d {
		out[idx] = dto.CreatedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

// TestDate returns a slice of all TestDate values from the collection
func (d CommentsDtos) TestDate() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(d))
	for idx, dto := range d {
		out[idx] = dto.TestDate
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (d CommentsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(d)*len(columns))
	for _, column := range columns {
		// Normalize column name for matching (lowercase)
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "post_id" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.PostId))
			}
		}

		if columnLower == "user_id" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.UserId))
			}
		}

		if columnLower == "content" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.Content))
			}
		}

		if columnLower == "is_approved" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.IsApproved))
			}
		}

		if columnLower == "created_at" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "test_date" {
			for _, dto := range d {
				im = append(im, ast.NewLiteralExpression(dto.TestDate))
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

func (d CommentsDtos) GetArgs(column ast.NamedExpression) ast.Expression {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(column.Name())

	if columnLower == "id" {
		out := d.Id()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "post_id" {
		out := d.PostId()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "user_id" {
		out := d.UserId()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "content" {
		out := d.Content()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "is_approved" {
		out := d.IsApproved()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "created_at" {
		out := d.CreatedAt()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "test_date" {
		out := d.TestDate()
		return ast.NewLiteralExpression(out)
	}

	return ast.NewErrorExpression(errors.New("unknown column"))
}
