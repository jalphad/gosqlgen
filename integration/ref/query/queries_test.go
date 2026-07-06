package query

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersDtoUpdateOneDefaultsToAllNonPrimaryKeyColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id:       &id,
		Username: "jane",
		Email:    "jane@example.com",
	}

	// Act
	sql, _, err := UsersDtoUpdateOne(nil, user).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET username = $1, email = $2, full_name = $3, created_at = $4, updated_at = $5, is_active = $6 WHERE users.id = $7", sql)
}

func TestUsersDtoUpdateOneUsesSelectedColumns(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id:    &id,
		Email: "jane@example.com",
	}

	// Act
	sql, _, err := UsersDtoUpdateOne(nil, user, UpdateColumns(users.Email())).ToSql()

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
	_, _, err := UsersDtoUpdateOne(nil, user, UpdateColumns()).ToSql()

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query.SetTo requires at least one field")
}

func TestUsersDtoUpdateManyDefaultsToUnnestWithSelectedColumns(t *testing.T) {
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
	sql, _, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns(users.Email(), users.FullName())).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = v.email, full_name = v.full_name FROM unnest(CAST($1 AS uuid[]), CAST($2 AS text[]), CAST($3 AS text[])) AS v(id, email, full_name) WHERE users.id = v.id", sql)
}

func TestUsersDtoUpdateManyCanUseValuesWithSelectedColumns(t *testing.T) {
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
	sql, _, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns(users.Email(), users.FullName()), UpdateWithValues()).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = v.email, full_name = v.full_name FROM (VALUES ($1, $2, $3), ($4, $5, $6)) AS v(id, email, full_name) WHERE users.id = v.id", sql)
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
	_, _, err := UsersDtoUpdateMany(nil, dtos, UpdateColumns()).ToSql()

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "UpdateSetList has no sets")
}
