package gosqlgen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jalphad/gosqlgen/parser"
)

// SQLGen is the main interface for the SQL DTO generator
type SQLGen struct {
	parser      *parser.Parser
	generator   *Generator
	packageName string
	outputPath  string
	packagePath string
}

// New creates a new SQLGen instance
func New() *SQLGen {
	p := parser.NewParser()
	generator := NewGenerator(p)

	return &SQLGen{
		parser:      p,
		generator:   generator,
		packageName: "models",
		outputPath:  "./models",
		packagePath: "example.local/example",
	}
}

// WithPackageName sets the package name for generated code
func (s *SQLGen) WithPackageName(name string) *SQLGen {
	s.packageName = name
	s.generator.SetPackageName(name)
	return s
}

// WithOutputPath sets the output path for generated files
func (s *SQLGen) WithOutputPath(path string) *SQLGen {
	s.outputPath = path
	return s
}

// WithPackagePath sets the package path for generated imports
func (s *SQLGen) WithPackagePath(path string) *SQLGen {
	s.packagePath = path
	s.generator.SetPackagePath(path)
	return s
}

// ParseFile parses SQL from a file
func (s *SQLGen) ParseFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return s.Parse(string(content))
}

// Parse parses SQL statements
func (s *SQLGen) Parse(sql string) error {
	return s.parser.Parse(sql)
}

// GenerateFiles generates Go code from parsed SQL as separate files
func (s *SQLGen) GenerateFiles() error {
	code, err := s.generator.GenerateFiles()
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	for filename, content := range code {
		// Create output directory if it doesn't exist
		dir := filepath.Dir(filepath.Join(s.outputPath, filename))
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Write generated code to file
		if err := os.WriteFile(filepath.Join(s.outputPath, filename), []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
	}

	return nil
}

// GetTables returns all parsed tables
func (s *SQLGen) GetTables() map[string]*parser.Table {
	return s.parser.GetTables()
}

// GetTable returns a specific parsed table
func (s *SQLGen) GetTable(name string) (*parser.Table, bool) {
	return s.parser.GetTable(name)
}

// Config holds configuration for the generator
type Config struct {
	PackageName    string
	OutputPath     string
	UseNullTypes   bool
	UseGORM        bool
	UseSquirrel    bool
	DatabaseDriver string // postgres, mysql, sqlite
}

// NewWithConfig creates a new SQLGen with configuration
func NewWithConfig(config Config) *SQLGen {
	gen := New()

	if config.PackageName != "" {
		gen.WithPackageName(config.PackageName)
	}

	if config.OutputPath != "" {
		gen.WithOutputPath(config.OutputPath)
	}

	// Additional configuration can be applied here

	return gen
}
