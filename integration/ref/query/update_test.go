package query_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/require"
)

func TestUpdateSetToDTOToSQL(t *testing.T) {
	fullName := "Update User"
	active := true
	id := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	user := &models.UsersDto{
		Id:       &id,
		Username: "update-user",
		Email:    "update@example.com",
		FullName: &fullName,
		IsActive: &active,
	}

	q := models.NewUsersQuery(nil).
		Update(query.SetTo(user,
			users.Username(),
			users.Email(),
			users.FullName(),
			users.IsActive(),
		)).
		Where(users.Id().Eq(query.Val(*user.Id)))

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "UPDATE users SET username = $1, email = $2, full_name = $3, is_active = $4 WHERE users.id = $5", sql)

	_, args, err := q.ToSqlArgs()
	require.NoError(t, err)
	require.Equal(t, []any{"update-user", "update@example.com", &fullName, &active, id}, args)
}

func TestUpdateSetToExpressionToSQL(t *testing.T) {
	q := models.NewUsersQuery(nil).
		Update(query.Set(users.Email()).To(query.Val("updated@example.com"))).
		Where(users.Username().Eq(query.Val("update-user")))

	sql, err := q.ToSql()

	require.NoError(t, err)
	require.Equal(t, "UPDATE users SET email = $1 WHERE users.username = $2", sql)
}
