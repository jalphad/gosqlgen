package gen

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/builder"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/users"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/wide_records"
)

// DB wraps the database connection pool
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates a new database wrapper
func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

// Users returns a query builder for users
func (db *DB) Users() *builder.KnownTableBuilder[models.UsersDto] {
	return NewUsersQuery(db.pool)
}

// WideRecords returns a query builder for wide_records
func (db *DB) WideRecords() *builder.KnownTableBuilder[models.WideRecordsDto] {
	return NewWideRecordsQuery(db.pool)
}

// Transaction executes a function within a transaction
func (db *DB) Transaction(ctx context.Context, fn func(*Tx) error) error {
	return pgx.BeginFunc(ctx, db.pool, func(pgxTx pgx.Tx) error {
		tx := NewTx(pgxTx)
		return fn(tx)
	})
}

// Tx provides transaction-aware query builders
type Tx struct {
	tx pgx.Tx
}

// NewTx creates transaction-aware query builders
func NewTx(tx pgx.Tx) *Tx {
	return &Tx{tx: tx}
}

// Users returns a query builder within the transaction
func (tx *Tx) Users() builder.KnownTableStartQuery[models.UsersDto] {
	q := NewUsersQuery(nil)
	return q.WithTx(tx.tx)
}

// WideRecords returns a query builder within the transaction
func (tx *Tx) WideRecords() builder.KnownTableStartQuery[models.WideRecordsDto] {
	q := NewWideRecordsQuery(nil)
	return q.WithTx(tx.tx)
}

// NewUsersQuery returns a query builder for users
func NewUsersQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return users.NewQuery(pool)
}

// NewWideRecordsQuery returns a query builder for wide_records
func NewWideRecordsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.WideRecordsDto] {
	return wide_records.NewQuery(pool)
}

func NewQuery[T any](pool *pgxpool.Pool, table *ast.TableSource) *builder.KnownTableBuilder[T] {
	return builder.NewKnownTableBuilder[T](pool, table)
}
