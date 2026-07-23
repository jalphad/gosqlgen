package ast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type contextTestValuesProvider struct {
	values [][]any
}

func (p contextTestValuesProvider) RowCount() int {
	return len(p.values)
}

func (p contextTestValuesProvider) ColumnCount() int {
	if len(p.values) == 0 {
		return 0
	}
	return len(p.values[0])
}

func (p contextTestValuesProvider) Value(row int, column int) (any, error) {
	if row < 0 || row >= len(p.values) || column < 0 || column >= len(p.values[row]) {
		return nil, fmt.Errorf("VALUES position (%d, %d) is out of range", row, column)
	}
	return p.values[row][column], nil
}

func TestQueryContextTypeAndPart(t *testing.T) {
	t.Run("SELECT query type is tracked", func(t *testing.T) {
		selectStmt := &SelectStatement{
			SelectList: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
			From: &TableSource{table: "users"},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(selectStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeSelect, ctx.Type)
	})

	t.Run("INSERT query type is tracked", func(t *testing.T) {
		insertStmt := &InsertStatement{
			Table:  "users",
			Into:   []NamedExpression{NewIntColumnExpression("users", "id")},
			Values: NewValuesTable(contextTestValuesProvider{values: [][]any{{1}}}),
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(insertStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeInsert, ctx.Type)
	})

	t.Run("UPDATE query type is tracked", func(t *testing.T) {
		updateStmt := &UpdateStatement{
			Table:   "users",
			SetList: []UpdateSetExpr{UpdateSet{Key: NewIntColumnExpression("users", "id"), Value: NewLiteralExpression(1)}},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(updateStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeUpdate, ctx.Type)
	})

	t.Run("DELETE query type is tracked", func(t *testing.T) {
		deleteStmt := &DeleteStatement{
			Table: "users",
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(deleteStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeDelete, ctx.Type)
	})

	t.Run("SELECT query parts are tracked", func(t *testing.T) {
		selectStmt := &SelectStatement{
			SelectList: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
			From:    &TableSource{table: "users"},
			Where:   NewSQLType(true),
			GroupBy: []Expression{NewIntColumnExpression("users", "id")},
			Having:  NewSQLType(true),
			OrderBy: []*OrderByItem{{Field: NewIntColumnExpression("users", "id"), Direction: Asc}},
			Limit:   &LimitClause{Limit: 10},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(selectStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		// After rendering, CurrentPart should be LIMIT (last part)
		assert.Equal(t, QueryTypeSelect, ctx.Type)
		assert.Equal(t, QueryPartLimit, ctx.CurrentPart)
	})

	t.Run("INSERT query parts are tracked", func(t *testing.T) {
		insertStmt := &InsertStatement{
			Table:  "users",
			Into:   []NamedExpression{NewIntColumnExpression("users", "id")},
			Values: NewValuesTable(contextTestValuesProvider{values: [][]any{{1}}}),
			Returning: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(insertStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeInsert, ctx.Type)
		assert.Equal(t, QueryPartReturning, ctx.CurrentPart)
	})

	t.Run("UPDATE query parts are tracked", func(t *testing.T) {
		updateStmt := &UpdateStatement{
			Table:   "users",
			SetList: []UpdateSetExpr{UpdateSet{Key: NewIntColumnExpression("users", "id"), Value: NewLiteralExpression(1)}},
			Where:   NewSQLType(true),
			Returning: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(updateStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeUpdate, ctx.Type)
		assert.Equal(t, QueryPartReturning, ctx.CurrentPart)
	})

	t.Run("DELETE query parts are tracked", func(t *testing.T) {
		deleteStmt := &DeleteStatement{
			Table: "users",
			Where: NewSQLType(true),
			Returning: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
		}

		ctx := &QueryContext{}
		_, err := RenderWithContext(deleteStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeDelete, ctx.Type)
		assert.Equal(t, QueryPartReturning, ctx.CurrentPart)
	})

	t.Run("JOIN part is tracked during rendering", func(t *testing.T) {
		selectStmt := &SelectStatement{
			SelectList: []NamedExpression{
				NewIntColumnExpression("users", "id"),
			},
			From: &TableSource{table: "users"},
		}
		selectStmt.From = selectStmt.From.Join(LeftJoin(&TableSource{table: "posts"}, On(NewSQLType(true))))

		ctx := &QueryContext{}
		_, err := RenderWithContext(selectStmt, &[]any{}, ctx)
		assert.NoError(t, err)

		assert.Equal(t, QueryTypeSelect, ctx.Type)
		// After rendering, should be in JOIN part (last thing rendered was the JOIN)
		assert.Equal(t, QueryPartJoin, ctx.CurrentPart)
	})
}

func TestNullPredicatesRenderAsBinaryIsExpressions(t *testing.T) {
	params := []any{}

	isNullSQL, err := RenderWithContext(NewStringColumnExpression("users", "email").IsNull(), &params, &QueryContext{})
	if err != nil {
		t.Fatal(err)
	}
	if isNullSQL != "users.email IS NULL" {
		t.Fatalf("expected users.email IS NULL, got %q", isNullSQL)
	}

	params = []any{}
	isNotNullSQL, err := RenderWithContext(NewStringColumnExpression("users", "email").IsNotNull(), &params, &QueryContext{})
	if err != nil {
		t.Fatal(err)
	}
	if isNotNullSQL != "users.email IS NOT NULL" {
		t.Fatalf("expected users.email IS NOT NULL, got %q", isNotNullSQL)
	}
}

func TestNot(t *testing.T) {
	t.Run("negates equality and preserves parameters", func(t *testing.T) {
		// Arrange
		params := []any{}
		predicate := NewStringColumnExpression("users", "name").Eq(NewSQLType("John"))

		// Act
		sql, err := RenderWithContext(Not(predicate), &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "NOT (users.name = $1)", sql)
		assert.Equal(t, []any{"John"}, params)
	})

	t.Run("negates like", func(t *testing.T) {
		// Arrange
		params := []any{}
		predicate := NewStringColumnExpression("users", "email").Like("%example.com")

		// Act
		sql, err := RenderWithContext(Not(predicate), &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "NOT (users.email LIKE $1)", sql)
		assert.Equal(t, []any{"%example.com"}, params)
	})

	t.Run("keeps chained negation boundaries explicit", func(t *testing.T) {
		// Arrange
		params := []any{}
		name := NewStringColumnExpression("users", "name").Eq(NewSQLType("John"))
		email := NewStringColumnExpression("users", "email").Like("%example.com")

		// Act
		sql, err := RenderWithContext(Not(name).And(Not(email)), &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "NOT (users.name = $1) AND NOT (users.email LIKE $2)", sql)
		assert.Equal(t, []any{"John", "%example.com"}, params)
	})

	t.Run("accepts any typed boolean expression", func(t *testing.T) {
		// Arrange
		params := []any{}
		expr := NewBoolColumnExpression("users", "active")

		// Act
		sql, err := RenderWithContext(Not(expr), &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "NOT (users.active)", sql)
		assert.Empty(t, params)
	})
}

func TestConcatNode(t *testing.T) {
	// Arrange
	params := []any{}
	expr := &ConcatNode{
		Left:  NewLiteralExpression("left"),
		Right: NewLiteralExpression("right"),
	}

	// Act
	sql, err := RenderWithContext(expr, &params, &QueryContext{})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "$1 $2", sql)
	assert.Equal(t, []any{"left", "right"}, params)
}

func TestBetween(t *testing.T) {
	t.Run("renders the receiver and both bounds", func(t *testing.T) {
		// Arrange
		params := []any{}
		predicate := NewIntColumnExpression("users", "age").Between(
			NewSQLType(int64(18)),
			NewSQLType(int64(65)),
		)

		// Act
		sql, err := RenderWithContext(predicate, &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "users.age BETWEEN $1 AND $2", sql)
		assert.Equal(t, []any{int64(18), int64(65)}, params)
	})

	t.Run("can be negated", func(t *testing.T) {
		// Arrange
		params := []any{}
		predicate := NewIntColumnExpression("users", "age").Between(
			NewSQLType(int64(18)),
			NewSQLType(int64(65)),
		)

		// Act
		sql, err := RenderWithContext(Not(predicate), &params, &QueryContext{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "NOT (users.age BETWEEN $1 AND $2)", sql)
		assert.Equal(t, []any{int64(18), int64(65)}, params)
	})
}
