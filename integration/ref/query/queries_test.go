package query

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type postWindowRow struct {
	Id      int64
	RowRank int64
}

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
	assert.Equal(t, "SELECT users.id, users.username, users.email, users.full_name, users.created_at, users.updated_at, users.is_active, users.attributes FROM users WHERE users.id = $1", sql)
}

func TestSelectWindowFunctionProjection(t *testing.T) {
	// Arrange
	rank := As("row_rank", RowNumber().
		Over().
		PartitionBy(posts.UserId()).
		OrderBy(Desc(posts.CreatedAt())))

	// Act
	sql, _, err := builder.NewKnownTableBuilder[postWindowRow](nil, posts.Table()).
		Select(
			Into(posts.Id(), func(row *postWindowRow) *int64 {
				return &row.Id
			}),
			Into(rank, func(row *postWindowRow) *int64 {
				return &row.RowRank
			}),
		).
		ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT posts.id, row_number() OVER (PARTITION BY posts.user_id ORDER BY posts.created_at DESC) AS row_rank FROM posts", sql)
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
	assert.Equal(t, "SELECT users.id, users.username, users.email, users.full_name, users.created_at, users.updated_at, users.is_active, users.attributes FROM users", sql)
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

func TestUsersDtoDeleteOneUsesPrimaryKey(t *testing.T) {
	// Arrange
	id := uuid.New()
	user := &models.UsersDto{
		Id: &id,
	}

	// Act
	sql, _, err := UsersDtoDeleteOne(nil, user).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "DELETE FROM users WHERE users.id = $1", sql)
}

func TestUsersDtoDeleteManyUsesPrimaryKeyValues(t *testing.T) {
	// Arrange
	id1 := uuid.New()
	id2 := uuid.New()
	dtos := models.UsersDtos{
		{Id: &id1},
		{Id: &id2},
	}

	// Act
	sql, _, err := UsersDtoDeleteMany(nil, dtos).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "DELETE FROM users USING (VALUES (CAST($1 AS uuid)), ($2)) AS v(id) WHERE users.id = v.id", sql)
}

func TestPostTagsDtoDeleteOneUsesCompositePrimaryKey(t *testing.T) {
	// Arrange
	dto := &models.PostTagsDto{
		PostId: 10,
		TagId:  20,
	}

	// Act
	sql, _, err := PostTagsDtoDeleteOne(nil, dto).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "DELETE FROM post_tags WHERE post_tags.post_id = $1 AND post_tags.tag_id = $2", sql)
}

func TestPostTagsDtoDeleteManyUsesCompositePrimaryKeyValues(t *testing.T) {
	// Arrange
	dtos := models.PostTagsDtos{
		{PostId: 10, TagId: 20},
		{PostId: 11, TagId: 21},
	}

	// Act
	sql, _, err := PostTagsDtoDeleteMany(nil, dtos).ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "DELETE FROM post_tags USING (VALUES (CAST($1 AS integer), CAST($2 AS integer)), ($3, $4)) AS v(post_id, tag_id) WHERE post_tags.post_id = v.post_id AND post_tags.tag_id = v.tag_id", sql)
}

func TestPostTagsDtoDeleteManyErrorsForEmptyList(t *testing.T) {
	// Arrange

	// Act
	_, _, err := PostTagsDtoDeleteMany(nil, models.PostTagsDtos{}).ToSql()

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VALUES table requires at least one row and one column")
}

func TestPostTagsDtoDeleteManyValueProviderErrorsForUnknownColumn(t *testing.T) {
	// Arrange
	provider := PostTagsDtoDeleteValuesProvider{
		dtos: models.PostTagsDtos{
			{PostId: 10, TagId: 20},
		},
	}

	// Act
	_, valueErr := provider.Value(0, 2)
	_, typeErr := provider.ColumnSQLType(2)

	// Assert
	require.Error(t, valueErr)
	assert.Contains(t, valueErr.Error(), "VALUES column 2 has no value")
	require.Error(t, typeErr)
	assert.Contains(t, typeErr.Error(), "VALUES column 2 has no SQL type")
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
	assert.Equal(t, "UPDATE users SET attributes = $1, created_at = $2, email = $3, full_name = $4, is_active = $5, updated_at = $6, username = $7 WHERE users.id = $8", sql)
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
