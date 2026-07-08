# Getting Started

This tutorial builds and renders a typed `SELECT` query using the current generated reference API.

The example uses the checked-in reference schema under `integration/ref`. In an application, these packages would be generated from your exported Postgres schema.

## Select Active Users

Create a query builder for the `users` table, select a few columns, add predicates, and render the SQL:

```go
qry := ref.NewQuery[models.UsersDto](nil, users.Table()).
	Select(
		users.Id(),
		users.Email(),
	).
	Where(
		users.Email().Like("%@corp.com").
			And(users.IsActive().IsTrue()),
	).
	OrderBy(q.Asc(users.Email())).
	Limit(10)

sql, params, err := qry.ToSql()
```

The rendered SQL stays easy to inspect:

```sql
SELECT users.id, users.email FROM users WHERE users.email LIKE $1 AND users.is_active = $2 ORDER BY users.email ASC LIMIT 10
```

The parameters are:

```go
[]any{"%@corp.com", true}
```

The full tutorial example is tested in [`docs/examples/getting_started_test.go`](../examples/getting_started_test.go), so the docs fail when the example stops compiling or the rendered SQL changes.
