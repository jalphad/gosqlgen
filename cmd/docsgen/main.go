package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type example struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Kind           string   `json:"kind"`
	Source         string   `json:"source"`
	ExpectedSQL    string   `json:"expectedSQL"`
	ExpectedParams []any    `json:"expectedParams"`
	Schema         string   `json:"schema"`
	PackageName    string   `json:"packageName"`
	OutputPath     string   `json:"outputPath"`
	PackagePath    string   `json:"packagePath"`
	Command        string   `json:"command"`
	Program        string   `json:"generatorProgram"`
	ExpectedFiles  []string `json:"expectedFiles"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root, err := repoRoot()
	if err != nil {
		return err
	}

	examples, err := loadExamples(filepath.Join(root, "docs-src", "examples.cue"))
	if err != nil {
		return err
	}
	if err := validateExampleMetadata(root, examples); err != nil {
		return err
	}
	if err := generateExampleTests(root, examples); err != nil {
		return err
	}
	if err := updateTutorial(root, examples); err != nil {
		return err
	}
	if err := validateExampleSources(root, examples); err != nil {
		return err
	}

	output := filepath.Join(root, "docs", "generated", "examples.md")
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return err
	}
	return os.WriteFile(output, renderExamples(examples), 0o644)
}

func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find go.mod")
		}
		dir = parent
	}
}

func loadExamples(path string) ([]example, error) {
	cmd := exec.Command("cue", "export", path, "-e", "examples", "--out", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("cue export %s: %w\n%s", path, err, strings.TrimSpace(string(output)))
	}

	var examples []example
	if err := json.Unmarshal(output, &examples); err != nil {
		return nil, err
	}
	return examples, nil
}

func validateExampleMetadata(root string, examples []example) error {
	seen := make(map[string]struct{}, len(examples))
	for _, example := range examples {
		if _, ok := seen[example.ID]; ok {
			return fmt.Errorf("duplicate docs example ID %q", example.ID)
		}
		seen[example.ID] = struct{}{}

		if !strings.HasPrefix(example.Source, "docs/examples/") {
			return fmt.Errorf("docs example %q source must be under docs/examples", example.ID)
		}
		switch example.Kind {
		case "query":
			if example.ExpectedSQL == "" {
				return fmt.Errorf("docs query example %q requires expectedSQL", example.ID)
			}
		case "generation":
			if example.Schema == "" {
				return fmt.Errorf("docs generation example %q requires schema", example.ID)
			}
			if _, err := os.Stat(filepath.Join(root, example.Schema)); err != nil {
				return fmt.Errorf("docs generation example %q schema %q: %w", example.ID, example.Schema, err)
			}
			if example.PackageName == "" || example.OutputPath == "" || example.PackagePath == "" {
				return fmt.Errorf("docs generation example %q requires packageName, outputPath, and packagePath", example.ID)
			}
			if example.Command == "" || example.Program == "" {
				return fmt.Errorf("docs generation example %q requires command and generatorProgram", example.ID)
			}
			if len(example.ExpectedFiles) == 0 {
				return fmt.Errorf("docs generation example %q requires expectedFiles", example.ID)
			}
		default:
			return fmt.Errorf("docs example %q has unknown kind %q", example.ID, example.Kind)
		}
	}
	return nil
}

func validateExampleSources(root string, examples []example) error {
	for _, example := range examples {
		if _, err := os.Stat(filepath.Join(root, example.Source)); err != nil {
			return fmt.Errorf("docs example %q source %q: %w", example.ID, example.Source, err)
		}
	}
	return nil
}

func generateExampleTests(root string, examples []example) error {
	for _, example := range examples {
		if example.Kind != "generation" {
			continue
		}
		output := filepath.Join(root, example.Source)
		if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(output, renderGenerationTest(example), 0o644); err != nil {
			return err
		}
		cmd := exec.Command("gofmt", "-w", output)
		if data, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("gofmt %s: %w\n%s", output, err, strings.TrimSpace(string(data)))
		}
	}
	return nil
}

func renderGenerationTest(example example) []byte {
	var buf bytes.Buffer
	buf.WriteString("package examples\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\t\"path/filepath\"\n")
	buf.WriteString("\t\"testing\"\n\n")
	buf.WriteString("\t\"github.com/jalphad/gosqlgen\"\n")
	buf.WriteString("\t\"github.com/stretchr/testify/require\"\n")
	buf.WriteString(")\n\n")
	fmt.Fprintf(&buf, "func Test%s(t *testing.T) {\n", toPascal(example.ID))
	buf.WriteString("\t// Arrange\n")
	buf.WriteString("\toutputPath := t.TempDir()\n")
	fmt.Fprintf(&buf, "\tschemaPath := filepath.Join(%s)\n", quotedPathParts("..", "..", example.Schema))
	buf.WriteString("\tgen := gosqlgen.New().\n")
	fmt.Fprintf(&buf, "\t\tWithPackageName(%s).\n", strconv.Quote(example.PackageName))
	buf.WriteString("\t\tWithOutputPath(outputPath).\n")
	fmt.Fprintf(&buf, "\t\tWithPackagePath(%s)\n\n", strconv.Quote(example.PackagePath))
	buf.WriteString("\t// Act\n")
	buf.WriteString("\terr := gen.ParseFile(schemaPath)\n\n")
	buf.WriteString("\t// Assert\n")
	buf.WriteString("\trequire.NoError(t, err)\n")
	buf.WriteString("\trequire.NoError(t, gen.GenerateFiles())\n")
	for _, file := range example.ExpectedFiles {
		fmt.Fprintf(&buf, "\trequire.FileExists(t, filepath.Join(outputPath, %s))\n", quotedPathParts(file))
	}
	buf.WriteString("}\n")
	return buf.Bytes()
}

func updateTutorial(root string, examples []example) error {
	var generation *example
	for i := range examples {
		if examples[i].Kind == "generation" && examples[i].ID == "getting-started-generate-from-schema" {
			generation = &examples[i]
			break
		}
	}
	if generation == nil {
		return nil
	}

	path := filepath.Join(root, "docs", "tutorials", "getting-started.md")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	updated, err := replaceMarkedSection(
		string(data),
		"<!-- docsgen:generation:start -->",
		"<!-- docsgen:generation:end -->",
		renderGenerationTutorial(*generation),
	)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(updated), 0o644)
}

func renderGenerationTutorial(example example) string {
	var buf bytes.Buffer
	buf.WriteString("Add a small generator command in your application, for example `cmd/gen/main.go`:\n\n")
	buf.WriteString("```go\n")
	buf.WriteString(strings.TrimSpace(example.Program))
	buf.WriteString("\n```\n\n")
	buf.WriteString("Run it:\n\n")
	buf.WriteString("```sh\n")
	buf.WriteString(example.Command)
	buf.WriteString("\n```\n\n")
	buf.WriteString("The generator writes these files:\n\n")
	files := append([]string(nil), example.ExpectedFiles...)
	sort.Strings(files)
	for _, file := range files {
		fmt.Fprintf(&buf, "- `%s`\n", file)
	}
	buf.WriteString("\n")
	fmt.Fprintf(
		&buf,
		"This generation example is defined in `docs-src/examples.cue` and tested by [`%s`](../examples/%s).\n",
		example.Source,
		filepath.Base(example.Source),
	)
	return buf.String()
}

