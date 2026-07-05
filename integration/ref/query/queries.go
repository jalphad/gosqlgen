package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/post_tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

// CommentsDtoInsertOne inserts a single comments DTO
func CommentsDtoInsertOne(pool *pgxpool.Pool, dto *models.CommentsDto) builder.InsertFinalizeQuery[models.CommentsDto] {
	return comments.NewQuery(pool).
		Insert(
			comments.PostId(),
			comments.UserId(),
			comments.Content(),
		).
		Returning(comments.Id()).
		Values(dto)
}

// CommentsDtoInsertMany inserts multiple comments DTOs
func CommentsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.CommentsDto) builder.InsertFinalizeQuery[models.CommentsDto] {
	return comments.NewQuery(pool).
		Insert(
			comments.PostId(),
			comments.UserId(),
			comments.Content(),
		).
		Returning(comments.Id()).
		Values(dtos...)
}

// PostTagsDtoInsertOne inserts a single post_tags DTO
func PostTagsDtoInsertOne(pool *pgxpool.Pool, dto *models.PostTagsDto) builder.InsertFinalizeQuery[models.PostTagsDto] {
	return post_tags.NewQuery(pool).
		Insert(
			post_tags.PostId(),
			post_tags.TagId(),
		).
		Returning(post_tags.PostId(), post_tags.TagId()).
		Values(dto)
}

// PostTagsDtoInsertMany inserts multiple post_tags DTOs
func PostTagsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.PostTagsDto) builder.InsertFinalizeQuery[models.PostTagsDto] {
	return post_tags.NewQuery(pool).
		Insert(
			post_tags.PostId(),
			post_tags.TagId(),
		).
		Returning(post_tags.PostId(), post_tags.TagId()).
		Values(dtos...)
}

// PostsDtoInsertOne inserts a single posts DTO
func PostsDtoInsertOne(pool *pgxpool.Pool, dto *models.PostsDto) builder.InsertFinalizeQuery[models.PostsDto] {
	return posts.NewQuery(pool).
		Insert(
			posts.UserId(),
			posts.Title(),
			posts.Content(),
			posts.PublishedAt(),
		).
		Returning(posts.Id()).
		Values(dto)
}

// PostsDtoInsertMany inserts multiple posts DTOs
func PostsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.PostsDto) builder.InsertFinalizeQuery[models.PostsDto] {
	return posts.NewQuery(pool).
		Insert(
			posts.UserId(),
			posts.Title(),
			posts.Content(),
			posts.PublishedAt(),
		).
		Returning(posts.Id()).
		Values(dtos...)
}

// TagsDtoInsertOne inserts a single tags DTO
func TagsDtoInsertOne(pool *pgxpool.Pool, dto *models.TagsDto) builder.InsertFinalizeQuery[models.TagsDto] {
	return tags.NewQuery(pool).
		Insert(
			tags.Name(),
			tags.Slug(),
		).
		Returning(tags.Id()).
		Values(dto)
}

// TagsDtoInsertMany inserts multiple tags DTOs
func TagsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.TagsDto) builder.InsertFinalizeQuery[models.TagsDto] {
	return tags.NewQuery(pool).
		Insert(
			tags.Name(),
			tags.Slug(),
		).
		Returning(tags.Id()).
		Values(dtos...)
}

// UsersDtoInsertOne inserts a single users DTO
func UsersDtoInsertOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dto)
}

// UsersDtoInsertMany inserts multiple users DTOs
func UsersDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dtos...)
}

func UpdateUser(pool *pgxpool.Pool, dto *models.UsersDto) builder.UpdateFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Update(
			SetTo(dto,
				users.Username(),
				users.Email(),
				users.FullName(),
				users.IsActive()),
		).Where(users.Id().Eq(Val(*dto.Id)))
}

func UpdateUsers(pool *pgxpool.Pool, in ...*models.UsersDto) builder.UpdateFinalizeQuery[models.UsersDto] {
	dtos := models.UsersDtos(in)
	v := ast.NewAlias("v",
		users.Id(),
		users.Username(),
		users.Email(),
		users.FullName())
	return users.NewQuery(pool).
		Update(
			Set(users.Username()).To(Rel(v, users.Username())),
			Set(users.Email()).To(Rel(v, users.Email())),
			Set(users.FullName()).To(Rel(v, users.FullName())),
		).
		From(
			Unnest(
				Cast(dtos.Id()).AsUUIDArray(),
				Cast(dtos.Username()).AsTextArray(),
				Cast(dtos.Email()).AsTextArray(),
				Cast(dtos.FullName()).AsTextArray(),
			).As(v),
		).
		Where(users.Id().Eq(Rel(v, users.Id())))
}

func DeleteUser(pool *pgxpool.Pool, userId uuid.UUID) builder.DeleteFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}
