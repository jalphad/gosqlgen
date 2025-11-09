package gosqlgen

import (
	"bytes"
	"fmt"
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
	g.generateCommonTypes(&commonBuf)

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
	g.generateCommonTypes(&buf)

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
func (g *Generator) generateCommonTypes(buf *bytes.Buffer) {
	buf.WriteString(`// OrderDirection represents sort order
type OrderDirection string

const ASC OrderDirection = "ASC"
const DESC OrderDirection = "DESC"

// ComparisonOp represents comparison operators
type ComparisonOp string

const EQ ComparisonOp = "="
const NEQ ComparisonOp = "!="
const GT ComparisonOp = ">"
const GTE ComparisonOp = ">="
const LT ComparisonOp = "<"
const LTE ComparisonOp = "<="
const LIKE ComparisonOp = "LIKE"
const IN ComparisonOp = "IN"

// LogicalOp represents logical operators
type LogicalOp string

const AND LogicalOp = "AND"
const OR LogicalOp = "OR"

// FieldRef represents a type-safe field reference
type FieldRef struct {
    Table  string
    Column string
    GoType string
    Alias  string
}

// As creates an aliased field reference
func (f *FieldRef) As(alias string) *FieldRef {
    return &FieldRef{
        Table:  f.Table,
        Column: f.Column,
        GoType: f.GoType,
        Alias:  alias,
    }
}

// String returns the SQL representation
func (f *FieldRef) String() string {
    if f.Alias != "" {
        if f.Table == "" {
            return fmt.Sprintf("%s AS %s", f.Column, f.Alias)
        }
        return fmt.Sprintf("%s.%s AS %s", f.Table, f.Column, f.Alias)
    }
    if f.Table == "" {
        return f.Column
    }
    return fmt.Sprintf("%s.%s", f.Table, f.Column)
}

// Condition represents a WHERE condition
type Condition struct {
    Field    *FieldRef
    Op       ComparisonOp
    Value    interface{}
    Logical  LogicalOp
    SubConds []Condition
}

// JoinClause represents a JOIN operation
type JoinClause struct {
    Type       string
    Table      string
    LeftField  *FieldRef
    RightField *FieldRef
}

// OrderClause represents an ORDER BY clause
type OrderClause struct {
    Field     *FieldRef
    Direction OrderDirection
}

`)
}

// generateTableStruct generates a struct for a table
func (g *Generator) generateTableStruct(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)

	fmt.Fprintf(buf, "// %s represents the %s table\n", structName, table.Name)
	fmt.Fprintf(buf, "type %s struct {\n", structName)

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
		fmt.Fprintf(buf, "\t%s %s %s\n", fieldName, goType, tags)
	}

	buf.WriteString("}\n\n")

	fmt.Fprintf(buf, "// TableName returns the table name for %s\n", structName)
	fmt.Fprintf(buf, "func (%s *%s) TableName() string {\n", strings.ToLower(structName[0:1]), structName)
	fmt.Fprintf(buf, "\treturn \"%s\"\n", table.Name)
	buf.WriteString("}\n\n")

	return nil
}

