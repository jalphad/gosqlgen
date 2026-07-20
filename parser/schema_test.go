package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_TracksSchemasAndAllowsDuplicateTableNames(t *testing.T) {
	// Arrange
	p := NewParser()
	sql := `
		CREATE TABLE auth.users (
			id UUID PRIMARY KEY
		);
		CREATE TABLE billing.users (
			id UUID PRIMARY KEY,
			auth_user_id UUID REFERENCES auth.users(id)
		);
		ALTER TABLE billing.users
			ADD CONSTRAINT billing_users_auth_user_fk
			FOREIGN KEY (auth_user_id) REFERENCES auth.users(id);
	`

	// Act
	err := p.Parse(sql)

	// Assert
	require.NoError(t, err)
	assert.Len(t, p.GetTables(), 2)

	authUsers, ok := p.GetTable("auth.users")
	require.True(t, ok)
	assert.Equal(t, "auth", authUsers.Schema)

	billingUsers, ok := p.GetTable("billing.users")
	require.True(t, ok)
	assert.Equal(t, "billing", billingUsers.Schema)
	require.NotEmpty(t, billingUsers.ForeignKeys)
	assert.Equal(t, "auth", billingUsers.ForeignKeys[0].ReferencedSchema)

	_, ok = p.GetTable("users")
	assert.False(t, ok, "bare lookup must be ambiguous")
}

func TestParser_TreatsUnqualifiedTablesAsPublic(t *testing.T) {
	// Arrange
	p := NewParser()

	// Act
	err := p.Parse(`CREATE TABLE users (id INTEGER PRIMARY KEY);`)

	// Assert
	require.NoError(t, err)
	users, ok := p.GetTable("users")
	require.True(t, ok)
	assert.Equal(t, "public", users.Schema)
}
