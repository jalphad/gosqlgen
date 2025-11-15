package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostsDto represents the posts table
type PostsDto struct {
	Id          *int64     `db:"id" json:"id"`
	UserId      string     `db:"user_id" json:"user_id"`
	Title       string     `db:"title" json:"title"`
	Content     *string    `db:"content" json:"content"`
	Status      *string    `db:"status" json:"status"`
	PublishedAt *time.Time `db:"published_at" json:"published_at"`
	ViewCount   *int64     `db:"view_count" json:"view_count"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`

	// Joined relationships (populated when corresponding Join method is called)
	User *UsersDto `joined:"users" fk:"user_id"`
}

// TableName returns the table name for PostsDto
func (p *PostsDto) TableName() string {
	return "posts"
}

// PostsFields provides type-safe field references for PostsDto
type PostsFields struct{}

// PostsTable returns field references for PostsDto
var PostsTable = PostsFields{}

// Id returns a field reference for PostsDto.Id
func (p PostsFields) Id() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "id",
		GoType: "int64",
	}
}

// UserId returns a field reference for PostsDto.UserId
func (p PostsFields) UserId() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "user_id",
		GoType: "string",
	}
}

// Title returns a field reference for PostsDto.Title
func (p PostsFields) Title() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "title",
		GoType: "string",
	}
}

// Content returns a field reference for PostsDto.Content
func (p PostsFields) Content() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "content",
		GoType: "string",
	}
}

// Status returns a field reference for PostsDto.Status
func (p PostsFields) Status() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "status",
		GoType: "string",
	}
}

// PublishedAt returns a field reference for PostsDto.PublishedAt
func (p PostsFields) PublishedAt() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "published_at",
		GoType: "time.Time",
	}
}

// ViewCount returns a field reference for PostsDto.ViewCount
func (p PostsFields) ViewCount() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "view_count",
		GoType: "int64",
	}
}

// CreatedAt returns a field reference for PostsDto.CreatedAt
func (p PostsFields) CreatedAt() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "created_at",
		GoType: "time.Time",
	}
}

// UpdatedAt returns a field reference for PostsDto.UpdatedAt
func (p PostsFields) UpdatedAt() *FieldRef {
	return &FieldRef{
		Table:  "posts",
		Column: "updated_at",
		GoType: "time.Time",
	}
}

// PostsQuery is a type-safe query builder for PostsDto
type PostsQuery struct {
	pool         *pgxpool.Pool
	tx           pgx.Tx
	selectFields []*FieldRef
	conditions   []Condition
	joins        []JoinClause
	activeJoins  map[string]bool // tracks which tables are joined for dynamic column selection
	orderBy      []OrderClause
	groupBy      []*FieldRef
	having       []Condition
	limit        *int
	offset       *int
	logical      LogicalOp
}

// NewPostsQuery creates a new query builder
func NewPostsQuery(pool *pgxpool.Pool) *PostsQuery {
	return &PostsQuery{
		pool:        pool,
		activeJoins: make(map[string]bool),
		logical:     AND,
	}
}

// WithTx sets the transaction
func (q *PostsQuery) WithTx(tx pgx.Tx) *PostsQuery {
	q.tx = tx
	return q
}

// Select specifies fields to select using type-safe field references
func (q *PostsQuery) Select(fields ...*FieldRef) *PostsQuery {
	q.selectFields = fields
	return q
}

// SelectAll selects all fields
func (q *PostsQuery) SelectAll() *PostsQuery {
	q.selectFields = []*FieldRef{
		PostsTable.Id(),
		PostsTable.UserId(),
		PostsTable.Title(),
		PostsTable.Content(),
		PostsTable.Status(),
		PostsTable.PublishedAt(),
		PostsTable.ViewCount(),
		PostsTable.CreatedAt(),
		PostsTable.UpdatedAt(),
	}
	return q
}

