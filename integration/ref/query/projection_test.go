package query_test

import (
	"strconv"
	"testing"

	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
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

func TestGeneratedFieldDescriptorProjectionToSQL(t *testing.T) {
	q := models.NewQuery[models.UsersDto](nil, users.Table()).
		Select(
			users.Id(),
			users.Email(),
			users.FullName(),
		)

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "SELECT users.id, users.email, users.full_name FROM users", sql)
}

func TestRelationshipProjectionToSQL(t *testing.T) {
	q := models.NewQuery[models.UsersDto](nil, users.Table()).
		Select(
			users.Into.Id(),
			users.Into.Posts(posts.Id(), posts.Title()),
		).
		Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id())).
		GroupBy(users.Id())

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "SELECT users.id, COALESCE(json_agg(DISTINCT jsonb_build_object('id', posts.id, 'title', posts.title)) FILTER (WHERE posts.id IS NOT NULL), CAST('[]' AS json)) AS posts FROM users LEFT JOIN posts ON posts.user_id = users.id GROUP BY users.id", sql)
}

func TestIntoCustomProjectionBinding(t *testing.T) {
	type customRow struct {
		Value int
	}

	projection := query.IntoCustom(
		query.Val("42"),
		func(r *customRow) *int { return &r.Value },
		strconv.Atoi,
	)
	var row customRow
	binding := projection.BindScan(&row)
	*binding.Destination().(*string) = "42"

	err := binding.Assign()

	require.NoError(t, err)
	require.Equal(t, 42, row.Value)
}

func TestIntoCustomProjectionBindingError(t *testing.T) {
	type customRow struct {
		Value int
	}

	projection := query.IntoCustom(
		query.Val("not an int"),
		func(r *customRow) *int { return &r.Value },
		strconv.Atoi,
	)
	var row customRow
	binding := projection.BindScan(&row)
	*binding.Destination().(*string) = "not an int"

	err := binding.Assign()

	require.Error(t, err)
	require.Zero(t, row.Value)
}

func TestInsertReturningProjectionToSQL(t *testing.T) {
	user := &models.UsersDto{
		Username: "janedoe",
		Email:    "jane@example.com",
	}

	q := models.NewUsersQuery(nil).
		Insert(users.Username(), users.Email()).
		Returning(users.Into.Id()).
		Values(user)

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING users.id", sql)
}
