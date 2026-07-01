package users

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/jalphad/gosqlgen/integration/output/models"
	"github.com/jalphad/gosqlgen/integration/output/query/ast"
	"github.com/jalphad/gosqlgen/integration/output/query/expr"
)

var Into = intoUsersDto{}

type intoUsersDto struct{}

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

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Username(),
		Email(),
		FullName(),
		CreatedAt(),
		UpdatedAt(),
		IsActive(),
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

func (intoUsersDto) AllColumns() []ast.Projection[models.UsersDto] {
	return []ast.Projection[models.UsersDto]{
		Into.Id(),
		Into.Username(),
		Into.Email(),
		Into.FullName(),
		Into.CreatedAt(),
		Into.UpdatedAt(),
		Into.IsActive(),
	}
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

func (intoUsersDto) Comments(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = defaultCommentsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("comments", expr.IsNotNull(ast.NewIntColumnExpression("comments", "id")), columns...),
		func(u *models.UsersDto) *[]models.CommentsDto {
			return &u.Comments
		},
	)
}

func defaultCommentsColumns() []ast.NamedExpression {
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

func (intoUsersDto) Posts(columns ...ast.NamedExpression) ast.Projection[models.UsersDto] {
	if len(columns) == 0 {
		columns = defaultPostsColumns()
	}
	return ast.NewJSONProjection(
		expr.JsonAggObject("posts", expr.IsNotNull(ast.NewIntColumnExpression("posts", "id")), columns...),
		func(u *models.UsersDto) *[]models.PostsDto {
			return &u.Posts
		},
	)
}

func defaultPostsColumns() []ast.NamedExpression {
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
