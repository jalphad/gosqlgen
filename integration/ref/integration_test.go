package ref

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	q "github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/post_tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	pgxPool  *pgxpool.Pool
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

type userCountRow struct {
	models.UsersDto
	Count int64
}

func (r *userCountRow) GetUsersDto() *models.UsersDto {
	return &r.UsersDto
}

func TestMain(m *testing.M) {
	var err error

	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	if err = pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err = pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15-alpine",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=testuser",
			"POSTGRES_DB=testdb",
			"listen_addresses='*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf("postgres://testuser:secret@%s/testdb?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseURL)

	resource.Expire(120)
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		pgxPool, err = pgxpool.New(context.Background(), databaseURL)
		if err != nil {
			return err
		}
		return pgxPool.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err := loadSchema(pgxPool); err != nil {
		log.Fatalf("Could not load schema: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func loadSchema(pool *pgxpool.Pool) error {
	schema, err := os.ReadFile("../schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = pool.Exec(context.Background(), string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}

func cleanupTables(t *testing.T) {
	t.Helper()

	for _, table := range []string{"post_tags", "comments", "posts", "tags", "users"} {
		_, err := pgxPool.Exec(context.Background(), fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		require.NoErrorf(t, err, "failed to truncate table %s", table)
	}
}

func insertUsers(t *testing.T, rows ...*models.UsersDto) {
	t.Helper()
	_, inserted, err := NewUsersQuery(pgxPool).
		Insert(users.Username(), users.Email(), users.FullName(), users.IsActive(), users.Attributes()).
		Values(rows...).
		Returning(users.Id()).
		Exec(context.Background())
	require.NoError(t, err)
	require.Len(t, inserted, len(rows))
	for i, row := range rows {
		row.Id = inserted[i].Id
		require.NotNil(t, row.Id)
	}
}

func insertPosts(t *testing.T, rows ...*models.PostsDto) {
	t.Helper()
	_, inserted, err := NewPostsQuery(pgxPool).
		Insert(posts.UserId(), posts.Title(), posts.Content(), posts.Status(), posts.ViewCount()).
		Values(rows...).
		Returning(posts.Id()).
		Exec(context.Background())
	require.NoError(t, err)
	require.Len(t, inserted, len(rows))
	for i, row := range rows {
		row.Id = inserted[i].Id
		require.NotNil(t, row.Id)
	}
}

func insertComments(t *testing.T, rows ...*models.CommentsDto) {
	t.Helper()
	_, inserted, err := NewCommentsQuery(pgxPool).
		Insert(comments.PostId(), comments.UserId(), comments.Content(), comments.IsApproved()).
		Values(rows...).
		Returning(comments.Id()).
		Exec(context.Background())
	require.NoError(t, err)
	require.Len(t, inserted, len(rows))
	for i, row := range rows {
		row.Id = inserted[i].Id
		require.NotNil(t, row.Id)
	}
}

func insertTags(t *testing.T, rows ...*models.TagsDto) {
	t.Helper()
	_, inserted, err := NewTagsQuery(pgxPool).
		Insert(tags.Into.Name(), tags.Into.Slug()).
		Values(rows...).
		Returning(tags.Into.Id()).
		Exec(context.Background())
	require.NoError(t, err)
	require.Len(t, inserted, len(rows))
	for i, row := range rows {
		row.Id = inserted[i].Id
		require.NotNil(t, row.Id)
	}
}

func findUserByID(t *testing.T, id uuid.UUID) models.UsersDto {
	t.Helper()
	user, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
		Select(users.Into.AllColumns()...).
		Where(users.Id().Eq(q.Val(id))).
		FindOne(context.Background())
	require.NoError(t, err)
	return user
}

func findPostsByUserID(t *testing.T, userID uuid.UUID) []models.PostsDto {
	t.Helper()
	rows, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
		Select(posts.Id(), posts.Title(), posts.UserId(), posts.Content(), posts.Status(), posts.PublishedAt(), posts.ViewCount()).
		Where(posts.UserId().Eq(q.Val(userID))).
		Find(context.Background())
	require.NoError(t, err)
	return rows
}

func TestIntegration_UserCRUD(t *testing.T) {
	cleanupTables(t)

	t.Run("Create User", func(t *testing.T) {
		attributes := map[string]string{"foo": "bar"}
		jsonAttr, _ := json.Marshal(attributes)
		user := &models.UsersDto{Username: "johndoe", Email: "john@example.com", FullName: new("John Doe"), IsActive: new(true), Attributes: (*json.RawMessage)(&jsonAttr)}

		insertUsers(t, user)

		assert.NotEqual(t, uuid.Nil, *user.Id)
	})

	t.Run("Create Users", func(t *testing.T) {
		user1 := &models.UsersDto{Username: "johndoe1", Email: "john1@example.com", FullName: new("John Doe"), IsActive: new(true)}
		user2 := &models.UsersDto{Username: "johndoe2", Email: "john2@example.com", FullName: new("John Doe 2"), IsActive: new(true)}

		insertUsers(t, user1, user2)

		assert.NotEqual(t, uuid.Nil, *user1.Id)
		assert.NotEqual(t, uuid.Nil, *user2.Id)
	})

	t.Run("Find User by ID", func(t *testing.T) {
		attributes := map[string]string{"foo": "bar"}
		jsonAttr, _ := json.Marshal(attributes)
		user := &models.UsersDto{Username: "janedoe", Email: "jane@example.com", Attributes: (*json.RawMessage)(&jsonAttr)}
		insertUsers(t, user)

		found := findUserByID(t, *user.Id)
		foundAttr := make(map[string]string)
		err := json.Unmarshal(*found.Attributes, &foundAttr)

		require.NoError(t, err)
		assert.Equal(t, user.Username, found.Username)
		assert.Equal(t, user.Email, found.Email)
		assert.Equal(t, attributes, foundAttr)
	})

	t.Run("Update User", func(t *testing.T) {
		user := &models.UsersDto{Username: "updateme", Email: "update@example.com"}
		insertUsers(t, user)
		user.Email = "updated@example.com"
		user.FullName = new("Updated Name")

		affected, _, err := NewUsersQuery(pgxPool).
			Update(q.SetTo(user, users.Email(), users.FullName())).
			Where(users.Id().Eq(q.Val(*user.Id))).
			Exec(context.Background())

		require.NoError(t, err)
		assert.Equal(t, int64(1), affected)
		found := findUserByID(t, *user.Id)
		assert.Equal(t, user.Email, found.Email)
		assert.Equal(t, user.FullName, found.FullName)
	})

	t.Run("Update Users", func(t *testing.T) {
		user1 := &models.UsersDto{Username: "updatemany1", Email: "many1@example.com"}
		user2 := &models.UsersDto{Username: "updatemany2", Email: "many2@example.com"}
		insertUsers(t, user1, user2)
		name1 := "Updated Name 1"
		name2 := "Updated Name 2"
		user1.Email = "updated1@example.com"
		user1.FullName = &name1
		user2.Email = "updated2@example.com"
		user2.FullName = &name2

		dtos := models.UsersDtos{user1, user2}
		v := ast.NewTableAlias("v", users.Id(), users.Email(), users.FullName())
		affected, _, err := NewUsersQuery(pgxPool).
			Update(
				q.Set(users.Email()).To(q.Rel(v, users.Email())),
				q.Set(users.FullName()).To(q.Rel(v, users.FullName())),
			).
			From(q.Unnest(
				q.Cast(dtos.Id()).AsUUIDArray(),
				q.Cast(dtos.Email()).AsTextArray(),
				q.Cast(dtos.FullName()).AsTextArray(),
			).As(v)).
			Where(users.Id().Eq(q.Rel(v, users.Id()))).
			Exec(context.Background())

		require.NoError(t, err)
		assert.Equal(t, int64(2), affected)
		assert.Equal(t, user1.Email, findUserByID(t, *user1.Id).Email)
		assert.Equal(t, user2.Email, findUserByID(t, *user2.Id).Email)
	})

	t.Run("Delete User", func(t *testing.T) {
		user := &models.UsersDto{Username: "deleteme", Email: "delete@example.com"}
		insertUsers(t, user)

		deleted, details, err := NewUsersQuery(pgxPool).
			Delete().
			Where(users.Id().Eq(q.Val(*user.Id))).
			Exec(context.Background())

		require.NoError(t, err)
		assert.EqualValues(t, 1, deleted)
		assert.Nil(t, details)
		_, err = NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Id()).
			Where(users.Id().Eq(q.Val(*user.Id))).
			FindOne(context.Background())
		require.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("Create User with Conflict", func(t *testing.T) {
		user := &models.UsersDto{Username: "conflictedjohndoe", Email: "john@conflict.example.com", FullName: new("John Doe"), IsActive: new(true)}
		insertUsers(t, user)
		user.Email = "updated@conflict.example.com"

		_, _, err := NewUsersQuery(pgxPool).
			Insert(users.Email(), users.Username(), users.FullName(), users.IsActive()).
			Values(user).
			OnConflict(users.Username()).Do(ast.Update(users.Email())).
			Returning(users.Id()).
			Exec(context.Background())

		require.NoError(t, err)
		found := findUserByID(t, *user.Id)
		assert.Equal(t, user.Email, found.Email)
	})

	t.Run("Update user returning full name", func(t *testing.T) {
		fullName := "Update Me"
		user := &models.UsersDto{Username: "updatemeandreturn", Email: "updateme@returning.example.com", FullName: &fullName}
		insertUsers(t, user)

		affected, details, err := NewUsersQuery(pgxPool).
			Update(
				q.Set(users.Email()).ToValue("updated@returning.example.com"),
				q.Set(users.Username()).ToValue("updatedname"),
			).
			Where(users.Id().Eq(q.Val(*user.Id))).
			Returning(users.FullName()).
			Exec(context.Background())

		require.NoError(t, err)
		assert.Equal(t, int64(1), affected)
		require.Len(t, details, 1)
		assert.Equal(t, fullName, *details[0].FullName)
	})

	t.Run("Delete User returning id", func(t *testing.T) {
		user := &models.UsersDto{Username: "deletereturn", Email: "delete-return@example.com"}
		insertUsers(t, user)

		deleted, details, err := NewUsersQuery(pgxPool).
			Delete().
			Where(users.Id().Eq(q.Val(*user.Id))).
			Returning(users.Id()).
			Exec(context.Background())

		require.NoError(t, err)
		assert.EqualValues(t, 1, deleted)
		require.Len(t, details, 1)
		assert.Equal(t, *user.Id, *details[0].Id)
	})
}

func TestIntegration_QueryBuilder(t *testing.T) {
	cleanupTables(t)

	seed := []*models.UsersDto{
		{Username: "alice", Email: "alice@example.com", FullName: new("Alice Smith"), IsActive: new(true)},
		{Username: "bob", Email: "bob@example.com", FullName: new("Bob Jones"), IsActive: new(true)},
		{Username: "charlie", Email: "charlie@example.com", FullName: new("Charlie Brown"), IsActive: new(false)},
		{Username: "david", Email: "david@example.com", FullName: new("David Wilson"), IsActive: new(true)},
	}
	insertUsers(t, seed...)

	t.Run("Where Equals", func(t *testing.T) {
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.Username().Eq(q.Val("alice"))).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "alice", results[0].Username)
	})

	t.Run("Where LIKE", func(t *testing.T) {
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.Username().Like("a%")).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 1)
	})

	t.Run("Where IN", func(t *testing.T) {
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.Username().Eq(q.Val("alice")).Or(users.Username().Eq(q.Val("bob")))).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("Where Boolean", func(t *testing.T) {
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.IsActive().IsTrue()).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 3)
	})

	t.Run("Order By", func(t *testing.T) {
		expected := []string{"alice", "bob", "charlie", "david"}
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			OrderBy(q.Asc(users.Username())).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 4)
		for i, result := range results {
			assert.Equal(t, expected[i], result.Username)
		}
	})

	t.Run("Limit and Offset", func(t *testing.T) {
		results, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			OrderBy(q.Asc(users.Username())).
			Limit(2).
			Offset(1).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 2)
		assert.Equal(t, "bob", results[0].Username)
		assert.Equal(t, "charlie", results[1].Username)
	})

	t.Run("Count", func(t *testing.T) {
		type countRow struct{ Count int64 }
		results, err := NewQuery[countRow](pgxPool, users.Table()).
			Select(q.Into(q.Count(users.Id()).As(ast.NewColumnAlias("count")), func(r *countRow) *int64 { return &r.Count })).
			Where(users.IsActive().IsTrue()).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.EqualValues(t, 3, results[0].Count)
	})

	t.Run("Count With Embedded DTO", func(t *testing.T) {
		count := q.Into(q.Count(users.Id()).As(ast.NewColumnAlias("count")), func(r *userCountRow) *int64 { return &r.Count })
		user := users.For((*userCountRow).GetUsersDto)
		results, err := NewQuery[userCountRow](pgxPool, users.Table()).
			Select(user.Id(), user.Username(), user.Email(), count).
			Where(users.IsActive().IsTrue()).
			GroupBy(users.Id(), users.Username(), users.Email()).
			OrderBy(q.Asc(users.Username())).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 3)
		assert.Equal(t, "alice", results[0].Username)
		assert.Equal(t, "bob", results[1].Username)
		assert.Equal(t, "david", results[2].Username)
		for _, result := range results {
			require.NotNil(t, result.Id)
			assert.NotEmpty(t, result.Email)
			assert.EqualValues(t, 1, result.Count)
		}
	})
}

