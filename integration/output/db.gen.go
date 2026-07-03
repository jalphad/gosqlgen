package output

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/output/models"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
	"github.com/jalphad/gosqlgen/integration/output/query/builder"
	"github.com/jalphad/gosqlgen/integration/output/query/comments"
	"github.com/jalphad/gosqlgen/integration/output/query/post_tags"
	"github.com/jalphad/gosqlgen/integration/output/query/posts"
	"github.com/jalphad/gosqlgen/integration/output/query/tags"
	"github.com/jalphad/gosqlgen/integration/output/query/users"
)

// DB wraps the database connection pool
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates a new database wrapper
func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

// Comments returns a query builder for comments
func (db *DB) Comments() *builder.KnownTableBuilder[models.CommentsDto] {
	return NewCommentsQuery(db.pool)
}

// PostTags returns a query builder for post_tags
func (db *DB) PostTags() *builder.KnownTableBuilder[models.PostTagsDto] {
	return NewPostTagsQuery(db.pool)
}

// Posts returns a query builder for posts
func (db *DB) Posts() *builder.KnownTableBuilder[models.PostsDto] {
	return NewPostsQuery(db.pool)
}

// Tags returns a query builder for tags
func (db *DB) Tags() *builder.KnownTableBuilder[models.TagsDto] {
	return NewTagsQuery(db.pool)
}

// Users returns a query builder for users
func (db *DB) Users() *builder.KnownTableBuilder[models.UsersDto] {
	return NewUsersQuery(db.pool)
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

// Comments returns a query builder within the transaction
func (tx *Tx) Comments() builder.KnownTableStartQuery[models.CommentsDto] {
	q := NewCommentsQuery(nil)
	return q.WithTx(tx.tx)
}

// PostTags returns a query builder within the transaction
func (tx *Tx) PostTags() builder.KnownTableStartQuery[models.PostTagsDto] {
	q := NewPostTagsQuery(nil)
	return q.WithTx(tx.tx)
}

// Posts returns a query builder within the transaction
func (tx *Tx) Posts() builder.KnownTableStartQuery[models.PostsDto] {
	q := NewPostsQuery(nil)
	return q.WithTx(tx.tx)
}

// Tags returns a query builder within the transaction
func (tx *Tx) Tags() builder.KnownTableStartQuery[models.TagsDto] {
	q := NewTagsQuery(nil)
	return q.WithTx(tx.tx)
}

// Users returns a query builder within the transaction
func (tx *Tx) Users() builder.KnownTableStartQuery[models.UsersDto] {
	q := NewUsersQuery(nil)
	return q.WithTx(tx.tx)
}

// NewCommentsQuery returns a query builder for comments
func NewCommentsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.CommentsDto] {
	return builder.NewKnownTableBuilder[models.CommentsDto](pool, comments.Table())
}

// NewPostTagsQuery returns a query builder for post_tags
func NewPostTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.PostTagsDto] {
	return builder.NewKnownTableBuilder[models.PostTagsDto](pool, post_tags.Table())
}

// NewPostsQuery returns a query builder for posts
func NewPostsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.PostsDto] {
	return builder.NewKnownTableBuilder[models.PostsDto](pool, posts.Table())
}

// NewTagsQuery returns a query builder for tags
func NewTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.TagsDto] {
	return builder.NewKnownTableBuilder[models.TagsDto](pool, tags.Table())
}

// NewUsersQuery returns a query builder for users
func NewUsersQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return builder.NewKnownTableBuilder[models.UsersDto](pool, users.Table())
}

func NewQuery[T any](pool *pgxpool.Pool, table *ast.TableSource) *builder.KnownTableBuilder[T] {
	return builder.NewKnownTableBuilder[T](pool, table)
}
