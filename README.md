# gosql

## Overview

gosql is a work-in-progress Go SQL library for Postgres schemas. It generates typed Go APIs from exported SQL schemas while keeping the generated SQL visible and easy to inspect.

The current query direction focuses on a SQL-like builder API backed by an AST. Generated table packages expose typed columns, projections, predicates, ordering helpers, and DTO bindings so common queries can be written without stringly typed SQL.

## Getting Started

Start with a Postgres schema export. The checked-in `integration/schema.sql` is the reference example; a small part of it looks like this:

```sql
CREATE TABLE users
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username   VARCHAR(50)  NOT NULL UNIQUE,
    email      VARCHAR(255) NOT NULL UNIQUE,
    full_name  VARCHAR(100),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    is_active  BOOLEAN          DEFAULT true,
    attributes JSONB
);
```

Add a small generator command in your application, for example `cmd/gen/main.go`:

```go
package main

import (
	"log"

	"github.com/jalphad/gosqlgen"
)

func main() {
	gen := gosqlgen.New().
		WithPackageName("db").
		WithOutputPath(".").
		WithPackagePath("example.com/myapp").
		WithSchemas("public")

	if err := gen.ParseFile("schema.sql"); err != nil {
		log.Fatalf("parse schema: %v", err)
	}
	if err := gen.GenerateFiles(); err != nil {
		log.Fatalf("generate files: %v", err)
	}
}
```

`WithSchemas` limits generation to the named Postgres schemas. Omit it, or
pass no schema names, to generate every parsed table from every schema. The
same setting is available as `Schemas []string` when using
`gosqlgen.NewWithConfig`.

Generated table names are always schema-qualified and quoted, such as
`"public"."users"`. If selected schemas contain tables with the same name,
their Go APIs are schema-prefixed (for example, `AuthUsersDto` and
`BillingUsersDto`).

Run it:

```sh
go run ./cmd/gen
```

That writes generated code under `./db`, including DTOs in `db/models`, table-specific query helpers in `db/query/<table>`, shared query helpers in `db/query`, and a `db.gen.go` facade.

After generation, write typed queries against the generated packages. The checked-in reference package under `integration/ref` shows the intended API shape:

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
SELECT "public"."users".id, "public"."users".email FROM "public"."users" WHERE "public"."users".email LIKE $1 AND "public"."users".is_active = $2 ORDER BY "public"."users".email ASC LIMIT 10
```

with params:

```go
[]any{"%@corp.com", true}
```

For the tested walkthrough, see `docs/tutorials/getting-started.md`.
