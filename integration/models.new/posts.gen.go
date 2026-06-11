package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type PostsDtos []*PostsDto

func (p PostsDtos) Id() ast.OfType[[]int64] {
	out := make([]*int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.Id
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (p PostsDtos) UserId() ast.OfType[[]uuid.UUID] {
	out := make([]uuid.UUID, len(p))
	for idx, dto := range p {
		out[idx] = dto.UserId
	}
	return ast.NewSQLType(out)
}

func (p PostsDtos) Title() ast.OfType[[]string] {
	out := make([]string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Title
	}
	return ast.NewSQLType(out)
}

func (p PostsDtos) Content() ast.OfType[[]string] {
	out := make([]*string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Content
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

func (p PostsDtos) Status() ast.OfType[[]string] {
	out := make([]*string, len(p))
	for idx, dto := range p {
		out[idx] = dto.Status
	}
	return ast.SetType[[]string](ast.NewLiteralExpression(out))
}

func (p PostsDtos) PublishedAt() ast.OfType[[]time.Time] {
	out := make([]*time.Time, len(p))
	for idx, dto := range p {
		out[idx] = dto.PublishedAt
	}
	return ast.SetType[[]time.Time](ast.NewLiteralExpression(out))
}

func (p PostsDtos) ViewCount() ast.OfType[[]int64] {
	out := make([]*int64, len(p))
	for idx, dto := range p {
		out[idx] = dto.ViewCount
	}
	return ast.SetType[[]int64](ast.NewLiteralExpression(out))
}

func (p PostsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(p)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "user_id" || columnLower == "userid" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.UserId))
			}
		}

		if columnLower == "title" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Title))
			}
		}

		if columnLower == "content" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Content))
			}
		}

		if columnLower == "status" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.Status))
			}
		}

		if columnLower == "published_at" || columnLower == "publishedat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.PublishedAt))
			}
		}

		if columnLower == "view_count" || columnLower == "viewcount" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.ViewCount))
			}
		}

		if columnLower == "created_at" || columnLower == "createdat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.CreatedAt))
			}
		}

		if columnLower == "updated_at" || columnLower == "updatedat" {
			for _, dto := range p {
				im = append(im, ast.NewLiteralExpression(dto.UpdatedAt))
			}
		}
	}

	values := make([]ast.Expression, 0, len(p))
	for i := 0; i < len(p); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(p)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

func (p PostsDtos) GetArgs(column ast.NamedExpression) ast.Expression {
	columnLower := strings.ToLower(column.Name())

	if columnLower == "id" {
		out := p.Id()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "user_id" || columnLower == "userid" {
		out := p.UserId()
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "title" {
		out := make([]string, len(p))
		for idx, dto := range p {
			out[idx] = dto.Title
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "content" {
		out := make([]*string, len(p))
		for idx, dto := range p {
			out[idx] = dto.Content
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "status" {
		out := make([]*string, len(p))
		for idx, dto := range p {
			out[idx] = dto.Status
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "published_at" || columnLower == "publishedat" {
		out := make([]*time.Time, len(p))
		for idx, dto := range p {
			out[idx] = dto.PublishedAt
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "view_count" || columnLower == "viewcount" {
		out := make([]*int64, len(p))
		for idx, dto := range p {
			out[idx] = dto.ViewCount
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		out := make([]*time.Time, len(p))
		for idx, dto := range p {
			out[idx] = dto.CreatedAt
		}
		return ast.NewLiteralExpression(out)
	}

	if columnLower == "updated_at" || columnLower == "updatedat" {
		out := make([]*time.Time, len(p))
		for idx, dto := range p {
			out[idx] = dto.UpdatedAt
		}
		return ast.NewLiteralExpression(out)
	}

	return nil
}

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

// GetArg returns an expression which will be converted to a parameter in the SQL query
func (p *PostsDto) GetArg(ref ast.NamedExpression) (ast.Expression, error) {
	// Normalize column name for matching (lowercase)
	columnLower := strings.ToLower(ref.Name())

	if columnLower == "id" {
		if p.Id == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.Id), nil
	}

	if columnLower == "user_id" || columnLower == "userid" {
		return ast.NewSQLType(p.UserId), nil
	}

	if columnLower == "title" {
		return ast.NewSQLType(p.Title), nil
	}

	if columnLower == "content" {
		if p.Content == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.Content), nil
	}

	if columnLower == "status" {
		if p.Status == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.Status), nil
	}

	if columnLower == "published_at" || columnLower == "publishedat" {
		if p.PublishedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.PublishedAt), nil
	}

	if columnLower == "view_count" || columnLower == "viewcount" {
		if p.ViewCount == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.ViewCount), nil
	}

	if columnLower == "created_at" || columnLower == "createdat" {
		if p.CreatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.CreatedAt), nil
	}

	if columnLower == "updated_at" || columnLower == "updatedat" {
		if p.UpdatedAt == nil {
			return ast.NewLiteralExpression(nil), nil
		}
		return ast.NewSQLType(*p.UpdatedAt), nil
	}

	return nil, errors.New("unknown column")
}

//
//// LoadTags loads associated tags through post_tags
//func (p *PostsDto) LoadTags(ctx context.Context, db *DB) error {
//	if p.Id == nil {
//		return nil
//	}
//
//	// Query junction table to get referenced IDs
//	qry := "SELECT tag_id FROM post_tags WHERE post_id = $1"
//
//	rows, err := db.pool.Query(ctx, qry, *p.Id)
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
//		p.Tags = []TagsDto{}
//		return nil
//	}
//
//	results, err := db.Tags().Where(tags.Id().In(ast.NewSQLType(ids))).Find(ctx)
//	if err != nil {
//		return err
//	}
//	p.Tags = results
//	return nil
//}
