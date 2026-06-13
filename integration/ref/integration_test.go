package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	pgxPool  *pgxpool.Pool
	testDB   *DB
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

func TestMain(m *testing.M) {
	var err error

	// Create a new pool for Docker
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Pull PostgreSQL docker image and run container
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
	databaseUrl := fmt.Sprintf("postgres://testuser:secret@%s/testdb?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	// Set expiry to 2 minutes to avoid hanging on failure
	resource.Expire(120)

	// Exponential backoff-retry to connect to database
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		pgxPool, err = pgxpool.New(context.Background(), databaseUrl)
		if err != nil {
			return err
		}
		return pgxPool.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Load schema
	if err := loadSchema(pgxPool); err != nil {
		log.Fatalf("Could not load schema: %s", err)
	}

	// Create DB wrapper
	testDB = NewDB(pgxPool)

	// Run tests
	code := m.Run()

	// Clean up
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

// Helper function to clean up tables between tests
func cleanupTables(t *testing.T) {
	t.Helper()

	tables := []string{"post_tags", "comments", "posts", "tags", "users"}
	for _, table := range tables {
		_, err := pgxPool.Exec(context.Background(), fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			t.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}
}

// TestIntegration_UserCRUD tests basic CRUD operations on Users
func TestIntegration_UserCRUD(t *testing.T) {
	cleanupTables(t)

	t.Run("Create User", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username:  "johndoe",
			Email:     "john@example.com",
			FullName:  new("John Doe"),
			IsActive:  new(true),
			CreatedAt: new(time.Now()),
			UpdatedAt: new(time.Now()),
		}

		// Act
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		//err := testDB.Users().Insert(context.Background(), user)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, user.Id, "expected user ID to be set after insert")
		assert.NotEmpty(t, *user.Id, "expected user ID to be set after insert")
	})

	t.Run("Create Users", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username:  "johndoe1",
			Email:     "john1@example.com",
			FullName:  new("John Doe"),
			IsActive:  new(true),
			CreatedAt: new(time.Now()),
			UpdatedAt: new(time.Now()),
		}

		user2 := &models.UsersDto{
			Username:  "johndoe2",
			Email:     "john2@example.com",
			FullName:  new("John Doe 2"),
			IsActive:  new(true),
			CreatedAt: new(time.Now()),
			UpdatedAt: new(time.Now()),
		}

		// Act
		err := query.UsersDtoInsertMany(testDB.pool, user, user2).Exec(context.Background())
		//err := testDB.Users().Insert(context.Background(), user)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, user.Id, "expected user ID to be set after insert")
		assert.NotEmpty(t, *user.Id, "expected user ID to be set after insert")
		require.NotNil(t, user2.Id, "expected user ID to be set after insert")
		assert.NotEmpty(t, *user2.Id, "expected user ID to be set after insert")
	})

	t.Run("Find User by ID", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "janedoe",
			Email:    "jane@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		// Act
		//found, err := testDB.Users().Where(
		//	UsersClause().Id.Eq(*user.Id),
		//).FindOne(context.Background())

		found, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(
				users.Id().Eq(query.Val(*user.Id))).
			FindOne(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Equal(t, user.Username, found.Username)
		assert.Equal(t, user.Email, found.Email)
	})

	t.Run("Update User", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "updateme",
			Email:    "update@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		user.Email = "updated@example.com"
		user.FullName = new("Updated Name")

		// Act
		affected, _, err := query.UpdateUser(testDB.pool, user).Exec(context.Background())

		//err = testDB.Users().Update(context.Background(), user)
		require.NoError(t, err)
		found, err := testDB.Users().
			Where(UsersClause().Id.Eq(*user.Id)).
			FindOne(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Equal(t, int64(1), affected)
		assert.Equal(t, user.Email, found.Email)
		assert.Equal(t, user.FullName, found.FullName)
	})

	t.Run("Update Users", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "updateme1",
			Email:    "update@example.com",
		}
		user2 := &models.UsersDto{
			Username: "updateme2",
			Email:    "update2@example.com",
		}
		err := query.UsersDtoInsertMany(testDB.pool, user, user2).Exec(context.Background())
		require.NoError(t, err)

		user.Email = "updated1@example.com"
		user.FullName = new("Updated Name")
		user2.Email = "updated2@example.com"
		user2.FullName = new("Updated Name 2")

		// Act
		affected, _, err := query.UpdateUsers(testDB.pool, user, user2).Exec(context.Background())

		//err = testDB.Users().Update(context.Background(), user)
		require.NoError(t, err)
		found, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(
				users.Id().Eq(query.Val(*user.Id))).
			FindOne(context.Background())
		require.NoError(t, err)
		found2, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(
				users.Id().Eq(query.Val(*user2.Id))).
			FindOne(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Equal(t, int64(2), affected)
		assert.Equal(t, user.Email, found.Email)
		assert.Equal(t, user.FullName, found.FullName)
		assert.Equal(t, user2.Email, found2.Email)
		assert.Equal(t, user2.FullName, found2.FullName)
	})

	t.Run("Delete User", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "deleteme",
			Email:    "delete@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		// Act
		deleted, details, err := query.DeleteUser(testDB.pool, *user.Id).Exec(context.Background())
		if err != nil {
			return
		}

		// Assert
		require.NoError(t, err)
		assert.EqualValues(t, 1, deleted)
		assert.Nil(t, details)

		_, err = models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.Id()).
			Where(users.Id().Eq(query.Val(*user.Id))).
			FindOne(context.Background())
		require.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("Create User with Conflict", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username:  "conflictedjohndoe",
			Email:     "john@conflict.example.com",
			FullName:  new("John Doe"),
			IsActive:  new(true),
			CreatedAt: new(time.Now()),
			UpdatedAt: new(time.Now()),
		}

		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		user.Email = "updated@conflict.example.com"
		conflict := models.NewUsersQuery(testDB.pool).
			Insert(
				users.Email(),
				users.Username(),
				users.FullName(),
				users.IsActive(),
			).
			OnConflict(users.Username()).Do(ast.Update(users.Email())).
			Returning(users.Into.Id()).
			Values(user)
		err = conflict.Exec(context.Background())
		require.NoError(t, err)

		// Act
		found, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.AllColumns()...).
			Where(
				users.Id().Eq(query.Val(*user.Id))).
			FindOne(context.Background())

		require.NoError(t, err)
		assert.Equal(t, user.Email, found.Email)
	})

	t.Run("Update user returning full name", func(t *testing.T) {
		// Arrange
		fullName := "Update Me"
		user := &models.UsersDto{
			Username: "updatemeandreturn",
			Email:    "updateme@returning.example.com",
			FullName: &fullName,
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		user.Email = "updated@returning.example.com"
		user.Username = "updatedname"

		// Act
		affected, details, err := models.NewUsersQuery(testDB.pool).
			Update(
				query.Set(users.Email()).To(query.Val(user.Email)),
				query.Set(users.Username()).To(query.Val(user.Username)),
			).
			Where(users.Id().Eq(query.Val(*user.Id))).
			Returning(users.Into.FullName()).
			Exec(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Equal(t, int64(1), affected)
		assert.Equal(t, fullName, *details[0].FullName)
	})

	t.Run("Delete User returning id", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "deleteme",
			Email:    "delete@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)

		// Act
		deleted, details, err := models.NewUsersQuery(testDB.pool).
			Delete().
			Where(users.Id().Eq(query.Val(*user.Id))).
			Returning(users.Into.Id()).
			Exec(context.Background())
		require.NoError(t, err)

		// Assert
		require.NoError(t, err)
		assert.EqualValues(t, 1, deleted)
		assert.Equal(t, *user.Id, *details[0].Id)

		_, err = models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(users.Into.Id()).
			Where(users.Id().Eq(query.Val(*user.Id))).FindOne(context.Background())
		require.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})
}

