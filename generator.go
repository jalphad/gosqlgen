package gosqlgen

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/jalphad/gosqlgen/templates"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

// Generator handles code generation with type-safe queries
type Generator struct {
	parser      *Parser
	packageName string
	imports     map[string]bool
}

// NewGenerator creates a new code generator
func NewGenerator(parser *Parser) *Generator {
	return &Generator{
		parser:      parser,
		packageName: "models",
		imports: map[string]bool{
			"context":                          true,
			"fmt":                              true,
			"time":                             true,
			"strings":                          true,
			"github.com/jackc/pgx/v5":          true,
			"github.com/jackc/pgx/v5/pgxpool":   true,
			"github.com/jackc/pgx/v5/pgconn":   true,
		},
	}
}

// SetPackageName sets the package name for generated code
func (g *Generator) SetPackageName(name string) {
	g.packageName = name
}

// GenerateFiles generates Go code for all parsed tables as separate files
// Returns a map of filename to file content
func (g *Generator) GenerateFiles() (map[string]string, error) {
	files := make(map[string]string)

	// Generate common.gen.go with utility types
	commonBuf := bytes.Buffer{}
	fmt.Fprintf(&commonBuf, "package %s\n\n", g.packageName)
	g.writeImports(&commonBuf)
	if err := g.generateCommonTypes(&commonBuf); err != nil {
		return nil, err
	}

	formatted, err := g.format("common.gen.go", commonBuf.Bytes())
	if err != nil {
		return nil, err
	}
	files["common.gen.go"] = string(formatted)

	// Generate a file for each table
	for _, table := range g.parser.GetTables() {
		tableBuf := bytes.Buffer{}

		// Package declaration
		fmt.Fprintf(&tableBuf, "package %s\n\n", g.packageName)

		// Imports for table file
		g.writeImports(&tableBuf)

		// Table struct
		if err := g.generateTableStruct(&tableBuf, table); err != nil {
			return nil, err
		}

		// Field references
		if err := g.generateFieldReferences(&tableBuf, table); err != nil {
			return nil, err
		}

		// Query builder
		if err := g.generateTypeSafeQueryBuilder(&tableBuf, table); err != nil {
			return nil, err
		}

		// Join builders
		if err := g.generateJoinBuilders(&tableBuf, table); err != nil {
			return nil, err
		}

		// Format the generated code
		filename := fmt.Sprintf("%s.gen.go", table.Name)
		formatted, err = g.format(filename, tableBuf.Bytes())
		if err != nil {
			return nil, err
		}
		files[filename] = string(formatted)
	}

	// Generate db.gen.go for database wrapper
	dbBuf := bytes.Buffer{}
	fmt.Fprintf(&dbBuf, "package %s\n\n", g.packageName)
	g.writeImports(&dbBuf)
	if err := g.generateDatabaseWrapper(&dbBuf); err != nil {
		return nil, err
	}

	formatted, err = g.format("db.gen.go", dbBuf.Bytes())
	if err != nil {
		return nil, err
	}
	files["db.gen.go"] = string(formatted)

	return files, nil
}

// Generate generates Go code for all parsed tables (backward compatibility)
// Returns all code in a single string
func (g *Generator) Generate() (string, error) {
	var buf bytes.Buffer

	// Write package declaration
	fmt.Fprintf(&buf, "package %s\n\n", g.packageName)

	// Write imports
	g.writeImports(&buf)

	// Generate common types
	if err := g.generateCommonTypes(&buf); err != nil {
		return "", err
	}

	// Generate structs for each table
	for _, table := range g.parser.GetTables() {
		if err := g.generateTableStruct(&buf, table); err != nil {
			return "", err
		}
	}

	// Generate field references for type safety
	for _, table := range g.parser.GetTables() {
		if err := g.generateFieldReferences(&buf, table); err != nil {
			return "", err
		}
	}

	// Generate query builders with type-safe methods
	for _, table := range g.parser.GetTables() {
		if err := g.generateTypeSafeQueryBuilder(&buf, table); err != nil {
			return "", err
		}
	}

	// Generate join builders
	for _, table := range g.parser.GetTables() {
		if err := g.generateJoinBuilders(&buf, table); err != nil {
			return "", err
		}
	}

	// Generate database wrapper
	if err := g.generateDatabaseWrapper(&buf); err != nil {
		return "", err
	}

	// Format the generated code
	fmtOptions := format.Options{}
	formatted, err := format.Source(buf.Bytes(), fmtOptions)
	if err != nil {
		return "", fmt.Errorf("failed to format generated code: %w", err)
	}

	return string(formatted), nil
}

