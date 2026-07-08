package examples

import (
	"testing"

	"github.com/jalphad/gosqlgen/integration/ref"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	q "github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGettingStartedSelectActiveUsers(t *testing.T) {
	// Arrange
	qry := ref.NewQuery[models.UsersDto](nil, users.Table()).
		Select(
			users.Id(),
			users.Email(),
		).
		Where(
			users.Email().Like("%@corp.com").
				And(users.IsActive().IsTrue()),
		).
		OrderBy(q.Asc(users.Email())).
		Limit(10)

	// Act
	sql, params, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT users.id, users.email FROM users WHERE users.email LIKE $1 AND users.is_active = $2 ORDER BY users.email ASC LIMIT 10", sql)
	assert.Equal(t, []any{"%@corp.com", true}, params)
}