// TestIntegration_QueryBuilder tests query builder features
func TestIntegration_QueryBuilder(t *testing.T) {
	cleanupTables(t)

	// Insert test data
	users := []*UsersDto{
		{Username: "alice", Email: "alice@example.com", FullName: new("Alice Smith"), IsActive: new(true)},
		{Username: "bob", Email: "bob@example.com", FullName: new("Bob Jones"), IsActive: new(true)},
		{Username: "charlie", Email: "charlie@example.com", FullName: new("Charlie Brown"), IsActive: new(false)},
		{Username: "david", Email: "david@example.com", FullName: new("David Wilson"), IsActive: new(true)},
	}

	for _, user := range users {
		if err := testDB.Users().Insert(context.Background(), user); err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	t.Run("Where Equals", func(t *testing.T) {
		// Act
		results, err := testDB.Users().Where(
			UsersClause().Username.Eq("alice"),
		).Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, "alice", results[0].Username)
	})

	t.Run("Where LIKE", func(t *testing.T) {
		// Act
		results, err := testDB.Users().Where(
			UsersClause().Username.Like("a%"),
		).Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 1)
	})

	t.Run("Where IN", func(t *testing.T) {
		// Act
		results, err := testDB.Users().Where(
			UsersClause().Username.In("alice", "bob"),
		).Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("Where Boolean", func(t *testing.T) {
		// Act
		results, err := testDB.Users().Where(
			UsersClause().IsActive.Eq(true),
		).Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 3)
	})

	t.Run("Order By", func(t *testing.T) {
		// Arrange
		expected := []string{"alice", "bob", "charlie", "david"}

		// Act
		results, err := testDB.Users().
			OrderByUsername(ASC).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 4)
		for i, result := range results {
			assert.Equal(t, expected[i], result.Username)
		}
	})

	t.Run("Limit and Offset", func(t *testing.T) {
		// Act
		results, err := testDB.Users().
			OrderByUsername(ASC).
			Limit(2).
			Offset(1).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, "bob", results[0].Username)
		assert.Equal(t, "charlie", results[1].Username)
	})

	t.Run("Count", func(t *testing.T) {
		// Act
		count, err := testDB.Users().Where(
			UsersClause().IsActive.Eq(true),
		).Count(context.Background())

		// Assert
		require.NoError(t, err)
		assert.EqualValues(t, 3, count)
	})
}

