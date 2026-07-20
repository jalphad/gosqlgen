package query

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type postWindowRow struct {
	Id      int64
	RowRank int64
}

type usersSelfJoinRow struct {
	Manager     models.UsersDto
	Subordinate models.UsersDto
}

func (r *usersSelfJoinRow) ManagerDto() *models.UsersDto {
	return &r.Manager
}

func (r *usersSelfJoinRow) SubordinateDto() *models.UsersDto {
	return &r.Subordinate
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

func TestUsersSelfJoinUsesAliasScopedDestinations(t *testing.T) {
	// Arrange
	manager := users.AliasFor(users.Table().As("manager"), (*usersSelfJoinRow).ManagerDto)
	subordinate := users.AliasFor(users.Table().As("subordinate"), (*usersSelfJoinRow).SubordinateDto)

	// Act
	sql, _, err := builder.NewKnownTableBuilder[usersSelfJoinRow](nil, manager.TableAlias).
		Select(manager.Id(), manager.Email(), subordinate.Id(), subordinate.Email()).
		Join(ast.JoinLeft, subordinate, subordinate.IsActive().Eq(manager.IsActive())).
		ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT manager.id, manager.email, subordinate.id, subordinate.email FROM users AS manager LEFT JOIN users AS subordinate ON subordinate.is_active = manager.is_active", sql)
}

func TestCTEAliasJoinDoesNotRenderColumnList(t *testing.T) {
	// Arrange
	activeUsers := users.As("active_users", users.Id(), users.Email())
	cteBody := builder.NewStatementBuilder(users.Table()).
		Select(users.Id(), users.Email()).
		Statement()

	// Act
	sql, _, err := builder.NewKnownTableBuilder[models.PostsDto](nil, posts.Table()).
		With(ast.NewCTE(activeUsers.TableAlias, cteBody)).
		Select(posts.Id()).
		Join(ast.JoinLeft, activeUsers, posts.UserId().Eq(activeUsers.Id())).
		ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "WITH active_users(id, email) AS (SELECT users.id, users.email FROM users) SELECT posts.id FROM posts LEFT JOIN active_users ON posts.user_id = active_users.id", sql)
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

func TestAggregateFunctionHelpers(t *testing.T) {
	// Arrange
	sumViews := As("sum_views", Sum(posts.ViewCount()))
	avgViews := As("avg_views", Avg(posts.ViewCount()))
	minViews := As("min_views", Min(posts.ViewCount()))
	maxViews := As("max_views", Max(posts.ViewCount()))

	// Act
	sql, _, err := builder.NewKnownTableBuilder[postWindowRow](nil, posts.Table()).
		Select(
			Into(sumViews, func(row *postWindowRow) *int64 {
				return &row.Id
			}),
			Into(avgViews, func(row *postWindowRow) *int64 {
				return &row.Id
			}),
			Into(minViews, func(row *postWindowRow) *int64 {
				return &row.Id
			}),
			Into(maxViews, func(row *postWindowRow) *int64 {
				return &row.Id
			}),
		).
		ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT sum(posts.view_count) AS sum_views, avg(posts.view_count) AS avg_views, min(posts.view_count) AS min_views, max(posts.view_count) AS max_views FROM posts", sql)
}

func TestStatementOnlySelectSupportsNamedExpressions(t *testing.T) {
	// Arrange
	postCount := Count(posts.Id()).As(ast.NewColumnAlias("post_count"))
	stmt := builder.NewStatementBuilder(posts.Table()).
		Select(posts.UserId(), postCount).
		Where(posts.Status().Eq(Val("published"))).
		GroupBy(posts.UserId()).
		OrderBy(Asc(posts.UserId())).
		Limit(10).
		Statement()
	params := make([]any, 0)

	// Act
	sql, err := ast.RenderWithContext(stmt, &params, &ast.QueryContext{PrimaryTable: posts.Table().Name()})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "SELECT posts.user_id, count(posts.id) AS post_count FROM posts WHERE posts.status = $1 GROUP BY posts.user_id ORDER BY posts.user_id ASC LIMIT 10", sql)
	assert.Equal(t, []any{"published"}, params)
}

