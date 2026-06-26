package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

func UsersDtoInsertOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewUsersQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dto)
}

func UsersDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewUsersQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dtos...)
}

func InsertPost(pool *pgxpool.Pool, dto *models.PostsDto) builder.InsertFinalizeQuery[models.PostsDto, *models.PostsDto] {
	return models.NewPostsQuery(pool).
		Insert(
			posts.UserId(),
			posts.Title(),
			posts.Content(),
		).
		Returning(posts.Id()).
		Values(dto)
}

func InsertComment(pool *pgxpool.Pool, dto *models.CommentsDto) builder.InsertFinalizeQuery[models.CommentsDto, *models.CommentsDto] {
	return models.NewCommentsQuery(pool).
		Insert(
			comments.PostId(),
			comments.UserId(),
			comments.Content(),
		).
		Returning(comments.Id()).
		Values(dto)
}

func UpdateUser(pool *pgxpool.Pool, dto *models.UsersDto) builder.UpdateFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewDTOQuery(pool, models.UsersDtos{}).
		Update(
			SetTo(dto,
				users.Username(),
				users.Email(),
				users.FullName(),
				users.IsActive()),
		).Where(users.Id().Eq(Val(*dto.Id)))
}

func UpdateUsers(pool *pgxpool.Pool, in ...*models.UsersDto) builder.UpdateFinalizeQuery[models.UsersDto, *models.UsersDto] {
	dtos := models.UsersDtos(in)
	v := ast.NewAlias("v")
	return models.NewUsersQuery(pool).
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
			).As(v, users.Id(), users.Username(), users.Email(), users.FullName()),
		).
		Where(users.Id().Eq(Rel(v, users.Id())))
}

func DeleteUser(pool *pgxpool.Pool, userId uuid.UUID) builder.DeleteFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewDTOQuery(pool, models.UsersDtos{}).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}
