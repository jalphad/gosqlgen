package gosqlgen

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	//parser "github.com/jalphad/gosqlgen/oldparser"
	"github.com/jalphad/gosqlgen/parser"
	"github.com/jalphad/gosqlgen/templates"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

// Generator handles code generation with type-safe queries
type Generator struct {
	parser      *parser.Parser
	packageName string
	imports     map[string]bool
}

// NewGenerator creates a new code generator
func NewGenerator(parser *parser.Parser) *Generator {
	return &Generator{
		parser:      parser,
		packageName: "models",
		imports: map[string]bool{
			"context":                         true,
			"fmt":                             true,
			"time":                            true,
			"strings":                         true,
			"github.com/jackc/pgx/v5":         true,
			"github.com/jackc/pgx/v5/pgxpool": true,
			"github.com/jackc/pgx/v5/pgconn":  true,
			"github.com/google/uuid":          true,
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
	var err error
	files := make(map[string]string)

	// Generate common.gen.go with utility types
	commonBuf := bytes.Buffer{}
	if _, err := fmt.Fprintf(&commonBuf, "package %s\n\n", g.packageName); err != nil {
		return nil, err
	}
	if err := g.writeImports(&commonBuf); err != nil {
		return nil, err
	}
	if err := g.generateCommonTypes(&commonBuf); err != nil {
		return nil, err
	}

	//formatted, err := g.format("common.gen.go", commonBuf.Bytes())
	//if err != nil {
	//	return nil, err
	//}
	//files["common.gen.go"] = string(formatted)

	// Generate a file for each table
	for _, table := range g.parser.GetTables() {
		tableBuf := bytes.Buffer{}

		// Package declaration
		if _, err = fmt.Fprintf(&tableBuf, "package %s\n\n", g.packageName); err != nil {
			return nil, err
		}

		// Imports for table file
		if err = g.writeImports(&tableBuf); err != nil {
			return nil, err
		}

		// Table struct
		if err := g.generateTableStruct(&tableBuf, table); err != nil {
			return nil, err
		}

		// Column references
		//if err := g.generateFieldReferences(&tableBuf, table); err != nil {
		//	return nil, err
		//}

		// Query builder
		//if err := g.generateTypeSafeQueryBuilder(&tableBuf, table); err != nil {
		//	return nil, err
		//}

		// Join builders
		//if err := g.generateJoinBuilders(&tableBuf, table); err != nil {
		//	return nil, err
		//}

		// Format the generated code
		filename := fmt.Sprintf("%s.gen.go", table.Name)
		formatted, err := g.format(filename, tableBuf.Bytes())
		if err != nil {
			return nil, err
		}
		files[filename] = string(formatted)
	}

	// Generate db.gen.go for database wrapper
	dbBuf := bytes.Buffer{}
	if _, err = fmt.Fprintf(&dbBuf, "package %s\n\n", g.packageName); err != nil {
		return nil, err
	}
	if err = g.writeImports(&dbBuf); err != nil {
		return nil, err
	}
	if err = g.generateDatabaseWrapper(&dbBuf); err != nil {
		return nil, err
	}

	formatted, err := g.format("db.gen.go", dbBuf.Bytes())
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
	if _, err := fmt.Fprintf(&buf, "package %s\n\n", g.packageName); err != nil {
		return "", err
	}

	// Write imports
	if err := g.writeImports(&buf); err != nil {
		return "", err
	}

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
	//for _, table := range g.parser.GetTables() {
	//	if err := g.generateFieldReferences(&buf, table); err != nil {
	//		return "", err
	//	}
	//}

	// Generate query builders with type-safe methods
	//for _, table := range g.parser.GetTables() {
	//	if err := g.generateTypeSafeQueryBuilder(&buf, table); err != nil {
	//		return "", err
	//	}
	//}

	// Generate join builders
	//for _, table := range g.parser.GetTables() {
	//	if err := g.generateJoinBuilders(&buf, table); err != nil {
	//		return "", err
	//	}
	//}

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
func (g *Generator) writeImports(buf *bytes.Buffer) error {
	buf.WriteString("import (\n")
	for imp := range g.imports {
		if _, err := fmt.Fprintf(buf, "\t\"%s\"\n", imp); err != nil {
			return err
		}
	}
	buf.WriteString(")\n\n")

	return nil
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
func (g *Generator) generateTableStruct(buf *bytes.Buffer, table *parser.Table) error {
	baseName := templates.ToPascalCase(table.Name)
	structName := baseName + "Dto"
	receiverName := strings.ToLower(baseName[0:1])

	// Prepare template data
	data := templates.TableStructData{
		StructName:         structName,
		TableName:          table.Name,
		ReceiverName:       receiverName,
		Fields:             make([]templates.StructField, 0, len(table.Columns)),
		JoinedFields:       make([]templates.JoinedField, 0, len(table.ForeignKeys)),
		ReverseRelFields:   g.reverseRelations(table),
		ManyToManyFields:   g.manyToManyRelations(table),
		NonSequenceColumns: make([]templates.StructField, 0),
		PrimaryKeys:        make([]string, 0),
		PrimaryKeyFields:   make([]string, 0),
	}

	for _, col := range table.Columns {
		if col.IsPrimary {
			data.PrimaryKeys = append(data.PrimaryKeys, col.Name)
			data.PrimaryKeyFields = append(data.PrimaryKeyFields, templates.ToPascalCase(col.Name))
		}

		var isPointer bool
		fieldName := templates.ToPascalCase(col.Name)
		goType := col.GoType

		// Make nullable pointer if:
		// 1. Column is SQL nullable (NULL constraint), OR
		// 2. Column has a DEFAULT value (can be omitted from INSERT), OR
		// 3. Column is a sequence (SERIAL/BIGSERIAL, can be omitted)
		if col.IsNullable || col.HasDefault || col.IsSequence {
			goType = "*" + goType
			isPointer = true
		}

		tags := g.buildStructTags(col)

		field := templates.StructField{
			FieldName:  fieldName,
			ColumnName: col.Name,
			GoType:     goType,
			SQLType:    col.SQLType,
			StructTags: tags,
			IsPointer:  isPointer,
		}

		if !col.IsSequence {
			data.NonSequenceColumns = append(data.NonSequenceColumns, field)
		}

		data.Fields = append(data.Fields, field)
	}

	// Add joined fields for each foreign key
	for _, fk := range table.ForeignKeys {
		// FK.Prefix was populated during validation
		joinedFieldName := templates.ToPascalCase(fk.Prefix)

		// Get referenced table to build struct type
		referencedTable, ok := g.parser.GetTable(fk.ReferencedTableName)
		if !ok {
			continue // Skip if referenced table doesn't exist
		}

		referencedStructName := templates.ToPascalCase(referencedTable.Name) + "Dto"

		var referencedColumns []templates.Column
		for _, col := range referencedTable.Columns {
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			referencedColumns = append(referencedColumns, templates.Column{
				ColumnName: col.Name,
				FieldName:  templates.ToPascalCase(col.Name),
				GoType:     col.GoType,
				IsNullable: col.IsNullable,
				IsPointer:  isPointer,
			})
		}

		data.JoinedFields = append(data.JoinedFields, templates.JoinedField{
			FieldName:         joinedFieldName,
			GoType:            "*" + referencedStructName,
			StructName:        referencedStructName,
			ReferencedTable:   fk.ReferencedTableName,
			FKColumn:          fk.Column,
			ReferencedColumn:  fk.ReferencedColumn,
			ReferencedColumns: referencedColumns,
		})
	}

	// Render the table struct using the template
	rendered, err := templates.RenderTableStruct(data)
	if err != nil {
		return err
	}

	buf.WriteString(rendered)
	buf.WriteString("\n")

	// Generate collection loader methods if there are reverse rels or M2M rels
	if len(table.ReverseRelationships) > 0 || len(table.ManyToManyRels) > 0 {
		// Determine if table has a primary key
		var primaryKeyField string
		var primaryKeyType string
		hasPrimaryKey := false
		for _, col := range table.Columns {
			if col.IsPrimary {
				primaryKeyField = templates.ToPascalCase(col.Name)
				primaryKeyType = col.GoType
				hasPrimaryKey = true
				break
			}
		}

		// Build loader data for reverse relationships
		reverseRelLoaderFields := make([]templates.ReverseRelLoaderField, 0)
		for _, reverseRel := range table.ReverseRelationships {
			fromTableBaseName := templates.ToPascalCase(reverseRel.FromTable.Name)
			fromStructName := fromTableBaseName + "Dto"
			fromTableMethod := fromTableBaseName // DB wrapper method uses table name
			fkFieldName := templates.ToPascalCase(reverseRel.FKColumn)

			reverseRelLoaderFields = append(reverseRelLoaderFields, templates.ReverseRelLoaderField{
				FieldName:       reverseRel.FieldName,
				FromTable:       reverseRel.FromTable.Name,
				FromStructName:  fromStructName,
				FromTableMethod: fromTableMethod,
				FKFieldName:     fkFieldName,
			})
		}

		// Build loader data for M2M relationships
		m2mLoaderFields := make([]templates.ManyToManyLoaderField, 0)
		for _, m2m := range table.ManyToManyRels {
			refTable, ok := g.parser.GetTable(m2m.ReferencedTable.Name)
			if !ok {
				continue
			}

			// Find the primary key of the referenced table
			var refPKField string
			var refPKType string
			for _, col := range refTable.Columns {
				if col.IsPrimary {
					refPKField = templates.ToPascalCase(col.Name)
					refPKType = col.GoType
					break
				}
			}

			refTableBaseName := templates.ToPascalCase(m2m.ReferencedTable.Name)
			refStructName := refTableBaseName + "Dto"
			refTableMethod := refTableBaseName // DB wrapper method uses table name

			m2mLoaderFields = append(m2mLoaderFields, templates.ManyToManyLoaderField{
				FieldName:             m2m.FieldName,
				JunctionTable:         m2m.JunctionTable.Name,
				LeftFKColumn:          m2m.LeftFKColumn,
				RightFKColumn:         m2m.RightFKColumn,
				ReferencedTable:       m2m.ReferencedTable.Name,
				ReferencedStructName:  refStructName,
				ReferencedTableMethod: refTableMethod,
				ReferencedPKField:     refPKField,
				ReferencedPKType:      refPKType,
			})
		}

		loaderData := templates.CollectionLoaderData{
			StructName:       structName,
			TableName:        table.Name,
			ReceiverName:     receiverName,
			HasPrimaryKey:    hasPrimaryKey,
			PrimaryKeyField:  primaryKeyField,
			PrimaryKeyType:   primaryKeyType,
			ReverseRelFields: reverseRelLoaderFields,
			ManyToManyFields: m2mLoaderFields,
		}

		loadersRendered, err := templates.RenderCollectionLoaders(loaderData)
		if err != nil {
			return err
		}

		buf.WriteString(loadersRendered)
		buf.WriteString("\n")
	}

	return nil
}

// generateFieldReferences generates type-safe field references
func (g *Generator) generateFieldReferences(buf *bytes.Buffer, table *parser.Table) error {
	baseName := templates.ToPascalCase(table.Name)
	structName := baseName + "Dto"
	fieldsTypeName := baseName + "Fields"
	receiverName := strings.ToLower(fieldsTypeName[0:1])

	// Prepare template data
	data := templates.FieldReferencesData{
		BaseName:         baseName,
		StructName:       structName,
		TableName:        table.Name,
		FieldsTypeName:   fieldsTypeName,
		ReceiverName:     receiverName,
		Fields:           make([]templates.FieldRefData, 0, len(table.Columns)),
		ReverseRelations: g.reverseRelations(table),
		ManyToManyRels:   g.manyToManyRelations(table),
	}

	for _, col := range table.Columns {
		fieldName := templates.ToPascalCase(col.Name)
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
func (g *Generator) generateTypeSafeQueryBuilder(buf *bytes.Buffer, table *parser.Table) error {
	baseName := templates.ToPascalCase(table.Name)
	structName := baseName + "Dto"
	builderName := baseName + "Query"

	// Get primary key info
	var primaryKey string
	var primaryKeyField string
	var primaryKeyType = "int64" // default

	for _, col := range table.Columns {
		if col.IsPrimary {
			primaryKey = col.Name
			primaryKeyField = templates.ToPascalCase(col.Name)
			primaryKeyType = col.GoType
			break
		}
	}

	// Prepare template data
	data := templates.QueryBuilderData{
		BaseName:        baseName,
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
			ColumnName: col.Name,
			FieldName:  templates.ToPascalCase(col.Name),
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

	// Add FK information for join handling
	for _, fk := range table.ForeignKeys {
		referencedTable, ok := g.parser.GetTable(fk.ReferencedTableName)
		if !ok {
			continue
		}

		referencedBaseName := templates.ToPascalCase(referencedTable.Name)
		referencedStructName := referencedBaseName + "Dto"
		joinedFieldName := templates.ToPascalCase(fk.Prefix)

		// Convert referenced table columns
		var referencedColumns []templates.Column
		for _, col := range referencedTable.Columns {
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			referencedColumns = append(referencedColumns, templates.Column{
				ColumnName: col.Name,
				FieldName:  templates.ToPascalCase(col.Name),
				GoType:     col.GoType,
				IsNullable: col.IsNullable,
				IsPointer:  isPointer,
			})
		}

		data.ForeignKeys = append(data.ForeignKeys, templates.ForeignKeyData{
			JoinedFieldName:      joinedFieldName,
			ReferencedTable:      fk.ReferencedTableName,
			ReferencedStructName: referencedStructName,
			ReferencedColumns:    referencedColumns,
		})
	}

	// Add reverse relationship information for scanInto
	for _, reverseRel := range table.ReverseRelationships {
		fromTableBaseName := templates.ToPascalCase(reverseRel.FromTable.Name)
		fromStructName := fromTableBaseName + "Dto"
		fromTableMethod := fromTableBaseName
		fkFieldName := templates.ToPascalCase(reverseRel.FKColumn)

		data.ReverseRelations = append(data.ReverseRelations, templates.ReverseRelLoaderField{
			FieldName:       reverseRel.FieldName,
			FromTable:       reverseRel.FromTable.Name,
			FromStructName:  fromStructName,
			FromTableMethod: fromTableMethod,
			FKFieldName:     fkFieldName,
		})
	}

	// Add M2M relationship information for scanInto
	for _, m2m := range table.ManyToManyRels {
		refTable, ok := g.parser.GetTable(m2m.ReferencedTable.Name)
		if !ok {
			continue
		}

		// Find the primary key of the referenced table
		var refPKField string
		var refPKType string
		for _, col := range refTable.Columns {
			if col.IsPrimary {
				refPKField = templates.ToPascalCase(col.Name)
				refPKType = col.GoType
				break
			}
		}

		refTableBaseName := templates.ToPascalCase(m2m.ReferencedTable.Name)
		refStructName := refTableBaseName + "Dto"
		refTableMethod := refTableBaseName

		data.ManyToManyRels = append(data.ManyToManyRels, templates.ManyToManyLoaderField{
			FieldName:             m2m.FieldName,
			JunctionTable:         m2m.JunctionTable.Name,
			LeftFKColumn:          m2m.LeftFKColumn,
			RightFKColumn:         m2m.RightFKColumn,
			ReferencedTable:       m2m.ReferencedTable.Name,
			ReferencedStructName:  refStructName,
			ReferencedTableMethod: refTableMethod,
			ReferencedPKField:     refPKField,
			ReferencedPKType:      refPKType,
		})
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
func (g *Generator) generateJoinBuilders(buf *bytes.Buffer, table *parser.Table) error {
	baseName := templates.ToPascalCase(table.Name)
	structName := baseName + "Dto"
	builderName := baseName + "Query"

	// Prepare template data
	data := templates.JoinBuildersData{
		BaseName:    baseName,
		BuilderName: builderName,
		StructName:  structName,
		TableName:   table.Name,
		Joins:       make([]templates.JoinData, 0, len(table.ForeignKeys)),
	}

	// Generate join methods for foreign keys
	for _, fk := range table.ForeignKeys {
		referencedTable, ok := g.parser.GetTable(fk.ReferencedTableName)
		if !ok {
			continue
		}

		referencedBaseName := templates.ToPascalCase(referencedTable.Name)
		referencedStructName := referencedBaseName + "Dto"
		joinMethodName := "Join" + referencedBaseName
		leftJoinMethodName := "LeftJoin" + referencedBaseName
		joinStructName := baseName + referencedBaseName + "Join"

		data.Joins = append(data.Joins, templates.JoinData{
			JoinMethodName:       joinMethodName,
			LeftJoinMethodName:   leftJoinMethodName,
			JoinStructName:       joinStructName,
			ReferencedTableName:  referencedTable.Name,
			ReferencedBaseName:   referencedBaseName,
			ReferencedStructName: referencedStructName,
			LeftFieldName:        templates.ToPascalCase(fk.Column),
			RightFieldName:       templates.ToPascalCase(fk.ReferencedColumn),
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
		tableName := templates.ToPascalCase(table.Name)
		methodName := tableName
		builderName := tableName + "Query"
		structName := tableName + "Dto"

		data.Tables = append(data.Tables, templates.TableMethod{
			MethodName:  methodName,
			BuilderName: builderName,
			TableName:   table.Name,
			StructName:  structName,
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
func (g *Generator) buildStructTags(col parser.Column) string {
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

// reverseRelations constructs the list of reverse relations on the table
func (g *Generator) reverseRelations(table *parser.Table) []templates.ReverseRelField {
	ret := make([]templates.ReverseRelField, 0, len(table.ReverseRelationships))
	for _, reverseRel := range table.ReverseRelationships {
		fromTableBaseName := templates.ToPascalCase(reverseRel.FromTable.Name)
		fromStructName := fromTableBaseName + "Dto"

		var fromPKField string
		for _, col := range reverseRel.FromTable.Columns {
			if col.IsPrimary {
				fromPKField = templates.ToPascalCase(col.Name)
				break
			}
		}

		ret = append(ret, templates.ReverseRelField{
			FieldName:   reverseRel.FieldName,
			GoType:      "[]" + fromStructName,
			StructName:  fromStructName,
			FromTable:   reverseRel.FromTable.Name,
			FKColumn:    reverseRel.FKColumn,
			FromPKField: fromPKField,
		})
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].FieldName < ret[j].FieldName
	})

	return ret
}

func (g *Generator) manyToManyRelations(table *parser.Table) []templates.ManyToManyField {
	ret := make([]templates.ManyToManyField, 0, len(table.ManyToManyRels))
	for _, m2m := range table.ManyToManyRels {
		refTableBaseName := templates.ToPascalCase(m2m.ReferencedTable.Name)
		refStructName := refTableBaseName + "Dto"

		var refPKField string
		for _, col := range m2m.ReferencedTable.Columns {
			if col.IsPrimary {
				refPKField = templates.ToPascalCase(col.Name)
				break
			}
		}

		ret = append(ret, templates.ManyToManyField{
			FieldName:         m2m.FieldName,
			GoType:            "[]" + refStructName,
			StructName:        refStructName,
			JunctionTable:     m2m.JunctionTable.Name,
			LeftFKColumn:      m2m.LeftFKColumn,
			RightFKColumn:     m2m.RightFKColumn,
			ReferencedTable:   m2m.ReferencedTable.Name,
			ReferencedPKField: refPKField,
		})
	}

	return ret
}
