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

type forUsersDto[E models.ExportsUsersDto[T], T any] struct{}

func For[E models.ExportsUsersDto[T], T any]() forUsersDto[E, T] {
	return forUsersDto[E, T]{}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("users")
}

func Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"users",
		"id",
		func(u *models.UsersDto) **uuid.UUID {
			return &u.Id
		},
	)
}

func Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"username",
		func(u *models.UsersDto) *string {
			return &u.Username
		},
	)
}

func Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"email",
		func(u *models.UsersDto) *string {
			return &u.Email
		},
	)
}

func FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	return ast.NewStringColumnProjection(
		"users",
		"full_name",
		func(u *models.UsersDto) **string {
			return &u.FullName
		},
	)
}

func CreatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"users",
		"created_at",
		func(u *models.UsersDto) **time.Time {
			return &u.CreatedAt
		},
	)
}

func UpdatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"users",
		"updated_at",
		func(u *models.UsersDto) **time.Time {
			return &u.UpdatedAt
		},
	)
}

func IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	return ast.NewBoolColumnProjection(
		"users",
		"is_active",
		func(u *models.UsersDto) **bool {
			return &u.IsActive
		},
	)
}

func Attributes() *ast.JsonColumnProjection[models.UsersDto, **json.RawMessage] {
	return ast.NewJsonColumnProjection(
		"users",
		"attributes",
		func(u *models.UsersDto) **json.RawMessage {
			return &u.Attributes
		},
	)
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

func (forUsersDto[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"users",
		"id",
		func(t *T) **uuid.UUID {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Id
		},
	)
}

func (forUsersDto[E, T]) Username() *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"username",
		func(t *T) *string {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Username
		},
	)
}

func (forUsersDto[E, T]) Email() *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		"users",
		"email",
		func(t *T) *string {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Email
		},
	)
}

func (forUsersDto[E, T]) FullName() *ast.StringColumnProjection[T, **string] {
	return ast.NewStringColumnProjection(
		"users",
		"full_name",
		func(t *T) **string {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.FullName
		},
	)
}

func (forUsersDto[E, T]) CreatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"users",
		"created_at",
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.CreatedAt
		},
	)
}

func (forUsersDto[E, T]) UpdatedAt() *ast.TimestampColumnProjection[T, **time.Time] {
	return ast.NewTimestampColumnProjection(
		"users",
		"updated_at",
		func(t *T) **time.Time {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.UpdatedAt
		},
	)
}

func (forUsersDto[E, T]) IsActive() *ast.BoolColumnProjection[T, **bool] {
	return ast.NewBoolColumnProjection(
		"users",
		"is_active",
		func(t *T) **bool {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.IsActive
		},
	)
}

func (forUsersDto[E, T]) Attributes() *ast.JsonColumnProjection[T, **json.RawMessage] {
	return ast.NewJsonColumnProjection(
		"users",
		"attributes",
		func(t *T) **json.RawMessage {
			e := E(t)
			dto := e.GetUsersDto()
			return &dto.Attributes
		},
	)
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

type Alias struct {
	*ast.Alias
}

func As(name string, columns ...ast.NamedExpression) *Alias {
	return &Alias{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias) Id() *ast.UUIDColumnProjection[models.UsersDto, **uuid.UUID] {
	column := Id()
	alias := ast.NewUUIDColumnProjection[models.UsersDto, **uuid.UUID](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **uuid.UUID {
			return &u.Id
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	alias.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Username() *ast.StringColumnProjection[models.UsersDto, *string] {
	column := Username()
	alias := ast.NewStringColumnProjection[models.UsersDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) *string {
			return &u.Username
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'username' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Email() *ast.StringColumnProjection[models.UsersDto, *string] {
	column := Email()
	alias := ast.NewStringColumnProjection[models.UsersDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) *string {
			return &u.Email
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'email' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) FullName() *ast.StringColumnProjection[models.UsersDto, **string] {
	column := FullName()
	alias := ast.NewStringColumnProjection[models.UsersDto, **string](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **string {
			return &u.FullName
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'full_name' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) CreatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	column := CreatedAt()
	alias := ast.NewTimestampColumnProjection[models.UsersDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **time.Time {
			return &u.CreatedAt
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'created_at' in alias %s", a.Alias.Name()))
	alias.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) UpdatedAt() *ast.TimestampColumnProjection[models.UsersDto, **time.Time] {
	column := UpdatedAt()
	alias := ast.NewTimestampColumnProjection[models.UsersDto, **time.Time](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **time.Time {
			return &u.UpdatedAt
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'updated_at' in alias %s", a.Alias.Name()))
	alias.TimestampColumnExpression = ast.NewTimestampColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) IsActive() *ast.BoolColumnProjection[models.UsersDto, **bool] {
	column := IsActive()
	alias := ast.NewBoolColumnProjection[models.UsersDto, **bool](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **bool {
			return &u.IsActive
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'is_active' in alias %s", a.Alias.Name()))
	alias.BoolColumnExpression = ast.NewBoolColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Attributes() *ast.JsonColumnProjection[models.UsersDto, **json.RawMessage] {
	column := Attributes()
	alias := ast.NewJsonColumnProjection[models.UsersDto, **json.RawMessage](
		a.Alias.Name(),
		column.Name(),
		func(u *models.UsersDto) **json.RawMessage {
			return &u.Attributes
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'attributes' in alias %s", a.Alias.Name()))
	alias.JsonColumnExpression = ast.NewJsonColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
