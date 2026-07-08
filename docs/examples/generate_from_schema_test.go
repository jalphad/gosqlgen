package examples

import (
	"path/filepath"
	"testing"

	"github.com/jalphad/gosqlgen"
	"github.com/stretchr/testify/require"
)

func TestGettingStartedGenerateFromSchema(t *testing.T) {
	// Arrange
	outputPath := t.TempDir()
	schemaPath := filepath.Join("..", "..", "integration", "schema.sql")
	gen := gosqlgen.New().
		WithPackageName("db").
		WithOutputPath(outputPath).
		WithPackagePath("example.com/myapp")

	// Act
	err := gen.ParseFile(schemaPath)

	// Assert
	require.NoError(t, err)
	require.NoError(t, gen.GenerateFiles())
	require.FileExists(t, filepath.Join(outputPath, "db", "db.gen.go"))
	require.FileExists(t, filepath.Join(outputPath, "db", "models", "users.gen.go"))
	require.FileExists(t, filepath.Join(outputPath, "db", "query", "users", "users.go"))
	require.FileExists(t, filepath.Join(outputPath, "db", "query", "helpers.gen.go"))
}