func TestIntegration_RelationshipProjection(t *testing.T) {
	cleanupTables(t)

	t.Run("Load posts into user DTO", func(t *testing.T) {
		user := &models.UsersDto{Username: "projectionuser", Email: "projection@example.com"}
		insertUsers(t, user)
		post1 := &models.PostsDto{UserId: *user.Id, Title: "Projection Post 1"}
		post2 := &models.PostsDto{UserId: *user.Id, Title: "Projection Post 2"}
		insertPosts(t, post1, post2)

		result, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Id(), users.Into.PostsByUserId(posts.Id(), posts.Title())).
			Join(ast.JoinLeft, posts.Table(), posts.UserId().Eq(users.Id())).
			Where(users.Id().Eq(q.Val(*user.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.PostsByUserId, 2)
		assert.ElementsMatch(t, []string{"Projection Post 1", "Projection Post 2"}, []string{result.PostsByUserId[0].Title, result.PostsByUserId[1].Title})
	})

	t.Run("Empty relationship becomes empty slice", func(t *testing.T) {
		user := &models.UsersDto{Username: "emptyprojectionuser", Email: "empty-projection@example.com"}
		insertUsers(t, user)

		result, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Id(), users.Into.PostsByUserId(posts.Id(), posts.Title())).
			Join(ast.JoinLeft, posts.Table(), posts.UserId().Eq(users.Id())).
			Where(users.Id().Eq(q.Val(*user.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.NotNil(t, result.PostsByUserId)
		assert.Empty(t, result.PostsByUserId)
	})
}

func TestIntegration_Joins(t *testing.T) {
	cleanupTables(t)

	user := &models.UsersDto{Username: "blogger", Email: "blogger@example.com"}
	insertUsers(t, user)
	postRows := []*models.PostsDto{
		{UserId: *user.Id, Title: "First Post", Content: new("Content 1"), Status: new("published")},
		{UserId: *user.Id, Title: "Second Post", Content: new("Content 2"), Status: new("draft")},
	}
	insertPosts(t, postRows...)

	t.Run("Join Users", func(t *testing.T) {
		type postWithUser struct {
			PostID   int64
			Title    string
			Username string
		}
		results, err := NewQuery[postWithUser](pgxPool, posts.Table()).
			Select(
				q.Into(posts.Id(), func(p *postWithUser) *int64 { return &p.PostID }),
				q.Into(posts.Title(), func(p *postWithUser) *string { return &p.Title }),
				q.Into(users.Username(), func(p *postWithUser) *string { return &p.Username }),
			).
			Join(ast.JoinLeft, users.Table(), posts.UserId().Eq(users.Id())).
			Where(posts.UserId().Eq(q.Val(*user.Id))).
			OrderBy(q.Asc(posts.Id())).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 2)
		assert.Equal(t, "blogger", results[0].Username)
	})

	t.Run("Filter by Related Table", func(t *testing.T) {
		results, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Title(), posts.UserId(), posts.Content(), posts.Status(), posts.PublishedAt(), posts.ViewCount()).
			Where(posts.Status().Eq(q.Val("published"))).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.Equal(t, postRows[0].Title, results[0].Title)
	})
}

func TestIntegration_Transactions(t *testing.T) {
	cleanupTables(t)

	t.Run("Successful Transaction", func(t *testing.T) {
		var user *models.UsersDto
		var post *models.PostsDto

		err := pgx.BeginFunc(context.Background(), pgxPool, func(tx pgx.Tx) error {
			user = &models.UsersDto{Username: "txuser", Email: "tx@example.com"}
			_, insertedUsers, err := NewUsersQuery(nil).WithTx(tx).
				Insert(users.Username(), users.Email()).
				Values(user).
				Returning(users.Id()).
				Exec(context.Background())
			if err != nil {
				return err
			}
			user.Id = insertedUsers[0].Id

			post = &models.PostsDto{UserId: *user.Id, Title: "Transaction Post"}
			_, insertedPosts, err := NewPostsQuery(nil).WithTx(tx).
				Insert(posts.UserId(), posts.Title()).
				Values(post).
				Returning(posts.Id()).
				Exec(context.Background())
			if err == nil {
				post.Id = insertedPosts[0].Id
			}
			return err
		})

		require.NoError(t, err)
		usersFound, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.Username().Eq(q.Val("txuser"))).
			Find(context.Background())
		require.NoError(t, err)
		assert.Len(t, usersFound, 1)
		postsFound := findPostsByUserID(t, *user.Id)
		require.Len(t, postsFound, 1)
		assert.Equal(t, post.Title, postsFound[0].Title)
	})

	t.Run("Failed Transaction Rollback", func(t *testing.T) {
		err := pgx.BeginFunc(context.Background(), pgxPool, func(tx pgx.Tx) error {
			user := &models.UsersDto{Username: "rollbackuser", Email: "rollback@example.com"}
			if _, _, err := NewUsersQuery(nil).WithTx(tx).
				Insert(users.Username(), users.Email()).
				Values(user).
				Returning(users.Id()).
				Exec(context.Background()); err != nil {
				return err
			}
			return assert.AnError
		})

		require.Error(t, err)
		usersFound, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(users.Username().Eq(q.Val("rollbackuser"))).
			Find(context.Background())
		require.NoError(t, err)
		assert.Empty(t, usersFound)
	})
}

