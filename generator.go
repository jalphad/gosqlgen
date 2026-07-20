package gosqlgen

import (
	"bytes"
	"fmt"
	"maps"
	"path"
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
	parser          *parser.Parser
	packageName     string
	packageRootPath string
	schemas         []string
	tables          []*parser.Table
	tableInfo       map[*parser.Table]generatedTableInfo
}

type generatedTableInfo struct {
	packageName string
	baseName    string
	plainName   string
	sqlName     string
}

// NewGenerator creates a new code generator
func NewGenerator(parser *parser.Parser) *Generator {
	return &Generator{
		parser:          parser,
		packageName:     "models",
		packageRootPath: "example.local/example",
	}
}

// SetPackageName sets the package name for generated code
func (g *Generator) SetPackageName(name string) {
	g.packageName = name
}

// SetPackagePath sets the import path root containing the generated package.
func (g *Generator) SetPackagePath(path string) {
	g.packageRootPath = path
}

// SetSchemas limits generation to tables in the provided schemas.
// An empty list generates tables from every parsed schema.
func (g *Generator) SetSchemas(schemas []string) {
	g.schemas = append([]string(nil), schemas...)
}

func (g *Generator) generatedPackagePath() string {
	if g.packageRootPath == "" {
		return g.packageName
	}
	if g.packageName == "" {
		return g.packageRootPath
	}
	return path.Join(g.packageRootPath, g.packageName)
}

// GenerateFiles generates Go code for all parsed tables as separate files
// Returns a map of filename to file content
func (g *Generator) GenerateFiles() (map[string]string, error) {
	if err := g.prepareTables(); err != nil {
		return nil, err
	}

	files := make(map[string]string)

	// Generate common.gen.go with utility types
	commonBuf := bytes.Buffer{}
	if err := g.generateCommonTypes(&commonBuf); err != nil {
		return nil, err
	}

	formatted, err := g.format("common.gen.go", commonBuf.Bytes())
	if err != nil {
		return nil, err
	}
	files["models/common.gen.go"] = string(formatted)

	// Generate a file for each table
	for _, table := range g.tables {
		tableBuf := bytes.Buffer{}
		// Table struct
		if err = g.generateTableStruct(&tableBuf, table); err != nil {
			return nil, err
		}

		// Format to generated code
		filename := fmt.Sprintf("models/%s.gen.go", g.info(table).packageName)
		formatted, err = g.format(filename, tableBuf.Bytes())
		if err != nil {
			return nil, err
		}
		files[filename] = string(formatted)
	}

	// Generate db.gen.go for the root database facade.
	dbBuf := bytes.Buffer{}
	if err = g.generateDatabaseWrapper(&dbBuf); err != nil {
		return nil, err
	}

	formatted, err = g.format("db.gen.go", dbBuf.Bytes())
	if err != nil {
		return nil, err
	}
	files["db.gen.go"] = string(formatted)

	// Generate table query packages
	if err = g.generateTableQueryPackages(files); err != nil {
		return nil, err
	}

	// Generate AST package
	astFiles, err := templates.RenderASTPackage()
	if err != nil {
		return nil, err
	}
	maps.Copy(files, astFiles)

	// Generate query Builder package
	packagePath := g.generatedPackagePath()

	builderFiles, err := templates.RenderQueryBuilderPackage(packagePath)
	if err != nil {
		return nil, err
	}
	maps.Copy(files, builderFiles)

	exprFiles, err := templates.RenderQueryExprPackage(packagePath)
	if err != nil {
		return nil, err
	}
	maps.Copy(files, exprFiles)

	// Generate query package helper files
	if err := g.generateQueryPackageHelpers(files); err != nil {
		return nil, err
	}

	return files, nil
}

func (g *Generator) prepareTables() error {
	requested := make(map[string]struct{}, len(g.schemas))
	for _, schema := range g.schemas {
		if schema != "" {
			requested[schema] = struct{}{}
		}
	}

	found := make(map[string]bool, len(requested))
	tables := make([]*parser.Table, 0, len(g.parser.GetTables()))
	for _, table := range g.parser.GetTables() {
		if len(requested) > 0 {
			if _, ok := requested[table.Schema]; !ok {
				continue
			}
			found[table.Schema] = true
		}
		tables = append(tables, table)
	}
	for schema := range requested {
		if !found[schema] {
			return fmt.Errorf("configured schema %q has no parsed tables", schema)
		}
	}

	sort.Slice(tables, func(i, j int) bool {
		if tables[i].Schema == tables[j].Schema {
			return tables[i].Name < tables[j].Name
		}
		return tables[i].Schema < tables[j].Schema
	})
	nameCounts := make(map[string]int, len(tables))
	for _, table := range tables {
		nameCounts[table.Name]++
	}

	g.tables = tables
	g.tableInfo = make(map[*parser.Table]generatedTableInfo, len(tables))
	for _, table := range tables {
		apiName := table.Name
		if nameCounts[table.Name] > 1 {
			apiName = table.Schema + "_" + table.Name
		}
		g.tableInfo[table] = generatedTableInfo{
			packageName: apiName,
			baseName:    templates.ToPascalCase(apiName),
			plainName:   table.Schema + "." + table.Name,
			sqlName:     quoteIdentifier(table.Schema) + "." + quoteIdentifier(table.Name),
		}
	}
	return nil
}

