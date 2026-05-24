package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/output/query/builder"
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
func (db *DB) Comments() *builder.KnownTableBuilder[CommentsDtos, CommentsDto, *CommentsDto] {
	return NewCommentsQuery(db.pool)
}

// PostTags returns a query builder for post_tags
func (db *DB) PostTags() *builder.KnownTableBuilder[PostTagsDtos, PostTagsDto, *PostTagsDto] {
	return NewPostTagsQuery(db.pool)
}

// Posts returns a query builder for posts
func (db *DB) Posts() *builder.KnownTableBuilder[PostsDtos, PostsDto, *PostsDto] {
	return NewPostsQuery(db.pool)
}

// Tags returns a query builder for tags
func (db *DB) Tags() *builder.KnownTableBuilder[TagsDtos, TagsDto, *TagsDto] {
	return NewTagsQuery(db.pool)
}

// Users returns a query builder for users
func (db *DB) Users() *builder.KnownTableBuilder[UsersDtos, UsersDto, *UsersDto] {
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
func (tx *Tx) Comments() builder.KnownTableStartQuery[CommentsDtos, CommentsDto, *CommentsDto] {
	q := NewCommentsQuery(nil)
	return q.WithTx(tx.tx)
}

// PostTags returns a query builder within the transaction
func (tx *Tx) PostTags() builder.KnownTableStartQuery[PostTagsDtos, PostTagsDto, *PostTagsDto] {
	q := NewPostTagsQuery(nil)
	return q.WithTx(tx.tx)
}

// Posts returns a query builder within the transaction
func (tx *Tx) Posts() builder.KnownTableStartQuery[PostsDtos, PostsDto, *PostsDto] {
	q := NewPostsQuery(nil)
	return q.WithTx(tx.tx)
}

// Tags returns a query builder within the transaction
func (tx *Tx) Tags() builder.KnownTableStartQuery[TagsDtos, TagsDto, *TagsDto] {
	q := NewTagsQuery(nil)
	return q.WithTx(tx.tx)
}

// Users returns a query builder within the transaction
func (tx *Tx) Users() builder.KnownTableStartQuery[UsersDtos, UsersDto, *UsersDto] {
	q := NewUsersQuery(nil)
	return q.WithTx(tx.tx)
}

// NewCommentsQuery returns a query builder for comments
func NewCommentsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[CommentsDtos, CommentsDto, *CommentsDto] {
	return builder.NewKnownTableBuilder(CommentsDtos{}, pool)
}

// NewPostTagsQuery returns a query builder for post_tags
func NewPostTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[PostTagsDtos, PostTagsDto, *PostTagsDto] {
	return builder.NewKnownTableBuilder(PostTagsDtos{}, pool)
}

// NewPostsQuery returns a query builder for posts
func NewPostsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[PostsDtos, PostsDto, *PostsDto] {
	return builder.NewKnownTableBuilder(PostsDtos{}, pool)
}

// NewTagsQuery returns a query builder for tags
func NewTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[TagsDtos, TagsDto, *TagsDto] {
	return builder.NewKnownTableBuilder(TagsDtos{}, pool)
}

// NewUsersQuery returns a query builder for users
func NewUsersQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[UsersDtos, UsersDto, *UsersDto] {
	return builder.NewKnownTableBuilder(UsersDtos{}, pool)
}