func TestIntegration_InsertFromSelect(t *testing.T) {
	cleanupTables(t)

	// Arrange
	insertUsers(t,
		&models.UsersDto{Username: "selected-user", Email: "selected@example.com"},
		&models.UsersDto{Username: "ignored-user", Email: "ignored@example.com"},
	)
	source := NewQuery[models.UsersDto](pgxPool, users.Table()).
		Select(users.Username(), users.Email()).
		Where(users.Username().Eq(q.Val("selected-user")))

	// Act
	affected, inserted, err := NewTagsQuery(pgxPool).
		Insert(tags.Name(), tags.Slug()).
		Select(source).
		Returning(tags.Id(), tags.Name(), tags.Slug()).
		Exec(context.Background())

	// Assert
	require.NoError(t, err)
	assert.EqualValues(t, 1, affected)
	require.Len(t, inserted, 1)
	require.NotNil(t, inserted[0].Id)
	assert.Equal(t, "selected-user", inserted[0].Name)
	assert.Equal(t, "selected@example.com", inserted[0].Slug)
}

func TestIntegration_NullableFields(t *testing.T) {
	cleanupTables(t)

	t.Run("Insert with NULL values", func(t *testing.T) {
		user := &models.UsersDto{Username: "nulltest", Email: "null@example.com", FullName: nil}
		insertUsers(t, user)

		found := findUserByID(t, *user.Id)

		assert.Nil(t, found.FullName)
	})

	t.Run("Update to NULL", func(t *testing.T) {
		user := &models.UsersDto{Username: "nullupdate", Email: "nullupdate@example.com", FullName: new("Initial Name")}
		insertUsers(t, user)
		user.FullName = nil

		affected, _, err := NewUsersQuery(pgxPool).
			Update(q.Set(users.FullName()).ToNullable(user.FullName)).
			Where(users.Id().Eq(q.Val(*user.Id))).
			Exec(context.Background())

		require.NoError(t, err)
		assert.Equal(t, int64(1), affected)
		assert.Nil(t, findUserByID(t, *user.Id).FullName)
	})
}

