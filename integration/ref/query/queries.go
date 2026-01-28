package query

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/post_tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

func UsersDtoInsertOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
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

func UsersDtoInsertMultiple(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto, *models.UsersDto] {
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
	return models.NewQuery(pool, models.UsersDtos{}).
		Update(
			Set(users.Username()).To(Val(dto.Username)),
			Set(users.Email()).To(Val(dto.Email)),
			Set(users.FullName()).To(NVal(dto.FullName)),
			Set(users.IsActive()).To(NVal(dto.IsActive)),
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
	return models.NewQuery(pool, models.UsersDtos{}).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}

func UsersDtoSelectById(pool *pgxpool.Pool, id uuid.UUID, opts ...func() *SelectOpt) builder.SelectFinalizeQuery[models.UsersDto] {
	var toSelect = []ast.NamedExpression{
		users.Id(),
	}
	for _, opt := range opts {
		toSelect = append(toSelect, opt().selectExpr)
	}

	partial := models.NewUsersQuery(pool).
		Select(
			toSelect...,
		)
	for _, opt := range opts {
		joinExpr := opt().joinExpr
		partial = partial.Join(joinExpr.Type, joinExpr.Table, joinExpr.Condition)
	}

	return partial.
		Where(users.Id().Eq(Val(id))).
		GroupBy(users.Id())
}

func UserWithComments() *SelectOpt {
	alias := ast.NewAlias("comments")
	return &SelectOpt{
		selectExpr: Coalesce(
			JsonAgg(
				Distinct(JsonbBuildObject(
					comments.Id(),
					comments.UserId(),
					comments.Content(),
					comments.CreatedAt(),
				)),
			),
		).As(alias),
		joinExpr: JoinExpr{
			Type:      ast.JoinLeft,
			Table:     "comments",
			Condition: comments.UserId().Eq(users.Id()),
		},
	}
}

func WithPosts() *SelectOpt {
	alias := ast.NewAlias("posts")
	return &SelectOpt{
		selectExpr: Coalesce(
			JsonAgg(
				Distinct(JsonbBuildObject(
					posts.Id(),
					posts.Title(),
				)),
			),
		).As(alias),
		joinExpr: JoinExpr{
			Type:      ast.JoinLeft,
			Table:     "posts",
			Condition: posts.UserId().Eq(users.Id()),
		},
	}
}

func RetrieveUserWithComments(id uuid.UUID, pool *pgxpool.Pool) builder.SelectFinalizeQuery[models.UsersDto] {
	c := ast.NewAlias("comments")
	return models.NewUsersQuery(pool).
		Select(
			users.Id(),
			Coalesce(
				JsonAgg(
					Distinct(JsonbBuildObject(
						comments.Id(),
						comments.UserId(),
						comments.Content(),
						comments.CreatedAt(),
					)),
				),
			).As(c)).
		Join(ast.JoinLeft, "comments", comments.UserId().Eq(users.Id())).
		Where(users.Id().Eq(Val(id))).
		GroupBy(users.Id())
}

func RetrievePostWithTags(id int64, pool *pgxpool.Pool) builder.SelectFinalizeQuery[models.PostsDto] {
	return models.NewPostsQuery(pool).
		Select(
			posts.Id(),
			posts.Title(),
			posts.UserId(),
			As("tags", Coalesce(
				JsonAgg(
					Distinct(JsonbBuildObject(
						tags.Id(),
						tags.Name(),
						tags.Slug(),
					)),
				),
			))).
		Join(ast.JoinLeft, "post_tags", post_tags.PostId().Eq(posts.Id())).
		Join(ast.JoinLeft, "tags", tags.Id().Eq(post_tags.TagId())).
		Where(posts.Id().Eq(Val(id))).
		GroupBy(posts.Id(), posts.Title(), posts.UserId())
}
