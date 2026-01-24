package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
)

// DB wraps the database connection pool
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates a new database wrapper
func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

//// Comments returns a query builder for comments
//func (db *DB) Comments() *builder.KnownTableBuilder[CommentsDto, *CommentsDto] {
//	return NewCommentsQuery(db.pool)
//}
//
//// PostTags returns a query builder for post_tags
//func (db *DB) PostTags() *builder.KnownTableBuilder[PostTagsDto, *PostTagsDto] {
//	return NewPostTagsQuery(db.pool)
//}
//
//// Posts returns a query builder for posts
//func (db *DB) Posts() *builder.KnownTableBuilder[PostsDto, *PostsDto] {
//	return NewPostsQuery(db.pool)
//}
//
//// Tags returns a query builder for tags
//func (db *DB) Tags() *builder.KnownTableBuilder[TagsDto, *TagsDto] {
//	return NewTagsQuery(db.pool)
//}

// Users returns a query builder for users
func (db *DB) Users() *builder.KnownTableBuilder[UsersDtos, UsersDto, *UsersDto] {
	return NewQuery(db.pool, UsersDtos{})
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

//// Comments returns a query builder within the transaction
//func (tx *Tx) Comments() builder.KnownTableStartQuery[CommentsDto, *CommentsDto] {
//	q := NewCommentsQuery(nil)
//	return q.WithTx(tx.tx)
//}
//
//// PostTags returns a query builder within the transaction
//func (tx *Tx) PostTags() builder.KnownTableStartQuery[PostTagsDto, *PostTagsDto] {
//	q := NewPostTagsQuery(nil)
//	return q.WithTx(tx.tx)
//}
//
//// Posts returns a query builder within the transaction
//func (tx *Tx) Posts() builder.KnownTableStartQuery[PostsDto, *PostsDto] {
//	q := NewPostsQuery(nil)
//	return q.WithTx(tx.tx)
//}
//
//// Tags returns a query builder within the transaction
//func (tx *Tx) Tags() builder.KnownTableStartQuery[TagsDto, *TagsDto] {
//	q := NewTagsQuery(nil)
//	return q.WithTx(tx.tx)
//}

// Users returns a query builder within the transaction
func (tx *Tx) Users() builder.KnownTableStartQuery[UsersDtos, UsersDto, *UsersDto] {
	q := NewQuery(nil, UsersDtos{})
	return q.WithTx(tx.tx)
}

//// NewCommentsQuery returns a query builder for comments
//func NewCommentsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[CommentsDto, *CommentsDto] {
//	return builder.NewKnownTableBuilder(&CommentsDto{}, pool)
//}
//
//// NewPostTagsQuery returns a query builder for post_tags
//func NewPostTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[PostTagsDto, *PostTagsDto] {
//	return builder.NewKnownTableBuilder(&PostTagsDto{}, pool)
//}
//
//// NewPostsQuery returns a query builder for posts
//func NewPostsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[PostsDto, *PostsDto] {
//	return builder.NewKnownTableBuilder(&PostsDto{}, pool)
//}
//
//// NewTagsQuery returns a query builder for tags
//func NewTagsQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[TagsDto, *TagsDto] {
//	return builder.NewKnownTableBuilder(&TagsDto{}, pool)
//}

// NewUsersQuery returns a query builder for users
func NewUsersQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[UsersDtos, UsersDto, *UsersDto] {
	return builder.NewKnownTableBuilder(UsersDtos{}, pool)
}

func NewQuery[S builder.DTOs[T, O], T any, O builder.DTO[T]](pool *pgxpool.Pool, s S) *builder.KnownTableBuilder[S, T, O] {
	return builder.NewKnownTableBuilder(s, pool)
}
