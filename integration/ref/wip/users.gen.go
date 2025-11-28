package wip

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

// UsersDto represents the users table
type UsersDto struct {
	Id        *string    `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	FullName  *string    `db:"full_name" json:"full_name"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	IsActive  *bool      `db:"is_active" json:"is_active"`

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

// UsersFields provides type-safe field references for UsersDto
type UsersFields struct{}

var UsersTable = UsersFields{}

// Id returns a field reference for UsersDto.Id
func (u UsersFields) Id() *ast.selectExpression {
	return &ast.selectExpression{
		Expression: ast.expression{
			Field: &ast.FieldRef{
				Table:  "users",
				Column: "id",
			},
		},
	}
}

// Username returns a field reference for UsersDto.Username
func (u UsersFields) Username() *ast.selectExpression {
	return &ast.selectExpression{
		Expression: ast.expression{
			Field: &ast.FieldRef{
				Table:  "users",
				Column: "username",
			},
		},
	}
}

// Email returns a field reference for UsersDto.Email
func (u UsersFields) Email() *ast.selectExpression {
	return &ast.selectExpression{
		Expression: ast.expression{
			Field: &ast.FieldRef{
				Table:  "users",
				Column: "email",
			},
		},
	}
}

// FullName returns a field reference for UsersDto.FullName
func (u UsersFields) FullName() *FieldRef {
	return &FieldRef{
		Table:      "users",
		Expression: "full_name",
	}
}

// CreatedAt returns a field reference for UsersDto.CreatedAt
func (u UsersFields) CreatedAt() *FieldRef {
	return &FieldRef{
		Table:      "users",
		Expression: "created_at",
	}
}

// UpdatedAt returns a field reference for UsersDto.UpdatedAt
func (u UsersFields) UpdatedAt() *FieldRef {
	return &FieldRef{
		Table:      "users",
		Expression: "updated_at",
	}
}

// IsActive returns a field reference for UsersDto.IsActive
func (u UsersFields) IsActive() *FieldRef {
	return &FieldRef{
		Table:      "users",
		Expression: "is_active",
	}
}

// AllFields returns all field references for UsersDto
func (u UsersFields) AllFields() []*FieldRef {
	return []*FieldRef{
		u.Id(),
		u.Username(),
		u.Email(),
		u.FullName(),
		u.CreatedAt(),
		u.UpdatedAt(),
		u.IsActive(),
	}
}

// UsersQuery is a type-safe query builder for UsersDto
type UsersQuery struct {
	pool  *pgxpool.Pool
	tx    pgx.Tx
	query ast.SelectStatement

	selectFields []*FieldRef
	where        Expression[*UsersClauses]
	joins        []JoinClause
	activeJoins  map[string]bool // tracks which tables are joined for dynamic column selection
	orderBy      []OrderClause
	groupBy      []*FieldRef
	having       Expression[*UsersClauses]
	limit        *int
	offset       *int
}

// NewUsersQuery creates a new query builder
func NewUsersQuery(pool *pgxpool.Pool) *UsersQuery {
	return &UsersQuery{
		pool:        pool,
		activeJoins: make(map[string]bool),
	}
}

// WithTx sets the transaction
func (q *UsersQuery) WithTx(tx pgx.Tx) *UsersQuery {
	q.tx = tx
	return q
}

// Select specifies fields to select using type-safe field references
func (q *UsersQuery) Select(fields ...*FieldRef) *UsersQuery {
	q.selectFields = fields
	return q
}

func (q *UsersQuery) Where(expr Expression[*UsersClauses]) *UsersQuery {
	q.where = expr
	return q
}

// OrderById adds ORDER BY for Id
func (q *UsersQuery) OrderById(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.Id(),
		Direction: dir,
	})
	return q
}

// OrderByUsername adds ORDER BY for Username
func (q *UsersQuery) OrderByUsername(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.Username(),
		Direction: dir,
	})
	return q
}

// OrderByEmail adds ORDER BY for Email
func (q *UsersQuery) OrderByEmail(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.Email(),
		Direction: dir,
	})
	return q
}

