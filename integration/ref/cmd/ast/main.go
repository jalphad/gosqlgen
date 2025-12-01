package main

import (
	"fmt"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

func main() {
	qry := query.NewUsersQuery().
		Select(
			users.Id(),
			users.FullName().Eq(Lit("foo")),
			users.Email()).
		Where(users.Email().Like("%@corp.com").
			And(users.IsActive().IsTrue()).
			And(users.Username().Eq(Lit("foo")))).
		GroupBy().
		OrderBy(
			query.Asc(users.Id()),
			query.Desc(users.Username())).
		Limit(20).Offset(20)

	sql, params := qry.ToSql()
	fmt.Println(sql)    // SELECT users.id, users.username FROM users AS users WHERE users.email LIKE $1 AND users.is_active = $2
	fmt.Println(params) // [%@corp.com true]

	//query = query.NewUsersQuery().Where(
	//	query.Field("users", "created_at").Gt(time.Now()).
	//		And(query.Field("users", "username").IsNotNull().
	//			Or(query.Field("users", "email").Like("%@corp.com")),
	//		),
	//)
	//
	//sql, params = query.ToSql()
	//fmt.Println(sql)    // SELECT users.id, users.username FROM users AS users WHERE users.created_at > $1 AND (users.username IS NOT NULL OR users.email LIKE $2)
	//fmt.Println(params) // params: [%@corp.com 2025-11-24 15:13:17 +0100 CET]
}

func Lit[T ast.MappedTypes](t T) ast.OfType[T] {
	return ast.NewSQLType(t)
}
