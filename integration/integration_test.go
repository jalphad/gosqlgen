package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	models2 "github.com/jalphad/gosqlgen/integration/models"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	pgxPool  *pgxpool.Pool
	testDB   *models2.DB
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
	testDB = models2.NewDB(pgxPool)

	// Run tests
	code := m.Run()

	// Clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func loadSchema(pool *pgxpool.Pool) error {
	schema, err := os.ReadFile("cmd/schema.sql")
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
		user := &models2.Users{
			Username:  "johndoe",
			Email:     "john@example.com",
			FullName:  strPtr("John Doe"),
			IsActive:  boolPtr(true),
			CreatedAt: timePtr(time.Now()),
			UpdatedAt: timePtr(time.Now()),
		}

		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		if user.Id == nil || *user.Id == "" {
			t.Error("Expected user ID to be set after insert")
		}

		t.Logf("Created user with ID: %s", *user.Id)
	})

	t.Run("Find User by ID", func(t *testing.T) {
		// Insert a user first
		user := &models2.Users{
			Username: "janedoe",
			Email:    "jane@example.com",
		}
		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		// Find the user
		found, err := testDB.Users().WhereIdEq(*user.Id).FindOne(context.Background())
		if err != nil {
			t.Fatalf("Failed to find user: %v", err)
		}

		if found.Username != "janedoe" {
			t.Errorf("Expected username 'janedoe', got '%s'", found.Username)
		}
		if found.Email != "jane@example.com" {
			t.Errorf("Expected email 'jane@example.com', got '%s'", found.Email)
		}
	})

	t.Run("Update User", func(t *testing.T) {
		// Insert a user
		user := &models2.Users{
			Username: "updateme",
			Email:    "update@example.com",
		}
		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		// Update the user
		user.Email = "updated@example.com"
		user.FullName = strPtr("Updated Name")

		err = testDB.Users().Update(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to update user: %v", err)
		}

		// Verify the update
		found, err := testDB.Users().WhereIdEq(*user.Id).FindOne(context.Background())
		if err != nil {
			t.Fatalf("Failed to find updated user: %v", err)
		}

		if found.Email != "updated@example.com" {
			t.Errorf("Expected email 'updated@example.com', got '%s'", found.Email)
		}
		if found.FullName == nil || *found.FullName != "Updated Name" {
			t.Error("Expected full name to be 'Updated Name'")
		}
	})

	t.Run("Delete User", func(t *testing.T) {
		// Insert a user
		user := &models2.Users{
			Username: "deleteme",
			Email:    "delete@example.com",
		}
		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		// Delete the user
		deleted, err := testDB.Users().WhereIdEq(*user.Id).Delete(context.Background())
		if err != nil {
			t.Fatalf("Failed to delete user: %v", err)
		}

		if deleted != 1 {
			t.Errorf("Expected 1 row deleted, got %d", deleted)
		}

		// Verify deletion
		_, err = testDB.Users().WhereIdEq(*user.Id).FindOne(context.Background())
		if err != pgx.ErrNoRows {
			t.Error("Expected ErrNoRows after deletion")
		}
	})
}

