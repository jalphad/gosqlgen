package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostTagsDto represents the post_tags table
type PostTagsDto struct {
	PostId int64 `db:"post_id" json:"post_id"`
	TagId  int64 `db:"tag_id" json:"tag_id"`
}

// TableName returns the table name for PostTagsDto
func (p *PostTagsDto) TableName() string {
	return "post_tags"
}

// PostTagsFields provides type-safe field references for PostTagsDto
type PostTagsFields struct{}

// PostTagsTable returns field references for PostTagsDto
var PostTagsTable = PostTagsFields{}

// PostId returns a field reference for PostTagsDto.PostId
func (p PostTagsFields) PostId() *FieldRef {
	return &FieldRef{
		Table:  "post_tags",
		Column: "post_id",
		GoType: "int64",
	}
}

// TagId returns a field reference for PostTagsDto.TagId
func (p PostTagsFields) TagId() *FieldRef {
	return &FieldRef{
		Table:  "post_tags",
		Column: "tag_id",
		GoType: "int64",
	}
}

// PostTagsQuery is a type-safe query builder for PostTagsDto
type PostTagsQuery struct {
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

// NewPostTagsQuery creates a new query builder
func NewPostTagsQuery(pool *pgxpool.Pool) *PostTagsQuery {
	return &PostTagsQuery{
		pool:    pool,
		logical: AND,
	}
}

// WithTx sets the transaction
func (q *PostTagsQuery) WithTx(tx pgx.Tx) *PostTagsQuery {
	q.tx = tx
	return q
}

// Select specifies fields to select using type-safe field references
func (q *PostTagsQuery) Select(fields ...*FieldRef) *PostTagsQuery {
	q.selectFields = fields
	return q
}

// SelectAll selects all fields
func (q *PostTagsQuery) SelectAll() *PostTagsQuery {
	q.selectFields = []*FieldRef{
		PostTagsTable.PostId(),
		PostTagsTable.TagId(),
	}
	return q
}

// WherePostId adds a condition for PostId
func (q *PostTagsQuery) WherePostId(op ComparisonOp, value int64) *PostTagsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostTagsTable.PostId(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WherePostIdEq adds an equality condition for PostId
func (q *PostTagsQuery) WherePostIdEq(value int64) *PostTagsQuery {
	return q.WherePostId(EQ, value)
}

// WherePostIdNotEq adds a not-equal condition for PostId
func (q *PostTagsQuery) WherePostIdNotEq(value int64) *PostTagsQuery {
	return q.WherePostId(NEQ, value)
}

// WherePostIdGt adds a greater-than condition for PostId
func (q *PostTagsQuery) WherePostIdGt(value int64) *PostTagsQuery {
	return q.WherePostId(GT, value)
}

// WherePostIdGte adds a greater-than-or-equal condition for PostId
func (q *PostTagsQuery) WherePostIdGte(value int64) *PostTagsQuery {
	return q.WherePostId(GTE, value)
}

// WherePostIdLt adds a less-than condition for PostId
func (q *PostTagsQuery) WherePostIdLt(value int64) *PostTagsQuery {
	return q.WherePostId(LT, value)
}

// WherePostIdLte adds a less-than-or-equal condition for PostId
func (q *PostTagsQuery) WherePostIdLte(value int64) *PostTagsQuery {
	return q.WherePostId(LTE, value)
}

// WherePostIdIn adds an IN condition for PostId
func (q *PostTagsQuery) WherePostIdIn(values ...int64) *PostTagsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostTagsTable.PostId(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByPostId adds ORDER BY for PostId
func (q *PostTagsQuery) OrderByPostId(dir OrderDirection) *PostTagsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostTagsTable.PostId(),
		Direction: dir,
	})
	return q
}

// WhereTagId adds a condition for TagId
func (q *PostTagsQuery) WhereTagId(op ComparisonOp, value int64) *PostTagsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostTagsTable.TagId(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereTagIdEq adds an equality condition for TagId
func (q *PostTagsQuery) WhereTagIdEq(value int64) *PostTagsQuery {
	return q.WhereTagId(EQ, value)
}

// WhereTagIdNotEq adds a not-equal condition for TagId
func (q *PostTagsQuery) WhereTagIdNotEq(value int64) *PostTagsQuery {
	return q.WhereTagId(NEQ, value)
}

// WhereTagIdGt adds a greater-than condition for TagId
func (q *PostTagsQuery) WhereTagIdGt(value int64) *PostTagsQuery {
	return q.WhereTagId(GT, value)
}

// WhereTagIdGte adds a greater-than-or-equal condition for TagId
func (q *PostTagsQuery) WhereTagIdGte(value int64) *PostTagsQuery {
	return q.WhereTagId(GTE, value)
}

// WhereTagIdLt adds a less-than condition for TagId
func (q *PostTagsQuery) WhereTagIdLt(value int64) *PostTagsQuery {
	return q.WhereTagId(LT, value)
}

// WhereTagIdLte adds a less-than-or-equal condition for TagId
func (q *PostTagsQuery) WhereTagIdLte(value int64) *PostTagsQuery {
	return q.WhereTagId(LTE, value)
}

// WhereTagIdIn adds an IN condition for TagId
func (q *PostTagsQuery) WhereTagIdIn(values ...int64) *PostTagsQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   PostTagsTable.TagId(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByTagId adds ORDER BY for TagId
func (q *PostTagsQuery) OrderByTagId(dir OrderDirection) *PostTagsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     PostTagsTable.TagId(),
		Direction: dir,
	})
	return q
}