// writeImports writes import statements
func (g *Generator) writeImports(buf *bytes.Buffer) {
	buf.WriteString("import (\n")
	for imp := range g.imports {
		fmt.Fprintf(buf, "\t\"%s\"\n", imp)
	}
	buf.WriteString(")\n\n")
}

// generateCommonTypes generates common types used across queries
func (g *Generator) generateCommonTypes(buf *bytes.Buffer) error {
	rendered, err := templates.RenderCommonTypes()
	if err != nil {
		return err
	}
	buf.WriteString(rendered)
	return nil
}

// generateTableStruct generates a struct for a table
func (g *Generator) generateTableStruct(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)
	receiverName := strings.ToLower(structName[0:1])

	// Prepare template data
	data := templates.TableStructData{
		StructName:   structName,
		TableName:    table.Name,
		ReceiverName: receiverName,
		Fields:       make([]templates.StructField, 0, len(table.Columns)),
	}

	for _, col := range table.Columns {
		fieldName := g.toPascalCase(col.Name)
		goType := col.GoType

		// Make nullable pointer if:
		// 1. Column is SQL nullable (NULL constraint), OR
		// 2. Column has a DEFAULT value (can be omitted from INSERT), OR
		// 3. Column is a sequence (SERIAL/BIGSERIAL, can be omitted)
		if col.IsNullable || col.HasDefault || col.IsSequence {
			goType = "*" + goType
		}

		tags := g.buildStructTags(col)

		data.Fields = append(data.Fields, templates.StructField{
			FieldName:  fieldName,
			GoType:     goType,
			StructTags: tags,
		})
	}

	// Render the table struct using the template
	rendered, err := templates.RenderTableStruct(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	buf.WriteString("\n")
	return nil
}

// generateFieldReferences generates type-safe field references
func (g *Generator) generateFieldReferences(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)
	fieldsTypeName := structName + "Fields"
	receiverName := strings.ToLower(fieldsTypeName[0:1])

	// Prepare template data
	data := templates.FieldReferencesData{
		StructName:     structName,
		TableName:      table.Name,
		FieldsTypeName: fieldsTypeName,
		ReceiverName:   receiverName,
		Fields:         make([]templates.FieldRefData, 0, len(table.Columns)),
	}

	for _, col := range table.Columns {
		fieldName := g.toPascalCase(col.Name)
		data.Fields = append(data.Fields, templates.FieldRefData{
			FieldName:  fieldName,
			ColumnName: col.Name,
			GoType:     col.GoType,
		})
	}

	// Render the field references using the template
	rendered, err := templates.RenderFieldReferences(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	buf.WriteString("\n")
	return nil
}

// generateTypeSafeQueryBuilder generates type-safe query builder
func (g *Generator) generateTypeSafeQueryBuilder(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)
	builderName := structName + "Query"

	// Get primary key info
	var primaryKey string
	var primaryKeyField string
	var primaryKeyType string = "int64" // default

	for _, col := range table.Columns {
		if col.IsPrimary {
			primaryKey = col.Name
			primaryKeyField = g.toPascalCase(col.Name)
			primaryKeyType = col.GoType
			break
		}
	}

	// Prepare template data
	data := templates.QueryBuilderData{
		BuilderName:     builderName,
		StructName:      structName,
		TableName:       table.Name,
		PrimaryKey:      primaryKey,
		PrimaryKeyField: primaryKeyField,
		PrimaryKeyType:  primaryKeyType,
	}

	// Convert columns for template
	for _, col := range table.Columns {
		// Determine if this field is a pointer type in Go
		isPointer := col.IsNullable || col.HasDefault || col.IsSequence

		tc := templates.Column{
			Name:       col.Name,
			FieldName:  g.toPascalCase(col.Name),
			GoType:     col.GoType,
			IsNullable: col.IsNullable,
			IsPointer:  isPointer,
		}
		data.Columns = append(data.Columns, tc)

		if !col.IsPrimary {
			data.NonPrimaryColumns = append(data.NonPrimaryColumns, tc)
		}

		if !col.IsSequence {
			data.NonSequenceColumns = append(data.NonSequenceColumns, tc)
		}
	}

	// Render the query builder using the templates package
	rendered, err := templates.RenderQueryBuilder(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	return nil
}