func TestIntegration_RelationshipProjection(t *testing.T) {
	cleanupTables(t)

	t.Run("Load posts into user DTO", func(t *testing.T) {
		user := &models.UsersDto{
			Username: "projectionuser",
			Email:    "projection@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, user.Id)

		post1 := &models.PostsDto{UserId: *user.Id, Title: "Projection Post 1"}
		post2 := &models.PostsDto{UserId: *user.Id, Title: "Projection Post 2"}
		err = query.InsertPost(testDB.pool, post1).Exec(context.Background())
		require.NoError(t, err)
		err = query.InsertPost(testDB.pool, post2).Exec(context.Background())
		require.NoError(t, err)

		result, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(
				users.Into.Id(),
				users.Into.Posts(posts.Id(), posts.Title()),
			).
			Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id())).
			Where(users.Id().Eq(query.Val(*user.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.NotNil(t, result.Posts)
		require.Len(t, result.Posts, 2)
		assert.ElementsMatch(t, []string{"Projection Post 1", "Projection Post 2"}, []string{
			result.Posts[0].Title,
			result.Posts[1].Title,
		})
	})

	t.Run("Empty relationship becomes empty slice", func(t *testing.T) {
		user := &models.UsersDto{
			Username: "emptyprojectionuser",
			Email:    "empty-projection@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, user.Id)

		result, err := models.NewQuery[models.UsersDto](testDB.pool, users.Table()).
			Select(
				users.Into.Id(),
				users.Into.Posts(posts.Id(), posts.Title()),
			).
			Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id())).
			Where(users.Id().Eq(query.Val(*user.Id))).
			GroupBy(users.Id()).
			FindOne(context.Background())

		require.NoError(t, err)
		require.NotNil(t, result.Posts)
		assert.Empty(t, result.Posts)
	})
}

// TestIntegration_Joins tests foreign key joins
func TestIntegration_Joins(t *testing.T) {
	cleanupTables(t)

	// Insert test data
	user := &UsersDto{
		Username: "blogger",
		Email:    "blogger@example.com",
	}
	if err := testDB.Users().Insert(context.Background(), user); err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	posts := []*PostsDto{
		{UserId: *user.Id, Title: "First Post", Content: new("Content 1"), Status: new("published")},
		{UserId: *user.Id, Title: "Second Post", Content: new("Content 2"), Status: new("draft")},
	}

	for _, post := range posts {
		if err := testDB.Posts().Insert(context.Background(), post); err != nil {
			t.Fatalf("Failed to insert post: %v", err)
		}
	}

	t.Run("Join Users", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().
			JoinUsers().
			WhereUserIdEq(*user.Id).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
		for _, post := range results {
			assert.Equal(t, *user.Id, post.UserId)
		}
	})

	t.Run("Filter by Related Table", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().
			WhereStatusEq("published").
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 1)
		assert.Equal(t, posts[0].Title, results[0].Title)
	})

	t.Run("Join Returns User Data", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().
			JoinUsers().
			WhereUserIdEq(*user.Id).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)

		// Verify User field is populated
		for _, post := range results {
			require.NotNil(t, post.User)
			require.NotNil(t, post.User.Id)
			assert.Equal(t, *user.Id, *post.User.Id)
			assert.Equal(t, user.Username, post.User.Username)
			assert.Equal(t, user.Email, post.User.Email)
		}
	})

	t.Run("No Join Means No User Data", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().
			WhereUserIdEq(*user.Id).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
		for _, post := range results {
			assert.Nil(t, post.User)
		}
	})
}