func replaceMarkedSection(content, startMarker, endMarker, replacement string) (string, error) {
	start := strings.Index(content, startMarker)
	if start < 0 {
		return "", fmt.Errorf("missing marker %s", startMarker)
	}
	end := strings.Index(content, endMarker)
	if end < 0 {
		return "", fmt.Errorf("missing marker %s", endMarker)
	}
	if end < start {
		return "", fmt.Errorf("marker %s appears before %s", endMarker, startMarker)
	}

	var buf strings.Builder
	buf.WriteString(content[:start+len(startMarker)])
	buf.WriteString("\n")
	buf.WriteString(strings.TrimRight(replacement, "\n"))
	buf.WriteString("\n")
	buf.WriteString(content[end:])
	return buf.String(), nil
}

func quotedPathParts(parts ...string) string {
	if len(parts) == 1 {
		return strings.Join(quoteSplitPath(filepath.ToSlash(parts[0])), ", ")
	}
	quoted := make([]string, 0, len(parts))
	for _, part := range parts {
		quoted = append(quoted, quoteSplitPath(filepath.ToSlash(part))...)
	}
	return strings.Join(quoted, ", ")
}

func quoteSplitPath(path string) []string {
	parts := strings.Split(path, "/")
	quoted := make([]string, 0, len(parts))
	for _, part := range parts {
		if part == "" {
			continue
		}
		quoted = append(quoted, strconv.Quote(part))
	}
	return quoted
}

func toPascal(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_' || r == ' '
	})
	var buf strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		buf.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			buf.WriteString(part[1:])
		}
	}
	return buf.String()
}

func renderExamples(examples []example) []byte {
	var buf bytes.Buffer
	buf.WriteString("# Examples\n\n")
	buf.WriteString("This page is generated by `go run ./cmd/docsgen`.\n\n")
	buf.WriteString("| ID | Title | Source |\n")
	buf.WriteString("| --- | --- | --- |\n")
	for _, example := range examples {
		source := strings.TrimPrefix(example.Source, "docs/")
		fmt.Fprintf(
			&buf,
			"| `%s` | %s | [`%s`](../%s) |\n",
			example.ID,
			example.Title,
			example.Source,
			source,
		)
	}
	return buf.Bytes()
}
