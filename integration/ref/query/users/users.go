package users

import (
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/expr"
)

// NewQuery returns a query builder for "public"."users"
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return builder.NewKnownTableBuilder[models.UsersDto](pool, Table())
}

var Into = intoUsersDto{}

type intoUsersDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource(`"public"."users"`)
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return newId(`"public"."users"`, func(u *models.UsersDto) **uuid.UUID {
		return &u.Id
	})
}

func Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	return newUsername(`"public"."users"`, func(u *models.UsersDto) *string {
		return &u.Username
	})
}

func Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	return newEmail(`"public"."users"`, func(u *models.UsersDto) *string {
		return &u.Email
	})
}

func FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	return newFullName(`"public"."users"`, func(u *models.UsersDto) **string {
		return &u.FullName
	})
}

func CreatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return newCreatedAt(`"public"."users"`, func(u *models.UsersDto) **time.Time {
		return &u.CreatedAt
	})
}

func UpdatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return newUpdatedAt(`"public"."users"`, func(u *models.UsersDto) **time.Time {
		return &u.UpdatedAt
	})
}

func IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	return newIsActive(`"public"."users"`, func(u *models.UsersDto) **bool {
		return &u.IsActive
	})
}

func Attributes() *ast.JsonColumnProjection[models.UsersDto, **json.RawMessage] {
	return newAttributes(`"public"."users"`, func(u *models.UsersDto) **json.RawMessage {
		return &u.Attributes
	})
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Username(),
		Email(),
		FullName(),
		CreatedAt(),
		UpdatedAt(),
		IsActive(),
		Attributes(),
	}
}

func (intoUsersDto) Id() ast.Projection[models.UsersDto] {
	return Id()
}

func (intoUsersDto) Username() ast.Projection[models.UsersDto] {
	return Username()
}

func (intoUsersDto) Email() ast.Projection[models.UsersDto] {
	return Email()
}

func (intoUsersDto) FullName() ast.Projection[models.UsersDto] {
	return FullName()
}

func (intoUsersDto) CreatedAt() ast.Projection[models.UsersDto] {
	return CreatedAt()
}

func (intoUsersDto) UpdatedAt() ast.Projection[models.UsersDto] {
	return UpdatedAt()
}

func (intoUsersDto) IsActive() ast.Projection[models.UsersDto] {
	return IsActive()
}

func (intoUsersDto) Attributes() ast.Projection[models.UsersDto] {
	return Attributes()
}

func (intoUsersDto) AllColumns() []ast.Projection[models.UsersDto] {
	return []ast.Projection[models.UsersDto]{
		Into.Id(),
		Into.Username(),
		Into.Email(),
		Into.FullName(),
		Into.CreatedAt(),
		Into.UpdatedAt(),
		Into.IsActive(),
		Into.Attributes(),
	}
}

type forUsersDto[T any] struct {
	dest func(*T) *models.UsersDto
}

func For[T any](dest func(*T) *models.UsersDto) forUsersDto[T] {
	return forUsersDto[T]{dest: dest}
}

func (f forUsersDto[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return newId(`"public"."users"`, func(t *T) **uuid.UUID {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forUsersDto[T]) Username() *ast.StringColumnProjection[T, *string] {
	return newUsername(`"public"."users"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Username
	})
}

func (f forUsersDto[T]) Email() *ast.StringColumnProjection[T, *string] {
	return newEmail(`"public"."users"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Email
	})
}

func (f forUsersDto[T]) FullName() *ast.StringColumnProjection[T, **string] {
	return newFullName(`"public"."users"`, func(t *T) **string {
		dto := f.dest(t)
		return &dto.FullName
	})
}

func (f forUsersDto[T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newCreatedAt(`"public"."users"`, func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.CreatedAt
	})
}

func (f forUsersDto[T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newUpdatedAt(`"public"."users"`, func(t *T) **time.Time {
		dto := f.dest(t)
		return &dto.UpdatedAt
	})
}

func (f forUsersDto[T]) IsActive() *ast.BoolColumnProjection[T, **bool] {
	return newIsActive(`"public"."users"`, func(t *T) **bool {
		dto := f.dest(t)
		return &dto.IsActive
	})
}

func (f forUsersDto[T]) Attributes() *ast.JsonColumnProjection[T, **json.RawMessage] {
	return newAttributes(`"public"."users"`, func(t *T) **json.RawMessage {
		dto := f.dest(t)
		return &dto.Attributes
	})
}

func (f forUsersDto[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
		f.Username(),
		f.Email(),
		f.FullName(),
		f.CreatedAt(),
		f.UpdatedAt(),
		f.IsActive(),
		f.Attributes(),
	}
}

func (intoUsersDto) CommentsByUserId(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = defaultCommentsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression(`"public"."comments"`, "id")), columns...),
		func(u *models.UsersDto) *[]models.CommentsDto {
			return &u.CommentsByUserId
		},
	)
}

func defaultCommentsByUserIdColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression(`"public"."comments"`, "id"),
		ast.NewIntColumnExpression(`"public"."comments"`, "post_id"),
		ast.NewUUIDColumnExpression(`"public"."comments"`, "user_id"),
		ast.NewStringColumnExpression(`"public"."comments"`, "content"),
		ast.NewBoolColumnExpression(`"public"."comments"`, "is_approved"),
		ast.NewTimestampColumnExpression(`"public"."comments"`, "created_at"),
		ast.NewDateColumnExpression(`"public"."comments"`, "test_date"),
	}
}