// TestIntegration_QueryBuilder tests query builder features
func TestIntegration_QueryBuilder(t *testing.T) {
	cleanupTables(t)

	// Insert test data
	users := []*models2.Users{
		{Username: "alice", Email: "alice@example.com", FullName: strPtr("Alice Smith"), IsActive: boolPtr(true)},
		{Username: "bob", Email: "bob@example.com", FullName: strPtr("Bob Jones"), IsActive: boolPtr(true)},
		{Username: "charlie", Email: "charlie@example.com", FullName: strPtr("Charlie Brown"), IsActive: boolPtr(false)},
		{Username: "david", Email: "david@example.com", FullName: strPtr("David Wilson"), IsActive: boolPtr(true)},
	}

	for _, user := range users {
		if err := testDB.Users().Insert(context.Background(), user); err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}
	}

	t.Run("Where Equals", func(t *testing.T) {
		results, err := testDB.Users().WhereUsernameEq("alice").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find users: %v", err)
		}

		if len(results) != 1 {
			t.Errorf("Expected 1 user, got %d", len(results))
		}
		if results[0].Username != "alice" {
			t.Errorf("Expected username 'alice', got '%s'", results[0].Username)
		}
	})

	t.Run("Where LIKE", func(t *testing.T) {
		results, err := testDB.Users().WhereUsernameLike("a%").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find users: %v", err)
		}

		if len(results) != 1 {
			t.Errorf("Expected 1 user with username starting with 'a', got %d", len(results))
		}
	})

	t.Run("Where IN", func(t *testing.T) {
		results, err := testDB.Users().WhereUsernameIn("alice", "bob").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find users: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 users, got %d", len(results))
		}
	})

	t.Run("Where Boolean", func(t *testing.T) {
		results, err := testDB.Users().WhereIsActiveEq(true).Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find active users: %v", err)
		}

		if len(results) != 3 {
			t.Errorf("Expected 3 active users, got %d", len(results))
		}
	})

	t.Run("Order By", func(t *testing.T) {
		results, err := testDB.Users().
			OrderByUsername(models2.ASC).
			Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find users: %v", err)
		}

		if len(results) != 4 {
			t.Fatalf("Expected 4 users, got %d", len(results))
		}

		// Verify ordering
		expected := []string{"alice", "bob", "charlie", "david"}
		for i, username := range expected {
			if results[i].Username != username {
				t.Errorf("Expected user %d to be '%s', got '%s'", i, username, results[i].Username)
			}
		}
	})

	t.Run("Limit and Offset", func(t *testing.T) {
		results, err := testDB.Users().
			OrderByUsername(models2.ASC).
			Limit(2).
			Offset(1).
			Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find users: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 users, got %d", len(results))
		}

		// Should get bob and charlie (offset 1, limit 2)
		if results[0].Username != "bob" {
			t.Errorf("Expected first user to be 'bob', got '%s'", results[0].Username)
		}
		if results[1].Username != "charlie" {
			t.Errorf("Expected second user to be 'charlie', got '%s'", results[1].Username)
		}
	})

	t.Run("Count", func(t *testing.T) {
		count, err := testDB.Users().WhereIsActiveEq(true).Count(context.Background())
		if err != nil {
			t.Fatalf("Failed to count users: %v", err)
		}

		if count != 3 {
			t.Errorf("Expected count of 3, got %d", count)
		}
	})
}

// TestIntegration_Joins tests foreign key joins
func TestIntegration_Joins(t *testing.T) {
	cleanupTables(t)

	// Insert test data
	user := &models2.Users{
		Username: "blogger",
		Email:    "blogger@example.com",
	}
	if err := testDB.Users().Insert(context.Background(), user); err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	posts := []*models2.Posts{
		{UserId: *user.Id, Title: "First Post", Content: strPtr("Content 1"), Status: strPtr("published")},
		{UserId: *user.Id, Title: "Second Post", Content: strPtr("Content 2"), Status: strPtr("draft")},
	}

	for _, post := range posts {
		if err := testDB.Posts().Insert(context.Background(), post); err != nil {
			t.Fatalf("Failed to insert post: %v", err)
		}
	}

	t.Run("Join Users", func(t *testing.T) {
		// Find posts with join to users
		results, err := testDB.Posts().
			JoinUsers().
			WhereUserIdEq(*user.Id).
			Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find posts: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 posts, got %d", len(results))
		}

		for _, post := range results {
			if post.UserId != *user.Id {
				t.Errorf("Expected user_id %s, got %s", *user.Id, post.UserId)
			}
		}
	})

	t.Run("Filter by Related Table", func(t *testing.T) {
		results, err := testDB.Posts().
			WhereStatusEq("published").
			Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find published posts: %v", err)
		}

		if len(results) != 1 {
			t.Errorf("Expected 1 published post, got %d", len(results))
		}

		if results[0].Title != "First Post" {
			t.Errorf("Expected 'First Post', got '%s'", results[0].Title)
		}
	})
}

// TestIntegration_Transactions tests transaction support
func TestIntegration_Transactions(t *testing.T) {
	cleanupTables(t)

	t.Run("Successful Transaction", func(t *testing.T) {
		err := testDB.Transaction(context.Background(), func(tx *models2.Tx) error {
			// Insert user
			user := &models2.Users{
				Username: "txuser",
				Email:    "tx@example.com",
			}
			if err := tx.Users().Insert(context.Background(), user); err != nil {
				return err
			}

			// Insert post
			post := &models2.Posts{
				UserId: *user.Id,
				Title:  "Transaction Post",
			}
			if err := tx.Posts().Insert(context.Background(), post); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			t.Fatalf("Transaction failed: %v", err)
		}

		// Verify data was committed
		users, err := testDB.Users().WhereUsernameEq("txuser").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find user: %v", err)
		}
		if len(users) != 1 {
			t.Error("Expected user to be committed")
		}

		posts, err := testDB.Posts().WhereTitleEq("Transaction Post").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to find post: %v", err)
		}
		if len(posts) != 1 {
			t.Error("Expected post to be committed")
		}
	})

	t.Run("Failed Transaction Rollback", func(t *testing.T) {
		err := testDB.Transaction(context.Background(), func(tx *models2.Tx) error {
			// Insert user
			user := &models2.Users{
				Username: "rollbackuser",
				Email:    "rollback@example.com",
			}
			if err := tx.Users().Insert(context.Background(), user); err != nil {
				return err
			}

			// Intentionally return error to trigger rollback
			return fmt.Errorf("intentional error")
		})

		if err == nil {
			t.Fatal("Expected transaction to fail")
		}

		// Verify data was rolled back
		users, err := testDB.Users().WhereUsernameEq("rollbackuser").Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to query users: %v", err)
		}
		if len(users) != 0 {
			t.Error("Expected user to be rolled back")
		}
	})
}

