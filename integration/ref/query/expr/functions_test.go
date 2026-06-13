package expr

import (
	"testing"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
	"github.com/stretchr/testify/require"
)

func TestInlineStringLiteral(t *testing.T) {
	params := []any{}

	sql, err := ast.RenderWithContext(InlineStringLiteral("[]"), &params, &ast.QueryContext{})

	require.NoError(t, err)
	require.Equal(t, "'[]'", sql)
	require.Empty(t, params)
}

func TestInlineStringLiteralEscapesQuotes(t *testing.T) {
	params := []any{}

	sql, err := ast.RenderWithContext(InlineStringLiteral("it's"), &params, &ast.QueryContext{})

	require.NoError(t, err)
	require.Equal(t, "'it''s'", sql)
	require.Empty(t, params)
}

func TestLiteralExpressionUsesBindParameter(t *testing.T) {
	params := []any{}

	sql, err := ast.RenderWithContext(ast.NewLiteralExpression("[]"), &params, &ast.QueryContext{})

	require.NoError(t, err)
	require.Equal(t, "$1", sql)
	require.Equal(t, []any{"[]"}, params)
}

func TestNull(t *testing.T) {
	params := []any{}

	sql, err := ast.RenderWithContext(Null(), &params, &ast.QueryContext{})

	require.NoError(t, err)
	require.Equal(t, "NULL", sql)
	require.Empty(t, params)
}

func TestNullPredicates(t *testing.T) {
	params := []any{}

	isNullSQL, err := ast.RenderWithContext(IsNull(ast.NewStringColumnExpression("users", "email")), &params, &ast.QueryContext{})
	require.NoError(t, err)
	require.Equal(t, "users.email IS NULL", isNullSQL)

	params = []any{}
	isNotNullSQL, err := ast.RenderWithContext(IsNotNull(ast.NewStringColumnExpression("users", "email")), &params, &ast.QueryContext{})
	require.NoError(t, err)
	require.Equal(t, "users.email IS NOT NULL", isNotNullSQL)
}

func TestCast(t *testing.T) {
	params := []any{}

	sql, err := ast.RenderWithContext(Cast(ast.SetType[string](InlineStringLiteral("[]"))).AsJson(), &params, &ast.QueryContext{})

	require.NoError(t, err)
	require.Equal(t, "CAST('[]' AS json)", sql)
	require.Empty(t, params)
}
