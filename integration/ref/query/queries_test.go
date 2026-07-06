package query

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersDtoSelectOneDefaultsToAllColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id: &id,
	}

	// Act
	qry, err := UsersDtoSelectOne(nil, user)
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT users.id, users.username, users.email, users.full_name, users.created_at, users.updated_at, users.is_active FROM users WHERE users.id = $1", sql)
}

func TestUsersDtoSelectOneUsesSelectedColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id: &id,
	}

	// Act
	qry, err := UsersDtoSelectOne(nil, user, SelectColumns(users.Id(), users.Email()))
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT users.id, users.email FROM users WHERE users.id = $1", sql)
}

func TestUsersDtoSelectManyDefaultsToAllColumns(t *testing.T) {
	// Arrange

	// Act
	qry, err := UsersDtoSelectMany(nil)
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT users.id, users.username, users.email, users.full_name, users.created_at, users.updated_at, users.is_active FROM users", sql)
}

func TestUsersDtoSelectManyUsesSelectedColumnsAndCanChainWhere(t *testing.T) {
	// Arrange
	id := uuid.New()

	// Act
	qry, err := UsersDtoSelectMany(nil, SelectColumns(users.Id(), users.Email()))
	require.NoError(t, err)
	sql, _, err := qry.Where(users.Id().Eq(Val(id))).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT users.id, users.email FROM users WHERE users.id = $1", sql)
}

func TestUsersDtoSelectManyErrorsForEmptySelectedColumns(t *testing.T) {
	// Arrange

	// Act
	_, err := UsersDtoSelectMany(nil, SelectColumns[models.UsersDto]())

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query.SelectColumns requires at least one column")
}

func TestUsersDtoSelectManyErrorsForNilSelectedColumn(t *testing.T) {
	// Arrange

	// Act
	_, err := UsersDtoSelectMany(nil, SelectColumns(users.Id(), nil))

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query.SelectColumns received a nil column")
}

func TestUsersDtoUpdateOneDefaultsToAllNonPrimaryKeyColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id:       &id,
		Username: "jane",
		Email:    "jane@example.com",
	}

	// Act
	qry, err := UsersDtoUpdateOne(nil, user)
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET created_at = $1, email = $2, full_name = $3, is_active = $4, updated_at = $5, username = $6 WHERE users.id = $7", sql)
}

func TestUsersDtoUpdateOneUsesSelectedColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id:    &id,
		Email: "jane@example.com",
	}

	// Act
	qry, err := UsersDtoUpdateOne(nil, user, UpdateColumns(users.Email()))
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = $1 WHERE users.id = $2", sql)
}

func TestUsersDtoUpdateOneErrorsForEmptySelectedColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id:    &id,
		Email: "jane@example.com",
	}

	// Act
	_, err := UsersDtoUpdateOne(nil, user, UpdateColumns())

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query.UpdateColumns requires at least one column")
}

func TestUsersDtoUpdateManyDefaultsToValuesWithSelectedColumns(t *testing.T) {
	// Arrange
	id1 := uuid.New()
	id2 := uuid.New()
	fullName1 := "Jane Doe"
	fullName2 := "Janet Doe"
	dtos := models.UsersDtos{
		{
			Id:       &id1,
			Email:    "jane@example.com",
			FullName: &fullName1,
		},
		{
			Id:       &id2,
			Email:    "janet@example.com",
			FullName: &fullName2,
		},
	}

	// Act
	qry, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns(users.Email(), users.FullName()))
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = v.email, full_name = v.full_name FROM (VALUES (CAST($1 AS uuid), CAST($2 AS text), CAST($3 AS text)), ($4, $5, $6)) AS v(id, email, full_name) WHERE users.id = v.id", sql)
}

func TestUsersDtoUpdateManyCanUseUnnestWithSelectedColumns(t *testing.T) {
	// Arrange
	id1 := uuid.New()
	id2 := uuid.New()
	fullName1 := "Jane Doe"
	fullName2 := "Janet Doe"
	dtos := models.UsersDtos{
		{
			Id:       &id1,
			Email:    "jane@example.com",
			FullName: &fullName1,
		},
		{
			Id:       &id2,
			Email:    "janet@example.com",
			FullName: &fullName2,
		},
	}

	// Act
	qry, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns(users.Email(), users.FullName()), UpdateWithUnnest())
	require.NoError(t, err)
	sql, _, err := qry.ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = v.email, full_name = v.full_name FROM unnest(CAST($1 AS uuid[]), CAST($2 AS text[]), CAST($3 AS text[])) AS v(id, email, full_name) WHERE users.id = v.id", sql)
}

func TestUsersDtoUpdateManyErrorsForEmptySelectedColumns(t *testing.T) {
	// Arrange
	id1 := uuid.New()
	id2 := uuid.New()
	dtos := models.UsersDtos{
		{
			Id:    &id1,
			Email: "jane@example.com",
		},
		{
			Id:    &id2,
			Email: "janet@example.com",
		},
	}

	// Act
	_, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns())

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query.UpdateColumns requires at least one column")
}
