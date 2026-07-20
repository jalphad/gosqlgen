package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertStatementSelectSource(t *testing.T) {
	// Arrange
	targetName := NewNamedExpression("name", NewColumnNode("tags", "name"))
	targetSlug := NewNamedExpression("slug", NewColumnNode("tags", "slug"))
	sourceName := NewNamedExpression("username", NewColumnNode("users", "username"))
	sourceSlug := NewNamedExpression("email", NewColumnNode("users", "email"))
	stmt := &InsertStatement{
		Table: "tags",
		Into:  []NamedExpression{targetName, targetSlug},
		Select: &SelectStatement{
			SelectList: []NamedExpression{sourceName, sourceSlug},
			From:       NewTableSource("users"),
			Where:      SetType[bool](NewLiteralExpression(true)),
		},
		Returning: []NamedExpression{NewNamedExpression("id", NewColumnNode("tags", "id"))},
	}
	params := make([]any, 0)
	ctx := &QueryContext{PrimaryTable: "tags"}

	// Act
	sql, err := RenderWithContext(stmt, &params, ctx)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "INSERT INTO tags(name, slug) SELECT users.username, users.email FROM users WHERE $1 RETURNING tags.id", sql)
	assert.Equal(t, []any{true}, params)
	assert.Equal(t, QueryTypeInsert, ctx.Type)
	assert.Equal(t, QueryPartReturning, ctx.CurrentPart)
}

func TestInsertStatementRequiresSource(t *testing.T) {
	// Arrange
	stmt := &InsertStatement{
		Table: "tags",
		Into:  []NamedExpression{NewNamedExpression("name", NewColumnNode("tags", "name"))},
	}
	params := make([]any, 0)

	// Act
	_, err := RenderWithContext(stmt, &params, &QueryContext{})

	// Assert
	require.EqualError(t, err, "INSERT requires a VALUES or SELECT source")
}

func TestInsertStatementRejectsSelectColumnCountMismatch(t *testing.T) {
	// Arrange
	stmt := &InsertStatement{
		Table: "tags",
		Into: []NamedExpression{
			NewNamedExpression("name", NewColumnNode("tags", "name")),
			NewNamedExpression("slug", NewColumnNode("tags", "slug")),
		},
		Select: &SelectStatement{
			SelectList: []NamedExpression{
				NewNamedExpression("username", NewColumnNode("users", "username")),
			},
			From: NewTableSource("users"),
		},
	}
	params := make([]any, 0)

	// Act
	_, err := RenderWithContext(stmt, &params, &QueryContext{})

	// Assert
	require.EqualError(t, err, "INSERT has 2 target columns but SELECT returns 1 columns")
}