// WhereId adds a condition for Id
func (q *PostsQuery) WhereId(op ComparisonOp, value int64) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Id(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIdEq adds an equality condition for Id
func (q *PostsQuery) WhereIdEq(value int64) *PostsQuery {
	return q.WhereId(EQ, value)
}

// WhereIdNotEq adds a not-equal condition for Id
func (q *PostsQuery) WhereIdNotEq(value int64) *PostsQuery {
	return q.WhereId(NEQ, value)
}

// WhereIdGt adds a greater-than condition for Id
func (q *PostsQuery) WhereIdGt(value int64) *PostsQuery {
	return q.WhereId(GT, value)
}

// WhereIdGte adds a greater-than-or-equal condition for Id
func (q *PostsQuery) WhereIdGte(value int64) *PostsQuery {
	return q.WhereId(GTE, value)
}

// WhereIdLt adds a less-than condition for Id
func (q *PostsQuery) WhereIdLt(value int64) *PostsQuery {
	return q.WhereId(LT, value)
}

// WhereIdLte adds a less-than-or-equal condition for Id
func (q *PostsQuery) WhereIdLte(value int64) *PostsQuery {
	return q.WhereId(LTE, value)
}

// WhereIdIn adds an IN condition for Id
func (q *PostsQuery) WhereIdIn(values ...int64) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Id(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderById adds ORDER BY for Id
func (q *PostsQuery) OrderById(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.Id(),
		Direction: dir,
	})
	return q
}