func (intoUsersDto) PostsByUserId(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = defaultPostsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("postsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression(`"public"."posts"`, "id")), columns...),
		func(u *models.UsersDto) *[]models.PostsDto {
			return &u.PostsByUserId
		},
	)
}

func defaultPostsByUserIdColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression(`"public"."posts"`, "id"),
		ast.NewUUIDColumnExpression(`"public"."posts"`, "user_id"),
		ast.NewStringColumnExpression(`"public"."posts"`, "title"),
		ast.NewStringColumnExpression(`"public"."posts"`, "content"),
		ast.NewStringColumnExpression(`"public"."posts"`, "status"),
		ast.NewTimestampColumnExpression(`"public"."posts"`, "published_at"),
		ast.NewIntColumnExpression(`"public"."posts"`, "view_count"),
		ast.NewTimestampColumnExpression(`"public"."posts"`, "created_at"),
		ast.NewTimestampColumnExpression(`"public"."posts"`, "updated_at"),
	}
}

func (f forUsersDto[T]) CommentsByUserId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultCommentsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression(`"public"."comments"`, "id")), columns...),
		func(t *T) *[]models.CommentsDto {
			dto := f.dest(t)
			return &dto.CommentsByUserId
		},
	)
}

func (f forUsersDto[T]) PostsByUserId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultPostsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("postsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression(`"public"."posts"`, "id")), columns...),
		func(t *T) *[]models.PostsDto {
			dto := f.dest(t)
			return &dto.PostsByUserId
		},
	)
}

type Alias[T any] struct {
	*ast.TableAlias
	forUsersDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.UsersDto] {
	return Alias[models.UsersDto]{
		TableAlias: Table().As(name, columns...),
		forUsersDto: forUsersDto[models.UsersDto]{
			dest: func(u *models.UsersDto) *models.UsersDto {
				return u
			},
		},
	}
}

func (a Alias[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	projection := newId(a.TableAlias.GetName(), func(t *T) **uuid.UUID {
		dto := a.dest(t)
		return &dto.Id
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.TableAlias.GetName()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Username() *ast.StringColumnProjection[T, *string] {
	projection := newUsername(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Username
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'username' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Email() *ast.StringColumnProjection[T, *string] {
	projection := newEmail(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Email
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'email' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) FullName() *ast.StringColumnProjection[T, **string] {
	projection := newFullName(a.TableAlias.GetName(), func(t *T) **string {
		dto := a.dest(t)
		return &dto.FullName
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'full_name' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newCreatedAt(a.TableAlias.GetName(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.CreatedAt
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.TableAlias.GetName()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newUpdatedAt(a.TableAlias.GetName(), func(t *T) **time.Time {
		dto := a.dest(t)
		return &dto.UpdatedAt
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", a.TableAlias.GetName()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) IsActive() *ast.BoolColumnProjection[T, **bool] {
	projection := newIsActive(a.TableAlias.GetName(), func(t *T) **bool {
		dto := a.dest(t)
		return &dto.IsActive
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_active' in alias %s", a.TableAlias.GetName()))
	projection.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Attributes() *ast.JsonColumnProjection[T, **json.RawMessage] {
	projection := newAttributes(a.TableAlias.GetName(), func(t *T) **json.RawMessage {
		dto := a.dest(t)
		return &dto.Attributes
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'attributes' in alias %s", a.TableAlias.GetName()))
	projection.JsonColumnExpression = ast.NewJsonColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
		a.Username(),
		a.Email(),
		a.FullName(),
		a.CreatedAt(),
		a.UpdatedAt(),
		a.IsActive(),
		a.Attributes(),
	}
}

func AliasFor[T any](alias *ast.TableAlias, dest func(*T) *models.UsersDto) Alias[T] {
	return Alias[T]{
		TableAlias:  alias,
		forUsersDto: forUsersDto[T]{dest: dest},
	}
}

func newId[T any](table string, ref func(*T) **uuid.UUID) *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		table,
		"id",
		ref,
	)
}

func newUsername[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"username",
		ref,
	)
}

func newEmail[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"email",
		ref,
	)
}

func newFullName[T any](table string, ref func(*T) **string) *ast.StringColumnProjection[T, **string] {
	return ast.NewStringColumnProjection(
		table,
		"full_name",
		ref,
	)
}

func newCreatedAt[T any](table string, ref func(*T) **time.Time) *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		table,
		"created_at",
		ref,
	)
}

func newUpdatedAt[T any](table string, ref func(*T) **time.Time) *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		table,
		"updated_at",
		ref,
	)
}

func newIsActive[T any](table string, ref func(*T) **bool) *ast.BoolColumnProjection[T, **bool] {
	return ast.NewBoolColumnProjection(
		table,
		"is_active",
		ref,
	)
}

func newAttributes[T any](table string, ref func(*T) **json.RawMessage) *ast.JsonColumnProjection[T, **json.RawMessage] {
	return ast.NewJsonColumnProjection(
		table,
		"attributes",
		ref,
	)
}
