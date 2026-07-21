package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJoinRendering(t *testing.T) {
	tests := []struct {
		name       string
		join       *JoinExpr
		wantSQL    string
		wantParams []any
	}{
		{
			name:       "inner join on",
			join:       InnerJoin(NewTableSource("posts"), On(NewSQLType(true))),
			wantSQL:    "users INNER JOIN posts ON $1",
			wantParams: []any{true},
		},
		{
			name:    "left join using",
			join:    LeftJoin(NewTableSource("posts"), Using(NewIntColumnExpression("", "id"), NewIntColumnExpression("", "tenant_id"))),
			wantSQL: "users LEFT JOIN posts USING (id, tenant_id)",
		},
		{
			name:    "right natural join",
			join:    RightJoin(NewTableSource("posts"), Natural()),
			wantSQL: "users NATURAL RIGHT JOIN posts",
		},
		{
			name:       "full join on",
			join:       FullJoin(NewTableSource("posts"), On(NewSQLType(false))),
			wantSQL:    "users FULL JOIN posts ON $1",
			wantParams: []any{false},
		},
		{
			name:    "cross join",
			join:    CrossJoin(NewTableSource("posts")),
			wantSQL: "users CROSS JOIN posts",
		},
		{
			name:    "cross join lateral",
			join:    CrossJoin(Lateral(NewTableSource("posts"))),
			wantSQL: "users CROSS JOIN LATERAL posts",
		},
		{
			name:       "left join lateral",
			join:       LeftJoin(Lateral(NewTableSource("posts")), On(NewSQLType(true))),
			wantSQL:    "users LEFT JOIN LATERAL posts ON $1",
			wantParams: []any{true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			table := NewTableSource("users").Join(tt.join)
			params := []any{}
			ctx := &QueryContext{}

			// Act
			sql, err := RenderWithContext(table, &params, ctx)

			// Assert
			require.NoError(t, err)
			assert.Equal(t, tt.wantSQL, sql)
			assert.Equal(t, tt.wantParams, params)
			assert.Equal(t, tt.join.joinType, ctx.JoinedTables["posts"])
			assert.Contains(t, ctx.AllTables, "posts")
		})
	}
}

func TestJoinRenderingErrors(t *testing.T) {
	tests := []struct {
		name string
		join *JoinExpr
	}{
		{name: "nil join", join: nil},
		{name: "missing right table", join: LeftJoin(nil, On(NewSQLType(true)))},
		{name: "missing qualifier", join: LeftJoin(NewTableSource("posts"), nil)},
		{name: "missing ON condition", join: LeftJoin(NewTableSource("posts"), On(nil))},
		{name: "empty USING columns", join: LeftJoin(NewTableSource("posts"), Using())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			table := NewTableSource("users").Join(tt.join)
			params := []any{}

			// Act
			_, err := RenderWithContext(table, &params, &QueryContext{})

			// Assert
			require.Error(t, err)
		})
	}
}
