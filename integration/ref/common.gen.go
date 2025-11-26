package models

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Query struct {
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

// buildQuery builds the SQL query
func (q *Query) buildQuery() (string, []interface{}) {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	// SELECT clause
	query.WriteString("SELECT ")
	if len(q.selectFields) == 0 {
		// Select all fields from main table
		query.WriteString("comments.*")

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

	query.WriteString(" FROM comments")

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
func (q *Query) buildConditions(conditions []OldCondition, argIndex *int) (string, []interface{}) {
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

// OrderDirection represents sort order
type OrderDirection string

const (
	ASC  OrderDirection = "ASC"
	DESC OrderDirection = "DESC"
)

// ComparisonOp represents comparison operators
type ComparisonOp string

const (
	EQ   ComparisonOp = "="
	NEQ  ComparisonOp = "!="
	GT   ComparisonOp = ">"
	GTE  ComparisonOp = ">="
	LT   ComparisonOp = "<"
	LTE  ComparisonOp = "<="
	LIKE ComparisonOp = "LIKE"
	IN   ComparisonOp = "IN"
)

// LogicalOp represents logical operators
type LogicalOp string

const (
	AND LogicalOp = "AND"
	OR  LogicalOp = "OR"
)

// JoinType represents SQL join types
type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	FullJoin  JoinType = "FULL JOIN"
)

// FieldRef represents a type-safe field reference or SQL expression
type FieldRef struct {
	Table      string
	Expression string // Can be a column name or any SQL expression
	Alias      string
}

// As creates an aliased field reference
func (f *FieldRef) As(alias string) *FieldRef {
	return &FieldRef{
		Table:      f.Table,
		Expression: f.Expression,
		Alias:      alias,
	}
}

// String returns the SQL representation
func (f *FieldRef) String() string {
	if f.Alias != "" {
		if f.Table == "" {
			return fmt.Sprintf("%s AS %s", f.Expression, f.Alias)
		}
		return fmt.Sprintf("%s.%s AS %s", f.Table, f.Expression, f.Alias)
	}
	if f.Table == "" {
		return f.Expression
	}
	return fmt.Sprintf("%s.%s", f.Table, f.Expression)
}

// ExpressionFromString creates a FieldRef from a raw SQL expression
func ExpressionFromString(expression string, alias string) *FieldRef {
	return &FieldRef{
		Table:      "",
		Expression: expression,
		Alias:      alias,
	}
}

// OldCondition represents a WHERE condition
type OldCondition struct {
	Field    *FieldRef
	Op       ComparisonOp
	Value    interface{}
	Logical  LogicalOp
	SubConds []OldCondition
}

// JoinClause represents a JOIN operation
type JoinClause struct {
	Type       JoinType
	Table      string
	LeftField  *FieldRef
	RightField *FieldRef
}

// OrderClause represents an ORDER BY clause
type OrderClause struct {
	Field     *FieldRef
	Direction OrderDirection
}

type DateWrapper struct {
	time.Time
}

func (w *DateWrapper) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(time.DateOnly, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type TimeWrapper struct {
	time.Time
}

func (w *TimeWrapper) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(time.TimeOnly, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type TimeWithTZWrapper struct {
	time.Time
}

func (w *TimeWithTZWrapper) UnmarshalJSON(b []byte) error {
	timeWithTZLayout := "15:04:05.000000Z07"
	t, err := time.Parse(timeWithTZLayout, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type NoTimezoneWrapper struct {
	time.Time
}

func (w *NoTimezoneWrapper) UnmarshalJSON(b []byte) error {
	b[len(b)-1] = 'Z'
	t, err := time.Parse(time.RFC3339, string(b[1:]))
	w.Time = t

	return err
}

// JsonAgg creates a JSON aggregation expression with COALESCE for empty arrays
func JsonAgg(alias string, distinct bool, filter *FieldRef, fields ...*FieldRef) *FieldRef {
	if len(fields) == 0 {
		panic("JsonAgg requires at least one field")
	}

	// Build jsonb_build_object pairs
	var pairs []string
	for _, f := range fields {
		fieldName := f.Expression
		if f.Table != "" {
			fieldName = f.Table + "." + f.Expression
		}
		pairs = append(pairs, fmt.Sprintf("'%s', %s", f.Expression, fieldName))
	}

	distinctClause := ""
	if distinct {
		distinctClause = "DISTINCT "
	}

	filterClause := ""
	if filter != nil {
		filterClause = fmt.Sprintf(" FILTER (WHERE %s IS NOT NULL)", filter.String())
	}

	expression := fmt.Sprintf(
		"COALESCE(json_agg(%sjsonb_build_object(%s))%s, '[]'::json)",
		distinctClause,
		strings.Join(pairs, ", "),
		filterClause,
	)

	return &FieldRef{
		Table:      "",
		Expression: expression,
		Alias:      alias,
	}
}

// replaceFunc returns a copy of the string s with the first n
// non-overlapping instances of old replaced by the return value of newFn.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func replaceFunc(s, old string, newFn func(i int) string, n int) string {
	if n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	//b.Grow(len(s) + n*(len(newFn)-len(old)))
	start := 0
	if len(old) > 0 {
		for i := range n {
			j := start + strings.Index(s[start:], old)
			b.WriteString(s[start:j])
			b.WriteString(newFn(i))
			start = j + len(old)
		}
	} else { // len(old) == 0
		b.WriteString(newFn(0))
		for i := range n - 1 {
			_, wid := utf8.DecodeRuneInString(s[start:])
			j := start + wid
			b.WriteString(s[start:j])
			b.WriteString(newFn(i + 1))
			start = j
		}
	}
	b.WriteString(s[start:])
	return b.String()
}
