package query_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/require"
)

func TestSelectCTEToSQLArgs(t *testing.T) {
	type userSummary struct {
		Email string
	}

	active := users.As("active_users", users.Id(), users.Email())
	userID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	cteBody := models.NewQuery[models.UsersDto](nil, users.Table()).
		Select(users.Id(), users.Email()).
		Where(users.Email().Like("%@corp.com")).
		Statement()

	q := models.NewQuery[userSummary](nil, query.Table(active.Alias)).
		With(query.CTE(active.Alias, cteBody)).
		Select(query.Into(active.Email(), func(s *userSummary) *string {
			return &s.Email
		})).
		Where(active.Id().Eq(query.Val(userID)))

	sql, args, err := q.ToSqlArgs()

	require.NoError(t, err)
	require.Equal(t, "WITH active_users(id, email) AS (SELECT users.id, users.email FROM users WHERE users.email LIKE $1) SELECT active_users.email FROM active_users WHERE active_users.id = $2", sql)
	require.Equal(t, []any{"%@corp.com", userID}, args)
}

func TestUpdateCTEWithSetToToSQLArgs(t *testing.T) {
	type userSummary struct {
		Email string
	}

	fullName := "Updated User"
	id := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	dto := &models.UsersDto{
		Id:       &id,
		Email:    "updated@example.com",
		FullName: &fullName,
	}
	updated := users.As("updated_users", users.Id(), users.Email(), users.FullName())
	cteBody := models.NewUsersQuery(nil).
		Update(query.SetTo(dto, users.Email(), users.FullName())).
		Where(users.Id().Eq(query.Val(*dto.Id))).
		Returning(users.Id(), users.Email(), users.FullName()).
		Statement()

	q := models.NewQuery[userSummary](nil, query.Table(updated.Alias)).
		With(query.CTE(updated.Alias, cteBody)).
		Select(query.Into(updated.Email(), func(s *userSummary) *string {
			return &s.Email
		}))

	sql, args, err := q.ToSqlArgs()

	require.NoError(t, err)
	require.Equal(t, "WITH updated_users(id, email, full_name) AS (UPDATE users SET email = $1, full_name = $2 WHERE users.id = $3 RETURNING users.id, users.email, users.full_name) SELECT updated_users.email FROM updated_users", sql)
	require.Equal(t, []any{"updated@example.com", &fullName, id}, args)
}
