package query_test

import (
	"testing"

	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/require"
)

type userSummary struct {
	Email     string
	PostCount int64
}

func TestProjectionQueryToSQL(t *testing.T) {
	q := models.NewQuery[userSummary](nil, users.Table()).
		Select(
			query.Into(users.Email(), func(s *userSummary) *string {
				return &s.Email
			}),
			query.Into(query.Count(posts.Id()), func(s *userSummary) *int64 {
				return &s.PostCount
			}),
		).
		GroupBy(users.Email())

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "SELECT users.email, count(posts.id) FROM users GROUP BY users.email", sql)
}

func TestGeneratedDTOProjectionToSQL(t *testing.T) {
	q := models.NewQuery[models.UsersDto](nil, users.Table()).
		Select(
			users.Into.Id(),
			users.Into.Email(),
			users.Into.FullName(),
		)

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "SELECT users.id, users.email, users.full_name FROM users", sql)
}