// generateJoinBuilders generates type-safe join builders
func (g *Generator) generateJoinBuilders(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)
	builderName := structName + "Query"

	// Prepare template data
	data := templates.JoinBuildersData{
		BuilderName: builderName,
		StructName:  structName,
		TableName:   table.Name,
		Joins:       make([]templates.JoinData, 0, len(table.ForeignKeys)),
	}

	// Generate join methods for foreign keys
	for _, fk := range table.ForeignKeys {
		referencedTable, ok := g.parser.GetTable(fk.ReferencedTable)
		if !ok {
			continue
		}

		referencedStructName := g.toPascalCase(referencedTable.Name)
		joinMethodName := "Join" + referencedStructName
		leftJoinMethodName := "LeftJoin" + referencedStructName
		joinStructName := structName + referencedStructName + "Join"

		data.Joins = append(data.Joins, templates.JoinData{
			JoinMethodName:       joinMethodName,
			LeftJoinMethodName:   leftJoinMethodName,
			JoinStructName:       joinStructName,
			ReferencedTableName:  referencedTable.Name,
			ReferencedStructName: referencedStructName,
			LeftFieldName:        g.toPascalCase(fk.Column),
			RightFieldName:       g.toPascalCase(fk.ReferencedColumn),
		})
	}

	// Render the join builders using the template
	rendered, err := templates.RenderJoinBuilders(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	buf.WriteString("\n")
	return nil
}

// generateDatabaseWrapper generates the main database wrapper
func (g *Generator) generateDatabaseWrapper(buf *bytes.Buffer) error {
	// Prepare template data
	data := templates.DBWrapperData{
		Tables: make([]templates.TableMethod, 0, len(g.parser.GetTables())),
	}

	// Generate methods for each table
	for _, table := range g.parser.GetTables() {
		structName := g.toPascalCase(table.Name)
		methodName := structName
		builderName := structName + "Query"

		data.Tables = append(data.Tables, templates.TableMethod{
			MethodName:  methodName,
			BuilderName: builderName,
			TableName:   table.Name,
		})
	}

	// Sort tables alphabetically by table name for deterministic output
	sort.Slice(data.Tables, func(i, j int) bool {
		return data.Tables[i].TableName < data.Tables[j].TableName
	})

	// Render the DB wrapper using the template
	rendered, err := templates.RenderDBWrapper(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	buf.WriteString("\n")
	return nil
}

// Helper methods remain the same
func (g *Generator) buildStructTags(col Column) string {
	if len(col.Tags) == 0 {
		return ""
	}

	// Get keys and sort them for deterministic output
	keys := make([]string, 0, len(col.Tags))
	for key := range col.Tags {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Build tags in sorted order
	tags := make([]string, 0, len(keys))
	for _, key := range keys {
		tags = append(tags, fmt.Sprintf(`%s:"%s"`, key, col.Tags[key]))
	}

	return "`" + strings.Join(tags, " ") + "`"
}

func (g *Generator) toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[0:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

func (g *Generator) format(fileName string, in []byte) ([]byte, error) {
	formatted, err := imports.Process(fileName, in, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to format %s: %w", fileName, err)
	}

	fmtOptions := format.Options{}
	formatted, err = format.Source(formatted, fmtOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to format %s: %w", fileName, err)
	}

	return formatted, nil
}
