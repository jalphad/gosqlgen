package ast

import (
	"errors"
	"strings"
	"testing"
)

func TestErrorExpression(t *testing.T) {
	t.Run("ErrorExpression sets error in context", func(t *testing.T) {
		err := NewErrorExpression(errors.New("test error"))
		ctx := &QueryContext{}

		var builder strings.Builder
		var params []any

		err.toSQL(&builder, &params, ctx)

		if ctx.Error == nil {
			t.Fatal("expected error to be set in context")
		}
		if ctx.Error.Error() != "test error" {
			t.Errorf("expected error 'test error', got '%v'", ctx.Error.Error())
		}
	})

	t.Run("GroupedExpression with no expressions sets error", func(t *testing.T) {
		grouped := NewGroupedExpression()
		ctx := &QueryContext{}

		var builder strings.Builder
		var params []any

		grouped.toSQL(&builder, &params, ctx)

		if ctx.Error == nil {
			t.Fatal("expected error for empty GroupedExpression")
		}
		if ctx.Error.Error() != "GroupedExpression has no expressions" {
			t.Errorf("unexpected error message: %v", ctx.Error.Error())
		}
	})

	t.Run("toSQL methods return early if error already set", func(t *testing.T) {
		ctx := &QueryContext{PrimaryTable: "users"}
		ctx.Error = errors.New("first error")

		var builder strings.Builder
		var params []any

		selectStmt := &SelectStatement{
			SelectList: []NamedExpression{
				NewIntColumnExpression("users", "username"),
			},
			From: &TableSource{Table: "users"},
		}

		selectStmt.toSQL(&builder, &params, ctx)

		// Verify error was not overwritten
		if ctx.Error.Error() != "first error" {
			t.Errorf("expected 'first error', got '%v'", ctx.Error.Error())
		}
		// Verify no SQL was written due to early return
		if builder.Len() > 0 {
			t.Errorf("expected empty builder, got '%s'", builder.String())
		}
	})

	t.Run("RenderWithContext returns error from context", func(t *testing.T) {
		grouped := NewGroupedExpression()

		ctx := &QueryContext{PrimaryTable: "users"}
		sql, err := RenderWithContext(grouped, &[]any{}, ctx)

		if err == nil {
			t.Fatal("expected error from RenderWithContext")
		}
		if err.Error() != "GroupedExpression has no expressions" {
			t.Errorf("expected error 'GroupedExpression has no expressions', got '%v'", err)
		}
		// SQL should be empty since GroupedExpression returned before writing anything
		if sql != "" {
			t.Errorf("expected empty SQL, got '%s'", sql)
		}
	})

	t.Run("KeywordNode rejects invalid tokens", func(t *testing.T) {
		ctx := &QueryContext{}
		sql, err := RenderWithContext(NewKeywordNode("NULL; DROP TABLE users"), &[]any{}, ctx)

		if err == nil {
			t.Fatal("expected error from invalid keyword token")
		}
		if sql != "" {
			t.Errorf("expected empty SQL, got '%s'", sql)
		}
	})
}
