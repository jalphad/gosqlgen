# Getting Started

This tutorial builds and renders a typed `SELECT` query using the current generated reference API.

The example uses the checked-in schema and generated reference package. In an application, you would follow the same flow with your own Postgres schema.

## Start With a Schema

Export or write a schema file containing `CREATE TABLE` statements. This tutorial uses `integration/schema.sql`; the `users` table starts like this:

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

## Generate Code

<!-- docsgen:generation:start -->
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
		WithPackagePath("example.com/myapp")

	if err := gen.ParseFile("schema.sql"); err != nil {
		log.Fatalf("parse schema: %v", err)
	}
	if err := gen.GenerateFiles(); err != nil {
		log.Fatalf("generate files: %v", err)
	}
}
```

Run it:

```sh
go run ./cmd/gen
```

The generator writes these files:

- `db/db.gen.go`
- `db/models/users.gen.go`
- `db/query/helpers.gen.go`
- `db/query/users/users.go`

This generation example is defined in `docs-src/examples.cue` and tested by [`docs/examples/generate_from_schema_test.go`](../examples/generate_from_schema_test.go).
<!-- docsgen:generation:end -->

## Query the Generated API

Create a query builder for the generated `users` table, select a few columns, add predicates, and render the SQL:

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