// TestIntegration_Transactions tests transaction support
func TestIntegration_Transactions(t *testing.T) {
	cleanupTables(t)

	t.Run("Successful Transaction", func(t *testing.T) {
		// Arrange
		var user *UsersDto
		var post *PostsDto

		// Act
		err := testDB.Transaction(context.Background(), func(tx *Tx) error {
			// Insert user
			user = &UsersDto{
				Username: "txuser",
				Email:    "tx@example.com",
			}
			if err := tx.Users().Insert(context.Background(), user); err != nil {
				return err
			}

			// Insert post
			post = &PostsDto{
				UserId: *user.Id,
				Title:  "Transaction Post",
			}
			if err := tx.Posts().Insert(context.Background(), post); err != nil {
				return err
			}

			return nil
		})

		// Assert
		require.NoError(t, err)

		users, err := testDB.Users().Where(
			UsersClause().Username.Eq("txuser"),
		).Find(context.Background())
		require.NoError(t, err)
		assert.Len(t, users, 1)

		posts, err := testDB.Posts().WhereTitleEq("Transaction Post").Find(context.Background())
		require.NoError(t, err)
		assert.Len(t, posts, 1)
		assert.Equal(t, post.UserId, posts[0].UserId)
	})

	t.Run("Failed Transaction Rollback", func(t *testing.T) {
		// Act
		err := testDB.Transaction(context.Background(), func(tx *Tx) error {
			// Insert user
			user := &UsersDto{
				Username: "rollbackuser",
				Email:    "rollback@example.com",
			}
			if err := tx.Users().Insert(context.Background(), user); err != nil {
				return err
			}

			// Intentionally return error to trigger rollback
			return assert.AnError
		})

		// Assert
		require.Error(t, err)

		users, err := testDB.Users().Where(
			UsersClause().Username.Eq("rollbackuser"),
		).Find(context.Background())
		require.NoError(t, err)
		assert.Len(t, users, 0)
	})
}

// TestIntegration_NullableFields tests nullable field handling
func TestIntegration_NullableFields(t *testing.T) {
	cleanupTables(t)

	t.Run("Insert with NULL values", func(t *testing.T) {
		// Arrange
		user := &UsersDto{
			Username: "nulltest",
			Email:    "null@example.com",
			FullName: nil, // NULL value
		}

		// Act
		err := testDB.Users().Insert(context.Background(), user)

		// Assert
		require.NoError(t, err)

		found, err := testDB.Users().Where(
			UsersClause().Id.Eq(uuid.MustParse(*user.Id)),
		).FindOne(context.Background())
		require.NoError(t, err)
		require.NotNil(t, found)
		assert.Nil(t, found.FullName)
	})

	t.Run("Update to NULL", func(t *testing.T) {
		// Arrange
		user := &UsersDto{
			Username: "nullupdate",
			Email:    "nullupdate@example.com",
			FullName: new("Initial Name"),
		}
		err := testDB.Users().Insert(context.Background(), user)
		require.NoError(t, err)

		user.FullName = nil

		// Act
		err = testDB.Users().Update(context.Background(), user)

		// Assert
		require.NoError(t, err)

		found, err := testDB.Users().Where(
			UsersClause().Id.Eq(uuid.MustParse(*user.Id)),
		).FindOne(context.Background())
		require.NoError(t, err)
		require.NotNil(t, found)
		assert.Nil(t, found.FullName)
	})
}