func TestStatementOnlyUpdateSupportsNamedReturning(t *testing.T) {
	// Arrange
	id := uuid.New()
	stmt := builder.NewStatementBuilder(users.Table()).
		Update(Set(users.Email()).ToValue("updated@example.com")).
		Where(users.Id().Eq(Val(id))).
		Returning(users.Id(), users.Email()).
		Statement()
	params := make([]any, 0)

	// Act
	sql, err := ast.RenderWithContext(stmt, &params, &ast.QueryContext{PrimaryTable: users.Table().Name()})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "UPDATE users SET email = $1 WHERE users.id = $2 RETURNING users.id, users.email", sql)
	assert.Equal(t, []any{"updated@example.com", id}, params)
}

func TestStatementOnlyDeleteSupportsNamedReturning(t *testing.T) {
	// Arrange
	stmt := builder.NewStatementBuilder(users.Table()).
		Delete().
		Where(users.Email().Eq(Val("deleted@example.com"))).
		Returning(users.Id()).
		Statement()
	params := make([]any, 0)

	// Act
	sql, err := ast.RenderWithContext(stmt, &params, &ast.QueryContext{PrimaryTable: users.Table().Name()})

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "DELETE FROM users WHERE users.email = $1 RETURNING users.id", sql)
	assert.Equal(t, []any{"deleted@example.com"}, params)
}

func TestInsertFromStatementSelectRendersInSQLOrder(t *testing.T) {
	// Arrange
	source := builder.NewStatementBuilder(users.Table()).
		Select(users.Username(), users.Email()).
		Where(users.Email().Eq(Val("source@example.com")))

	// Act
	sql, params, err := builder.NewKnownTableBuilder[models.TagsDto](nil, tags.Table()).
		Insert(tags.Name(), tags.Slug()).
		Select(source).
		Returning(tags.Id()).
		ToSql()

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "INSERT INTO tags(name, slug) SELECT users.username, users.email FROM users WHERE users.email = $1 RETURNING tags.id", sql)
	assert.Equal(t, []any{"source@example.com"}, params)
}

func TestInsertValuesRejectsEmptyAndNilRows(t *testing.T) {
	// Arrange
	var nilTag *models.TagsDto

	// Act
	_, _, emptyErr := builder.NewKnownTableBuilder[models.TagsDto](nil, tags.Table()).
		Insert(tags.Name()).
		Values().
		ToSql()
	_, _, nilErr := builder.NewKnownTableBuilder[models.TagsDto](nil, tags.Table()).
		Insert(tags.Name()).
		Values(nilTag).
		ToSql()

	// Assert
	require.EqualError(t, emptyErr, "VALUES table requires at least one row and one column")
	require.EqualError(t, nilErr, "INSERT VALUES row 0 is nil")
}

func TestInsertValuesRejectsNonProjectionColumns(t *testing.T) {
	// Arrange
	tag := &models.TagsDto{Name: "invalid"}
	column := ast.NewNamedExpression("name", ast.NewColumnNode("tags", "name"))

	// Act
	_, _, err := builder.NewKnownTableBuilder[models.TagsDto](nil, tags.Table()).
		Insert(column).
		Values(tag).
		ToSql()

	// Assert
	require.EqualError(t, err, `INSERT column "name" is not a projection for the inserted type`)
}

func TestInsertValuesRejectsNilColumns(t *testing.T) {
	// Arrange
	tag := &models.TagsDto{Name: "invalid"}

	// Act
	_, _, err := builder.NewKnownTableBuilder[models.TagsDto](nil, tags.Table()).
		Insert(nil).
		Values(tag).
		ToSql()

	// Assert
	require.EqualError(t, err, "INSERT column 0 is nil")
}

func TestStatementOnlyFinalizeInterfaceExposesOnlyStatement(t *testing.T) {
	// Arrange
	finalize := reflect.TypeOf((*builder.StatementFinalizeQuery)(nil)).Elem()

	// Act
	method, ok := finalize.MethodByName("Statement")

	// Assert
	require.True(t, ok)
	assert.Equal(t, "Statement", method.Name)
	assert.Equal(t, 1, finalize.NumMethod())
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
