package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
func (db *DB) Comments() *CommentsQuery {
	return NewCommentsQuery(db.pool)
}

// PostTags returns a query builder for post_tags
func (db *DB) PostTags() *PostTagsQuery {
	return NewPostTagsQuery(db.pool)
}

// Posts returns a query builder for posts
func (db *DB) Posts() *PostsQuery {
	return NewPostsQuery(db.pool)
}

// Tags returns a query builder for tags
func (db *DB) Tags() *TagsQuery {
	return NewTagsQuery(db.pool)
}

// Users returns a query builder for users
func (db *DB) Users() *UsersQuery {
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
func (tx *Tx) Comments() *CommentsQuery {
	q := NewCommentsQuery(nil)
	q.WithTx(tx.tx)
	return q
}

// PostTags returns a query builder within the transaction
func (tx *Tx) PostTags() *PostTagsQuery {
	q := NewPostTagsQuery(nil)
	q.WithTx(tx.tx)
	return q
}

// Posts returns a query builder within the transaction
func (tx *Tx) Posts() *PostsQuery {
	q := NewPostsQuery(nil)
	q.WithTx(tx.tx)
	return q
}

// Tags returns a query builder within the transaction
func (tx *Tx) Tags() *TagsQuery {
	q := NewTagsQuery(nil)
	q.WithTx(tx.tx)
	return q
}

// Users returns a query builder within the transaction
func (tx *Tx) Users() *UsersQuery {
	q := NewUsersQuery(nil)
	q.WithTx(tx.tx)
	return q
}