func TestIntegration_ComplexQueries(t *testing.T) {
	cleanupTables(t)

	user1 := &models.UsersDto{Username: "user1", Email: "user1@example.com"}
	user2 := &models.UsersDto{Username: "user2", Email: "user2@example.com"}
	insertUsers(t, user1, user2)
	postRows := []*models.PostsDto{
		{UserId: *user1.Id, Title: "User1 Post 1", ViewCount: new(int64(100))},
		{UserId: *user1.Id, Title: "User1 Post 2", ViewCount: new(int64(200))},
		{UserId: *user2.Id, Title: "User2 Post 1", ViewCount: new(int64(50))},
	}
	insertPosts(t, postRows...)

	t.Run("Greater Than Query", func(t *testing.T) {
		results, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Title(), posts.UserId(), posts.Content(), posts.Status(), posts.PublishedAt(), posts.ViewCount()).
			Where(posts.ViewCount().Gt(q.Val(int64(75)))).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("Multiple Conditions", func(t *testing.T) {
		results, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Title(), posts.UserId(), posts.Content(), posts.Status(), posts.PublishedAt(), posts.ViewCount()).
			Where(posts.UserId().Eq(q.Val(*user1.Id)).And(posts.ViewCount().Gte(q.Val(int64(100))))).
			Find(context.Background())

		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("UpdateFields", func(t *testing.T) {
		affected, _, err := NewPostsQuery(pgxPool).
			Update(q.Set(posts.ViewCount()).ToValue(int64(999))).
			Where(posts.UserId().Eq(q.Val(*user1.Id))).
			Exec(context.Background())

		require.NoError(t, err)
		assert.EqualValues(t, 2, affected)
		for _, post := range findPostsByUserID(t, *user1.Id) {
			require.NotNil(t, post.ViewCount)
			assert.EqualValues(t, 999, *post.ViewCount)
		}
	})
}

func TestIntegration_ReverseRelationships(t *testing.T) {
	cleanupTables(t)

	user1 := &models.UsersDto{Username: "author1", Email: "author1@example.com"}
	user2 := &models.UsersDto{Username: "author2", Email: "author2@example.com"}
	insertUsers(t, user1, user2)
	post1 := &models.PostsDto{UserId: *user1.Id, Title: "Post 1", Content: new("Content 1")}
	post2 := &models.PostsDto{UserId: *user1.Id, Title: "Post 2", Content: new("Content 2")}
	post3 := &models.PostsDto{UserId: *user2.Id, Title: "Post 3", Content: new("Content 3")}
	insertPosts(t, post1, post2, post3)
	comment1 := &models.CommentsDto{PostId: *post1.Id, UserId: *user1.Id, Content: "Comment 1"}
	comment2 := &models.CommentsDto{PostId: *post1.Id, UserId: *user2.Id, Content: "Comment 2"}
	insertComments(t, comment1, comment2)

	t.Run("Load Posts for User", func(t *testing.T) {
		result, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Id(), users.Into.PostsByUserId(posts.Id(), posts.Title())).
			Join(ast.JoinLeft, posts.Table(), posts.UserId().Eq(users.Id())).
			Where(users.Id().Eq(q.Val(*user1.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.PostsByUserId, 2)
		assert.ElementsMatch(t, []string{post1.Title, post2.Title}, []string{result.PostsByUserId[0].Title, result.PostsByUserId[1].Title})
	})

	t.Run("Load Comments for Post", func(t *testing.T) {
		result, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Into.CommentsByPostId(comments.Id(), comments.Content())).
			Join(ast.JoinLeft, comments.Table(), comments.PostId().Eq(posts.Id())).
			Where(posts.Id().Eq(q.Val(*post1.Id))).
			GroupBy(posts.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.CommentsByPostId, 2)
		assert.ElementsMatch(t, []string{comment1.Content, comment2.Content}, []string{result.CommentsByPostId[0].Content, result.CommentsByPostId[1].Content})
	})

	t.Run("Load Comments for User", func(t *testing.T) {
		result, err := NewQuery[models.UsersDto](pgxPool, users.Table()).
			Select(users.Id(), users.Into.CommentsByUserId(comments.Id(), comments.Content())).
			Join(ast.JoinLeft, comments.Table(), comments.UserId().Eq(users.Id())).
			Where(users.Id().Eq(q.Val(*user1.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.CommentsByUserId, 1)
		assert.Equal(t, comment1.Content, result.CommentsByUserId[0].Content)
	})
}

func TestIntegration_ManyToMany(t *testing.T) {
	cleanupTables(t)

	user := &models.UsersDto{Username: "blogger-many", Email: "blogger-many@example.com"}
	insertUsers(t, user)
	post1 := &models.PostsDto{UserId: *user.Id, Title: "Go Programming", Content: new("Learn Go")}
	post2 := &models.PostsDto{UserId: *user.Id, Title: "SQL Optimization", Content: new("Optimize queries")}
	insertPosts(t, post1, post2)
	tag1 := &models.TagsDto{Name: "programming", Slug: "programming"}
	tag2 := &models.TagsDto{Name: "database", Slug: "database"}
	tag3 := &models.TagsDto{Name: "golang", Slug: "golang"}
	insertTags(t, tag1, tag2, tag3)

	_, err := pgxPool.Exec(context.Background(), "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)", *post1.Id, *tag1.Id, *post1.Id, *tag3.Id, *post2.Id, *tag1.Id, *post2.Id, *tag2.Id)
	require.NoError(t, err)

	t.Run("Load Tags for Post", func(t *testing.T) {
		result, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Into.Tags(tags.Id(), tags.Name())).
			Join(ast.JoinLeft, post_tags.Table(), posts.Id().Eq(post_tags.PostId())).
			Join(ast.JoinLeft, tags.Table(), post_tags.TagId().Eq(tags.Id())).
			Where(posts.Id().Eq(q.Val(*post1.Id))).
			GroupBy(posts.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.Tags, 2)
		assert.ElementsMatch(t, []string{tag1.Name, tag3.Name}, []string{result.Tags[0].Name, result.Tags[1].Name})
	})

	t.Run("Load Posts for Tag", func(t *testing.T) {
		result, err := NewQuery[models.TagsDto](pgxPool, tags.Table()).
			Select(tags.Into.Id(), tags.Into.Posts(posts.Id(), posts.Title())).
			Join(ast.JoinLeft, post_tags.Table(), tags.Id().Eq(post_tags.TagId())).
			Join(ast.JoinLeft, posts.Table(), post_tags.PostId().Eq(posts.Id())).
			Where(tags.Name().Eq(q.Val("programming"))).
			GroupBy(tags.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.Len(t, result.Posts, 2)
		assert.ElementsMatch(t, []string{post1.Title, post2.Title}, []string{result.Posts[0].Title, result.Posts[1].Title})
	})

	t.Run("Load Empty Collection", func(t *testing.T) {
		post3 := &models.PostsDto{UserId: *user.Id, Title: "Untagged Post"}
		insertPosts(t, post3)

		result, err := NewQuery[models.PostsDto](pgxPool, posts.Table()).
			Select(posts.Id(), posts.Into.Tags(tags.Id(), tags.Name())).
			Join(ast.JoinLeft, post_tags.Table(), posts.Id().Eq(post_tags.PostId())).
			Join(ast.JoinLeft, tags.Table(), post_tags.TagId().Eq(tags.Id())).
			Where(posts.Id().Eq(q.Val(*post3.Id))).
			GroupBy(posts.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.NotNil(t, result.Tags)
		assert.Empty(t, result.Tags)
	})
}

func TestIntegration_ExpressionHelpers(t *testing.T) {
	cleanupTables(t)

	user := &models.UsersDto{Username: "testuser", Email: "test@example.com"}
	insertUsers(t, user)
	insertPosts(t,
		&models.PostsDto{UserId: *user.Id, Title: "Post 1", ViewCount: new(int64(100))},
		&models.PostsDto{UserId: *user.Id, Title: "Post 2", ViewCount: new(int64(200))},
		&models.PostsDto{UserId: *user.Id, Title: "Post 3", ViewCount: new(int64(300))},
	)

	t.Run("Count expression", func(t *testing.T) {
		type countRow struct {
			models.PostsDto
			Count int64
		}
		count := q.Into(q.Count(posts.Id()),
			func(r *countRow) *int64 { return &r.Count })
		userId := q.Into(posts.UserId(),
			func(r *countRow) *uuid.UUID { return &r.UserId })
		results, err := NewQuery[countRow](pgxPool, posts.Table()).
			Select(
				userId,
				count).
			Where(posts.UserId().Eq(q.Val(*user.Id))).
			GroupBy(posts.UserId()).
			Find(context.Background())

		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.EqualValues(t, 3, results[0].Count)
	})
}

func TestContextTracking(t *testing.T) {
	t.Run("Simple select without joins", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{SelectList: []ast.NamedExpression{users.Id(), users.Username()}, From: users.Table()}
		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)

		require.NoError(t, err)
		assert.Contains(t, sql, "SELECT")
		assert.Contains(t, sql, "FROM users")
		assert.Equal(t, "users", ctx.PrimaryTable)
		assert.Nil(t, ctx.JoinedTables)
		assert.Equal(t, []string{"users"}, ctx.AllTables)
		assert.Len(t, ctx.ColumnReferences, 2)
	})

	t.Run("Select with LEFT JOIN", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{SelectList: []ast.NamedExpression{users.Username(), posts.Title()}, From: users.Table()}
		selectStmt.From = selectStmt.From.Join(ast.JoinLeft, posts.Table(), posts.UserId().Eq(users.Id()))
		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)

		require.NoError(t, err)
		assert.Contains(t, sql, "LEFT JOIN posts")
		assert.Equal(t, ast.JoinLeft, ctx.JoinedTables["posts"])
		assert.Len(t, ctx.AllTables, 2)
	})

	t.Run("Select with multiple JOINs", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{SelectList: []ast.NamedExpression{users.Username(), posts.Title(), comments.Content()}, From: users.Table()}
		selectStmt.From = selectStmt.From.Join(ast.JoinLeft, posts.Table(), posts.UserId().Eq(users.Id()))
		selectStmt.From = selectStmt.From.Join(ast.JoinLeft, comments.Table(), comments.UserId().Eq(users.Id()))
		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)

		require.NoError(t, err)
		assert.Contains(t, sql, "LEFT JOIN posts")
		assert.Contains(t, sql, "LEFT JOIN comments")
		assert.Len(t, ctx.JoinedTables, 2)
	})
}

