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

// Comments represents the comments table
type Comments struct {
	Id         *int64     `db:"id" json:"id"`
	PostId     int64      `db:"post_id" json:"post_id"`
	UserId     string     `db:"user_id" json:"user_id"`
	Content    string     `db:"content" json:"content"`
	IsApproved *bool      `db:"is_approved" json:"is_approved"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
}

// TableName returns the table name for Comments
func (c *Comments) TableName() string {
	return "comments"
}

// CommentsFields provides type-safe field references for Comments
type CommentsFields struct{}

// Fields returns field references for Comments
var CommentsTable = CommentsFields{}

// Id returns a field reference for Comments.Id
func (c CommentsFields) Id() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "id",
		GoType: "int64",
	}
}

// PostId returns a field reference for Comments.PostId
func (c CommentsFields) PostId() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "post_id",
		GoType: "int64",
	}
}

// UserId returns a field reference for Comments.UserId
func (c CommentsFields) UserId() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "user_id",
		GoType: "string",
	}
}

// Content returns a field reference for Comments.Content
func (c CommentsFields) Content() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "content",
		GoType: "string",
	}
}

// IsApproved returns a field reference for Comments.IsApproved
func (c CommentsFields) IsApproved() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "is_approved",
		GoType: "bool",
	}
}

// CreatedAt returns a field reference for Comments.CreatedAt
func (c CommentsFields) CreatedAt() *FieldRef {
	return &FieldRef{
		Table:  "comments",
		Column: "created_at",
		GoType: "time.Time",
	}
}

// CommentsQuery is a type-safe query builder for Comments
type CommentsQuery struct {
	pool         *pgxpool.Pool
	tx           pgx.Tx
	selectFields []*FieldRef
	conditions   []Condition
	joins        []JoinClause
	orderBy      []OrderClause
	groupBy      []*FieldRef
	having       []Condition
	limit        *int
	offset       *int
	logical      LogicalOp
}

// NewCommentsQuery creates a new query builder
func NewCommentsQuery(pool *pgxpool.Pool) *CommentsQuery {
	return &CommentsQuery{
		pool:    pool,
		logical: AND,
	}
}

// WithTx sets the transaction
func (q *CommentsQuery) WithTx(tx pgx.Tx) *CommentsQuery {
	q.tx = tx
	return q
}

// Select specifies fields to select using type-safe field references
func (q *CommentsQuery) Select(fields ...*FieldRef) *CommentsQuery {
	q.selectFields = fields
	return q
}

// SelectAll selects all fields
func (q *CommentsQuery) SelectAll() *CommentsQuery {
	q.selectFields = []*FieldRef{
		CommentsTable.Id(),
		CommentsTable.PostId(),
		CommentsTable.UserId(),
		CommentsTable.Content(),
		CommentsTable.IsApproved(),
		CommentsTable.CreatedAt(),
	}
	return q
}

// WhereId adds a condition for Id
func (q *CommentsQuery) WhereId(op ComparisonOp, value int64) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.Id(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIdEq adds an equality condition for Id
func (q *CommentsQuery) WhereIdEq(value int64) *CommentsQuery {
	return q.WhereId(EQ, value)
}

// WhereIdNotEq adds a not-equal condition for Id
func (q *CommentsQuery) WhereIdNotEq(value int64) *CommentsQuery {
	return q.WhereId(NEQ, value)
}

// WhereIdGt adds a greater-than condition for Id
func (q *CommentsQuery) WhereIdGt(value int64) *CommentsQuery {
	return q.WhereId(GT, value)
}

// WhereIdGte adds a greater-than-or-equal condition for Id
func (q *CommentsQuery) WhereIdGte(value int64) *CommentsQuery {
	return q.WhereId(GTE, value)
}

// WhereIdLt adds a less-than condition for Id
func (q *CommentsQuery) WhereIdLt(value int64) *CommentsQuery {
	return q.WhereId(LT, value)
}

// WhereIdLte adds a less-than-or-equal condition for Id
func (q *CommentsQuery) WhereIdLte(value int64) *CommentsQuery {
	return q.WhereId(LTE, value)
}

// WhereIdIn adds an IN condition for Id
func (q *CommentsQuery) WhereIdIn(values ...int64) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.Id(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderById adds ORDER BY for Id
func (q *CommentsQuery) OrderById(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.Id(),
		Direction: dir,
	})
	return q
}

// WherePostId adds a condition for PostId
func (q *CommentsQuery) WherePostId(op ComparisonOp, value int64) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.PostId(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WherePostIdEq adds an equality condition for PostId
func (q *CommentsQuery) WherePostIdEq(value int64) *CommentsQuery {
	return q.WherePostId(EQ, value)
}

// WherePostIdNotEq adds a not-equal condition for PostId
func (q *CommentsQuery) WherePostIdNotEq(value int64) *CommentsQuery {
	return q.WherePostId(NEQ, value)
}

// WherePostIdGt adds a greater-than condition for PostId
func (q *CommentsQuery) WherePostIdGt(value int64) *CommentsQuery {
	return q.WherePostId(GT, value)
}

// WherePostIdGte adds a greater-than-or-equal condition for PostId
func (q *CommentsQuery) WherePostIdGte(value int64) *CommentsQuery {
	return q.WherePostId(GTE, value)
}

// WherePostIdLt adds a less-than condition for PostId
func (q *CommentsQuery) WherePostIdLt(value int64) *CommentsQuery {
	return q.WherePostId(LT, value)
}

// WherePostIdLte adds a less-than-or-equal condition for PostId
func (q *CommentsQuery) WherePostIdLte(value int64) *CommentsQuery {
	return q.WherePostId(LTE, value)
}

// WherePostIdIn adds an IN condition for PostId
func (q *CommentsQuery) WherePostIdIn(values ...int64) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.PostId(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByPostId adds ORDER BY for PostId
func (q *CommentsQuery) OrderByPostId(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.PostId(),
		Direction: dir,
	})
	return q
}

// WhereUserId adds a condition for UserId
func (q *CommentsQuery) WhereUserId(op ComparisonOp, value string) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.UserId(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereUserIdEq adds an equality condition for UserId
func (q *CommentsQuery) WhereUserIdEq(value string) *CommentsQuery {
	return q.WhereUserId(EQ, value)
}

// WhereUserIdNotEq adds a not-equal condition for UserId
func (q *CommentsQuery) WhereUserIdNotEq(value string) *CommentsQuery {
	return q.WhereUserId(NEQ, value)
}

// WhereUserIdLike adds a LIKE condition for UserId
func (q *CommentsQuery) WhereUserIdLike(pattern string) *CommentsQuery {
	return q.WhereUserId(LIKE, pattern)
}

// WhereUserIdIn adds an IN condition for UserId
func (q *CommentsQuery) WhereUserIdIn(values ...string) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.UserId(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByUserId adds ORDER BY for UserId
func (q *CommentsQuery) OrderByUserId(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.UserId(),
		Direction: dir,
	})
	return q
}

// WhereContent adds a condition for Content
func (q *CommentsQuery) WhereContent(op ComparisonOp, value string) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.Content(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereContentEq adds an equality condition for Content
func (q *CommentsQuery) WhereContentEq(value string) *CommentsQuery {
	return q.WhereContent(EQ, value)
}

// WhereContentNotEq adds a not-equal condition for Content
func (q *CommentsQuery) WhereContentNotEq(value string) *CommentsQuery {
	return q.WhereContent(NEQ, value)
}

// WhereContentLike adds a LIKE condition for Content
func (q *CommentsQuery) WhereContentLike(pattern string) *CommentsQuery {
	return q.WhereContent(LIKE, pattern)
}

// WhereContentIn adds an IN condition for Content
func (q *CommentsQuery) WhereContentIn(values ...string) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.Content(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByContent adds ORDER BY for Content
func (q *CommentsQuery) OrderByContent(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.Content(),
		Direction: dir,
	})
	return q
}

// WhereIsApproved adds a condition for IsApproved
func (q *CommentsQuery) WhereIsApproved(op ComparisonOp, value bool) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.IsApproved(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIsApprovedEq adds an equality condition for IsApproved
func (q *CommentsQuery) WhereIsApprovedEq(value bool) *CommentsQuery {
	return q.WhereIsApproved(EQ, value)
}

// WhereIsApprovedNotEq adds a not-equal condition for IsApproved
func (q *CommentsQuery) WhereIsApprovedNotEq(value bool) *CommentsQuery {
	return q.WhereIsApproved(NEQ, value)
}

// WhereIsApprovedIn adds an IN condition for IsApproved
func (q *CommentsQuery) WhereIsApprovedIn(values ...bool) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.IsApproved(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereIsApprovedIsNull adds an IS NULL condition for IsApproved
func (q *CommentsQuery) WhereIsApprovedIsNull() *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.IsApproved(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereIsApprovedIsNotNull adds an IS NOT NULL condition for IsApproved
func (q *CommentsQuery) WhereIsApprovedIsNotNull() *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.IsApproved(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByIsApproved adds ORDER BY for IsApproved
func (q *CommentsQuery) OrderByIsApproved(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.IsApproved(),
		Direction: dir,
	})
	return q
}

// WhereCreatedAt adds a condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAt(op ComparisonOp, value time.Time) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.CreatedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtEq adds an equality condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtEq(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(EQ, value)
}

// WhereCreatedAtNotEq adds a not-equal condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtNotEq(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(NEQ, value)
}

// WhereCreatedAtGt adds a greater-than condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtGt(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(GT, value)
}

// WhereCreatedAtGte adds a greater-than-or-equal condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtGte(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(GTE, value)
}

// WhereCreatedAtLt adds a less-than condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtLt(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(LT, value)
}

// WhereCreatedAtLte adds a less-than-or-equal condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtLte(value time.Time) *CommentsQuery {
	return q.WhereCreatedAt(LTE, value)
}

// WhereCreatedAtIn adds an IN condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtIn(values ...time.Time) *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.CreatedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNull adds an IS NULL condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtIsNull() *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.CreatedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNotNull adds an IS NOT NULL condition for CreatedAt
func (q *CommentsQuery) WhereCreatedAtIsNotNull() *CommentsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   CommentsTable.CreatedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// OrderByCreatedAt adds ORDER BY for CreatedAt
func (q *CommentsQuery) OrderByCreatedAt(dir OrderDirection) *CommentsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     CommentsTable.CreatedAt(),
		Direction: dir,
	})
	return q
}

// And sets the logical operator to AND for subsequent conditions
func (q *CommentsQuery) And() *CommentsQuery {
	q.logical = AND
	return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *CommentsQuery) Or() *CommentsQuery {
	q.logical = OR
	return q
}

// WhereGroup adds a group of conditions
func (q *CommentsQuery) WhereGroup(fn func(*CommentsQuery)) *CommentsQuery {
	subQuery := &CommentsQuery{logical: AND}
	fn(subQuery)

	q.conditions = append(q.conditions, Condition{
		Logical:  q.logical,
		SubConds: subQuery.conditions,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *CommentsQuery) GroupBy(fields ...*FieldRef) *CommentsQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *CommentsQuery) Having(field *FieldRef, op ComparisonOp, value interface{}) *CommentsQuery {
	q.having = append(q.having, Condition{
		Field: field,
		Op:    op,
		Value: value,
	})
	return q
}

// Limit sets the LIMIT
func (q *CommentsQuery) Limit(limit int) *CommentsQuery {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET
func (q *CommentsQuery) Offset(offset int) *CommentsQuery {
	q.offset = &offset
	return q
}

// buildQuery builds the SQL query
func (q *CommentsQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		query.WriteString("comments.*")
	} else {
		fields := make([]string, len(q.selectFields))
		for i, field := range q.selectFields {
			fields[i] = field.String()
		}
		query.WriteString(strings.Join(fields, ", "))
	}

	query.WriteString(" FROM comments")

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
func (q *CommentsQuery) buildConditions(conditions []Condition, argIndex *int) (string, []interface{}) {
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
func (q *CommentsQuery) Find(ctx context.Context) ([]*Comments, error) {
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

	var results []*Comments
	for rows.Next() {
		result := &Comments{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *CommentsQuery) FindOne(ctx context.Context) (*Comments, error) {
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
func (q *CommentsQuery) Count(ctx context.Context) (int64, error) {
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
func (q *CommentsQuery) scanInto(rows pgx.Rows, dest *Comments) error {
	// If custom fields selected, use dynamic scanning
	if len(q.selectFields) > 0 && q.selectFields[0].Table != "" {
		// This would need more complex implementation for custom field scanning
		// For now, scan all fields
	}

	return rows.Scan(
		&dest.Id,
		&dest.PostId,
		&dest.UserId,
		&dest.Content,
		&dest.IsApproved,
		&dest.CreatedAt,
	)
}

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (q *CommentsQuery) Insert(ctx context.Context, record *Comments) error {
	var columns []string
	var placeholders []string
	var args []interface{}
	argIdx := 1

	// Dynamically build column list based on non-nil values

	// Required field (not nullable, no default)
	columns = append(columns, "post_id")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.PostId)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "user_id")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.UserId)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "content")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Content)
	argIdx++

	if record.IsApproved != nil {
		columns = append(columns, "is_approved")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.IsApproved)
		argIdx++
	}

	if record.CreatedAt != nil {
		columns = append(columns, "created_at")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.CreatedAt)
		argIdx++
	}

	if len(columns) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	query := fmt.Sprintf("INSERT INTO comments (%s) VALUES (%s) RETURNING id",
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
func (q *CommentsQuery) InsertBatch(ctx context.Context, records []*Comments) error {
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
		columns = append(columns, "post_id")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.PostId)
		argIdx++

		// Required field (not nullable, no default)
		columns = append(columns, "user_id")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.UserId)
		argIdx++

		// Required field (not nullable, no default)
		columns = append(columns, "content")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Content)
		argIdx++

		if record.IsApproved != nil {
			columns = append(columns, "is_approved")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.IsApproved)
			argIdx++
		}

		if record.CreatedAt != nil {
			columns = append(columns, "created_at")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.CreatedAt)
			argIdx++
		}

		if len(columns) > 0 {
			query := fmt.Sprintf("INSERT INTO comments (%s) VALUES (%s) RETURNING id",
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
func (q *CommentsQuery) Update(ctx context.Context, record *Comments) error {
	query := "UPDATE comments SET " +
		"post_id = $2, user_id = $3, content = $4, is_approved = $5, created_at = $6" +
		" WHERE id = $1"

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query,
			record.Id,
			record.PostId,
			record.UserId,
			record.Content,
			record.IsApproved,
			record.CreatedAt,
		)
	} else {
		tag, err = q.pool.Exec(ctx, query,
			record.Id,
			record.PostId,
			record.UserId,
			record.Content,
			record.IsApproved,
			record.CreatedAt,
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
func (q *CommentsQuery) UpdateFields(ctx context.Context, updates map[*FieldRef]interface{}) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE comments SET ")

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
func (q *CommentsQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("DELETE FROM comments")

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
} // JoinPosts performs a type-safe inner join with posts
func (q *CommentsQuery) JoinPosts() *CommentsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "INNER JOIN",
		Table:      "posts",
		LeftField:  CommentsTable.PostId(),
		RightField: PostsTable.Id(),
	})
	return q
}

// LeftJoinPosts performs a type-safe left join with posts
func (q *CommentsQuery) LeftJoinPosts() *CommentsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "LEFT JOIN",
		Table:      "posts",
		LeftField:  CommentsTable.PostId(),
		RightField: PostsTable.Id(),
	})
	return q
}

// CommentsPostsJoin represents a join result between comments and posts
type CommentsPostsJoin struct {
	Comments *Comments
	Posts    *Posts
}

// JoinUsers performs a type-safe inner join with users
func (q *CommentsQuery) JoinUsers() *CommentsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "INNER JOIN",
		Table:      "users",
		LeftField:  CommentsTable.UserId(),
		RightField: UsersTable.Id(),
	})
	return q
}

// LeftJoinUsers performs a type-safe left join with users
func (q *CommentsQuery) LeftJoinUsers() *CommentsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "LEFT JOIN",
		Table:      "users",
		LeftField:  CommentsTable.UserId(),
		RightField: UsersTable.Id(),
	})
	return q
}

// CommentsUsersJoin represents a join result between comments and users
type CommentsUsersJoin struct {
	Comments *Comments
	Users    *Users
}

// JoinOn performs a custom join with type-safe field references
func (q *CommentsQuery) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *CommentsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