// generateFieldReferences generates type-safe field references
func (g *Generator) generateFieldReferences(buf *bytes.Buffer, table *Table) error {
	structName := g.toPascalCase(table.Name)
	fieldsTypeName := structName + "Fields"

	// Generate fields struct
	fmt.Fprintf(buf, "// %s provides type-safe field references for %s\n", fieldsTypeName, structName)
	fmt.Fprintf(buf, "type %s struct{}\n\n", fieldsTypeName)

	fmt.Fprintf(buf, "// Fields returns field references for %s\n", structName)
	fmt.Fprintf(buf, "var %sTable = %s{}\n\n", structName, fieldsTypeName)

	// Generate field accessor methods
	for _, col := range table.Columns {
		fieldName := g.toPascalCase(col.Name)

		// Generate field reference type
		fmt.Fprintf(buf, "// %s returns a field reference for %s.%s\n", fieldName, structName, fieldName)
		fmt.Fprintf(buf, "func (%s %s) %s() *FieldRef {\n", strings.ToLower(fieldsTypeName[0:1]), fieldsTypeName, fieldName)
		fmt.Fprintf(buf, "\treturn &FieldRef{\n")
		fmt.Fprintf(buf, "\t\tTable:  \"%s\",\n", table.Name)
		fmt.Fprintf(buf, "\t\tColumn: \"%s\",\n", col.Name)
		fmt.Fprintf(buf, "\t\tGoType: \"%s\",\n", col.GoType)
		fmt.Fprintf(buf, "\t}\n")
		buf.WriteString("}\n\n")
	}

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

	// Generate join methods for foreign keys
	for _, fk := range table.ForeignKeys {
		referencedTable, ok := g.parser.GetTable(fk.ReferencedTable)
		if !ok {
			continue
		}

		referencedStructName := g.toPascalCase(referencedTable.Name)
		joinMethodName := "Join" + referencedStructName
		leftJoinMethodName := "LeftJoin" + referencedStructName

		// InnerJoin method
		fmt.Fprintf(buf, "// %s performs a type-safe inner join with %s\n", joinMethodName, referencedTable.Name)
		fmt.Fprintf(buf, "func (q *%s) %s() *%s {\n", builderName, joinMethodName, builderName)
		fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
		fmt.Fprintf(buf, "\t\tType:       \"INNER JOIN\",\n")
		fmt.Fprintf(buf, "\t\tTable:      \"%s\",\n", referencedTable.Name)
		fmt.Fprintf(buf, "\t\tLeftField:  %sTable.%s(),\n", structName, g.toPascalCase(fk.Column))
		fmt.Fprintf(buf, "\t\tRightField: %sTable.%s(),\n", referencedStructName, g.toPascalCase(fk.ReferencedColumn))
		fmt.Fprintf(buf, "\t})\n")
		fmt.Fprintf(buf, "\treturn q\n")
		fmt.Fprintf(buf, "}\n\n")

		// LeftJoin method
		fmt.Fprintf(buf, "// %s performs a type-safe left join with %s\n", leftJoinMethodName, referencedTable.Name)
		fmt.Fprintf(buf, "func (q *%s) %s() *%s {\n", builderName, leftJoinMethodName, builderName)
		fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
		fmt.Fprintf(buf, "\t\tType:       \"LEFT JOIN\",\n")
		fmt.Fprintf(buf, "\t\tTable:      \"%s\",\n", referencedTable.Name)
		fmt.Fprintf(buf, "\t\tLeftField:  %sTable.%s(),\n", structName, g.toPascalCase(fk.Column))
		fmt.Fprintf(buf, "\t\tRightField: %sTable.%s(),\n", referencedStructName, g.toPascalCase(fk.ReferencedColumn))
		fmt.Fprintf(buf, "\t})\n")
		fmt.Fprintf(buf, "\treturn q\n")
		fmt.Fprintf(buf, "}\n\n")

		// Generate join result struct
		joinStructName := structName + referencedStructName + "Join"
		fmt.Fprintf(buf, "// %s represents a join result between %s and %s\n",
			joinStructName, table.Name, referencedTable.Name)
		fmt.Fprintf(buf, "type %s struct {\n", joinStructName)
		fmt.Fprintf(buf, "\t%s *%s\n", structName, structName)
		fmt.Fprintf(buf, "\t%s *%s\n", referencedStructName, referencedStructName)
		fmt.Fprintf(buf, "}\n\n")
	}

	// Generate custom join method for arbitrary tables
	fmt.Fprintf(buf, "// JoinOn performs a custom join with type-safe field references\n")
	fmt.Fprintf(buf, "func (q *%s) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *%s {\n", builderName, builderName)
	fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
	fmt.Fprintf(buf, "\t\tType:       joinType,\n")
	fmt.Fprintf(buf, "\t\tTable:      table,\n")
	fmt.Fprintf(buf, "\t\tLeftField:  leftField,\n")
	fmt.Fprintf(buf, "\t\tRightField: rightField,\n")
	fmt.Fprintf(buf, "\t})\n")
	fmt.Fprintf(buf, "\treturn q\n")
	fmt.Fprintf(buf, "}\n\n")

	return nil
}

// generateDatabaseWrapper generates the main database wrapper
func (g *Generator) generateDatabaseWrapper(buf *bytes.Buffer) error {
	buf.WriteString("// DB wraps the database connection pool\n")
	buf.WriteString("type DB struct {\n")
	buf.WriteString("\tpool *pgxpool.Pool\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// NewDB creates a new database wrapper\n")
	buf.WriteString("func NewDB(pool *pgxpool.Pool) *DB {\n")
	buf.WriteString("\treturn &DB{pool: pool}\n")
	buf.WriteString("}\n\n")

	// Generate methods for each table
	for _, table := range g.parser.GetTables() {
		structName := g.toPascalCase(table.Name)
		methodName := structName
		builderName := structName + "Query"

		fmt.Fprintf(buf, "// %s returns a query builder for %s\n", methodName, table.Name)
		fmt.Fprintf(buf, "func (db *DB) %s() *%s {\n", methodName, builderName)
		fmt.Fprintf(buf, "\treturn New%s(db.pool)\n", builderName)
		buf.WriteString("}\n\n")
	}

	// Transaction support
	buf.WriteString("// Transaction executes a function within a transaction\n")
	buf.WriteString("func (db *DB) Transaction(ctx context.Context, fn func(*Tx) error) error {\n")
	buf.WriteString("\treturn pgx.BeginFunc(ctx, db.pool, func(pgxTx pgx.Tx) error {\n")
	buf.WriteString("\t\ttx := NewTx(pgxTx)\n")
	buf.WriteString("\t\treturn fn(tx)\n")
	buf.WriteString("\t})\n")
	buf.WriteString("}\n\n")

	// Transaction methods for each table
	buf.WriteString("// Tx provides transaction-aware query builders\n")
	buf.WriteString("type Tx struct {\n")
	buf.WriteString("\ttx pgx.Tx\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// NewTx creates transaction-aware query builders\n")
	buf.WriteString("func NewTx(tx pgx.Tx) *Tx {\n")
	buf.WriteString("\treturn &Tx{tx: tx}\n")
	buf.WriteString("}\n\n")

	for _, table := range g.parser.GetTables() {
		structName := g.toPascalCase(table.Name)
		methodName := structName
		builderName := structName + "Query"

		fmt.Fprintf(buf, "// %s returns a query builder within the transaction\n", methodName)
		fmt.Fprintf(buf, "func (tx *Tx) %s() *%s {\n", methodName, builderName)
		fmt.Fprintf(buf, "\tq := New%s(nil)\n", builderName)
		fmt.Fprintf(buf, "\tq.WithTx(tx.tx)\n")
		fmt.Fprintf(buf, "\treturn q\n")
		buf.WriteString("}\n\n")
	}

	return nil
}

// Helper methods remain the same
func (g *Generator) buildStructTags(col Column) string {
	var tags []string
	for key, value := range col.Tags {
		tags = append(tags, fmt.Sprintf(`%s:"%s"`, key, value))
	}
	if len(tags) == 0 {
		return ""
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