func (g *Generator) info(table *parser.Table) generatedTableInfo {
	return g.tableInfo[table]
}

func (g *Generator) selectedReferencedTable(fk parser.ForeignKey) (*parser.Table, bool) {
	schema := fk.ReferencedSchema
	if schema == "" {
		schema = "public"
	}
	table, ok := g.parser.GetTable(schema + "." + fk.ReferencedTableName)
	if !ok {
		return nil, false
	}
	_, selected := g.tableInfo[table]
	return table, selected
}

func quoteIdentifier(identifier string) string {
	return `"` + strings.ReplaceAll(identifier, `"`, `""`) + `"`
}

// generateQueryPackageHelpers generates functions.gen.go, grammar.gen.go, helpers.gen.go, queries.gen.go
func (g *Generator) generateQueryPackageHelpers(files map[string]string) error {
	// Generate functions.gen.go
	funcData := templates.QueryFunctionsData{
		PackagePath: g.generatedPackagePath(),
	}
	content, err := templates.RenderQueryFunctions(funcData)
	if err != nil {
		return fmt.Errorf("failed to render query functions: %w", err)
	}
	fileName := "functions.gen.go"
	content, err = g.format(fileName, content)
	if err != nil {
		return err
	}
	files["query/functions.gen.go"] = string(content)

	// Generate grammar.gen.go
	grammarData := templates.QueryGrammarData{
		PackagePath: g.generatedPackagePath(),
	}
	content, err = templates.RenderQueryGrammar(grammarData)
	if err != nil {
		return fmt.Errorf("failed to render query grammar: %w", err)
	}
	files["query/grammar.gen.go"] = string(content)

	// Generate helpers.gen.go
	helpersData := templates.QueryHelpersData{
		PackagePath: g.generatedPackagePath(),
	}
	content, err = templates.RenderQueryHelpers(helpersData)
	if err != nil {
		return fmt.Errorf("failed to render query helpers: %w", err)
	}
	files["query/helpers.gen.go"] = string(content)

	// Generate queries.gen.go
	if err = g.generateQueryFunctions(files); err != nil {
		return err
	}

	return nil
}

// generateQueryFunctions generates InsertOne and InsertMany functions for each table
func (g *Generator) generateQueryFunctions(files map[string]string) error {
	tables := make([]templates.QueryTableData, 0, len(g.tables))

	for _, table := range g.tables {
		info := g.info(table)
		baseName := info.baseName
		structName := baseName + "Dto"
		packageName := info.packageName
		receiverName := strings.ToLower(baseName[0:1])

		// Collect ALL primary key columns (support composite PKs)
		primaryKeyFields := make([]string, 0)
		primaryKeyColumns := make([]templates.QueryColumn, 0)
		updateColumns := make([]templates.QueryColumn, 0)
		for _, col := range table.Columns {
			fieldName := templates.ToPascalCase(col.Name)
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			if col.IsPrimary {
				primaryKeyFields = append(primaryKeyFields, fieldName)
				primaryKeyColumns = append(primaryKeyColumns, templates.QueryColumn{
					FieldName:  fieldName,
					ColumnName: col.Name,
					GoType:     col.GoType,
					SQLType:    col.SQLType,
					IsPointer:  isPointer,
					IsPrimary:  true,
					CastType:   templates.ToCastSQLType(col.GoType, col.SQLType),
					CastArray:  templates.ToCastArrayMethod(col.GoType, col.SQLType),
				})
				continue
			}
			updateColumns = append(updateColumns, templates.QueryColumn{
				FieldName:  fieldName,
				ColumnName: col.Name,
				GoType:     col.GoType,
				SQLType:    col.SQLType,
				IsPointer:  isPointer,
				CastType:   templates.ToCastSQLType(col.GoType, col.SQLType),
				CastArray:  templates.ToCastArrayMethod(col.GoType, col.SQLType),
			})
		}

		// Filter columns for INSERT:
		// - Exclude sequence columns (SERIAL/BIGSERIAL)
		// - Exclude columns with default values
		// - INCLUDE nullable columns (nil values will insert NULL)
		insertColumns := make([]templates.QueryColumn, 0)
		for _, col := range table.Columns {
			// Skip if:
			// 1. Column is a sequence (SERIAL/BIGSERIAL)
			if col.IsSequence {
				continue
			}
			// 2. Column has a DEFAULT value
			if col.HasDefault {
				continue
			}

			fieldName := templates.ToPascalCase(col.Name)
			insertColumns = append(insertColumns, templates.QueryColumn{
				FieldName:  fieldName,
				ColumnName: col.Name,
				GoType:     col.GoType,
				SQLType:    col.SQLType,
				IsPointer:  col.IsNullable || col.HasDefault || col.IsSequence,
				IsPrimary:  col.IsPrimary,
				CastArray:  templates.ToCastArrayMethod(col.GoType, col.SQLType),
			})
		}

		tables = append(tables, templates.QueryTableData{
			PackageName:       packageName,
			TableName:         info.sqlName,
			StructName:        structName,
			ReceiverName:      receiverName,
			InsertColumns:     insertColumns,
			UpdateColumns:     updateColumns,
			PrimaryKeyFields:  primaryKeyFields,
			PrimaryKeyColumns: primaryKeyColumns,
		})
	}
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].StructName < tables[j].StructName
	})

	data := templates.QueryQueriesData{
		PackagePath: g.generatedPackagePath(),
		Tables:      tables,
	}

	content, err := templates.RenderQueryQueries(data)
	if err != nil {
		return fmt.Errorf("failed to render queries package: %w", err)
	}
	fileName := "queries.go"
	content, err = g.format(fileName, content)
	if err != nil {
		return fmt.Errorf("failed to format %s: %w", fileName, err)
	}
	files["query/"+fileName] = string(content)

	return nil
}