// WhereUserId adds a condition for UserId
func (q *PostsQuery) WhereUserId(op ComparisonOp, value string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UserId(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereUserIdEq adds an equality condition for UserId
func (q *PostsQuery) WhereUserIdEq(value string) *PostsQuery {
	return q.WhereUserId(EQ, value)
}

// WhereUserIdNotEq adds a not-equal condition for UserId
func (q *PostsQuery) WhereUserIdNotEq(value string) *PostsQuery {
	return q.WhereUserId(NEQ, value)
}

// WhereUserIdLike adds a LIKE condition for UserId
func (q *PostsQuery) WhereUserIdLike(pattern string) *PostsQuery {
	return q.WhereUserId(LIKE, pattern)
}

// WhereUserIdIn adds an IN condition for UserId
func (q *PostsQuery) WhereUserIdIn(values ...string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UserId(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByUserId adds ORDER BY for UserId
func (q *PostsQuery) OrderByUserId(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.UserId(),
		Direction: dir,
	})
	return q
}

// WhereTitle adds a condition for Title
func (q *PostsQuery) WhereTitle(op ComparisonOp, value string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Title(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereTitleEq adds an equality condition for Title
func (q *PostsQuery) WhereTitleEq(value string) *PostsQuery {
	return q.WhereTitle(EQ, value)
}

// WhereTitleNotEq adds a not-equal condition for Title
func (q *PostsQuery) WhereTitleNotEq(value string) *PostsQuery {
	return q.WhereTitle(NEQ, value)
}

// WhereTitleLike adds a LIKE condition for Title
func (q *PostsQuery) WhereTitleLike(pattern string) *PostsQuery {
	return q.WhereTitle(LIKE, pattern)
}

// WhereTitleIn adds an IN condition for Title
func (q *PostsQuery) WhereTitleIn(values ...string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Title(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByTitle adds ORDER BY for Title
func (q *PostsQuery) OrderByTitle(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.Title(),
		Direction: dir,
	})
	return q
}

// WhereContent adds a condition for Content
func (q *PostsQuery) WhereContent(op ComparisonOp, value string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Content(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereContentEq adds an equality condition for Content
func (q *PostsQuery) WhereContentEq(value string) *PostsQuery {
	return q.WhereContent(EQ, value)
}

// WhereContentNotEq adds a not-equal condition for Content
func (q *PostsQuery) WhereContentNotEq(value string) *PostsQuery {
	return q.WhereContent(NEQ, value)
}

// WhereContentLike adds a LIKE condition for Content
func (q *PostsQuery) WhereContentLike(pattern string) *PostsQuery {
	return q.WhereContent(LIKE, pattern)
}

// WhereContentIn adds an IN condition for Content
func (q *PostsQuery) WhereContentIn(values ...string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Content(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereContentIsNull adds an IS NULL condition for Content
func (q *PostsQuery) WhereContentIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Content(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereContentIsNotNull adds an IS NOT NULL condition for Content
func (q *PostsQuery) WhereContentIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Content(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByContent adds ORDER BY for Content
func (q *PostsQuery) OrderByContent(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.Content(),
		Direction: dir,
	})
	return q
}

// WhereStatus adds a condition for Status
func (q *PostsQuery) WhereStatus(op ComparisonOp, value string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Status(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereStatusEq adds an equality condition for Status
func (q *PostsQuery) WhereStatusEq(value string) *PostsQuery {
	return q.WhereStatus(EQ, value)
}

// WhereStatusNotEq adds a not-equal condition for Status
func (q *PostsQuery) WhereStatusNotEq(value string) *PostsQuery {
	return q.WhereStatus(NEQ, value)
}

// WhereStatusLike adds a LIKE condition for Status
func (q *PostsQuery) WhereStatusLike(pattern string) *PostsQuery {
	return q.WhereStatus(LIKE, pattern)
}

// WhereStatusIn adds an IN condition for Status
func (q *PostsQuery) WhereStatusIn(values ...string) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Status(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereStatusIsNull adds an IS NULL condition for Status
func (q *PostsQuery) WhereStatusIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Status(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereStatusIsNotNull adds an IS NOT NULL condition for Status
func (q *PostsQuery) WhereStatusIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.Status(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByStatus adds ORDER BY for Status
func (q *PostsQuery) OrderByStatus(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.Status(),
		Direction: dir,
	})
	return q
}

// WherePublishedAt adds a condition for PublishedAt
func (q *PostsQuery) WherePublishedAt(op ComparisonOp, value time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.PublishedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WherePublishedAtEq adds an equality condition for PublishedAt
func (q *PostsQuery) WherePublishedAtEq(value time.Time) *PostsQuery {
	return q.WherePublishedAt(EQ, value)
}

// WherePublishedAtNotEq adds a not-equal condition for PublishedAt
func (q *PostsQuery) WherePublishedAtNotEq(value time.Time) *PostsQuery {
	return q.WherePublishedAt(NEQ, value)
}

// WherePublishedAtGt adds a greater-than condition for PublishedAt
func (q *PostsQuery) WherePublishedAtGt(value time.Time) *PostsQuery {
	return q.WherePublishedAt(GT, value)
}

// WherePublishedAtGte adds a greater-than-or-equal condition for PublishedAt
func (q *PostsQuery) WherePublishedAtGte(value time.Time) *PostsQuery {
	return q.WherePublishedAt(GTE, value)
}

// WherePublishedAtLt adds a less-than condition for PublishedAt
func (q *PostsQuery) WherePublishedAtLt(value time.Time) *PostsQuery {
	return q.WherePublishedAt(LT, value)
}

// WherePublishedAtLte adds a less-than-or-equal condition for PublishedAt
func (q *PostsQuery) WherePublishedAtLte(value time.Time) *PostsQuery {
	return q.WherePublishedAt(LTE, value)
}

// WherePublishedAtIn adds an IN condition for PublishedAt
func (q *PostsQuery) WherePublishedAtIn(values ...time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.PublishedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WherePublishedAtIsNull adds an IS NULL condition for PublishedAt
func (q *PostsQuery) WherePublishedAtIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.PublishedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WherePublishedAtIsNotNull adds an IS NOT NULL condition for PublishedAt
func (q *PostsQuery) WherePublishedAtIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.PublishedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByPublishedAt adds ORDER BY for PublishedAt
func (q *PostsQuery) OrderByPublishedAt(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.PublishedAt(),
		Direction: dir,
	})
	return q
}

// WhereViewCount adds a condition for ViewCount
func (q *PostsQuery) WhereViewCount(op ComparisonOp, value int64) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.ViewCount(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereViewCountEq adds an equality condition for ViewCount
func (q *PostsQuery) WhereViewCountEq(value int64) *PostsQuery {
	return q.WhereViewCount(EQ, value)
}

// WhereViewCountNotEq adds a not-equal condition for ViewCount
func (q *PostsQuery) WhereViewCountNotEq(value int64) *PostsQuery {
	return q.WhereViewCount(NEQ, value)
}

// WhereViewCountGt adds a greater-than condition for ViewCount
func (q *PostsQuery) WhereViewCountGt(value int64) *PostsQuery {
	return q.WhereViewCount(GT, value)
}

// WhereViewCountGte adds a greater-than-or-equal condition for ViewCount
func (q *PostsQuery) WhereViewCountGte(value int64) *PostsQuery {
	return q.WhereViewCount(GTE, value)
}

// WhereViewCountLt adds a less-than condition for ViewCount
func (q *PostsQuery) WhereViewCountLt(value int64) *PostsQuery {
	return q.WhereViewCount(LT, value)
}

// WhereViewCountLte adds a less-than-or-equal condition for ViewCount
func (q *PostsQuery) WhereViewCountLte(value int64) *PostsQuery {
	return q.WhereViewCount(LTE, value)
}

// WhereViewCountIn adds an IN condition for ViewCount
func (q *PostsQuery) WhereViewCountIn(values ...int64) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.ViewCount(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereViewCountIsNull adds an IS NULL condition for ViewCount
func (q *PostsQuery) WhereViewCountIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.ViewCount(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereViewCountIsNotNull adds an IS NOT NULL condition for ViewCount
func (q *PostsQuery) WhereViewCountIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.ViewCount(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByViewCount adds ORDER BY for ViewCount
func (q *PostsQuery) OrderByViewCount(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.ViewCount(),
		Direction: dir,
	})
	return q
}

// WhereCreatedAt adds a condition for CreatedAt
func (q *PostsQuery) WhereCreatedAt(op ComparisonOp, value time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.CreatedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtEq adds an equality condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtEq(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(EQ, value)
}

// WhereCreatedAtNotEq adds a not-equal condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtNotEq(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(NEQ, value)
}

// WhereCreatedAtGt adds a greater-than condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtGt(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(GT, value)
}

// WhereCreatedAtGte adds a greater-than-or-equal condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtGte(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(GTE, value)
}

// WhereCreatedAtLt adds a less-than condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtLt(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(LT, value)
}

// WhereCreatedAtLte adds a less-than-or-equal condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtLte(value time.Time) *PostsQuery {
	return q.WhereCreatedAt(LTE, value)
}

// WhereCreatedAtIn adds an IN condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtIn(values ...time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.CreatedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNull adds an IS NULL condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.CreatedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNotNull adds an IS NOT NULL condition for CreatedAt
func (q *PostsQuery) WhereCreatedAtIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.CreatedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByCreatedAt adds ORDER BY for CreatedAt
func (q *PostsQuery) OrderByCreatedAt(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.CreatedAt(),
		Direction: dir,
	})
	return q
}

// WhereUpdatedAt adds a condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAt(op ComparisonOp, value time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UpdatedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtEq adds an equality condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtEq(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(EQ, value)
}

// WhereUpdatedAtNotEq adds a not-equal condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtNotEq(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(NEQ, value)
}

// WhereUpdatedAtGt adds a greater-than condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtGt(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(GT, value)
}

// WhereUpdatedAtGte adds a greater-than-or-equal condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtGte(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(GTE, value)
}

// WhereUpdatedAtLt adds a less-than condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtLt(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(LT, value)
}

// WhereUpdatedAtLte adds a less-than-or-equal condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtLte(value time.Time) *PostsQuery {
	return q.WhereUpdatedAt(LTE, value)
}

// WhereUpdatedAtIn adds an IN condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtIn(values ...time.Time) *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UpdatedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtIsNull adds an IS NULL condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtIsNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UpdatedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtIsNotNull adds an IS NOT NULL condition for UpdatedAt
func (q *PostsQuery) WhereUpdatedAtIsNotNull() *PostsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostsTable.UpdatedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByUpdatedAt adds ORDER BY for UpdatedAt
func (q *PostsQuery) OrderByUpdatedAt(dir OrderDirection) *PostsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostsTable.UpdatedAt(),
		Direction: dir,
	})
	return q
}

