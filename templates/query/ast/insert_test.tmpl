package ast

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConflictUpdateRendersUpdateSets(t *testing.T) {
	// Arrange
	email := NewStringColumnExpression("users", "email")
	excludedEmail := String(NewColumnNode("EXCLUDED", "email"))
	action := Update(
		UpdateSet{Key: email, Value: excludedEmail},
		UpdateSet{Key: NewStringColumnExpression("users", "full_name"), Value: NewSQLType("Updated")},
	).Where(email.IsNotNull())
	params := []any{}

	// Act
	sql, err := RenderWithContext(action, &params, &QueryContext{})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, " UPDATE SET email = EXCLUDED.email, full_name = $1 WHERE users.email IS NOT NULL", sql)
	assert.Equal(t, []any{"Updated"}, params)
}

func TestConflictUpdateRejectsEmptyAssignments(t *testing.T) {
	// Arrange
	params := []any{}

	// Act
	_, err := RenderWithContext(Update(), &params, &QueryContext{})

	// Assert
	require.EqualError(t, err, "ON CONFLICT DO UPDATE requires at least one SET assignment")
}

func TestConflictUpdatePropagatesUpdateSetErrors(t *testing.T) {
	// Arrange
	params := []any{}
	action := Update(NewUpdateSetError(errors.New("invalid assignment")))

	// Act
	_, err := RenderWithContext(action, &params, &QueryContext{})

	// Assert
	require.EqualError(t, err, "invalid assignment")
}
