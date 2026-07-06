package main

import (
	"fmt"

	"github.com/jalphad/gosqlgen/integration/ref"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	q "github.com/jalphad/gosqlgen/integration/ref/query"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

func main() {
	qry := ref.NewQuery[models.UsersDto](nil, users.Table()).
		Select(
			users.Id(),
			users.FullName(),
			users.Email()).
		Where(users.Email().Like("%@corp.com").
			And(users.IsActive().IsTrue()).
			And(users.Username().Eq(q.Val("foo")))).
		GroupBy().
		Having(q.Val(true)).
		OrderBy(
			q.Asc(users.Id()),
			q.Desc(users.Username())).
		Limit(20).Offset(20)

	sql, _, _ := qry.ToSql()
	fmt.Println(sql) // SELECT users.id, users.username FROM users AS users WHERE users.email LIKE $1 AND users.is_active = $2
	//fmt.Println(params) // [%@corp.com true]
}
