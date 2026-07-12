package ast

import "testing"

func TestWindowFunctionRendering(t *testing.T) {
	t.Run("renders empty over clause", func(t *testing.T) {
		// Arrange
		expr := NewFunction[int64](NewFunctionNode("count", []Expression{NewIntColumnExpression("posts", "id")}, nil)).
			Over()

		// Act
		sql, err := RenderWithContext(expr, &[]any{}, &QueryContext{})

		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if sql != "count(posts.id) OVER ()" {
			t.Fatalf("expected empty OVER clause, got %q", sql)
		}
	})

	t.Run("renders partition by", func(t *testing.T) {
		// Arrange
		expr := NewFunction[int64](NewFunctionNode("count", []Expression{NewIntColumnExpression("posts", "id")}, nil)).
			Over().
			PartitionBy(NewIntColumnExpression("posts", "user_id"))

		// Act
		sql, err := RenderWithContext(expr, &[]any{}, &QueryContext{})

		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if sql != "count(posts.id) OVER (PARTITION BY posts.user_id)" {
			t.Fatalf("expected PARTITION BY clause, got %q", sql)
		}
	})

	t.Run("renders order by", func(t *testing.T) {
		// Arrange
		expr := NewFunction[int64](NewFunctionNode("count", []Expression{NewIntColumnExpression("posts", "id")}, nil)).
			Over().
			OrderBy(&OrderByItem{Field: NewTimestampColumnExpression("posts", "created_at"), Direction: Desc})

		// Act
		sql, err := RenderWithContext(expr, &[]any{}, &QueryContext{})

		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if sql != "count(posts.id) OVER (ORDER BY posts.created_at DESC)" {
			t.Fatalf("expected ORDER BY clause, got %q", sql)
		}
	})

	t.Run("renders partition by with order by", func(t *testing.T) {
		// Arrange
		expr := NewFunction[int64](NewFunctionNode("count", []Expression{NewIntColumnExpression("posts", "id")}, nil)).
			Over().
			PartitionBy(NewIntColumnExpression("posts", "user_id")).
			OrderBy(&OrderByItem{Field: NewTimestampColumnExpression("posts", "created_at"), Direction: Desc})

		// Act
		sql, err := RenderWithContext(expr, &[]any{}, &QueryContext{})

		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if sql != "count(posts.id) OVER (PARTITION BY posts.user_id ORDER BY posts.created_at DESC)" {
			t.Fatalf("expected combined window clause, got %q", sql)
		}
	})

	t.Run("renders aliased window function", func(t *testing.T) {
		// Arrange
		expr := NewFunction[int64](NewFunctionNode("count", []Expression{NewIntColumnExpression("posts", "id")}, nil)).
			Over().
			PartitionBy(NewIntColumnExpression("posts", "user_id")).
			As(NewColumnAlias("post_count"))

		// Act
		sql, err := RenderWithContext(expr, &[]any{}, &QueryContext{})

		// Assert
		if err != nil {
			t.Fatal(err)
		}
		if sql != "count(posts.id) OVER (PARTITION BY posts.user_id) AS post_count" {
			t.Fatalf("expected aliased window function, got %q", sql)
		}
	})
}
