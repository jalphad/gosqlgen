package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jalphad/gosqlgen"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Example 1: Generate DTOs from SQL schema
	generateModels()

	// Example 2: Use the generated models
	useGeneratedModels()
}

func generateModels() {
	// Define your SQL schema
	schema := `
    -- Users table
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) NOT NULL UNIQUE,
        email VARCHAR(255) NOT NULL UNIQUE,
        full_name VARCHAR(100),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        is_active BOOLEAN DEFAULT true
    );
    
    -- Posts table with foreign key to users
    CREATE TABLE posts (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        status VARCHAR(20) DEFAULT 'draft',
        published_at TIMESTAMP,
        view_count INTEGER DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
    
    -- Comments table with foreign keys
    CREATE TABLE comments (
        id SERIAL PRIMARY KEY,
        post_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        is_approved BOOLEAN DEFAULT false,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
    
    -- Tags table
    CREATE TABLE tags (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL UNIQUE,
        slug VARCHAR(50) NOT NULL UNIQUE
    );
    
    -- Post tags junction table
    CREATE TABLE post_tags (
        post_id INTEGER NOT NULL,
        tag_id INTEGER NOT NULL,
        PRIMARY KEY (post_id, tag_id),
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
    );
    `

	// Create generator
	gen := gosqlgen.New().
		WithPackageName("models").
		WithOutputPath("./generated/models")

	// Parse the schema
	if err := gen.Parse(schema); err != nil {
		log.Fatal("Failed to parse schema:", err)
	}

	// Generate Go code
	code, err := gen.Generate()
	if err != nil {
		log.Fatal("Failed to generate code:", err)
	}

	fmt.Println("Generated code preview:")
	fmt.Println(code[:500], "...") // Show first 500 chars

	// Save to file
	if err := gen.GenerateToFile("./generated/models/models.go"); err != nil {
		log.Fatal("Failed to save generated code:", err)
	}

	fmt.Println("\nModels generated successfully to ./generated/models/models.go")
}

func useGeneratedModels() {
	// This demonstrates how you would use the generated models
	// Note: This is pseudo-code as the actual generated models would be in a separate package

	fmt.Println("\n=== Example Usage of Generated Models ===\n")

	// Connect to database
	dsn := "postgresql://user:password@localhost/mydb?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// The following shows how the generated code would be used:
	demonstrateUsage()
}

func demonstrateUsage() {
	fmt.Println("1. Simple Query Example:")
	fmt.Println(`
    // Query users
    users, err := db.Users().
        Where("is_active = ?", true).
        OrderBy("created_at DESC").
        Limit(10).
        Find()
    `)

	fmt.Println("\n2. Inner Join Example - Get posts with user information:")
	fmt.Println(`
    // Join posts with users
    results, err := db.Posts().
        InnerJoin("users", "posts.user_id = users.id").
        Select("posts.*, users.username, users.email").
        Where("posts.status = ?", "published").
        Where("posts.published_at < ?", time.Now()).
        OrderBy("posts.published_at DESC").
        Limit(20).
        Find()
    `)

	fmt.Println("\n3. Complex Join Example - Get posts with comments and users:")
	fmt.Println(`
    // Multiple joins
    results, err := db.Posts().
        InnerJoin("users", "posts.user_id = users.id").
        InnerJoin("comments", "posts.id = comments.post_id").
        Select("posts.title", "users.username", "COUNT(comments.id) as comment_count").
        Where("posts.status = ?", "published").
        GroupBy("posts.id, posts.title, users.username").
        Having("COUNT(comments.id) > ?", 5).
        OrderBy("comment_count DESC").
        Find()
    `)

	fmt.Println("\n4. Typed Join Example (using generated helper methods):")
	fmt.Println(`
    // Type-safe join using generated methods
    userPosts, err := db.Users().
        JoinPosts().                              // Generated method
        WhereUser("users.is_active = ?", true).   // Type-safe where
        WherePost("posts.status = ?", "published").
        FindUserPosts()                           // Returns typed result
    
    // Access joined data with type safety
    for _, up := range userPosts {
        fmt.Printf("User: %s, Post: %s\n", 
            up.User.Username,  // Typed access
            up.Post.Title)     // Typed access
    }
    `)

	fmt.Println("\n5. Transaction Example:")
	fmt.Println(`
    // Use transactions
    err := db.Transaction(func(tx *Tx) error {
        // Create user
        user := &User{
            Username: "johndoe",
            Email:    "john@example.com",
            FullName: "John Doe",
        }
        if err := tx.Users().Insert(user); err != nil {
            return err
        }
        
        // Create post for the user
        post := &Post{
            UserId:  user.Id,
            Title:   "My First Post",
            Content: "Hello, World!",
            Status:  "published",
        }
        if err := tx.Posts().Insert(post); err != nil {
            return err
        }
        
        return nil
    })
    `)

	fmt.Println("\n6. Many-to-Many Join Example (posts with tags):")
	fmt.Println(`
    // Get posts with their tags
    results, err := db.Posts().
        InnerJoin("post_tags", "posts.id = post_tags.post_id").
        InnerJoin("tags", "post_tags.tag_id = tags.id").
        Select("posts.*, GROUP_CONCAT(tags.name) as tag_names").
        Where("posts.status = ?", "published").
        GroupBy("posts.id").
        Find()
    `)

	fmt.Println("\n7. Subquery Example (users with post count):")
	fmt.Println(`
    // This would be implemented as a custom query method
    query := ` + "`" + `
        SELECT u.*, 
               (SELECT COUNT(*) FROM posts p WHERE p.user_id = u.id) as post_count
        FROM users u
        WHERE u.is_active = true
        HAVING post_count > 0
        ORDER BY post_count DESC
    ` + "`" + `
    results, err := db.Raw(query).Find()
    `)

	fmt.Println("\n8. Update Example:")
	fmt.Println(`
    // Update a user
    user.Email = "newemail@example.com"
    user.UpdatedAt = time.Now()
    err := db.Users().Update(user)
    
    // Bulk update
    affected, err := db.Posts().
        Where("status = ?", "draft").
        Where("created_at < ?", time.Now().AddDate(0, -1, 0)).
        UpdateColumns(map[string]interface{}{
            "status": "archived",
            "updated_at": time.Now(),
        })
    `)

	fmt.Println("\n9. Delete Example:")
	fmt.Println(`
    // Delete with conditions
    deleted, err := db.Comments().
        Where("is_approved = ?", false).
        Where("created_at < ?", time.Now().AddDate(0, 0, -30)).
        Delete()
    
    fmt.Printf("Deleted %d unapproved comments older than 30 days\n", deleted)
    `)
}
