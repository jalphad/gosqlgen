# gosql

## Overview

gosql is a work-in-progress Go SQL library for Postgres schemas. It generates typed Go APIs from exported SQL schemas while keeping the generated SQL visible and easy to inspect.

The current query direction focuses on a SQL-like builder API backed by an AST. Generated table packages expose typed columns, projections, predicates, ordering helpers, and DTO bindings so common queries can be written without stringly typed SQL.

## Getting Started

The checked-in reference schema under `integration/ref` shows the intended generated API shape. A minimal typed `SELECT` query looks like this:

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

It renders:

```sql
SELECT users.id, users.email FROM users WHERE users.email LIKE $1 AND users.is_active = $2 ORDER BY users.email ASC LIMIT 10
```

with params:

```go
[]any{"%@corp.com", true}
```

For the tested walkthrough, see `docs/tutorials/getting-started.md`.
