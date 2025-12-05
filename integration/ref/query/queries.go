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

func UserWithComments(id uuid.UUID, pool *pgxpool.Pool) builder.CompleteQuery[models.UsersDto] {
	return models.NewUsersQuery(pool).
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
		Where(users.Id().Eq(Lit(id))).
		GroupBy(users.Id())
}

func PostsWithTags(id int64, pool *pgxpool.Pool) builder.CompleteQuery[models.PostsDto] {
	return models.NewPostsQuery(pool).
		Select(
			posts.Id(),
			posts.Title(),
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
		Where(posts.Id().Eq(Lit(id))).
		GroupBy(posts.Id(), posts.Title())
}