func TestIntegration_CTEs(t *testing.T) {
	t.Run("Insert post with two tags in one query", func(t *testing.T) {
		cleanupTables(t)

		// Arrange
		user := &models.UsersDto{Username: "cte-insert-user", Email: "cte-insert@example.com"}
		insertUsers(t, user)

		post := &models.PostsDto{UserId: *user.Id, Title: "Post with tags"}
		tagRows := []*models.TagsDto{
			{Name: "First tag", Slug: "first-tag"},
			{Name: "Second tag", Slug: "second-tag"},
		}

		insertedPost := posts.As("inserted_post", posts.Id())
		insertPost := NewPostsQuery(nil).
			Insert(posts.UserId(), posts.Title()).
			Values(post).
			Returning(posts.Id()).
			Statement()

		insertedTags := tags.As("inserted_tags", tags.Id())
		insertTags := NewTagsQuery(nil).
			Insert(tags.Name(), tags.Slug()).
			Values(tagRows...).
			Returning(tags.Id()).
			Statement()

		linkSource := NewStatementQuery(insertedPost.TableAlias).
			Select(insertedPost.Id(), insertedTags.Id()).
			Join(ast.JoinInner, insertedTags.TableAlias, q.Val(true))
		insertedPostTags := post_tags.As(
			"inserted_post_tags",
			post_tags.PostId(),
			post_tags.TagId(),
		)
		insertPostTags := NewPostTagsQuery(nil).
			Insert(post_tags.PostId(), post_tags.TagId()).
			Select(linkSource).
			Returning(post_tags.PostId(), post_tags.TagId()).
			Statement()

		type insertedIDs struct {
			PostID int64
			TagID  int64
		}
		queryBuilder := NewQuery[insertedIDs](pgxPool, insertedPostTags.TableAlias).
			With(
				q.CTE(insertedPost.TableAlias, insertPost),
				q.CTE(insertedTags.TableAlias, insertTags),
				q.CTE(insertedPostTags.TableAlias, insertPostTags),
			).
			Select(
				q.Into(insertedPostTags.PostId(), func(row *insertedIDs) *int64 {
					return &row.PostID
				}),
				q.Into(insertedPostTags.TagId(), func(row *insertedIDs) *int64 {
					return &row.TagID
				}),
			).
			OrderBy(q.Asc(insertedPostTags.TagId()))

		// Act
		sql, args, err := queryBuilder.ToSql()
		require.NoError(t, err)
		results, err := queryBuilder.Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "WITH inserted_post(id) AS (INSERT INTO posts(user_id, title) VALUES ($1, $2) RETURNING posts.id), inserted_tags(id) AS (INSERT INTO tags(name, slug) VALUES ($3, $4), ($5, $6) RETURNING tags.id), inserted_post_tags(post_id, tag_id) AS (INSERT INTO post_tags(post_id, tag_id) SELECT inserted_post.id, inserted_tags.id FROM inserted_post INNER JOIN inserted_tags ON $7 RETURNING post_tags.post_id, post_tags.tag_id) SELECT inserted_post_tags.post_id, inserted_post_tags.tag_id FROM inserted_post_tags ORDER BY inserted_post_tags.tag_id ASC", sql)
		assert.Equal(t, []any{post.UserId, post.Title, tagRows[0].Name, tagRows[0].Slug, tagRows[1].Name, tagRows[1].Slug, true}, args)
		require.Len(t, results, 2)
		assert.Equal(t, results[0].PostID, results[1].PostID)
		assert.NotEqual(t, results[0].TagID, results[1].TagID)
	})

	t.Run("Select CTE preserves parameter order and executes", func(t *testing.T) {
		cleanupTables(t)

		corpUser := &models.UsersDto{Username: "corp-user", Email: "corp@example.com", IsActive: new(true)}
		otherUser := &models.UsersDto{Username: "other-user", Email: "other@example.org", IsActive: new(true)}
		insertUsers(t, corpUser, otherUser)

		type userSummary struct{ Email string }
		activeUsers := users.As("active_users", users.Id(), users.Email())
		cteBody := NewQuery[models.UsersDto](nil, users.Table()).
			Select(users.Id(), users.Email()).
			Where(users.Email().Like("%@example.com")).
			Statement()

		queryBuilder := NewQuery[userSummary](pgxPool, activeUsers.TableAlias).
			With(q.CTE(activeUsers.TableAlias, cteBody)).
			Select(q.Into(activeUsers.Email(), func(s *userSummary) *string { return &s.Email })).
			Where(activeUsers.Id().Eq(q.Val(*corpUser.Id)))

		sql, args, err := queryBuilder.ToSql()
		require.NoError(t, err)
		assert.Equal(t, "WITH active_users(id, email) AS (SELECT users.id, users.email FROM users WHERE users.email LIKE $1) SELECT active_users.email FROM active_users WHERE active_users.id = $2", sql)
		assert.Equal(t, []any{"%@example.com", *corpUser.Id}, args)

		results, err := queryBuilder.Find(context.Background())
		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.Equal(t, "corp@example.com", results[0].Email)
	})

	t.Run("Update CTE with SetTo returns updated rows", func(t *testing.T) {
		cleanupTables(t)

		user := &models.UsersDto{Username: "cte-update-user", Email: "old@example.com", FullName: new("Old Name"), IsActive: new(true)}
		insertUsers(t, user)
		newName := "Updated Name"
		updateDTO := &models.UsersDto{Id: user.Id, Email: "new@example.com", FullName: &newName}
		type updatedUser struct {
			Email    string
			FullName *string
		}

		updatedUsers := users.As("updated_users", users.Id(), users.Email(), users.FullName())
		cteBody := NewUsersQuery(nil).
			Update(q.SetTo(updateDTO, users.Email(), users.FullName())).
			Where(users.Id().Eq(q.Val(*updateDTO.Id))).
			Returning(users.Id(), users.Email(), users.FullName()).
			Statement()

		queryBuilder := NewQuery[updatedUser](pgxPool, updatedUsers.TableAlias).
			With(q.CTE(updatedUsers.TableAlias, cteBody)).
			Select(
				q.Into(updatedUsers.Email(), func(u *updatedUser) *string { return &u.Email }),
				q.Into(updatedUsers.FullName(), func(u *updatedUser) **string { return &u.FullName }),
			)

		sql, args, err := queryBuilder.ToSql()
		require.NoError(t, err)
		assert.Equal(t, "WITH updated_users(id, email, full_name) AS (UPDATE users SET email = $1, full_name = $2 WHERE users.id = $3 RETURNING users.id, users.email, users.full_name) SELECT updated_users.email, updated_users.full_name FROM updated_users", sql)
		assert.Equal(t, []any{"new@example.com", &newName, *user.Id}, args)

		results, err := queryBuilder.Find(context.Background())
		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.Equal(t, "new@example.com", results[0].Email)
		require.NotNil(t, results[0].FullName)
		assert.Equal(t, "Updated Name", *results[0].FullName)
		assert.Equal(t, "new@example.com", findUserByID(t, *user.Id).Email)
	})

	t.Run("Unknown alias column returns render error", func(t *testing.T) {
		activeUsers := users.As("active_users", users.Id())
		cteBody := NewQuery[models.UsersDto](nil, users.Table()).Select(users.Id()).Statement()
		type userSummary struct{ Email string }

		queryBuilder := NewQuery[userSummary](nil, activeUsers.TableAlias).
			With(q.CTE(activeUsers.TableAlias, cteBody)).
			Select(q.Into(activeUsers.Email(), func(s *userSummary) *string { return &s.Email }))

		sql, args, err := queryBuilder.ToSql()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown column 'email' in alias active_users")
		assert.Empty(t, args)
		assert.Contains(t, sql, "WITH active_users(id) AS")
	})

	t.Run("Alias-aware For projects CTE columns into custom row", func(t *testing.T) {
		cleanupTables(t)

		corpUser := &models.UsersDto{Username: "cte-for-user", Email: "cte-for@example.com", IsActive: new(true)}
		insertUsers(t, corpUser)

		countExpr := q.Count(users.Id()).As(ast.NewColumnAlias("count"))
		activeUsers := users.As("active_users", users.Id(), users.Email(), countExpr)
		activeUsersForRow := users.AliasFor(activeUsers.TableAlias, (*userCountRow).GetUsersDto)

		cteBody := NewStatementQuery(users.Table()).
			Select(users.Id(), users.Email(), countExpr).
			Where(users.Email().Eq(q.Val(corpUser.Email))).
			GroupBy(users.Id(), users.Email()).
			Statement()

		queryBuilder := NewQuery[userCountRow](pgxPool, activeUsers.TableAlias).
			With(q.CTE(activeUsers.TableAlias, cteBody)).
			Select(
				activeUsersForRow.Id(),
				activeUsersForRow.Email(),
				q.Into(q.Rel(activeUsers.TableAlias, countExpr), func(r *userCountRow) *int64 { return &r.Count }),
			)

		sql, args, err := queryBuilder.ToSql()
		require.NoError(t, err)
		assert.Equal(t, "WITH active_users(id, email, count) AS (SELECT users.id, users.email, count(users.id) AS count FROM users WHERE users.email = $1 GROUP BY users.id, users.email) SELECT active_users.id, active_users.email, active_users.count FROM active_users", sql)
		assert.Equal(t, []any{corpUser.Email}, args)

		results, err := queryBuilder.Find(context.Background())
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.NotNil(t, results[0].Id)
		assert.Equal(t, *corpUser.Id, *results[0].Id)
		assert.Equal(t, corpUser.Email, results[0].Email)
		assert.EqualValues(t, 1, results[0].Count)
	})

	t.Run("Alias-aware For unknown column returns render error", func(t *testing.T) {
		activeUsers := users.As("active_users", users.Id())
		activeUsersForRow := users.AliasFor(activeUsers.TableAlias, (*userCountRow).GetUsersDto)

		queryBuilder := NewQuery[userCountRow](nil, activeUsers.TableAlias).
			Select(activeUsersForRow.Email())

		sql, args, err := queryBuilder.ToSql()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown column 'email' in alias active_users")
		assert.Empty(t, args)
		assert.Contains(t, sql, "SELECT ")
	})
}
