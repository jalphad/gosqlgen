package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCTEAliasRendersAsDeleteUsingReference(t *testing.T) {
	// Arrange
	id := NewIntColumnExpression("users", "id")
	newUser := NewTableSource("users").As("new_user", id)
	cteBody := &SelectStatement{SelectList: []NamedExpression{id}, From: NewTableSource("users")}
	stmt := &DeleteStatement{
		With:  []*CTE{NewCTE(newUser, cteBody)},
		Table: "users",
		Using: []NamedExpression{newUser},
	}
	params := []any{}

	// Act
	sql, err := RenderWithContext(stmt, &params, &QueryContext{})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "WITH new_user(id) AS (SELECT users.id FROM users) DELETE FROM users USING new_user", sql)
}

func TestSourceOwnedAliasRendering(t *testing.T) {
	// Arrange
	id := NewIntColumnExpression("users", "id")
	physicalAlias := NewTableSource("users").As("u", id)
	functionAlias := NewTableAlias("items", NewIntColumnExpression("items", "id"))
	function := NewSetReturningFunction(
		NewFunctionNode("unnest", []Expression{NewLiteralExpression([]int64{1, 2})}, nil),
	).As(functionAlias)

	// Act
	physicalSQL, physicalErr := RenderWithContext(physicalAlias, &[]any{}, &QueryContext{})
	params := []any{}
	functionSQL, functionErr := RenderWithContext(function, &params, &QueryContext{})

	// Assert
	require.NoError(t, physicalErr)
	require.NoError(t, functionErr)
	assert.Equal(t, "users AS u", physicalSQL)
	assert.Equal(t, "unnest($1) AS items(id)", functionSQL)
	assert.Equal(t, []any{[]int64{1, 2}}, params)
}

func TestLateralSetReturningFunctionKeepsAliasDefinition(t *testing.T) {
	// Arrange
	items := NewTableAlias("items", NewIntColumnExpression("items", "id"))
	function := NewSetReturningFunction(
		NewFunctionNode("unnest", []Expression{NewLiteralExpression([]int64{1, 2})}, nil),
	).As(items)
	table := NewTableSource("users").Join(CrossJoin(Lateral(function)))
	params := []any{}

	// Act
	sql, err := RenderWithContext(table, &params, &QueryContext{})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "users CROSS JOIN LATERAL unnest($1) AS items(id)", sql)
	assert.Equal(t, []any{[]int64{1, 2}}, params)
}

func TestCTEScopesExposePreviouslyDeclaredNamedRelations(t *testing.T) {
	// Arrange
	id := NewIntColumnExpression("users", "id")
	first := NewTableSource("users").As("first_user", id)
	second := NewTableSource("users").As("second_user", id)
	firstBody := &SelectStatement{SelectList: []NamedExpression{id}, From: NewTableSource("users")}
	secondBody := &SelectStatement{
		SelectList: []NamedExpression{NewIntColumnExpression("first_user", "id")},
		From:       first,
	}
	stmt := &SelectStatement{
		With:       []*CTE{NewCTE(first, firstBody), NewCTE(second, secondBody)},
		SelectList: []NamedExpression{NewIntColumnExpression("second_user", "id")},
		From:       second,
	}
	params := []any{}

	// Act
	sql, err := RenderWithContext(stmt, &params, &QueryContext{})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "WITH first_user(id) AS (SELECT users.id FROM users), second_user(id) AS (SELECT first_user.id FROM first_user) SELECT second_user.id FROM second_user", sql)
	assert.Same(t, firstBody.GetQueryContext().parent, secondBody.GetQueryContext().parent)
	assert.NotSame(t, firstBody.GetQueryContext(), secondBody.GetQueryContext())
}

func TestChildScopeRangeVisibilityPolicy(t *testing.T) {
	// Arrange
	alias := NewTableSource("users").As("u")
	parent := &QueryContext{}
	parent.registerRangeVariable(alias)
	isolated := newChildQueryContext(parent, false)
	correlated := newChildQueryContext(parent, true)

	// Act
	isolatedCanSeeAlias := isolated.lookupRangeVariable(alias)
	correlatedCanSeeAlias := correlated.lookupRangeVariable(alias)

	// Assert
	assert.False(t, isolatedCanSeeAlias)
	assert.True(t, correlatedCanSeeAlias)
}
