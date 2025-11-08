package gosqlgen

//import (
//    "database/sql"
//    "fmt"
//    "log"
//    "time"
//
//    _ "github.com/lib/pq" // PostgreSQL driver
//)
//
//func main() {
//    // Example showing the new type-safe API
//    demonstrateTypeSafeAPI()
//}
//
//func demonstrateTypeSafeAPI() {
//    // Connect to database
//    dsn := "postgresql://user:password@localhost/mydb?sslmode=disable"
//    db, err := sql.Open("postgres", dsn)
//    if err != nil {
//        log.Fatal("Failed to connect to database:", err)
//    }
//    defer db.Close()
//
//    // Initialize the generated models
//    models := NewDB(db)
//
//    fmt.Println("=== Type-Safe Query Examples ===\n")
//
//    // Example 1: Simple type-safe queries
//    example1SimpleQueries(models)
//
//    // Example 2: Complex conditions with type safety
//    example2ComplexConditions(models)
//
//    // Example 3: Type-safe joins
//    example3TypeSafeJoins(models)
//
//    // Example 4: Field selection with type safety
//    example4FieldSelection(models)
//
//    // Example 5: Updates with type safety
//    example5TypeSafeUpdates(models)
//
//    // Example 6: Transactions
//    example6Transactions(models)
//}
//
//func example1SimpleQueries(db *DB) {
//    fmt.Println("1. Simple Type-Safe Queries:\n")
//
//    // Find all active users - no string column names!
//    activeUsers, err := db.Users().
//        WhereIsActiveEq(true).
//        OrderByCreatedAt(DESC).
//        Limit(10).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Println("Active users query - completely type-safe, no strings!")
//
//    // Find user by email - compile-time checked
//    user, err := db.Users().
//        WhereEmailEq("john@example.com").
//        FindOne()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found user: %s\n", user.Username)
//
//    // Find users created after a date
//    recentUsers, err := db.Users().
//        WhereCreatedAtGt(time.Now().AddDate(0, -1, 0)). // Last month
//        OrderByCreatedAt(DESC).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d recent users\n", len(recentUsers))
//
//    // Find users with username matching pattern
//    johnUsers, err := db.Users().
//        WhereUsernameLike("john%").
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d users named John\n", len(johnUsers))
//}
//
//func example2ComplexConditions(db *DB) {
//    fmt.Println("\n2. Complex Type-Safe Conditions:\n")
//
//    // Multiple conditions with AND (default)
//    posts, err := db.Posts().
//        WhereStatusEq("published").
//        WherePublishedAtNotEq(nil).                    // Not null check
//        WhereViewCountGte(100).                        // >= 100 views
//        OrderByPublishedAt(DESC).
//        Limit(20).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d popular published posts\n", len(posts))
//
//    // Using OR conditions
//    urgentPosts, err := db.Posts().
//        WhereStatusEq("urgent").
//        Or().
//        WhereViewCountGt(1000).                        // OR high view count
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d urgent or popular posts\n", len(urgentPosts))
//
//    // Complex grouped conditions: (status = 'draft' AND created > date) OR (status = 'published' AND views > 100)
//    complexPosts, err := db.Posts().
//        WhereGroup(func(q *PostsQuery) {
//            q.WhereStatusEq("draft").
//                WhereCreatedAtGt(time.Now().AddDate(0, 0, -7)) // Last week
//        }).
//        Or().
//        WhereGroup(func(q *PostsQuery) {
//            q.WhereStatusEq("published").
//                WhereViewCountGt(100)
//        }).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d posts matching complex criteria\n", len(complexPosts))
//
//    // Using IN clause - type-safe
//    specificUsers, err := db.Users().
//        WhereIdIn(1, 2, 3, 4, 5).                     // Type-safe IN clause
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d specific users\n", len(specificUsers))
//}
//
//func example3TypeSafeJoins(db *DB) {
//    fmt.Println("\n3. Type-Safe Joins:\n")
//
//    // Type-safe join using foreign key relationship
//    // The JoinUsers() method is generated based on the foreign key
//    postsWithUsers, err := db.Posts().
//        JoinUsers().                                    // Generated from FK
//        WhereStatusEq("published").
//        SelectAll().                                    // Select all fields from posts
//        OrderByPublishedAt(DESC).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d posts with user data\n", len(postsWithUsers))
//
//    // Left join to include posts without comments
//    postsWithComments, err := db.Posts().
//        LeftJoinComments().                            // Generated left join
//        WhereStatusEq("published").
//        GroupBy(PostsTable.Id()).
//        Having(PostsTable.Id(), GT, 0).               // Type-safe having
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d posts with optional comments\n", len(postsWithComments))
//
//    // Custom join with type-safe field references
//    customJoin, err := db.Posts().
//        JoinOn("INNER JOIN", "tags",
//            PostsTable.Id(),                           // Type-safe field reference
//            TagsTable.PostId()).                       // Type-safe field reference
//        WhereStatusEq("published").
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d posts with custom join\n", len(customJoin))
//}
//
//func example4FieldSelection(db *DB) {
//    fmt.Println("\n4. Type-Safe Field Selection:\n")
//
//    // Select specific fields with type safety
//    userSummaries, err := db.Users().
//        Select(
//            UsersTable.Id(),
//            UsersTable.Username(),
//            UsersTable.Email(),
//        ).
//        WhereIsActiveEq(true).
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d user summaries\n", len(userSummaries))
//
//    // Select with field aliases
//    aliasedResults, err := db.Posts().
//        Select(
//            PostsTable.Title().As("post_title"),
//            PostsTable.ViewCount().As("views"),
//        ).
//        WhereStatusEq("published").
//        Find()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Found %d posts with aliased fields\n", len(aliasedResults))
//
//    // Aggregate queries with type safety
//    count, err := db.Posts().
//        WhereStatusEq("published").
//        WherePublishedAtGte(time.Now().AddDate(0, -1, 0)).
//        Count()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Published posts in last month: %d\n", count)
//}
//
//func example5TypeSafeUpdates(db *DB) {
//    fmt.Println("\n5. Type-Safe Updates:\n")
//
//    // Update specific fields with type safety
//    affected, err := db.Posts().
//        WhereStatusEq("draft").
//        WhereCreatedAtLt(time.Now().AddDate(0, -1, 0)). // Older than a month
//        UpdateFields(map[*FieldRef]interface{}{
//            PostsTable.Status():    "archived",
//            PostsTable.UpdatedAt(): time.Now(),
//        })
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Archived %d old draft posts\n", affected)
//
//    // Delete with type-safe conditions
//    deleted, err := db.Comments().
//        WhereIsApprovedEq(false).
//        WhereCreatedAtLt(time.Now().AddDate(0, 0, -30)). // Older than 30 days
//        Delete()
//
//    if err != nil {
//        log.Printf("Error: %v", err)
//        return
//    }
//
//    fmt.Printf("Deleted %d unapproved old comments\n", deleted)
//}
//
//func example6Transactions(db *DB) {
//    fmt.Println("\n6. Type-Safe Transactions:\n")
//
//    err := db.Transaction(func(tx *Tx) error {
//        // Create a user - all type-safe
//        user := &Users{
//            Username:  "jane_doe",
//            Email:     "jane@example.com",
//            FullName:  "Jane Doe",
//            IsActive:  true,
//            CreatedAt: time.Now(),
//            UpdatedAt: time.Now(),
//        }
//
//        // Insert with transaction
//        if err := tx.Users().Insert(user); err != nil {
//            return fmt.Errorf("failed to insert user: %w", err)
//        }
//
//        // Create posts for the user - type-safe reference to user.Id
//        post := &Posts{
//            UserId:      user.Id,
//            Title:       "My First Post",
//            Content:     "Hello, World!",
//            Status:      "published",
//            PublishedAt: &time.Now(),
//            ViewCount:   0,
//            CreatedAt:   time.Now(),
//            UpdatedAt:   time.Now(),
//        }
//
//        if err := tx.Posts().Insert(post); err != nil {
//            return fmt.Errorf("failed to insert post: %w", err)
//        }
//
//        // Update user's post count - completely type-safe
//        _, err := tx.Users().
//            WhereIdEq(user.Id).
//            UpdateFields(map[*FieldRef]interface{}{
//                UsersTable.PostCount(): 1,
//                UsersTable.UpdatedAt(): time.Now(),
//            })
//
//        if err != nil {
//            return fmt.Errorf("failed to update user post count: %w", err)
//        }
//
//        return nil // Commit transaction
//    })
//
//    if err != nil {
//        log.Printf("Transaction failed: %v", err)
//        return
//    }
//
//    fmt.Println("Transaction completed successfully")
//}
//
//// Example showing the problems that the type-safe API prevents
//func demonstrateTypeS
