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

// Users represents records in the users table
type Users struct {
	Id        *string    `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	FullName  *string    `db:"full_name" json:"full_name"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	IsActive  *bool      `db:"is_active" json:"is_active"`
}

// TableName returns the table name for Users
func (u *Users) TableName() string {
	return "users"
}

// UsersFields provides type-safe field references for Users
type UsersFields struct{}

// Fields returns field references for Users
var UsersTable = UsersFields{}

// Id returns a field reference for Users.Id
func (u UsersFields) Id() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "id",
		GoType: "string",
	}
}

// Username returns a field reference for Users.Username
func (u UsersFields) Username() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "username",
		GoType: "string",
	}
}

// Email returns a field reference for Users.Email
func (u UsersFields) Email() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "email",
		GoType: "string",
	}
}

// FullName returns a field reference for Users.FullName
func (u UsersFields) FullName() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "full_name",
		GoType: "string",
	}
}

// CreatedAt returns a field reference for Users.CreatedAt
func (u UsersFields) CreatedAt() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "created_at",
		GoType: "time.Time",
	}
}

// UpdatedAt returns a field reference for Users.UpdatedAt
func (u UsersFields) UpdatedAt() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "updated_at",
		GoType: "time.Time",
	}
}

// IsActive returns a field reference for Users.IsActive
func (u UsersFields) IsActive() *FieldRef {
	return &FieldRef{
		Table:  "users",
		Column: "is_active",
		GoType: "bool",
	}
}