// And sets the logical operator to AND for subsequent conditions
func (q *PostTagsQuery) And() *PostTagsQuery {
	q.logical = AND
	return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *PostTagsQuery) Or() *PostTagsQuery {
	q.logical = OR
	return q
}

// WhereGroup adds a group of conditions
func (q *PostTagsQuery) WhereGroup(fn func(*PostTagsQuery)) *PostTagsQuery {
	subQuery := &PostTagsQuery{logical: AND}
	fn(subQuery)

	q.conditions = append(q.conditions, Condition{
		Logical:  q.logical,
		SubConds: subQuery.conditions,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *PostTagsQuery) GroupBy(fields ...*FieldRef) *PostTagsQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *PostTagsQuery) Having(field *FieldRef, op ComparisonOp, value interface{}) *PostTagsQuery {
	q.having = append(q.having, Condition{
		Field: field,
		Op:    op,
		Value: value,
	})
	return q
}

// Limit sets the LIMIT
func (q *PostTagsQuery) Limit(limit int) *PostTagsQuery {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET
func (q *PostTagsQuery) Offset(offset int) *PostTagsQuery {
	q.offset = &offset
	return q
}

// buildQuery builds the SQL query
func (q *PostTagsQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		query.WriteString("post_tags.*")
	} else {
		fields := make([]string, len(q.selectFields))
		for i, field := range q.selectFields {
			fields[i] = field.String()
		}
		query.WriteString(strings.Join(fields, ", "))
	}

	query.WriteString(" FROM post_tags")

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
func (q *PostTagsQuery) buildConditions(conditions []Condition, argIndex *int) (string, []interface{}) {
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
func (q *PostTagsQuery) Find(ctx context.Context) ([]*PostTagsDto, error) {
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

	var results []*PostTagsDto
	for rows.Next() {
		result := &PostTagsDto{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *PostTagsQuery) FindOne(ctx context.Context) (*PostTagsDto, error) {
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
func (q *PostTagsQuery) Count(ctx context.Context) (int64, error) {
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
func (q *PostTagsQuery) scanInto(rows pgx.Rows, dest *PostTagsDto) error {
	// If custom fields selected, use dynamic scanning
	if len(q.selectFields) > 0 && q.selectFields[0].Table != "" {
		// This would need more complex implementation for custom field scanning
		// For now, scan all fields
	}

	return rows.Scan(
		&dest.PostId,
		&dest.TagId,
	)
}

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (q *PostTagsQuery) Insert(ctx context.Context, record *PostTagsDto) error {
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
	columns = append(columns, "tag_id")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.TagId)
	argIdx++

	if len(columns) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	query := fmt.Sprintf("INSERT INTO post_tags (%s) VALUES (%s)",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	var err error
	if q.tx != nil {
		_, err = q.tx.Exec(ctx, query, args...)
	} else {
		_, err = q.pool.Exec(ctx, query, args...)
	}
	return err
}

// InsertBatch inserts multiple records efficiently using pgx batch
func (q *PostTagsQuery) InsertBatch(ctx context.Context, records []*PostTagsDto) error {
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
		columns = append(columns, "tag_id")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.TagId)
		argIdx++

		if len(columns) > 0 {
			query := fmt.Sprintf("INSERT INTO post_tags (%s) VALUES (%s)",
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

	// Just execute all batched queries
	for range records {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}

	return br.Close()
}

// Update updates a record using its primary key
func (q *PostTagsQuery) Update(ctx context.Context, record *PostTagsDto) error {
	return fmt.Errorf("table post_tags has no primary key")
}

// UpdateFields updates specific fields for matching records
func (q *PostTagsQuery) UpdateFields(ctx context.Context, updates map[*FieldRef]interface{}) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE post_tags SET ")

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
func (q *PostTagsQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("DELETE FROM post_tags")

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

// JoinPosts performs a type-safe inner join with posts
func (q *PostTagsQuery) JoinPosts() *PostTagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "INNER JOIN",
		Table:      "posts",
		LeftField:  PostTagsTable.PostId(),
		RightField: PostsTable.Id(),
	})
	return q
}

// LeftJoinPosts performs a type-safe left join with posts
func (q *PostTagsQuery) LeftJoinPosts() *PostTagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "LEFT JOIN",
		Table:      "posts",
		LeftField:  PostTagsTable.PostId(),
		RightField: PostsTable.Id(),
	})
	return q
}

// PostTagsPostsJoin represents a join result between post_tags and posts
type PostTagsPostsJoin struct {
	PostTagsDto *PostTagsDto
	PostsDto    *PostsDto
}

// JoinTags performs a type-safe inner join with tags
func (q *PostTagsQuery) JoinTags() *PostTagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "INNER JOIN",
		Table:      "tags",
		LeftField:  PostTagsTable.TagId(),
		RightField: TagsTable.Id(),
	})
	return q
}

// LeftJoinTags performs a type-safe left join with tags
func (q *PostTagsQuery) LeftJoinTags() *PostTagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       "LEFT JOIN",
		Table:      "tags",
		LeftField:  PostTagsTable.TagId(),
		RightField: TagsTable.Id(),
	})
	return q
}

// PostTagsTagsJoin represents a join result between post_tags and tags
type PostTagsTagsJoin struct {
	PostTagsDto *PostTagsDto
	TagsDto     *TagsDto
}

// JoinOn performs a custom join with type-safe field references
func (q *PostTagsQuery) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *PostTagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