// TestIntegration_NullableFields tests nullable field handling
func TestIntegration_NullableFields(t *testing.T) {
	cleanupTables(t)

	t.Run("Insert with NULL values", func(t *testing.T) {
		user := &models2.Users{
			Username: "nulltest",
			Email:    "null@example.com",
			FullName: nil, // NULL value
		}

		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user with NULL: %v", err)
		}

		// Verify NULL was stored
		found, err := testDB.Users().WhereIdEq(*user.Id).FindOne(context.Background())
		if err != nil {
			t.Fatalf("Failed to find user: %v", err)
		}

		if found.FullName != nil {
			t.Error("Expected FullName to be NULL")
		}
	})

	t.Run("Update to NULL", func(t *testing.T) {
		user := &models2.Users{
			Username: "nullupdate",
			Email:    "nullupdate@example.com",
			FullName: strPtr("Initial Name"),
		}

		err := testDB.Users().Insert(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		// Update to NULL
		user.FullName = nil
		err = testDB.Users().Update(context.Background(), user)
		if err != nil {
			t.Fatalf("Failed to update user: %v", err)
		}

		// Verify NULL was stored
		found, err := testDB.Users().WhereIdEq(*user.Id).FindOne(context.Background())
		if err != nil {
			t.Fatalf("Failed to find user: %v", err)
		}

		if found.FullName != nil {
			t.Error("Expected FullName to be NULL after update")
		}
	})
}

// TestIntegration_ComplexQueries tests more complex query scenarios
func TestIntegration_ComplexQueries(t *testing.T) {
	cleanupTables(t)

	// Setup test data
	user1 := &models2.Users{Username: "user1", Email: "user1@example.com"}
	user2 := &models2.Users{Username: "user2", Email: "user2@example.com"}

	testDB.Users().Insert(context.Background(), user1)
	testDB.Users().Insert(context.Background(), user2)

	// Create posts for both users
	posts := []*models2.Posts{
		{UserId: *user1.Id, Title: "User1 Post 1", ViewCount: int64Ptr(100)},
		{UserId: *user1.Id, Title: "User1 Post 2", ViewCount: int64Ptr(200)},
		{UserId: *user2.Id, Title: "User2 Post 1", ViewCount: int64Ptr(50)},
	}

	for _, post := range posts {
		testDB.Posts().Insert(context.Background(), post)
	}

	t.Run("Greater Than Query", func(t *testing.T) {
		results, err := testDB.Posts().WhereViewCountGt(75).Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to query posts: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 posts with view_count > 75, got %d", len(results))
		}
	})

	t.Run("Multiple Conditions", func(t *testing.T) {
		results, err := testDB.Posts().
			WhereUserIdEq(*user1.Id).
			And().
			WhereViewCountGte(100).
			Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to query posts: %v", err)
		}

		if len(results) != 2 {
			t.Errorf("Expected 2 posts, got %d", len(results))
		}
	})

	t.Run("UpdateFields", func(t *testing.T) {
		// Update view_count for all user1's posts
		affected, err := testDB.Posts().
			WhereUserIdEq(*user1.Id).
			UpdateFields(context.Background(), map[*models2.FieldRef]interface{}{
				models2.PostsTable.ViewCount(): int64(999),
			})

		if err != nil {
			t.Fatalf("Failed to update posts: %v", err)
		}

		if affected != 2 {
			t.Errorf("Expected 2 rows affected, got %d", affected)
		}

		// Verify updates
		results, err := testDB.Posts().WhereUserIdEq(*user1.Id).Find(context.Background())
		if err != nil {
			t.Fatalf("Failed to query posts: %v", err)
		}

		for _, post := range results {
			if post.ViewCount == nil || *post.ViewCount != 999 {
				t.Error("Expected view_count to be updated to 999")
			}
		}
	})
}

// Helper functions
func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func int64Ptr(i int64) *int64 {
	return &i
}