// UsersQuery is a type-safe query builder for Users
type UsersQuery struct {
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

// NewUsersQuery creates a new query builder
func NewUsersQuery(pool *pgxpool.Pool) *UsersQuery {
	return &UsersQuery{
		pool:    pool,
		logical: AND,
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

// SelectAll selects all fields
func (q *UsersQuery) SelectAll() *UsersQuery {
	q.selectFields = []*FieldRef{
		UsersTable.Id(),
		UsersTable.Username(),
		UsersTable.Email(),
		UsersTable.FullName(),
		UsersTable.CreatedAt(),
		UsersTable.UpdatedAt(),
		UsersTable.IsActive(),
	}
	return q
}

// WhereId adds a condition for Id
func (q *UsersQuery) WhereId(op ComparisonOp, value string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Id(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIdEq adds an equality condition for Id
func (q *UsersQuery) WhereIdEq(value string) *UsersQuery {
	return q.WhereId(EQ, value)
}

// WhereIdNotEq adds a not-equal condition for Id
func (q *UsersQuery) WhereIdNotEq(value string) *UsersQuery {
	return q.WhereId(NEQ, value)
}

// WhereIdLike adds a LIKE condition for Id
func (q *UsersQuery) WhereIdLike(pattern string) *UsersQuery {
	return q.WhereId(LIKE, pattern)
}

// WhereIdIn adds an IN condition for Id
func (q *UsersQuery) WhereIdIn(values ...string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Id(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
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

// WhereUsername adds a condition for Username
func (q *UsersQuery) WhereUsername(op ComparisonOp, value string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Username(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereUsernameEq adds an equality condition for Username
func (q *UsersQuery) WhereUsernameEq(value string) *UsersQuery {
	return q.WhereUsername(EQ, value)
}

// WhereUsernameNotEq adds a not-equal condition for Username
func (q *UsersQuery) WhereUsernameNotEq(value string) *UsersQuery {
	return q.WhereUsername(NEQ, value)
}

// WhereUsernameLike adds a LIKE condition for Username
func (q *UsersQuery) WhereUsernameLike(pattern string) *UsersQuery {
	return q.WhereUsername(LIKE, pattern)
}

// WhereUsernameIn adds an IN condition for Username
func (q *UsersQuery) WhereUsernameIn(values ...string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Username(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
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

// WhereEmail adds a condition for Email
func (q *UsersQuery) WhereEmail(op ComparisonOp, value string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Email(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereEmailEq adds an equality condition for Email
func (q *UsersQuery) WhereEmailEq(value string) *UsersQuery {
	return q.WhereEmail(EQ, value)
}

// WhereEmailNotEq adds a not-equal condition for Email
func (q *UsersQuery) WhereEmailNotEq(value string) *UsersQuery {
	return q.WhereEmail(NEQ, value)
}

// WhereEmailLike adds a LIKE condition for Email
func (q *UsersQuery) WhereEmailLike(pattern string) *UsersQuery {
	return q.WhereEmail(LIKE, pattern)
}

// WhereEmailIn adds an IN condition for Email
func (q *UsersQuery) WhereEmailIn(values ...string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.Email(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
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

// WhereFullName adds a condition for FullName
func (q *UsersQuery) WhereFullName(op ComparisonOp, value string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.FullName(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereFullNameEq adds an equality condition for FullName
func (q *UsersQuery) WhereFullNameEq(value string) *UsersQuery {
	return q.WhereFullName(EQ, value)
}

// WhereFullNameNotEq adds a not-equal condition for FullName
func (q *UsersQuery) WhereFullNameNotEq(value string) *UsersQuery {
	return q.WhereFullName(NEQ, value)
}

// WhereFullNameLike adds a LIKE condition for FullName
func (q *UsersQuery) WhereFullNameLike(pattern string) *UsersQuery {
	return q.WhereFullName(LIKE, pattern)
}

// WhereFullNameIn adds an IN condition for FullName
func (q *UsersQuery) WhereFullNameIn(values ...string) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.FullName(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereFullNameIsNull adds an IS NULL condition for FullName
func (q *UsersQuery) WhereFullNameIsNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.FullName(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereFullNameIsNotNull adds an IS NOT NULL condition for FullName
func (q *UsersQuery) WhereFullNameIsNotNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.FullName(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
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

// WhereCreatedAt adds a condition for CreatedAt
func (q *UsersQuery) WhereCreatedAt(op ComparisonOp, value time.Time) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.CreatedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtEq adds an equality condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtEq(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(EQ, value)
}

// WhereCreatedAtNotEq adds a not-equal condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtNotEq(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(NEQ, value)
}

// WhereCreatedAtGt adds a greater-than condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtGt(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(GT, value)
}

// WhereCreatedAtGte adds a greater-than-or-equal condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtGte(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(GTE, value)
}

// WhereCreatedAtLt adds a less-than condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtLt(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(LT, value)
}

// WhereCreatedAtLte adds a less-than-or-equal condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtLte(value time.Time) *UsersQuery {
	return q.WhereCreatedAt(LTE, value)
}

// WhereCreatedAtIn adds an IN condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtIn(values ...time.Time) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.CreatedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNull adds an IS NULL condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtIsNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.CreatedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereCreatedAtIsNotNull adds an IS NOT NULL condition for CreatedAt
func (q *UsersQuery) WhereCreatedAtIsNotNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.CreatedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
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

// WhereUpdatedAt adds a condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAt(op ComparisonOp, value time.Time) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.UpdatedAt(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtEq adds an equality condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtEq(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(EQ, value)
}

// WhereUpdatedAtNotEq adds a not-equal condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtNotEq(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(NEQ, value)
}

// WhereUpdatedAtGt adds a greater-than condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtGt(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(GT, value)
}

// WhereUpdatedAtGte adds a greater-than-or-equal condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtGte(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(GTE, value)
}

// WhereUpdatedAtLt adds a less-than condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtLt(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(LT, value)
}

// WhereUpdatedAtLte adds a less-than-or-equal condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtLte(value time.Time) *UsersQuery {
	return q.WhereUpdatedAt(LTE, value)
}

// WhereUpdatedAtIn adds an IN condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtIn(values ...time.Time) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.UpdatedAt(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtIsNull adds an IS NULL condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtIsNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.UpdatedAt(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereUpdatedAtIsNotNull adds an IS NOT NULL condition for UpdatedAt
func (q *UsersQuery) WhereUpdatedAtIsNotNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.UpdatedAt(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
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

// WhereIsActive adds a condition for IsActive
func (q *UsersQuery) WhereIsActive(op ComparisonOp, value bool) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.IsActive(),
		Op:      op,
		Value:   value,
		Logical: q.logical,
	})
	return q
}

// WhereIsActiveEq adds an equality condition for IsActive
func (q *UsersQuery) WhereIsActiveEq(value bool) *UsersQuery {
	return q.WhereIsActive(EQ, value)
}

// WhereIsActiveNotEq adds a not-equal condition for IsActive
func (q *UsersQuery) WhereIsActiveNotEq(value bool) *UsersQuery {
	return q.WhereIsActive(NEQ, value)
}

// WhereIsActiveIn adds an IN condition for IsActive
func (q *UsersQuery) WhereIsActiveIn(values ...bool) *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.IsActive(),
		Op:      IN,
		Value:   values,
		Logical: q.logical,
	})
	return q
}

// WhereIsActiveIsNull adds an IS NULL condition for IsActive
func (q *UsersQuery) WhereIsActiveIsNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.IsActive(),
		Op:      EQ,
		Value:   nil,
		Logical: q.logical,
	})
	return q
}

// WhereIsActiveIsNotNull adds an IS NOT NULL condition for IsActive
func (q *UsersQuery) WhereIsActiveIsNotNull() *UsersQuery {
	q.conditions = append(q.conditions, Condition{
		Field:   UsersTable.IsActive(),
		Op:      NEQ,
		Value:   nil,
		Logical: q.logical,
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

// And sets the logical operator to AND for subsequent conditions
func (q *UsersQuery) And() *UsersQuery {
	q.logical = AND
	return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *UsersQuery) Or() *UsersQuery {
	q.logical = OR
	return q
}

// WhereGroup adds a group of conditions
func (q *UsersQuery) WhereGroup(fn func(*UsersQuery)) *UsersQuery {
	subQuery := &UsersQuery{logical: AND}
	fn(subQuery)

	q.conditions = append(q.conditions, Condition{
		Logical:  q.logical,
		SubConds: subQuery.conditions,
	})
	return q
}

// GroupBy adds GROUP BY clause
func (q *UsersQuery) GroupBy(fields ...*FieldRef) *UsersQuery {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Having adds HAVING clause
func (q *UsersQuery) Having(field *FieldRef, op ComparisonOp, value interface{}) *UsersQuery {
	q.having = append(q.having, Condition{
		Field: field,
		Op:    op,
		Value: value,
	})
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

// buildQuery builds the SQL query
func (q *UsersQuery) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		query.WriteString("users.*")
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
func (q *UsersQuery) buildConditions(conditions []Condition, argIndex *int) (string, []interface{}) {
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
func (q *UsersQuery) Find(ctx context.Context) ([]*Users, error) {
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

	var results []*Users
	for rows.Next() {
		result := &Users{}
		if err := q.scanInto(rows, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

// FindOne returns a single result
func (q *UsersQuery) FindOne(ctx context.Context) (*Users, error) {
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
func (q *UsersQuery) scanInto(rows pgx.Rows, dest *Users) error {
	// If custom fields selected, use dynamic scanning
	if len(q.selectFields) > 0 && q.selectFields[0].Table != "" {
		// This would need more complex implementation for custom field scanning
		// For now, scan all fields
	}

	return rows.Scan(
		&dest.Id,
		&dest.Username,
		&dest.Email,
		&dest.FullName,
		&dest.CreatedAt,
		&dest.UpdatedAt,
		&dest.IsActive,
	)
}

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (q *UsersQuery) Insert(ctx context.Context, record *Users) error {
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
func (q *UsersQuery) InsertBatch(ctx context.Context, records []*Users) error {
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
func (q *UsersQuery) Update(ctx context.Context, record *Users) error {
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
func (q *UsersQuery) Delete(ctx context.Context) (int64, error) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("DELETE FROM users")

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
} // JoinOn performs a custom join with type-safe field references
func (q *UsersQuery) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *UsersQuery {
	q.joins = append(q.joins, JoinClause{
		Type:       joinType,
		Table:      table,
		LeftField:  leftField,
		RightField: rightField,
	})
	return q
}