// TestIntegration_ComplexQueries tests more complex query scenarios
func TestIntegration_ComplexQueries(t *testing.T) {
	cleanupTables(t)

	// Setup test data
	user1 := &UsersDto{Username: "user1", Email: "user1@example.com"}
	user2 := &UsersDto{Username: "user2", Email: "user2@example.com"}

	testDB.Users().Insert(context.Background(), user1)
	testDB.Users().Insert(context.Background(), user2)

	// Create posts for both users
	posts := []*PostsDto{
		{UserId: *user1.Id, Title: "User1 Post 1", ViewCount: new(int64(100))},
		{UserId: *user1.Id, Title: "User1 Post 2", ViewCount: new(int64(200))},
		{UserId: *user2.Id, Title: "User2 Post 1", ViewCount: new(int64(50))},
	}

	for _, post := range posts {
		testDB.Posts().Insert(context.Background(), post)
	}

	t.Run("Greater Than Query", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().WhereViewCountGt(75).Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("Multiple Conditions", func(t *testing.T) {
		// Act
		results, err := testDB.Posts().
			WhereUserIdEq(*user1.Id).
			And().
			WhereViewCountGte(100).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("UpdateFields", func(t *testing.T) {
		// Act
		affected, err := testDB.Posts().
			WhereUserIdEq(*user1.Id).
			UpdateFields(context.Background(), map[*FieldRef]interface{}{
				PostsTable.ViewCount(): int64(999),
			})

		// Assert
		require.NoError(t, err)
		assert.EqualValues(t, 2, affected)

		results, err := testDB.Posts().WhereUserIdEq(*user1.Id).Find(context.Background())
		require.NoError(t, err)
		for _, post := range results {
			require.NotNil(t, post.ViewCount)
			assert.EqualValues(t, 999, *post.ViewCount)
		}
	})
}

// TestIntegration_ReverseRelationships tests one-to-many reverse relationships
func TestIntegration_ReverseRelationships(t *testing.T) {
	cleanupTables(t)

	// Setup test data
	user1 := &models.UsersDto{Username: "author1", Email: "author1@example.com"}
	user2 := &UsersDto{Username: "author2", Email: "author2@example.com"}
	query.UsersDtoInsertOne(testDB.pool, user1).Exec(context.Background())
	//testDB.Users().Insert(context.Background(), user1)
	testDB.Users().Insert(context.Background(), user2)

	// Create posts for user1
	post1 := &PostsDto{UserId: user1.Id.String(), Title: "Post 1", Content: new("Content 1")}
	post2 := &PostsDto{UserId: user1.Id.String(), Title: "Post 2", Content: new("Content 2")}
	post3 := &PostsDto{UserId: *user2.Id, Title: "Post 3", Content: new("Content 3")}
	testDB.Posts().Insert(context.Background(), post1)
	testDB.Posts().Insert(context.Background(), post2)
	testDB.Posts().Insert(context.Background(), post3)

	// Create comments
	comment1 := &CommentsDto{PostId: *post1.Id, UserId: user1.Id.String(), Content: "Comment 1"}
	comment2 := &CommentsDto{PostId: *post1.Id, UserId: *user2.Id, Content: "Comment 2"}
	testDB.Comments().Insert(context.Background(), comment1)
	testDB.Comments().Insert(context.Background(), comment2)

	t.Run("Load Posts for User", func(t *testing.T) {
		// Arrange
		user, err := testDB.Users().Where(
			UsersClause().Id.Eq(*user1.Id),
		).FindOne(context.Background())
		require.NoError(t, err)
		expectedTitles := []string{post1.Title, post2.Title}

		// Act
		err = user.LoadPosts(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		assert.Len(t, user.Posts, 2)
		assert.ElementsMatch(t, expectedTitles, func() []string {
			ret := make([]string, 0, len(user.Posts))
			for _, post := range user.Posts {
				ret = append(ret, post.Title)
			}
			return ret
		}())
	})

	t.Run("Load Comments for Post", func(t *testing.T) {
		// Arrange
		post, err := testDB.Posts().WhereIdEq(*post1.Id).FindOne(context.Background())
		require.NoError(t, err)
		expectedComments := []string{comment1.Content, comment2.Content}

		// Act
		err = post.LoadComments(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		assert.Len(t, post.Comments, 2)
		assert.ElementsMatch(t, expectedComments, func() []string {
			ret := make([]string, 0, len(post.Comments))
			for _, comment := range post.Comments {
				ret = append(ret, comment.Content)
			}
			return ret
		}())
	})

	t.Run("Load Comments for User", func(t *testing.T) {
		// Arrange
		user, err := testDB.Users().Where(
			UsersClause().Id.Eq(*user1.Id),
		).FindOne(context.Background())
		require.NoError(t, err)

		// Act
		err = user.LoadComments(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		assert.Len(t, user.Comments, 1)
		assert.Equal(t, comment1.Content, user.Comments[0].Content)
	})

}

// TestIntegration_ManyToMany tests many-to-many relationships through junction tables
func TestIntegration_ManyToMany(t *testing.T) {
	cleanupTables(t)

	// Setup test data
	user := &UsersDto{Username: "blogger", Email: "blogger@example.com"}
	testDB.Users().Insert(context.Background(), user)

	post1 := &PostsDto{UserId: *user.Id, Title: "Go Programming", Content: new("Learn Go")}
	post2 := &PostsDto{UserId: *user.Id, Title: "SQL Optimization", Content: new("Optimize queries")}
	testDB.Posts().Insert(context.Background(), post1)
	testDB.Posts().Insert(context.Background(), post2)

	tag1 := &TagsDto{Name: "programming", Slug: "programming"}
	tag2 := &TagsDto{Name: "database", Slug: "database"}
	tag3 := &TagsDto{Name: "golang", Slug: "golang"}
	testDB.Tags().Insert(context.Background(), tag1)
	testDB.Tags().Insert(context.Background(), tag2)
	testDB.Tags().Insert(context.Background(), tag3)

	// Create many-to-many associations
	// post1 -> programming, golang
	// post2 -> programming, database
	pgxPool.Exec(context.Background(), "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)", *post1.Id, *tag1.Id)
	pgxPool.Exec(context.Background(), "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)", *post1.Id, *tag3.Id)
	pgxPool.Exec(context.Background(), "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)", *post2.Id, *tag1.Id)
	pgxPool.Exec(context.Background(), "INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)", *post2.Id, *tag2.Id)

	t.Run("Load Tags for Post", func(t *testing.T) {
		// Arrange
		post, err := testDB.Posts().WhereIdEq(*post1.Id).FindOne(context.Background())
		require.NoError(t, err)

		// Act
		err = post.LoadTags(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		assert.Len(t, post.Tags, 2)
		assert.ElementsMatch(t, []string{tag1.Name, tag3.Name}, func() []string {
			ret := make([]string, 0, len(post.Tags))
			for _, tag := range post.Tags {
				ret = append(ret, tag.Name)
			}
			return ret
		}())
	})

	t.Run("Load Posts for Tag", func(t *testing.T) {
		// Arrange
		tag, err := testDB.Tags().WhereNameEq("programming").FindOne(context.Background())
		require.NoError(t, err)

		// Act
		err = tag.LoadPosts(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		assert.Len(t, tag.Posts, 2)
		assert.ElementsMatch(t, []string{post1.Title, post2.Title}, func() []string {
			ret := make([]string, 0, len(tag.Posts))
			for _, post := range tag.Posts {
				ret = append(ret, post.Title)
			}
			return ret
		}())
	})

	t.Run("Load Empty Collection", func(t *testing.T) {
		// Arrange
		post3 := &PostsDto{UserId: *user.Id, Title: "Untagged Post"}
		testDB.Posts().Insert(context.Background(), post3)

		post, err := testDB.Posts().WhereIdEq(*post3.Id).FindOne(context.Background())
		require.NoError(t, err)

		// Act
		err = post.LoadTags(context.Background(), testDB)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, post.Tags)
		assert.Len(t, post.Tags, 0)
	})
}

// TestIntegration_ExpressionFromString tests custom SQL expressions
func TestIntegration_ExpressionFromString(t *testing.T) {
	cleanupTables(t)

	// Setup test data
	user := &UsersDto{Username: "testuser", Email: "test@example.com"}
	testDB.Users().Insert(context.Background(), user)

	post1 := &PostsDto{UserId: *user.Id, Title: "Post 1", ViewCount: new(int64(100))}
	post2 := &PostsDto{UserId: *user.Id, Title: "Post 2", ViewCount: new(int64(200))}
	post3 := &PostsDto{UserId: *user.Id, Title: "Post 3", ViewCount: new(int64(300))}
	testDB.Posts().Insert(context.Background(), post1)
	testDB.Posts().Insert(context.Background(), post2)
	testDB.Posts().Insert(context.Background(), post3)

	t.Run("Custom COUNT expression", func(t *testing.T) {
		// Arrange
		alias := "total_count"
		expression := "COUNT(*)"
		expectedSQL := "COUNT(*) AS total_count"

		// Act
		countExpr := ExpressionFromString(expression, alias)

		// Assert
		assert.Equal(t, expression, countExpr.Expression)
		assert.Equal(t, alias, countExpr.Alias)
		assert.Empty(t, countExpr.Table)
		assert.Equal(t, expectedSQL, countExpr.String())
	})

	t.Run("Custom AVG expression", func(t *testing.T) {
		// Arrange
		expectedSQL := "AVG(view_count) AS avg_views"

		// Act
		avgExpr := ExpressionFromString("AVG(view_count)", "avg_views")

		// Assert
		assert.Equal(t, expectedSQL, avgExpr.String())
	})

}

// TestContextTracking tests that QueryContext properly tracks joins during SQL generation
func TestContextTracking(t *testing.T) {
	t.Run("Simple select without joins", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{
			SelectList: []ast.NamedExpression{
				users.Id(),
				users.Username(),
			},
			From: &ast.TableSource{Table: "users"},
		}

		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)
		require.NoError(t, err)

		assert.Contains(t, sql, "SELECT")
		assert.Contains(t, sql, "FROM users")
		assert.Equal(t, "users", ctx.PrimaryTable)
		assert.Nil(t, ctx.JoinedTables, "expected no joined tables")
		assert.Equal(t, []string{"users"}, ctx.AllTables)
		assert.Len(t, ctx.ColumnReferences, 2, "expected 2 column references")
	})

	t.Run("Select with LEFT JOIN", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{
			SelectList: []ast.NamedExpression{
				users.Username(),
				posts.Title(),
			},
			From: &ast.TableSource{Table: "users"},
		}
		selectStmt.From.Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id()))

		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)
		require.NoError(t, err)

		assert.Contains(t, sql, "SELECT")
		assert.Contains(t, sql, "FROM users")
		assert.Contains(t, sql, "LEFT JOIN posts")
		assert.Contains(t, sql, "ON")
		assert.Equal(t, "users", ctx.PrimaryTable)
		assert.NotNil(t, ctx.JoinedTables, "expected joined tables map")
		assert.Len(t, ctx.JoinedTables, 1, "expected 1 joined table")
		assert.Contains(t, ctx.JoinedTables, "posts", "expected posts in joined tables")
		assert.Equal(t, ast.JoinLeft, ctx.JoinedTables["posts"])
		assert.Len(t, ctx.AllTables, 2, "expected 2 tables total")
	})

	t.Run("Select with multiple JOINs", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{
			SelectList: []ast.NamedExpression{
				users.Username(),
				posts.Title(),
				comments.Content(),
			},
			From: &ast.TableSource{Table: "users"},
		}
		selectStmt.From.Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id()))
		selectStmt.From.Join(ast.JoinLeft, "comments", comments.UserId().Eq(users.Id()))

		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)
		require.NoError(t, err)

		assert.Contains(t, sql, "SELECT")
		assert.Contains(t, sql, "FROM users")
		assert.Contains(t, sql, "LEFT JOIN posts")
		assert.Contains(t, sql, "LEFT JOIN comments")
		assert.Equal(t, "users", ctx.PrimaryTable)
		assert.NotNil(t, ctx.JoinedTables, "expected joined tables map")
		assert.Len(t, ctx.JoinedTables, 2, "expected 2 joined tables")
		assert.Contains(t, ctx.JoinedTables, "posts")
		assert.Contains(t, ctx.JoinedTables, "comments")
		assert.Len(t, ctx.AllTables, 3, "expected 3 tables total")
	})

	t.Run("Verify SQL includes proper table prefixes based on context", func(t *testing.T) {
		selectStmt := &ast.SelectStatement{
			SelectList: []ast.NamedExpression{
				users.Username(),
				posts.Title(),
			},
			From: &ast.TableSource{Table: "users"},
		}
		selectStmt.From.Join(ast.JoinLeft, "posts", posts.UserId().Eq(users.Id()))

		ctx := &ast.QueryContext{PrimaryTable: "users"}
		sql, err := ast.RenderWithContext(selectStmt, &[]any{}, ctx)
		require.NoError(t, err)

		// The SELECT list should have proper prefixes (both primary and joined)
		assert.Contains(t, sql, "users.username", "expected users.username with table prefix")
		assert.Contains(t, sql, "posts.title", "expected posts.title with table prefix (joined table)")

		// Verify context tracking
		assert.Equal(t, "users", ctx.PrimaryTable)
		assert.Len(t, ctx.JoinedTables, 1)
		assert.Contains(t, ctx.JoinedTables, "posts")
	})

	t.Run("Execute query with JOIN and verify joined data", func(t *testing.T) {
		cleanupTables(t)

		// Arrange - Create user
		user := &models.UsersDto{
			Username: "joinuser",
			Email:    "join@example.com",
			IsActive: new(true),
		}
		err := models.NewUsersQuery(testDB.pool).
			Insert(
				users.Username(),
				users.Email(),
				users.FullName(),
				users.IsActive(),
			).
			Returning(users.Into.Id()).
			Values(user).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, user.Id)
		t.Logf("User ID after insert: %v", user.Id)

		// Create post
		post := &models.PostsDto{
			UserId:  *user.Id,
			Title:   "Test Post",
			Content: new("Test Content"),
		}
		err = models.NewDTOQuery(testDB.pool, models.PostsDtos{}).
			Insert(
				posts.UserId(),
				posts.Title(),
				posts.Content(),
			).
			Returning(posts.Into.Id()).
			Values(post).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, post.Id)

		type postWithUser struct {
			PostID   int64
			Title    string
			Username string
			Email    string
		}

		// Act - Select explicit projections from posts and the joined users table.
		results, err := models.NewQuery[postWithUser](testDB.pool, posts.Table()).
			Select(
				query.Into(posts.Id(), func(p *postWithUser) *int64 { return &p.PostID }),
				query.Into(posts.Title(), func(p *postWithUser) *string { return &p.Title }),
				query.Into(users.Username(), func(p *postWithUser) *string { return &p.Username }),
				query.Into(users.Email(), func(p *postWithUser) *string { return &p.Email }),
			).
			Join(ast.JoinLeft, "users", posts.UserId().Eq(users.Id())).
			Where(posts.Id().Eq(query.Val(*post.Id))).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		require.Len(t, results, 1, "expected 1 result")

		postResult := results[0]
		assert.Equal(t, *post.Id, postResult.PostID)
		assert.Equal(t, "Test Post", postResult.Title)
		assert.Equal(t, "joinuser", postResult.Username)
		assert.Equal(t, "join@example.com", postResult.Email)
	})

	t.Run("Execute query with multiple JOINs", func(t *testing.T) {
		cleanupTables(t)

		// Arrange - Create user using new API
		user := &models.UsersDto{
			Username: "multijoinuser",
			Email:    "multi@example.com",
			IsActive: new(true),
		}
		err := models.NewUsersQuery(testDB.pool).
			Insert(
				users.Username(),
				users.Email(),
				users.FullName(),
				users.IsActive(),
			).
			Returning(users.Into.Id()).
			Values(user).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, user.Id)

		post := &models.PostsDto{
			UserId:  *user.Id,
			Title:   "Multi Join Post",
			Content: new("Multi Join Content"),
		}
		err = models.NewPostsQuery(testDB.pool).
			Insert(
				posts.UserId(),
				posts.Title(),
				posts.Content(),
			).
			Returning(posts.Into.Id()).
			Values(post).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, post.Id)

		// Create comment using builder API
		comment := &models.CommentsDto{
			UserId:  *user.Id,
			PostId:  *post.Id,
			Content: "Test Comment",
		}
		err = models.NewCommentsQuery(testDB.pool).
			Insert(
				comments.PostId(),
				comments.UserId(),
				comments.Content(),
			).
			Returning(comments.Into.Id()).
			Values(comment).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, comment.Id)

		type commentWithPostAndUser struct {
			CommentID  int64
			Content    string
			PostTitle  string
			PostUserID uuid.UUID
			Username   string
			Email      string
		}

		// Act - Select explicit projections from comments and both joined tables.
		results, err := models.NewQuery[commentWithPostAndUser](testDB.pool, comments.Table()).
			Select(
				query.Into(comments.Id(), func(c *commentWithPostAndUser) *int64 { return &c.CommentID }),
				query.Into(comments.Content(), func(c *commentWithPostAndUser) *string { return &c.Content }),
				query.Into(posts.Title(), func(c *commentWithPostAndUser) *string { return &c.PostTitle }),
				query.Into(posts.UserId(), func(c *commentWithPostAndUser) *uuid.UUID { return &c.PostUserID }),
				query.Into(users.Username(), func(c *commentWithPostAndUser) *string { return &c.Username }),
				query.Into(users.Email(), func(c *commentWithPostAndUser) *string { return &c.Email }),
			).
			Join(ast.JoinLeft, "posts", comments.PostId().Eq(posts.Id())).
			Join(ast.JoinLeft, "users", comments.UserId().Eq(users.Id())).
			Where(comments.Id().Eq(query.Val(*comment.Id))).
			Find(context.Background())

		// Assert
		require.NoError(t, err)
		require.Len(t, results, 1, "expected 1 result")

		commentResult := results[0]
		assert.Equal(t, *comment.Id, commentResult.CommentID)
		assert.Equal(t, "Test Comment", commentResult.Content)
		assert.Equal(t, "Multi Join Post", commentResult.PostTitle)
		assert.Equal(t, *user.Id, commentResult.PostUserID)
		assert.Equal(t, "multijoinuser", commentResult.Username)
		assert.Equal(t, "multi@example.com", commentResult.Email)
	})

	t.Run("Verify SQL generation with actual query", func(t *testing.T) {
		// Arrange
		user := &models.UsersDto{
			Username: "sqlverify",
			Email:    "sqlverify@example.com",
		}
		err := query.UsersDtoInsertOne(testDB.pool, user).Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, user.Id)

		post := &models.PostsDto{
			UserId: *user.Id,
			Title:  "SQL Verify Post",
		}
		err = models.NewDTOQuery(testDB.pool, models.PostsDtos{}).
			Insert(
				posts.UserId(),
				posts.Title(),
			).
			Returning(posts.Into.Id()).
			Values(post).
			Exec(context.Background())
		require.NoError(t, err)
		require.NotNil(t, post.Id)

		// Act - Build query and get SQL
		type sqlVerifyRow struct {
			Title    string
			Username string
		}

		queryBuilder := models.NewQuery[sqlVerifyRow](testDB.pool, posts.Table()).
			Select(
				query.Into(posts.Title(), func(r *sqlVerifyRow) *string { return &r.Title }),
				query.Into(users.Username(), func(r *sqlVerifyRow) *string { return &r.Username }),
			).
			Join(ast.JoinLeft, "users", posts.UserId().Eq(users.Id())).
			Where(posts.Id().Eq(query.Val(*post.Id)))

		sql, err := queryBuilder.ToSql()

		// Assert
		require.NoError(t, err)
		assert.Contains(t, sql, "SELECT")
		assert.Contains(t, sql, "FROM posts")
		assert.Contains(t, sql, "LEFT JOIN users")
		assert.Contains(t, sql, "posts.title", "expected posts.title with table prefix")
		assert.Contains(t, sql, "username", "expected username without table prefix (primary table of users)")
		assert.Contains(t, sql, "ON posts.user_id = users.id")

		// Verify it actually executes
		results, err := queryBuilder.Find(context.Background())
		require.NoError(t, err)
		require.Len(t, results, 1)
		assert.Equal(t, "SQL Verify Post", results[0].Title)
		assert.Equal(t, "sqlverify", results[0].Username)
	})
}
