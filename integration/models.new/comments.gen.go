package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type CommentsDtos []*CommentsDto

func (c CommentsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(c))
	for idx, dto := range c {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) PostId() ast.OfType[[]int64] {
	out := make([]int64, len(c))
	for idx, dto := range c {
		out[idx] = dto.PostId
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) UserId() ast.OfType[[]uuid.UUID] {
	out := make([]uuid.UUID, len(c))
	for idx, dto := range c {
		out[idx] = dto.UserId
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) Content() ast.OfType[[]string] {
	out := make([]string, len(c))
	for idx, dto := range c {
		out[idx] = dto.Content
	}
	return ast.NewSQLType(out)
}

func (c CommentsDtos) IsApproved() ast.OfType[[]bool] {
	out := make([]*bool, len(c))
	for idx, dto := range c {
		out[idx] = dto.IsApproved
	}
	return ast.SetType[[]bool](ast.NewLiteralExpression(out))
}

func (c CommentsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(c)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "post_id" || columnLower == "postid" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.PostId))
			}
		}

		if columnLower == "user_id" || columnLower == "userid" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.UserId))
			}
		}

		if columnLower == "content" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.Content))
			}
		}

		if columnLower == "is_approved" || columnLower == "isapproved" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.IsApproved))
			}
		}

		if columnLower == "created_at" || columnLower == "createdat" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "test_date" || columnLower == "testdate" {
			for _, dto := range c {
				im = append(im, ast.NewLiteralExpression(dto.TestDate))
			}
		}
	}

	values := make([]ast.Expression, 0, len(c))
	for i := 0; i < len(c); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(c)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

func (c CommentsDtos) GetArgs(column ast.NamedExpression) ast.Expression {
	columnLower := strings.ToLower(column.Name())

	if columnLower == "id" {
		out := c.Id()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "post_id" || columnLower == "postid" {
		out := c.PostId()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "user_id" || columnLower == "userid" {
		out := c.UserId()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "content" {
		out := make([]string, len(c))
		for idx, dto := range c {
			out[idx] = dto.Content
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "is_approved" || columnLower == "isapproved" {
		out := make([]*bool, len(c))
		for idx, dto := range c {
			out[idx] = dto.IsApproved
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		out := make([]*time.Time, len(c))
		for idx, dto := range c {
			out[idx] = dto.CreatedAt
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "test_date" || columnLower == "testdate" {
		out := make([]*time.Time, len(c))
		for idx, dto := range c {
			out[idx] = dto.TestDate
		}
		return ast.NewLiteralExpression(out)
	}

	return nil
}

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
			destPtr, unmarshalFunc := c.getScanDestForField(field, stmt)
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

		// Add joined table columns if active (using context)
		ctx := stmt.GetQueryContext()
		if ctx != nil && ctx.JoinedTables != nil {
			if _, hasPosts := ctx.JoinedTables["posts"]; hasPosts {
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
			if _, hasUsers := ctx.JoinedTables["users"]; hasUsers {
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
func (c *CommentsDto) getScanDestForField(ref ast.NamedExpression, stmt ast.SqlStatement) (any, func() error) {
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
	// Check if it matches a joined table column (using context)
	ctx := stmt.GetQueryContext()
	if ctx != nil && ctx.JoinedTables != nil {
		if refTable == "posts" {
			_, hasPosts := ctx.JoinedTables["posts"]
			if hasPosts && c.Post == nil {
				c.Post = &PostsDto{}
			}

			if aliasLower == "id" {
				return &c.Post.Id, nil
			}

			if aliasLower == "user_id" || aliasLower == "userid" {
				return &c.Post.UserId, nil
			}

			if aliasLower == "title" {
				return &c.Post.Title, nil
			}

			if aliasLower == "content" {
				return &c.Post.Content, nil
			}

			if aliasLower == "status" {
				return &c.Post.Status, nil
			}

			if aliasLower == "published_at" || aliasLower == "publishedat" {
				return &c.Post.PublishedAt, nil
			}

			if aliasLower == "view_count" || aliasLower == "viewcount" {
				return &c.Post.ViewCount, nil
			}

			if aliasLower == "created_at" || aliasLower == "createdat" {
				return &c.Post.CreatedAt, nil
			}

			if aliasLower == "updated_at" || aliasLower == "updatedat" {
				return &c.Post.UpdatedAt, nil
			}
		}
		if refTable == "users" {
			_, hasUsers := ctx.JoinedTables["users"]
			if hasUsers && c.User == nil {
				c.User = &UsersDto{}
			}

			if aliasLower == "id" {
				return &c.User.Id, nil
			}

			if aliasLower == "username" {
				return &c.User.Username, nil
			}

			if aliasLower == "email" {
				return &c.User.Email, nil
			}

			if aliasLower == "full_name" || aliasLower == "fullname" {
				return &c.User.FullName, nil
			}

			if aliasLower == "created_at" || aliasLower == "createdat" {
				return &c.User.CreatedAt, nil
			}

			if aliasLower == "updated_at" || aliasLower == "updatedat" {
				return &c.User.UpdatedAt, nil
			}

			if aliasLower == "is_active" || aliasLower == "isactive" {
				return &c.User.IsActive, nil
			}
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
