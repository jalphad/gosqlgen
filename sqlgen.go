package gosqlgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SQLGen is the main interface for the SQL DTO generator
type SQLGen struct {
	parser      *Parser
	generator   *Generator
	packageName string
	outputPath  string
}

// New creates a new SQLGen instance
func New() *SQLGen {
	parser := NewParser()
	generator := NewGenerator(parser)

	return &SQLGen{
		parser:      parser,
		generator:   generator,
		packageName: "models",
		outputPath:  "./models",
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

// ParseFile parses SQL from a file
func (s *SQLGen) ParseFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return s.Parse(string(content))
}

// Parse parses SQL statements
func (s *SQLGen) Parse(sql string) error {
	return s.parser.Parse(sql)
}

// Generate generates Go code from parsed SQL
func (s *SQLGen) Generate() (string, error) {
	return s.generator.Generate()
}

// GenerateToFile generates Go code and writes it to a file
func (s *SQLGen) GenerateToFile(filename string) error {
	code, err := s.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Create output directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write generated code to file
	if err := ioutil.WriteFile(filename, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// GetTables returns all parsed tables
func (s *SQLGen) GetTables() map[string]*Table {
	return s.parser.GetTables()
}

// GetTable returns a specific parsed table
func (s *SQLGen) GetTable(name string) (*Table, bool) {
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
