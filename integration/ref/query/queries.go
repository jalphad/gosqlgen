package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

func InsertUser(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewUsersQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
			users.IsActive(),
		).
		Returning(users.Id()).
		Values(dto)
}

func InsertUsers(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewUsersQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
			users.IsActive(),
		).
		Returning(users.Id()).
		Values(dtos...)
}

func UpdateUser(pool *pgxpool.Pool, dto *models.UsersDto) builder.UpdateFinalizeQuery[models.UsersDto, *models.UsersDto] {

	return models.NewQuery(pool, models.UsersDtos{}).
		Update(
			Set(users.Username()).To(Val(dto.Username)),
			Set(users.Email()).To(Val(dto.Email)),
			Set(users.FullName()).To(NVal(dto.FullName)),
			Set(users.IsActive()).To(NVal(dto.IsActive)),
		).Where(users.Id().Eq(Val(*dto.Id)))
}

func UpdateUsers(pool *pgxpool.Pool, dtos models.UsersDtos) builder.UpdateFinalizeQuery[models.UsersDto, *models.UsersDto] {
	//UPDATE users
	//SET username = v.u
	//FROM UNNEST ($1::int[], $2::text[]) AS v(id,u)
	//WHERE users.id = v.id;

	v := ast.NewRelation("v")
	return models.NewUsersQuery(pool).
		Update(
			Set(users.Username()).To(Rel(v, users.Username())),
		).From(
		Unnest(
			Cast(dtos.Id()).AsUUIDArray(),
			Cast(dtos.Username()).AsTextArray(),
		).As(v, users.Id(), users.Username()),
	).Where(users.Id().Eq(Rel(v, users.Id())))
}

func DeleteUser(pool *pgxpool.Pool, userId uuid.UUID) builder.DeleteFinalizeQuery[models.UsersDto, *models.UsersDto] {
	return models.NewQuery(pool, models.UsersDtos{}).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}

func RetrieveUserWithComments(id uuid.UUID, pool *pgxpool.Pool) builder.SelectFinalizeQuery[models.UsersDto] {
	return models.NewQuery(pool, models.UsersDtos{}).
		Select(
			users.Id(),
			As("comments", Coalesce(
				JsonAgg(
					Distinct(JsonbBuildObject(
						comments.Id(),
						comments.UserId(),
						comments.Content(),
					)),
				),
			)),
		).
		Join(ast.JoinLeft, "comments", comments.UserId().Eq(users.Id())).
		Where(users.Id().Eq(Val(id))).
		GroupBy(users.Id())
}

//func RetrievePostWithTags(id int64, pool *pgxpool.Pool) builder.SelectFinalizeQuery[models.PostsDto] {
//	return models.NewPostsQuery(pool).
//		Select(
//			posts.Id(),
//			posts.Title(),
//			posts.UserId(),
//			As("tags", Coalesce(
//				JsonAgg(
//					Distinct(JsonbBuildObject(
//						tags.Id(),
//						tags.Name(),
//						tags.Slug(),
//					)),
//				),
//			))).
//		Join(ast.JoinLeft, "post_tags", post_tags.PostId().Eq(posts.Id())).
//		Join(ast.JoinLeft, "tags", tags.Id().Eq(post_tags.TagId())).
//		Where(posts.Id().Eq(Val(id))).
//		GroupBy(posts.Id(), posts.Title(), posts.UserId())
//}