// generateTableQueryPackages generates per-table query packages
func (g *Generator) generateTableQueryPackages(files map[string]string) error {
	for _, table := range g.tables {
		// Prepare template data
		fields := make([]templates.StructField, 0, len(table.Columns))
		for _, col := range table.Columns {
			fieldName := templates.ToPascalCase(col.Name)
			goType := col.GoType
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			if isPointer {
				goType = "*" + goType
			}
			fields = append(fields, templates.StructField{
				FieldName:  fieldName,
				ColumnName: col.Name,
				GoType:     goType,
				SQLType:    col.SQLType,
				IsPointer:  isPointer,
			})
		}

		info := g.info(table)
		baseName := info.baseName
		reverseRelations := g.reverseRelations(table)
		manyToManyRels := g.manyToManyRelations(table)
		data := templates.TableStructData{
			StructName:       baseName + "Dto",
			PackageName:      info.packageName,
			TableName:        info.sqlName,
			ReceiverName:     strings.ToLower(baseName[:1]),
			Fields:           fields,
			ReverseRelFields: reverseRelations,
			ManyToManyFields: manyToManyRels,
			PackagePath:      g.generatedPackagePath(),
		}

		// Render template
		content, err := templates.RenderColumnExpressions(data)
		if err != nil {
			return fmt.Errorf("failed to render column expressions for table %s: %w", table.Name, err)
		}

		// Add to files map with query package path
		filename := fmt.Sprintf("query/%s/%s.go", info.packageName, info.packageName)
		formatted, err := g.format(filename, []byte(content))
		if err != nil {
			return fmt.Errorf("failed to format %s: %w", filename, err)
		}
		files[filename] = string(formatted)
	}

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
	info := g.info(table)
	baseName := info.baseName
	structName := baseName + "Dto"
	receiverName := strings.ToLower(baseName[0:1])

	// Prepare template data
	data := templates.TableStructData{
		StructName:         structName,
		PackageName:        info.packageName,
		TableName:          info.sqlName,
		ReceiverName:       receiverName,
		Fields:             make([]templates.StructField, 0, len(table.Columns)),
		JoinedFields:       make([]templates.JoinedField, 0, len(table.ForeignKeys)),
		ReverseRelFields:   g.reverseRelations(table),
		ManyToManyFields:   g.manyToManyRelations(table),
		NonSequenceColumns: make([]templates.StructField, 0),
		PrimaryKeys:        make([]string, 0),
		PrimaryKeyFields:   make([]string, 0),
		PackagePath:        g.generatedPackagePath(),
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
		joinedFieldName := templates.ToPascalCase(fk.Column) + "Ref"

		// Get referenced table to build struct type
		referencedTable, ok := g.selectedReferencedTable(fk)
		if !ok {
			continue // Skip if referenced table doesn't exist
		}

		referencedInfo := g.info(referencedTable)
		referencedStructName := referencedInfo.baseName + "Dto"

		var referencedColumns []templates.Column
		for _, col := range referencedTable.Columns {
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			referencedColumns = append(referencedColumns, templates.Column{
				ColumnName: col.Name,
				FieldName:  templates.ToPascalCase(col.Name),
				GoType:     col.GoType,
				SQLType:    col.SQLType,
				IsNullable: col.IsNullable,
				IsPointer:  isPointer,
			})
		}

		data.JoinedFields = append(data.JoinedFields, templates.JoinedField{
			FieldName:          joinedFieldName,
			GoType:             "*" + referencedStructName,
			StructName:         referencedStructName,
			ReferencedTable:    referencedInfo.sqlName,
			ReferencedTableTag: referencedInfo.plainName,
			FKColumn:           fk.Column,
			ReferencedColumn:   fk.ReferencedColumn,
			ReferencedColumns:  referencedColumns,
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

// generateDatabaseWrapper generates the main database wrapper
func (g *Generator) generateDatabaseWrapper(buf *bytes.Buffer) error {
	// Prepare template data
	data := templates.DBWrapperData{
		Tables:      make([]templates.TableMethod, 0, len(g.tables)),
		PackagePath: g.generatedPackagePath(),
		PackageName: g.packageName,
	}

	// Generate methods for each table
	for _, table := range g.tables {
		info := g.info(table)
		tableName := info.baseName
		methodName := tableName
		builderName := tableName + "Query"
		structName := tableName + "Dto"

		data.Tables = append(data.Tables, templates.TableMethod{
			MethodName:  methodName,
			BuilderName: builderName,
			PackageName: info.packageName,
			TableName:   info.sqlName,
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
		fromInfo, selected := g.tableInfo[reverseRel.FromTable]
		if !selected {
			continue
		}
		fromTableBaseName := fromInfo.baseName
		fromStructName := fromTableBaseName + "Dto"
		fieldName := reverseRel.FieldName
		if fromInfo.baseName != templates.ToPascalCase(reverseRel.FromTable.Name) {
			fieldName = fromInfo.baseName + "By" + templates.ToPascalCase(reverseRel.FKColumn)
		}

		var fromPKField string
		fromColumns := make([]templates.Column, len(reverseRel.FromTable.Columns))
		for i, col := range reverseRel.FromTable.Columns {
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			fromColumns[i] = templates.Column{
				ColumnName: col.Name,
				FieldName:  templates.ToPascalCase(col.Name),
				GoType:     col.GoType,
				SQLType:    col.SQLType,
				IsNullable: col.IsNullable,
				IsPointer:  isPointer,
			}
			if col.IsPrimary {
				fromPKField = templates.ToPascalCase(col.Name)
			}
		}

		ret = append(ret, templates.ReverseRelField{
			FieldName:    fieldName,
			GoType:       "[]" + fromStructName,
			StructName:   fromStructName,
			FromTable:    fromInfo.sqlName,
			FromTableTag: fromInfo.plainName,
			FKColumn:     reverseRel.FKColumn,
			FromPKField:  fromPKField,
			FromColumns:  fromColumns,
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
		refInfo, refSelected := g.tableInfo[m2m.ReferencedTable]
		junctionInfo, junctionSelected := g.tableInfo[m2m.JunctionTable]
		if !refSelected || !junctionSelected {
			continue
		}
		refTableBaseName := refInfo.baseName
		refStructName := refTableBaseName + "Dto"
		fieldName := m2m.FieldName
		if refInfo.baseName != templates.ToPascalCase(m2m.ReferencedTable.Name) {
			fieldName = refInfo.baseName
		}

		var refPKField string
		referencedColumns := make([]templates.Column, len(m2m.ReferencedTable.Columns))
		for i, col := range m2m.ReferencedTable.Columns {
			isPointer := col.IsNullable || col.HasDefault || col.IsSequence
			referencedColumns[i] = templates.Column{
				ColumnName: col.Name,
				FieldName:  templates.ToPascalCase(col.Name),
				GoType:     col.GoType,
				SQLType:    col.SQLType,
				IsNullable: col.IsNullable,
				IsPointer:  isPointer,
			}
			if col.IsPrimary {
				refPKField = templates.ToPascalCase(col.Name)
			}
		}

		ret = append(ret, templates.ManyToManyField{
			FieldName:         fieldName,
			GoType:            "[]" + refStructName,
			StructName:        refStructName,
			JunctionTable:     junctionInfo.sqlName,
			JunctionTableTag:  junctionInfo.plainName,
			LeftFKColumn:      m2m.LeftFKColumn,
			RightFKColumn:     m2m.RightFKColumn,
			ReferencedTable:   refInfo.sqlName,
			ReferencedPKField: refPKField,
			ReferencedColumns: referencedColumns,
		})
	}

	return ret
}
