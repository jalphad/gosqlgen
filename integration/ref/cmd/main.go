package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/ref"
)

func main() {
	// Example 1: Simple equality
	query1 := models.NewUsersQuery(nil).Where(models.UsersClause().Id.Eq(uuid.New()).And().Email.Like("%@ing.com"))
	sql1, args1 := query1.ToSql()
	fmt.Println("Query 1:")
	fmt.Println(sql1)
	fmt.Printf("Args: %v\n\n", args1)

	// Example 2: Chained conditions with AND
	testTime := time.Now().Add(-24 * time.Hour)
	query2 := models.NewUsersQuery(nil).Where(
		models.UsersClause().CreatedAt.Gt(testTime).And().IsActive.Eq(true),
	)
	sql2, args2 := query2.ToSql()
	fmt.Println("Query 2:")
	fmt.Println(sql2)
	fmt.Printf("Args: %v\n\n", args2)

	// Example 3: Complex query with functions and precedence
	// (LENGTH(username) > 5 OR email LIKE '%@corp.com') AND created_at > 'timestamp'
	query3 := models.NewUsersQuery(nil).Where(
		models.GroupClause(
			models.UsersClause().Username.IsNotNull().Or().Email.Like("%@corp.com"),
		).And().CreatedAt.Gt(testTime),
	)
	sql3, args3 := query3.ToSql()
	fmt.Println("Query 3:")
	fmt.Println(sql3)
	fmt.Printf("Args: %v\n\n", args3)

	// Example 3: Complex query with functions and precedence
	// created_at > 'timestamp' AND (LENGTH(username) > 5 OR email LIKE '%@corp.com')
	query31 := models.NewUsersQuery(nil).Where(
		models.UsersClause().CreatedAt.Gt(testTime).And().
			GroupClause(
				models.UsersClause().Username.IsNotNull().Or().Email.Like("%@corp.com")),
	)
	sql31, args31 := query31.ToSql()
	fmt.Println("Query 3.1:")
	fmt.Println(sql31)
	fmt.Printf("Args: %v\n\n", args31)

	// Example 4: Multiple OR conditions
	query4 := models.NewUsersQuery(nil).Where(
		models.UsersClause().Email.Like("%@example.com").Or().
			Email.Like("%test.com").Or().
			IsActive.IsTrue(),
	)
	sql4, args4 := query4.ToSql()
	fmt.Println("Query 4:")
	fmt.Println(sql4)
	fmt.Printf("Args: %v\n\n", args4)

	// Example 5: Using IN clause
	query5 := models.NewUsersQuery(nil).Where(
		models.UsersClause().Username.In("alice", "bob", "charlie"),
	)
	sql5, args5 := query5.ToSql()
	fmt.Println("Query 5:")
	fmt.Println(sql5)
	fmt.Printf("Args: %v\n\n", args5)
}