// OrderByFullName adds ORDER BY for FullName
func (q *UsersQuery) OrderByFullName(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.FullName(),
		Direction: dir,
	})
	return q
}

// OrderByCreatedAt adds ORDER BY for CreatedAt
func (q *UsersQuery) OrderByCreatedAt(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.CreatedAt(),
		Direction: dir,
	})
	return q
}

// OrderByUpdatedAt adds ORDER BY for UpdatedAt
func (q *UsersQuery) OrderByUpdatedAt(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.UpdatedAt(),
		Direction: dir,
	})
	return q
}

// OrderByIsActive adds ORDER BY for IsActive
func (q *UsersQuery) OrderByIsActive(dir OrderDirection) *UsersQuery {
	q.orderBy = append(q.orderBy, OrderClause{
		Field:     UsersTable.IsActive(),
		Direction: dir,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *UsersQuery) GroupBy(fields ...*FieldRef) *UsersQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *UsersQuery) Having(expr Expression[*UsersClauses]) *UsersQuery {
	q.having = expr
	return q
}

// Limit sets the LIMIT
func (q *UsersQuery) Limit(limit int) *UsersQuery {
	q.limit = &limit
	return q
}

// Offset sets the OFFSET
func (q *UsersQuery) Offset(offset int) *UsersQuery {
	q.offset = &offset
	return q
}

func (q *UsersQuery) ToSql() (string, []any) {
	return q.buildQuery()
}

// buildQuery builds the SQL query
func (q *UsersQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		// Select all fields from main table
		query.WriteString("users.*")

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

	query.WriteString(" FROM users")

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
	if q.where != nil {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.where.ToSQL()
		whereStr = replaceFunc(whereStr, "$?", func(i int) string {
			return fmt.Sprintf("$%d", i+1)
		}, len(whereArgs))
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
	if q.having != nil {
		query.WriteString(" HAVING ")
		havingStr, havingArgs := q.having.ToSQL()
		havingStr = replaceFunc(havingStr, "$?", func(i int) string {
			return fmt.Sprintf("$%d", i+1)
		}, len(havingArgs))
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
func (q *UsersQuery) buildConditions(conditions []OldCondition, argIndex *int) (string, []interface{}) {
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
func (q *UsersQuery) Find(ctx context.Context) ([]*UsersDto, error) {
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

	var results []*UsersDto
	for rows.Next() {
		result := &UsersDto{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *UsersQuery) FindOne(ctx context.Context) (*UsersDto, error) {
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
func (q *UsersQuery) Count(ctx context.Context) (int64, error) {
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
func (q *UsersQuery) scanInto(rows pgx.Rows, dest *UsersDto) error {
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
		scanDest = append(scanDest, &dest.Username)
		scanDest = append(scanDest, &dest.Email)
		scanDest = append(scanDest, &dest.FullName)
		scanDest = append(scanDest, &dest.CreatedAt)
		scanDest = append(scanDest, &dest.UpdatedAt)
		scanDest = append(scanDest, &dest.IsActive)

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
func (q *UsersQuery) getScanDestForField(field *FieldRef, dest *UsersDto) (interface{}, func() error) {
	alias := field.Alias
	if alias == "" {
		// Use expression as alias if no explicit alias
		alias = field.Expression
	}

	// Normalize alias for matching (lowercase)
	aliasLower := strings.ToLower(alias)
	// Check if alias matches a reverse relationship collection field
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
	if aliasLower == "comments" {
		var jsonData []byte
		unmarshalFunc := func() error {
			if len(jsonData) > 0 && string(jsonData) != "null" {
				var items []*CommentsDto
				if err := json.Unmarshal(jsonData, &items); err != nil {
					return fmt.Errorf("field Comments: %w", err)
				}
				dest.Comments = items
			} else {
				dest.Comments = []*CommentsDto{}
			}
			return nil
		}
		return &jsonData, unmarshalFunc
	}

	// Check if alias matches a regular table column
	if field.Table == "users" {

		if aliasLower == "id" {
			return &dest.Id, nil
		}

		if aliasLower == "username" {
			return &dest.Username, nil
		}

		if aliasLower == "email" {
			return &dest.Email, nil
		}

		if aliasLower == "full_name" || aliasLower == "fullname" {
			return &dest.FullName, nil
		}

		if aliasLower == "created_at" || aliasLower == "createdat" {
			return &dest.CreatedAt, nil
		}

		if aliasLower == "updated_at" || aliasLower == "updatedat" {
			return &dest.UpdatedAt, nil
		}

		if aliasLower == "is_active" || aliasLower == "isactive" {
			return &dest.IsActive, nil
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
func (q *UsersQuery) Insert(ctx context.Context, record *UsersDto) error {
	var columns []string
	var placeholders []string
	var args []interface{}
	argIdx := 1

	// Dynamically build column list based on non-nil values

	if record.Id != nil {
		columns = append(columns, "id")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Id)
		argIdx++
	}

	// Required field (not nullable, no default)
	columns = append(columns, "username")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Username)
	argIdx++

	// Required field (not nullable, no default)
	columns = append(columns, "email")
	placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
	args = append(args, record.Email)
	argIdx++

	if record.FullName != nil {
		columns = append(columns, "full_name")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.FullName)
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

	if record.IsActive != nil {
		columns = append(columns, "is_active")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.IsActive)
		argIdx++
	}

	if len(columns) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	query := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s) RETURNING id",
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
func (q *UsersQuery) InsertBatch(ctx context.Context, records []*UsersDto) error {
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

		if record.Id != nil {
			columns = append(columns, "id")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.Id)
			argIdx++
		}

		// Required field (not nullable, no default)
		columns = append(columns, "username")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Username)
		argIdx++

		// Required field (not nullable, no default)
		columns = append(columns, "email")
		placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
		args = append(args, record.Email)
		argIdx++

		if record.FullName != nil {
			columns = append(columns, "full_name")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.FullName)
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

		if record.IsActive != nil {
			columns = append(columns, "is_active")
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
			args = append(args, record.IsActive)
			argIdx++
		}

		if len(columns) > 0 {
			query := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s) RETURNING id",
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
func (q *UsersQuery) Update(ctx context.Context, record *UsersDto) error {
	query := "UPDATE users SET " +
		"username = $2, email = $3, full_name = $4, created_at = $5, updated_at = $6, is_active = $7" +
		" WHERE id = $1"

	var tag pgconn.CommandTag
	var err error

	if q.tx != nil {
		tag, err = q.tx.Exec(ctx, query,
			record.Id,
			record.Username,
			record.Email,
			record.FullName,
			record.CreatedAt,
			record.UpdatedAt,
			record.IsActive,
		)
	} else {
		tag, err = q.pool.Exec(ctx, query,
			record.Id,
			record.Username,
			record.Email,
			record.FullName,
			record.CreatedAt,
			record.UpdatedAt,
			record.IsActive,
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
func (q *UsersQuery) UpdateFields(ctx context.Context, updates map[*FieldRef]interface{}) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE users SET ")

	setClauses := make([]string, 0, len(updates))
	for field, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field.Expression, argIndex))
		args = append(args, value)
		argIndex++
	}
	query.WriteString(strings.Join(setClauses, ", "))

	// Add WHERE conditions
	if q.where != nil {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.where.ToSQL()
		whereStr = replaceFunc(whereStr, "$?", func(i int) string {
			return fmt.Sprintf("$%d", i+1)
		}, len(whereArgs))
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
func (q *UsersQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}

	query.WriteString("DELETE FROM users")

	// Add WHERE conditions
	if q.where != nil {
		query.WriteString(" WHERE ")
		whereStr, whereArgs := q.where.ToSQL()
		whereStr = replaceFunc(whereStr, "$?", func(i int) string {
			return fmt.Sprintf("$%d", i+1)
		}, len(whereArgs))
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
func (q *UsersQuery) JoinOn(joinType JoinType, table string, leftField, rightField *FieldRef) *UsersQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