// And sets the logical operator to AND for subsequent conditions
func (q *PostsQuery) And() *PostsQuery {
	q.logical = AND
	return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *PostsQuery) Or() *PostsQuery {
	q.logical = OR
	return q
}

// WhereGroup adds a group of conditions
func (q *PostsQuery) WhereGroup(fn func(*PostsQuery)) *PostsQuery {
	subQuery := &PostsQuery{logical: AND}
	fn(subQuery)

	q.conditions = append(q.conditions, Condition{
		Logical:  q.logical,
		SubConds: subQuery.conditions,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *PostsQuery) GroupBy(fields ...*FieldRef) *PostsQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *PostsQuery) Having(field *FieldRef, op ComparisonOp, value interface{}) *PostsQuery {
	q.having = append(q.having, Condition{
		Field: field,
		Op:    op,
		Value: value,
	})
	return q
}

// Limit sets the LIMIT
func (q *PostsQuery) Limit(limit int) *PostsQuery {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET
func (q *PostsQuery) Offset(offset int) *PostsQuery {
	q.offset = &offset
	return q
}

// buildQuery builds the SQL query
func (q *PostsQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		// Select all fields from main table
		query.WriteString("posts.*")

		// Select all fields from joined tables
		for tableName := range q.activeJoins {
			query.WriteString(", ")
			query.WriteString(tableName)
			query.WriteString(".*")
		}
	} else {
		fields := make([]string, len(q.selectFields))
		for i, field := range q.selectFields {
			fields[i] = field.String()
		}
		query.WriteString(strings.Join(fields, ", "))
	}

	query.WriteString(" FROM posts")

	// JOIN clauses
	for _, join := range q.joins {
		query.WriteString(" ")
		query.WriteString(join.Type)
		query.WriteString(" ")
		query.WriteString(join.Table)
		query.WriteString(" ON ")
		query.WriteString(join.LeftField.String())
		query.WriteString(" = ")
		query.WriteString(join.RightField.String())
	}

	// WHERE clause
	if len(q.conditions) > 0 {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
		query.WriteString(whereStr)
		args = append(args, whereArgs...)
	}

	// GROUP BY clause
	if len(q.groupBy) > 0 {
		query.WriteString(" GROUP BY ")
		groupFields := make([]string, len(q.groupBy))
		for i, field := range q.groupBy {
			groupFields[i] = field.String()
		}
		query.WriteString(strings.Join(groupFields, ", "))
	}

	// HAVING clause
	if len(q.having) > 0 {
		query.WriteString(" HAVING ")
		havingStr, havingArgs := q.buildConditions(q.having, &argIndex)
		query.WriteString(havingStr)
		args = append(args, havingArgs...)
	}

	// ORDER BY clause
	if len(q.orderBy) > 0 {
		query.WriteString(" ORDER BY ")
		orderClauses := make([]string, len(q.orderBy))
		for i, order := range q.orderBy {
			orderClauses[i] = fmt.Sprintf("%s %s", order.Field.String(), order.Direction)
		}
		query.WriteString(strings.Join(orderClauses, ", "))
	}

	// LIMIT clause
	if q.limit != nil {
		query.WriteString(fmt.Sprintf(" LIMIT %d", *q.limit))
	}

	// OFFSET clause
	if q.offset != nil {
		query.WriteString(fmt.Sprintf(" OFFSET %d", *q.offset))
	}

	return query.String(), args
}

// buildConditions builds WHERE/HAVING conditions
func (q *PostsQuery) buildConditions(conditions []Condition, argIndex *int) (string, []interface{}) {
	var parts []string
	var args []interface{}

	for i, cond := range conditions {
		if len(cond.SubConds) > 0 {
			subStr, subArgs := q.buildConditions(cond.SubConds, argIndex)
			if i > 0 {
				parts = append(parts, fmt.Sprintf("%s (%s)", cond.Logical, subStr))
			} else {
				parts = append(parts, fmt.Sprintf("(%s)", subStr))
			}
			args = append(args, subArgs...)
		} else {
			var condStr string
			if cond.Value == nil {
				if cond.Op == EQ {
					condStr = fmt.Sprintf("%s IS NULL", cond.Field.String())
				} else {
					condStr = fmt.Sprintf("%s IS NOT NULL", cond.Field.String())
				}
			} else if cond.Op == IN {
				// Handle IN clause with reflection to support any slice type
				v := cond.Value
				switch vals := v.(type) {
				case []int64:
					placeholders := make([]string, len(vals))
					for j, val := range vals {
						placeholders[j] = fmt.Sprintf("$%d", *argIndex)
						*argIndex++
						args = append(args, val)
					}
					condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
				case []string:
					placeholders := make([]string, len(vals))
					for j, val := range vals {
						placeholders[j] = fmt.Sprintf("$%d", *argIndex)
						*argIndex++
						args = append(args, val)
					}
					condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
				case []float64:
					placeholders := make([]string, len(vals))
					for j, val := range vals {
						placeholders[j] = fmt.Sprintf("$%d", *argIndex)
						*argIndex++
						args = append(args, val)
					}
					condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
				case []bool:
					placeholders := make([]string, len(vals))
					for j, val := range vals {
						placeholders[j] = fmt.Sprintf("$%d", *argIndex)
						*argIndex++
						args = append(args, val)
					}
					condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
				default:
					// Fallback for other types - this shouldn't normally happen
					condStr = fmt.Sprintf("%s IN (?)", cond.Field.String())
				}
			} else {
				condStr = fmt.Sprintf("%s %s $%d", cond.Field.String(), cond.Op, *argIndex)
				*argIndex++
				args = append(args, cond.Value)
			}

			if i > 0 {
				parts = append(parts, fmt.Sprintf("%s %s", cond.Logical, condStr))
			} else {
				parts = append(parts, condStr)
			}
		}
	}

	return strings.Join(parts, " "), args
}

// Find executes the query and returns results
func (q *PostsQuery) Find(ctx context.Context) ([]*PostsDto, error) {
	query, args := q.buildQuery()

	var rows pgx.Rows
	var err error

	if q.tx != nil {
		rows, err = q.tx.Query(ctx, query, args...)
	} else {
		rows, err = q.pool.Query(ctx, query, args...)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*PostsDto
	for rows.Next() {
		result := &PostsDto{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *PostsQuery) FindOne(ctx context.Context) (*PostsDto, error) {
	q.Limit(1)
	results, err := q.Find(ctx)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, pgx.ErrNoRows
	}
	return results[0], nil
}

// Count returns the count
func (q *PostsQuery) Count(ctx context.Context) (int64, error) {
	// Save and restore select fields
	originalSelect := q.selectFields
	countField := &FieldRef{Table: "", Column: "COUNT(*)", GoType: "int64"}
	q.selectFields = []*FieldRef{countField}
	defer func() { q.selectFields = originalSelect }()

	query, args := q.buildQuery()

	var count int64
	if q.tx != nil {
		err := q.tx.QueryRow(ctx, query, args...).Scan(&count)
		return count, err
	}
	err := q.pool.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

// scanInto scans a row into a struct
func (q *PostsQuery) scanInto(rows pgx.Rows, dest *PostsDto) error {
	// Build scan destinations dynamically based on active joins
	var scanDest []interface{}

	// Add main table columns
	scanDest = append(scanDest, &dest.Id)
	scanDest = append(scanDest, &dest.UserId)
	scanDest = append(scanDest, &dest.Title)
	scanDest = append(scanDest, &dest.Content)
	scanDest = append(scanDest, &dest.Status)
	scanDest = append(scanDest, &dest.PublishedAt)
	scanDest = append(scanDest, &dest.ViewCount)
	scanDest = append(scanDest, &dest.CreatedAt)
	scanDest = append(scanDest, &dest.UpdatedAt)

	// Add joined table columns if active
	if q.activeJoins["users"] {
		joinedUser := &UsersDto{}
		scanDest = append(scanDest, &joinedUser.Id)
		scanDest = append(scanDest, &joinedUser.Username)
		scanDest = append(scanDest, &joinedUser.Email)
		scanDest = append(scanDest, &joinedUser.FullName)
		scanDest = append(scanDest, &joinedUser.CreatedAt)
		scanDest = append(scanDest, &joinedUser.UpdatedAt)
		scanDest = append(scanDest, &joinedUser.IsActive)
		dest.User = joinedUser
	}

	return rows.Scan(scanDest...)
}

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (q *PostsQuery) Insert(ctx context.Context, record *PostsDto) error {
	var columns []string
	var placeholders []string
	var args []interface{}
	argIdx := 1

	// Dynamically build column list based on non-nil values

	// Required field (not nullable, no default)
	columns = append(columns, "user_id")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.UserId)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "title")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Title)
	argIdx++

	if record.Content != nil {
		columns = append(columns, "content")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Content)
		argIdx++
	}

	if record.Status != nil {
		columns = append(columns, "status")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Status)
		argIdx++
	}

	if record.PublishedAt != nil {
		columns = append(columns, "published_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.PublishedAt)
		argIdx++
	}

	if record.ViewCount != nil {
		columns = append(columns, "view_count")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.ViewCount)
		argIdx++
	}

	if record.CreatedAt != nil {
		columns = append(columns, "created_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.CreatedAt)
		argIdx++
	}

	if record.UpdatedAt != nil {
		columns = append(columns, "updated_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.UpdatedAt)
		argIdx++
	}

	if len(columns) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	query := fmt.Sprintf("INSERT INTO posts (%s) VALUES (%s) RETURNING id",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	var err error
	if q.tx != nil {
		err = q.tx.QueryRow(ctx, query, args...).Scan(&record.Id)
	} else {
		err = q.pool.QueryRow(ctx, query, args...).Scan(&record.Id)
	}
	return err
}

// InsertBatch inserts multiple records efficiently using pgx batch
func (q *PostsQuery) InsertBatch(ctx context.Context, records []*PostsDto) error {
	if len(records) == 0 {
		return nil
	}

	batch := &pgx.Batch{}

	for _, record := range records {
		var columns []string
		var placeholders []string
		var args []interface{}
		argIdx := 1

		// Dynamically build column list based on non-nil values

		// Required field (not nullable, no default)
		columns = append(columns, "user_id")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.UserId)
		argIdx++

		// Required field (not nullable, no default)
		columns = append(columns, "title")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Title)
		argIdx++

		if record.Content != nil {
			columns = append(columns, "content")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.Content)
			argIdx++
		}

		if record.Status != nil {
			columns = append(columns, "status")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.Status)
			argIdx++
		}

		if record.PublishedAt != nil {
			columns = append(columns, "published_at")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.PublishedAt)
			argIdx++
		}

		if record.ViewCount != nil {
			columns = append(columns, "view_count")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.ViewCount)
			argIdx++
		}

		if record.CreatedAt != nil {
			columns = append(columns, "created_at")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.CreatedAt)
			argIdx++
		}

		if record.UpdatedAt != nil {
			columns = append(columns, "updated_at")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.UpdatedAt)
			argIdx++
		}

		if len(columns) > 0 {
			query := fmt.Sprintf("INSERT INTO posts (%s) VALUES (%s) RETURNING id",
				strings.Join(columns, ", "),
				strings.Join(placeholders, ", "))
			batch.Queue(query, args...)
		}
	}

	var br pgx.BatchResults

	if q.tx != nil {
		br = q.tx.SendBatch(ctx, batch)
	} else {
		br = q.pool.SendBatch(ctx, batch)
	}
	defer br.Close()

	// Scan returned IDs back into records
	for i := range records {
		if err := br.QueryRow().Scan(&records[i].Id); err != nil {
			return fmt.Errorf("failed to scan batch result %d: %w", i, err)
		}
	}

	return br.Close()
}

