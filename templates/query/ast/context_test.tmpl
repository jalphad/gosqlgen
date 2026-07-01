package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			Values: []Expression{NewLiteralExpression(1)},
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
			Values: []Expression{NewLiteralExpression(1)},
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
		selectStmt.From.Join(JoinLeft, &TableSource{table: "posts"}, NewSQLType(true))

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
