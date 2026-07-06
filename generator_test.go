package gosqlgen

import (
	"testing"

	"github.com/jalphad/gosqlgen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerator_GeneratedPackagePath(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		packageRoot string
		packageName string
		want        string
	}{
		{
			name:        "joins package root and package name",
			packageRoot: "example.local/app",
			packageName: "generated",
			want:        "example.local/app/generated",
		},
		{
			name:        "cleans extra slash",
			packageRoot: "example.local/app/",
			packageName: "generated",
			want:        "example.local/app/generated",
		},
		{
			name:        "uses package name when root is empty",
			packageRoot: "",
			packageName: "generated",
			want:        "generated",
		},
		{
			name:        "uses root when package name is empty",
			packageRoot: "example.local/app",
			packageName: "",
			want:        "example.local/app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			generator := NewGenerator(parser.NewParser())
			generator.SetPackagePath(tt.packageRoot)
			generator.SetPackageName(tt.packageName)

			// Act
			got := generator.generatedPackagePath()

			// Assert
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGenerator_GenerateFilesUsesPackageRootAndPackageNameForImports(t *testing.T) {
	// Arrange
	p := parser.NewParser()
	err := p.Parse(`
		CREATE TABLE users (
			id UUID PRIMARY KEY,
			email TEXT NOT NULL
		);
	`)
	require.NoError(t, err)

	generator := NewGenerator(p)
	generator.SetPackagePath("example.local/app")
	generator.SetPackageName("accounts")

	// Act
	files, err := generator.GenerateFiles()
	require.NoError(t, err)

	// Assert
	dbFile := requireGeneratedContent(t, files, "db.gen.go")
	assert.Contains(t, dbFile, "package accounts")
	assert.Contains(t, dbFile, `"example.local/app/accounts/models"`)
	assert.NotContains(t, dbFile, `"example.local/app/models"`)

	modelFile := requireGeneratedContent(t, files, "models/users.gen.go")
	assert.Contains(t, modelFile, `"example.local/app/accounts/query/ast"`)

	queryFile := requireGeneratedContent(t, files, "query/users/users.go")
	assert.Contains(t, queryFile, `"example.local/app/accounts/models"`)
	assert.Contains(t, queryFile, `"example.local/app/accounts/query/builder"`)

	builderFile := requireGeneratedContent(t, files, "query/builder/builder.gen.go")
	assert.Contains(t, builderFile, `"example.local/app/accounts/query/ast"`)
}

func TestGenerator_GenerateFilesSkipsUpdateQueriesForPrimaryKeyOnlyTables(t *testing.T) {
	// Arrange
	p := parser.NewParser()
	err := p.Parse(`
		CREATE TABLE post_tags (
			post_id INTEGER NOT NULL,
			tag_id INTEGER NOT NULL,
			PRIMARY KEY (post_id, tag_id)
		);
	`)
	require.NoError(t, err)

	generator := NewGenerator(p)
	generator.SetPackagePath("example.local/app")
	generator.SetPackageName("blog")

	// Act
	files, err := generator.GenerateFiles()
	require.NoError(t, err)

	// Assert
	queryFile := requireGeneratedContent(t, files, "query/queries.go")
	assert.Contains(t, queryFile, "func PostTagsDtoInsertOne")
	assert.Contains(t, queryFile, "func PostTagsDtoSelectOne")
	assert.Contains(t, queryFile, "func PostTagsDtoSelectMany")
	assert.NotContains(t, queryFile, "func PostTagsDtoUpdateOne")
	assert.NotContains(t, queryFile, "func PostTagsDtoUpdateMany")
}

func requireGeneratedContent(t *testing.T, files map[string]string, filename string) string {
	t.Helper()

	content, ok := files[filename]
	require.Truef(t, ok, "Generated file %q not found", filename)

	return content
}
