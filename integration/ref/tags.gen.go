package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TagsDto represents the tags table
type TagsDto struct {
	Id   *int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`

	// Many-to-many relationships (populated via LoadXxx methods)
	Posts []*PostsDto `manytomany:"post_tags"`

	// FromExpressions stores custom aggregations and expressions not mapped to fields
	FromExpressions map[string]interface{} `json:"from_expressions,omitempty"`
}

// TableName returns the table name for TagsDto
func (t *TagsDto) TableName() string {
	return "tags"
}

// LoadPosts loads associated posts through post_tags
func (t *TagsDto) LoadPosts(ctx context.Context, db *DB) error {
	if t.Id == nil {
		return nil
	}

	// Query junction table to get referenced IDs
	query := "SELECT post_id FROM post_tags WHERE tag_id = $1"

	rows, err := db.pool.Query(ctx, query, *t.Id)
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
		t.Posts = []*PostsDto{}
		return nil
	}

	results, err := db.Posts().WhereIdIn(ids...).Find(ctx)
	if err != nil {
		return err
	}
	t.Posts = results
	return nil
}

// TagsFields provides type-safe field references for TagsDto
type TagsFields struct{}

// TagsTable returns field references for TagsDto
var TagsTable = TagsFields{}

// Id returns a field reference for TagsDto.Id
func (t TagsFields) Id() *FieldRef {
	return &FieldRef{
		Table:      "tags",
		Expression: "id",
	}
}

// Name returns a field reference for TagsDto.Name
func (t TagsFields) Name() *FieldRef {
	return &FieldRef{
		Table:      "tags",
		Expression: "name",
	}
}

// Slug returns a field reference for TagsDto.Slug
func (t TagsFields) Slug() *FieldRef {
	return &FieldRef{
		Table:      "tags",
		Expression: "slug",
	}
}

// AllFields returns all field references for TagsDto
func (t TagsFields) AllFields() []*FieldRef {
	return []*FieldRef{
		t.Id(),
		t.Name(),
		t.Slug(),
	}
}

// AggregatePosts returns a JSON aggregation for Posts (many-to-many)
// Note: Requires manual joins through post_tags
func (t TagsFields) AggregatePosts(fields ...*FieldRef) *FieldRef {
	if len(fields) == 0 {
		fields = PostsTable.AllFields()
	}

	filterField := PostsTable.Id()

	return JsonAgg("posts", true, filterField, fields...)
}

// TagsQuery is a type-safe query builder for TagsDto
type TagsQuery struct {
	pool         *pgxpool.Pool
	tx           pgx.Tx
	selectFields []*FieldRef
	conditions   []OldCondition
	joins        []JoinClause
	activeJoins  map[string]bool // tracks which tables are joined for dynamic column selection
	orderBy      []OrderClause
	groupBy      []*FieldRef
	having       []OldCondition
	limit        *int
	offset       *int
	logical      LogicalOp
}

// NewTagsQuery creates a new query builder
func NewTagsQuery(pool *pgxpool.Pool) *TagsQuery {
	return &TagsQuery{
		pool:        pool,
		activeJoins: make(map[string]bool),
		logical:     AND,
	}
}

// WithTx sets the transaction
func (q *TagsQuery) WithTx(tx pgx.Tx) *TagsQuery {
	q.tx = tx
	return q
}

// Select specifies fields to select using type-safe field references
func (q *TagsQuery) Select(fields ...*FieldRef) *TagsQuery {
	q.selectFields = fields
	return q
}

// WhereId adds a condition for Id
func (q *TagsQuery) WhereId(op ComparisonOp, value int64) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Id(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIdEq adds an equality condition for Id
func (q *TagsQuery) WhereIdEq(value int64) *TagsQuery {
	return q.WhereId(EQ, value)
}

// WhereIdNotEq adds a not-equal condition for Id
func (q *TagsQuery) WhereIdNotEq(value int64) *TagsQuery {
	return q.WhereId(NEQ, value)
}

// WhereIdGt adds a greater-than condition for Id
func (q *TagsQuery) WhereIdGt(value int64) *TagsQuery {
	return q.WhereId(GT, value)
}

// WhereIdGte adds a greater-than-or-equal condition for Id
func (q *TagsQuery) WhereIdGte(value int64) *TagsQuery {
	return q.WhereId(GTE, value)
}

// WhereIdLt adds a less-than condition for Id
func (q *TagsQuery) WhereIdLt(value int64) *TagsQuery {
	return q.WhereId(LT, value)
}

// WhereIdLte adds a less-than-or-equal condition for Id
func (q *TagsQuery) WhereIdLte(value int64) *TagsQuery {
	return q.WhereId(LTE, value)
}

// WhereIdIn adds an IN condition for Id
func (q *TagsQuery) WhereIdIn(values ...int64) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Id(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderById adds ORDER BY for Id
func (q *TagsQuery) OrderById(dir OrderDirection) *TagsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     TagsTable.Id(),
		Direction: dir,
	})
	return q
}

// WhereName adds a condition for Name
func (q *TagsQuery) WhereName(op ComparisonOp, value string) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Name(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereNameEq adds an equality condition for Name
func (q *TagsQuery) WhereNameEq(value string) *TagsQuery {
	return q.WhereName(EQ, value)
}

// WhereNameNotEq adds a not-equal condition for Name
func (q *TagsQuery) WhereNameNotEq(value string) *TagsQuery {
	return q.WhereName(NEQ, value)
}

// WhereNameLike adds a LIKE condition for Name
func (q *TagsQuery) WhereNameLike(pattern string) *TagsQuery {
	return q.WhereName(LIKE, pattern)
}

// WhereNameIn adds an IN condition for Name
func (q *TagsQuery) WhereNameIn(values ...string) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Name(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderByName adds ORDER BY for Name
func (q *TagsQuery) OrderByName(dir OrderDirection) *TagsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     TagsTable.Name(),
		Direction: dir,
	})
	return q
}

// WhereSlug adds a condition for Slug
func (q *TagsQuery) WhereSlug(op ComparisonOp, value string) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Slug(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereSlugEq adds an equality condition for Slug
func (q *TagsQuery) WhereSlugEq(value string) *TagsQuery {
	return q.WhereSlug(EQ, value)
}

// WhereSlugNotEq adds a not-equal condition for Slug
func (q *TagsQuery) WhereSlugNotEq(value string) *TagsQuery {
	return q.WhereSlug(NEQ, value)
}

// WhereSlugLike adds a LIKE condition for Slug
func (q *TagsQuery) WhereSlugLike(pattern string) *TagsQuery {
	return q.WhereSlug(LIKE, pattern)
}

// WhereSlugIn adds an IN condition for Slug
func (q *TagsQuery) WhereSlugIn(values ...string) *TagsQuery {
	q.conditions = append(q.conditions, OldCondition{
		Field:   TagsTable.Slug(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// OrderBySlug adds ORDER BY for Slug
func (q *TagsQuery) OrderBySlug(dir OrderDirection) *TagsQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     TagsTable.Slug(),
		Direction: dir,
	})
	return q
}

// And sets the logical operator to AND for subsequent conditions
func (q *TagsQuery) And() *TagsQuery {
	q.logical = AND
	return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *TagsQuery) Or() *TagsQuery {
	q.logical = OR
	return q
}

// WhereGroup adds a group of conditions
func (q *TagsQuery) WhereGroup(fn func(*TagsQuery)) *TagsQuery {
	subQuery := &TagsQuery{logical: AND}
	fn(subQuery)

	q.conditions = append(q.conditions, OldCondition{
		Logical:  q.logical,
		SubConds: subQuery.conditions,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *TagsQuery) GroupBy(fields ...*FieldRef) *TagsQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *TagsQuery) Having(field *FieldRef, op ComparisonOp, value interface{}) *TagsQuery {
	q.having = append(q.having, OldCondition{
		Field: field,
		Op:    op,
		Value: value,
	})
	return q
}

// Limit sets the LIMIT
func (q *TagsQuery) Limit(limit int) *TagsQuery {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET
func (q *TagsQuery) Offset(offset int) *TagsQuery {
	q.offset = &offset
	return q
}

// buildQuery builds the SQL query
func (q *TagsQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		// Select all fields from main table
		query.WriteString("tags.*")

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

	query.WriteString(" FROM tags")

	// JOIN clauses
	for _, join := range q.joins {
		query.WriteString(" ")
		query.WriteString(string(join.Type))
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
func (q *TagsQuery) buildConditions(conditions []OldCondition, argIndex *int) (string, []interface{}) {
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
func (q *TagsQuery) Find(ctx context.Context) ([]*TagsDto, error) {
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

	var results []*TagsDto
	for rows.Next() {
		result := &TagsDto{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *TagsQuery) FindOne(ctx context.Context) (*TagsDto, error) {
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
func (q *TagsQuery) Count(ctx context.Context) (int64, error) {
	// Save and restore select fields
	originalSelect := q.selectFields
	countField := &FieldRef{Table: "", Expression: "COUNT(*)", Alias: ""}
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
func (q *TagsQuery) scanInto(rows pgx.Rows, dest *TagsDto) error {
	var scanDest []interface{}
	var jsonUnmarshalFuncs []func() error

	if len(q.selectFields) > 0 {
		// Use selectFields - scan in exact order
		for _, field := range q.selectFields {
			destPtr, unmarshalFunc := q.getScanDestForField(field, dest)
			scanDest = append(scanDest, destPtr)
			if unmarshalFunc != nil {
				jsonUnmarshalFuncs = append(jsonUnmarshalFuncs, unmarshalFunc)
			}
		}
	} else {
		// Default behavior - scan all table columns + joined tables
		scanDest = append(scanDest, &dest.Id)
		scanDest = append(scanDest, &dest.Name)
		scanDest = append(scanDest, &dest.Slug)

		// Add joined table columns if active
	}

	// Perform scan
	if err := rows.Scan(scanDest...); err != nil {
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
func (q *TagsQuery) getScanDestForField(field *FieldRef, dest *TagsDto) (interface{}, func() error) {
	alias := field.Alias
	if alias == "" {
		// Use expression as alias if no explicit alias
		alias = field.Expression
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(alias)
	// Check if alias matches a many-to-many collection field
	if aliasLower == "posts" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []*PostsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Posts: %w", err)
				}
				dest.Posts = items
			} else {
				dest.Posts = []*PostsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}

	// Check if alias matches a regular table column
	if field.Table == "tags" {

		if aliasLower == "id" {
			return &dest.Id, nil
		}

		if aliasLower == "name" {
			return &dest.Name, nil
		}

		if aliasLower == "slug" {
			return &dest.Slug, nil
		}
	}

	// No match - store in FromExpressions as interface{}
	if dest.FromExpressions == nil {
		dest.FromExpressions = make(map[string]interface{})
	}
	var value interface{}
	// Store a pointer to value that we'll populate after scan
	unmarshalFunc := func() error {
		dest.FromExpressions[alias] = value
		return nil
	}
	return &value, unmarshalFunc
}

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (q *TagsQuery) Insert(ctx context.Context, record *TagsDto) error {
	var columns []string
	var placeholders []string
	var args []interface{}
	argIdx := 1

	// Dynamically build column list based on non-nil values

	// Required field (not nullable, no default)
	columns = append(columns, "name")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Name)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "slug")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Slug)
	argIdx++

	if len(columns) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	query := fmt.Sprintf("INSERT INTO tags (%s) VALUES (%s) RETURNING id",
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
func (q *TagsQuery) InsertBatch(ctx context.Context, records []*TagsDto) error {
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
		columns = append(columns, "name")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Name)
		argIdx++

		// Required field (not nullable, no default)
		columns = append(columns, "slug")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Slug)
		argIdx++

		if len(columns) > 0 {
			query := fmt.Sprintf("INSERT INTO tags (%s) VALUES (%s) RETURNING id",
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
func (q *TagsQuery) Update(ctx context.Context, record *TagsDto) error {
	query := "UPDATE tags SET " +
		"name = $2, slug = $3" +
		" WHERE id = $1"

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query,
			record.Id,
			record.Name,
			record.Slug,
		)
	} else {
		tag, err = q.pool.Exec(ctx, query,
			record.Id,
			record.Name,
			record.Slug,
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
func (q *TagsQuery) UpdateFields(ctx context.Context, updates map[*FieldRef]interface{}) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE tags SET ")

	setClauses := make([]string, 0, len(updates))
	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field.Expression, argIndex))
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
func (q *TagsQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("DELETE FROM tags")

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

// JoinOn performs a custom join with type-safe field references
func (q *TagsQuery) JoinOn(joinType JoinType, table string, leftField, rightField *FieldRef) *TagsQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