// Update updates a record using its primary key
func (q *PostsQuery) Update(ctx context.Context, record *PostsDto) error {
	query := "UPDATE posts SET " +
		"user_id = $2, title = $3, content = $4, status = $5, published_at = $6, view_count = $7, created_at = $8, updated_at = $9" +
		" WHERE id = $1"

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query,
			record.Id,
			record.UserId,
			record.Title,
			record.Content,
			record.Status,
			record.PublishedAt,
			record.ViewCount,
			record.CreatedAt,
			record.UpdatedAt,
		)
	} else {
		tag, err = q.pool.Exec(ctx, query,
			record.Id,
			record.UserId,
			record.Title,
			record.Content,
			record.Status,
			record.PublishedAt,
			record.ViewCount,
			record.CreatedAt,
			record.UpdatedAt,
		)
	}

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// UpdateFields updates specific fields for matching records
func (q *PostsQuery) UpdateFields(ctx context.Context, updates map[*FieldRef]interface{}) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE posts SET ")

	setClauses := make([]string, 0, len(updates))
	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field.Column, argIndex))
		args = append(args, value)
		argIndex++
	}
	query.WriteString(strings.Join(setClauses, ", "))

	// Add WHERE conditions
	if len(q.conditions) > 0 {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
		query.WriteString(whereStr)
		args = append(args, whereArgs...)
	}

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query.String(), args...)
	} else {
		tag, err = q.pool.Exec(ctx, query.String(), args...)
	}

	if err != nil {
		return 0, err
	}

	return tag.RowsAffected(), nil
}

// Delete deletes matching records
func (q *PostsQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("DELETE FROM posts")

	// Add WHERE conditions
	if len(q.conditions) > 0 {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
		query.WriteString(whereStr)
		args = append(args, whereArgs...)
	}

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query.String(), args...)
	} else {
		tag, err = q.pool.Exec(ctx, query.String(), args...)
	}

	if err != nil {
		return 0, err
	}

	return tag.RowsAffected(), nil
}

// JoinUsers performs a type-safe inner join with users
func (q *PostsQuery) JoinUsers() *PostsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "INNER JOIN",
		Table:      "users",
		LeftField:  PostsTable.UserId(),
		RightField: UsersTable.Id(),
	})
	q.activeJoins["users"] = true
	return q
}

// LeftJoinUsers performs a type-safe left join with users
func (q *PostsQuery) LeftJoinUsers() *PostsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "LEFT JOIN",
		Table:      "users",
		LeftField:  PostsTable.UserId(),
		RightField: UsersTable.Id(),
	})
	q.activeJoins["users"] = true
	return q
}

// JoinOn performs a custom join with type-safe field references
func (q *PostsQuery) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *PostsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
