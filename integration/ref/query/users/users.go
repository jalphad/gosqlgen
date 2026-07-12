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

// NewQuery returns a query builder for users
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.UsersDto] {
	return builder.NewKnownTableBuilder[models.UsersDto](pool, Table())
}

var Into = intoUsersDto{}

type intoUsersDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("users")
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return newId("users", func(u *models.UsersDto) **uuid.UUID {
		return &u.Id
	})
}

func Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	return newUsername("users", func(u *models.UsersDto) *string {
		return &u.Username
	})
}

func Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	return newEmail("users", func(u *models.UsersDto) *string {
		return &u.Email
	})
}

func FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	return newFullName("users", func(u *models.UsersDto) **string {
		return &u.FullName
	})
}

func CreatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return newCreatedAt("users", func(u *models.UsersDto) **time.Time {
		return &u.CreatedAt
	})
}

func UpdatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return newUpdatedAt("users", func(u *models.UsersDto) **time.Time {
		return &u.UpdatedAt
	})
}

func IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	return newIsActive("users", func(u *models.UsersDto) **bool {
		return &u.IsActive
	})
}

func Attributes() *ast.JsonColumnProjection[models.UsersDto, **json.RawMessage] {
	return newAttributes("users", func(u *models.UsersDto) **json.RawMessage {
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

type forUsersDto[E models.ExportsUsersDto[T], T any] struct{}

func For[E models.ExportsUsersDto[T], T any]() forUsersDto[E, T] {
	return forUsersDto[E, T]{}
}

func (f forUsersDto[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return newId("users", func(t *T) **uuid.UUID {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Id
	})
}

func (f forUsersDto[E, T]) Username() *ast.StringColumnProjection[T, *string] {
	return newUsername("users", func(t *T) *string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Username
	})
}

func (f forUsersDto[E, T]) Email() *ast.StringColumnProjection[T, *string] {
	return newEmail("users", func(t *T) *string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Email
	})
}

func (f forUsersDto[E, T]) FullName() *ast.StringColumnProjection[T, **string] {
	return newFullName("users", func(t *T) **string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.FullName
	})
}

func (f forUsersDto[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newCreatedAt("users", func(t *T) **time.Time {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.CreatedAt
	})
}

func (f forUsersDto[E, T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return newUpdatedAt("users", func(t *T) **time.Time {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.UpdatedAt
	})
}

func (f forUsersDto[E, T]) IsActive() *ast.BoolColumnProjection[T, **bool] {
	return newIsActive("users", func(t *T) **bool {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.IsActive
	})
}

func (f forUsersDto[E, T]) Attributes() *ast.JsonColumnProjection[T, **json.RawMessage] {
	return newAttributes("users", func(t *T) **json.RawMessage {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Attributes
	})
}

func (f forUsersDto[E, T]) AllColumns() []ast.Projection[T] {
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
		expr.JsonAggObject("commentsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(u *models.UsersDto) *[]models.CommentsDto {
			return &u.CommentsByUserId
		},
	)
}

func defaultCommentsByUserIdColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("comments", "id"),
		ast.NewIntColumnExpression("comments", "post_id"),
		ast.NewUUIDColumnExpression("comments", "user_id"),
		ast.NewStringColumnExpression("comments", "content"),
		ast.NewBoolColumnExpression("comments", "is_approved"),
		ast.NewTimestampColumnExpression("comments", "created_at"),
		ast.NewDateColumnExpression("comments", "test_date"),
	}
}

func (intoUsersDto) PostsByUserId(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = defaultPostsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("postsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(u *models.UsersDto) *[]models.PostsDto {
			return &u.PostsByUserId
		},
	)
}

func defaultPostsByUserIdColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		ast.NewIntColumnExpression("posts", "id"),
		ast.NewUUIDColumnExpression("posts", "user_id"),
		ast.NewStringColumnExpression("posts", "title"),
		ast.NewStringColumnExpression("posts", "content"),
		ast.NewStringColumnExpression("posts", "status"),
		ast.NewTimestampColumnExpression("posts", "published_at"),
		ast.NewIntColumnExpression("posts", "view_count"),
		ast.NewTimestampColumnExpression("posts", "created_at"),
		ast.NewTimestampColumnExpression("posts", "updated_at"),
	}
}

func (forUsersDto[E, T]) CommentsByUserId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultCommentsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("commentsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(t *T) *[]models.CommentsDto {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.CommentsByUserId
		},
	)
}

func (forUsersDto[E, T]) PostsByUserId(columns ...ast.NamedExpression) ast.Projection[T] {
	if len(columns) == 0 {
		columns = defaultPostsByUserIdColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("postsbyuserid", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(t *T) *[]models.PostsDto {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.PostsByUserId
		},
	)
}

type Alias[E models.ExportsUsersDto[T], T any] struct {
	*ast.Alias
	forUsersDto[E, T]
}

func As[E models.ExportsUsersDto[T], T any](name string, columns ...ast.NamedExpression) Alias[E, T] {
	return Alias[E, T]{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	projection := newId(a.Alias.Name(), func(t *T) **uuid.UUID {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Id
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) Username() *ast.StringColumnProjection[T, *string] {
	projection := newUsername(a.Alias.Name(), func(t *T) *string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Username
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'username' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) Email() *ast.StringColumnProjection[T, *string] {
	projection := newEmail(a.Alias.Name(), func(t *T) *string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Email
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'email' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) FullName() *ast.StringColumnProjection[T, **string] {
	projection := newFullName(a.Alias.Name(), func(t *T) **string {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.FullName
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'full_name' in alias %s", a.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newCreatedAt(a.Alias.Name(), func(t *T) **time.Time {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.CreatedAt
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	projection := newUpdatedAt(a.Alias.Name(), func(t *T) **time.Time {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.UpdatedAt
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", a.Alias.Name()))
	projection.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) IsActive() *ast.BoolColumnProjection[T, **bool] {
	projection := newIsActive(a.Alias.Name(), func(t *T) **bool {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.IsActive
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_active' in alias %s", a.Alias.Name()))
	projection.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
}

func (a *Alias[E, T]) Attributes() *ast.JsonColumnProjection[T, **json.RawMessage] {
	projection := newAttributes(a.Alias.Name(), func(t *T) **json.RawMessage {
		e := E(t)
		dto := e.GetUsersDto()
		return &dto.Attributes
	})
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == projection.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'attributes' in alias %s", a.Alias.Name()))
	projection.JsonColumnExpression = ast.NewJsonColumnExpressionFromExpr(projection.Name(), errExpr)
	return projection
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
