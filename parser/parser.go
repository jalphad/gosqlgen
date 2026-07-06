package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// Column represents a database column
type Column struct {
	Name         string
	SQLType      string
	GoType       string
	IsPrimary    bool
	IsNullable   bool
	IsUnique     bool
	IsSequence   bool // true for SERIAL, BIGSERIAL (sequence-based auto-increment)
	HasDefault   bool // true for any DEFAULT clause
	DefaultValue string
	Tags         map[string]string
}

// ForeignKey represents a foreign key constraint
type ForeignKey struct {
	Column              string
	ReferencedTableName string
	ReferencedColumn    string
}

// ReverseRelation represents a one-to-many relationship from the "one" side
type ReverseRelation struct {
	FromTable *Table // Table with the FK
	FKColumn  string // FK column name (e.g., "user_id")
	FieldName string // Collection field name (e.g., "Posts")
}

// ManyToManyRelation represents a many-to-many relationship through a junction table
type ManyToManyRelation struct {
	FieldName       string // "Tags"
	JunctionTable   *Table // "post_tags"
	LeftFKColumn    string // "post_id"
	RightFKColumn   string // "tag_id"
	ReferencedTable *Table // "tags"
}

// JunctionTableInfo contains information about a detected junction table
type JunctionTableInfo struct {
	LeftTable     *Table
	LeftFKColumn  string
	RightTable    *Table
	RightFKColumn string
}

// Table represents a database table
type Table struct {
	Name                 string
	Columns              []Column
	PrimaryKeys          []string
	ForeignKeys          []ForeignKey
	Indexes              []string
	SkipRelConfigs       []string             // From @skip-rel comment
	NoJunction           bool                 // From @no-junction comment
	ReverseRelationships []ReverseRelation    // Computed reverse rels
	ManyToManyRels       []ManyToManyRelation // Computed M2M rels
	JunctionInfo         *JunctionTableInfo   // nil if not junction
}

// Parser handles SQL parsing
type Parser struct {
	sequences map[string]string
	tables    map[string]*Table
}

// NewParser creates a new SQL parser
func NewParser() *Parser {
	return &Parser{
		sequences: make(map[string]string),
		tables:    make(map[string]*Table),
	}
}

// Parse parses CREATE TABLE statements
func (p *Parser) Parse(sql string) error {
	// Clean and normalize the SQL
	sql = p.cleanSQL(sql)

	// Parse all sequence statements
	p.parseSequences(sql)

	// Extract all CREATE TABLE statements
	statements := p.extractCreateTableStatements(sql)

	// Parse each CREATE TABLE statement
	for _, stmt := range statements {
		if err := p.parseCreateTable(stmt); err != nil {
			return fmt.Errorf("error parsing CREATE TABLE: %w", err)
		}
	}

	// Extract and parse ALTER TABLE statements for constraints
	alterStatements := p.extractAlterTableStatements(sql)
	for _, stmt := range alterStatements {
		if err := p.parseAlterTable(stmt); err != nil {
			return fmt.Errorf("error parsing ALTER TABLE: %w", err)
		}
	}

	// Validate foreign keys after all tables are parsed
	if err := p.validateForeignKeys(); err != nil {
		return err
	}

	// Validate skip-rel configurations
	if err := p.validateSkipRelConfigs(); err != nil {
		return err
	}

	// Detect junction tables
	if err := p.detectJunctionTables(); err != nil {
		return err
	}

	// Build reverse relationships
	if err := p.buildReverseRelationships(); err != nil {
		return err
	}

	// Build many-to-many relationships
	if err := p.buildManyToManyRelationships(); err != nil {
		return err
	}

	return nil
}

func (p *Parser) GetTable(table string) (*Table, bool) {
	ret, ok := p.tables[table]
	return ret, ok
}

// GetTables returns all parsed tables
func (p *Parser) GetTables() map[string]*Table {
	return p.tables
}

// cleanSQL removes comments and normalizes whitespace
func (p *Parser) cleanSQL(sql string) string {
	// Remove single-line comments but preserve special comments like @skip-rel
	reComments := regexp.MustCompile(`(--[^']*)$`)
	lines := strings.Split(sql, "\n")
	var cleaned []string
	for _, line := range lines {
		// Keep special directive comments
		if strings.Contains(line, "@skip-rel") || strings.Contains(line, "@no-junction") {
			cleaned = append(cleaned, line)
			continue
		}
		// Remove regular SQL comments
		match := reComments.FindAllStringIndex(line, -1)
		if match != nil {
			idx := match[0][0]
			line = line[:idx]
		}
		if strings.TrimSpace(line) != "" {
			cleaned = append(cleaned, line)
		}
	}
	return strings.Join(cleaned, "\n")
}

func (p *Parser) parseSequences(sql string) {
	// Regular expressions for parsing
	reCreateSeq := regexp.MustCompile(`CREATE SEQUENCE\s+(?:(\w+)\.)?(\w+);`)
	reAlterSeq := regexp.MustCompile(`ALTER SEQUENCE\s+(?:(\w+)\.)?(\w+)\s+OWNED BY\s+(?:(\w+)\.)?(\w+)\.(\w+);`)

	// First pass: find CREATE SEQUENCE statements
	sequenceNames := make(map[string]struct{})
	matches := reCreateSeq.FindAllStringSubmatch(sql, -1)
	for _, match := range matches {
		seqName := match[2]
		sequenceNames[seqName] = struct{}{}
	}

	matches = reAlterSeq.FindAllStringSubmatch(sql, -1)
	for _, match := range matches {
		seqName := match[2]
		tableName := match[4]
		columnName := match[5]

		p.sequences[fmt.Sprintf("%s.%s", tableName, columnName)] = seqName
	}
}

// extractCreateTableStatements extracts individual CREATE TABLE statements
func (p *Parser) extractCreateTableStatements(sql string) []string {
	var statements []string
	re := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?(\w+\.)?(\w+)\s*\(((?:[^;]|\n)*?)\);`)
	matches := re.FindAllStringSubmatch(sql, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			tableName := match[2]
			tableBody := match[3]
			statements = append(statements, fmt.Sprintf("CREATE TABLE %s (%s);", tableName, tableBody))
		}
	}

	return statements
}

// extractAlterTableStatements extracts ALTER TABLE statements
func (p *Parser) extractAlterTableStatements(sql string) []string {
	var statements []string
	re := regexp.MustCompile(`(?i)ALTER\s+TABLE\s+(?:ONLY\s+)?(?:public\.)?(\w+)\s+(.*?);`)
	matches := re.FindAllString(sql, -1)
	statements = append(statements, matches...)
	return statements
}

// parseCreateTable parses a single CREATE TABLE statement
func (p *Parser) parseCreateTable(stmt string) error {
	// Extract table name
	tableNameRe := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(?:(?:IF\s+NOT\s+EXISTS|public\.)\s+)?(\w+)`)
	nameMatch := tableNameRe.FindStringSubmatch(stmt)
	if len(nameMatch) < 2 {
		return fmt.Errorf("could not extract table name from: %s", stmt)
	}

	tableName := nameMatch[1]
	table := &Table{
		Name:        tableName,
		Columns:     []Column{},
		PrimaryKeys: []string{},
		ForeignKeys: []ForeignKey{},
		Indexes:     []string{},
	}

	// Extract table body (content between parentheses)
	bodyRe := regexp.MustCompile(`\(((?:[^()]|\([^)]*\))*)\)`)
	bodyMatch := bodyRe.FindStringSubmatch(stmt)
	if len(bodyMatch) < 2 {
		return fmt.Errorf("could not extract table body from: %s", stmt)
	}

	body := bodyMatch[1]

	// Split by comma, but not commas inside parentheses
	items := p.splitTableItems(body)

	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		// Check for PRIMARY KEY constraint
		if matched, _ := regexp.MatchString(`(?i)^\s*PRIMARY\s+KEY`, item); matched {
			p.parsePrimaryKeyConstraint(table, item)
			continue
		}

		// Check for FOREIGN KEY constraint
		if matched, _ := regexp.MatchString(`(?i)^\s*FOREIGN\s+KEY`, item); matched {
			p.parseForeignKeyConstraint(table, item)
			continue
		}

		// Check for CONSTRAINT with PRIMARY KEY
		if matched, _ := regexp.MatchString(`(?i)^\s*CONSTRAINT\s+\w+\s+PRIMARY\s+KEY`, item); matched {
			p.parsePrimaryKeyConstraint(table, item)
			continue
		}

		// Check for CONSTRAINT with FOREIGN KEY
		if matched, _ := regexp.MatchString(`(?i)^\s*CONSTRAINT\s+\w+\s+FOREIGN\s+KEY`, item); matched {
			p.parseForeignKeyConstraint(table, item)
			continue
		}

		// Check for UNIQUE constraint
		if matched, _ := regexp.MatchString(`(?i)^\s*(?:CONSTRAINT\s+\w+\s+)?UNIQUE`, item); matched {
			continue // Skip table-level UNIQUE constraints for now
		}

		// Otherwise, it's a column definition
		if err := p.parseColumnDefinition(table, item); err != nil {
			return err
		}
	}

	p.tables[tableName] = table
	return nil
}

// splitTableItems splits table definition items by comma, respecting parentheses
func (p *Parser) splitTableItems(body string) []string {
	var items []string
	var current strings.Builder
	depth := 0

	var inComment bool
	for _, char := range body {
		switch {
		case char == '\'':
			inComment = !inComment
			current.WriteRune(char)
		case char == '(' && !inComment:
			depth++
			current.WriteRune(char)
		case char == ')' && !inComment:
			depth--
			current.WriteRune(char)
		case char == ',' && !inComment:
			if depth == 0 {
				items = append(items, current.String())
				current.Reset()
			} else {
				current.WriteRune(char)
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		items = append(items, current.String())
	}

	return items
}

// parseColumnDefinition parses a column definition
func (p *Parser) parseColumnDefinition(table *Table, def string) error {
	def = strings.TrimSpace(def)

	// Extract column name (first word)
	parts := strings.Fields(def)
	if len(parts) < 2 {
		return fmt.Errorf("invalid column definition: %s", def)
	}

	col := Column{
		Name:       parts[0],
		IsNullable: true, // Default to nullable
		Tags:       make(map[string]string),
	}

	// Check if column has an associated sequence
	if _, ok := p.sequences[fmt.Sprintf("%s.%s", table.Name, col.Name)]; ok {
		col.IsSequence = true
		col.HasDefault = true
	}

	// Parse the rest of the definition
	remainingDef := strings.Join(parts[1:], " ")

	// Extract SQL type
	col.SQLType = p.extractSQLType(remainingDef)

	// Map SQL type to Go type
	col.GoType = p.mapSQLTypeToGo(parts[1])

	// Check if column has an associated sequence or is of SERIAL or BIGSERIAL type
	if _, ok := p.sequences[fmt.Sprintf("%s.%s", table.Name, col.Name)]; ok {
		col.IsSequence = true
		col.HasDefault = true
	} else if matched, _ := regexp.MatchString(`(?i)\b(SERIAL|BIGSERIAL)\b`, remainingDef); matched {
		col.IsSequence = true
		col.HasDefault = true
		if strings.Contains(strings.ToUpper(remainingDef), "BIGSERIAL") {
			col.SQLType = "BIGSERIAL"
		} else {
			col.SQLType = "SERIAL"
		}
	}

	// Check for PRIMARY KEY
	if matched, _ := regexp.MatchString(`(?i)\bPRIMARY\s+KEY\b`, remainingDef); matched {
		col.IsPrimary = true
		table.PrimaryKeys = append(table.PrimaryKeys, col.Name)
	}

	// Check for NOT NULL
	if matched, _ := regexp.MatchString(`(?i)\bNOT\s+NULL\b`, remainingDef); matched {
		col.IsNullable = false
	}

	// Check for UNIQUE
	if matched, _ := regexp.MatchString(`(?i)\bUNIQUE\b`, remainingDef); matched {
		col.IsUnique = true
	}

	// Check for DEFAULT
	defaultRe := regexp.MustCompile(`(?i)\bDEFAULT\s+('.*'|[^,\s]+)`)
	if defaultMatch := defaultRe.FindStringSubmatch(remainingDef); len(defaultMatch) > 1 {
		col.HasDefault = true
		col.DefaultValue = strings.TrimSpace(defaultMatch[1])
	}

	// Check for inline REFERENCES (foreign key)
	referencesRe := regexp.MustCompile(`(?i)\bREFERENCES\s+(\w+)\s*\((\w+)\)`)
	if refMatch := referencesRe.FindStringSubmatch(remainingDef); len(refMatch) > 2 {
		fk := ForeignKey{
			Column:              col.Name,
			ReferencedTableName: refMatch[1],
			ReferencedColumn:    refMatch[2],
		}
		table.ForeignKeys = append(table.ForeignKeys, fk)
	}

	// Set struct tags
	col.Tags["db"] = col.Name
	col.Tags["json"] = toSnakeCase(col.Name)

	table.Columns = append(table.Columns, col)
	return nil
}

// extractSQLType extracts the SQL type from a column definition
func (p *Parser) extractSQLType(def string) string {
	// Match type with optional size/precision
	typeRe := regexp.MustCompile(`(?i)^((?:CHARACTER\s+VARYING|VARCHAR|INTEGER|BIGINT|SERIAL|BIGSERIAL|TEXT|BOOLEAN|TIMESTAMP(?:\s+(?:WITH|WITHOUT)\s+TIME\s+ZONE)?|DATE|TIME(?:\s+(?:WITH|WITHOUT)\s+TIME\s+ZONE)?|UUID|NUMERIC|DECIMAL|REAL|DOUBLE\s+PRECISION|SMALLINT|BIGINT)(?:\s*\(\d+(?:\s*,\s*\d+)?\))?)`)
	match := typeRe.FindStringSubmatch(def)
	if len(match) > 1 {
		return strings.ToUpper(strings.TrimSpace(match[1]))
	}

	// Fallback: return first word
	parts := strings.Fields(def)
	if len(parts) > 0 {
		return strings.ToUpper(parts[0])
	}

	return "TEXT"
}

// parsePrimaryKeyConstraint parses PRIMARY KEY constraint
func (p *Parser) parsePrimaryKeyConstraint(table *Table, constraint string) {
	// Extract column names from PRIMARY KEY (col1, col2, ...)
	pkRe := regexp.MustCompile(`(?i)PRIMARY\s+KEY\s*\(([^)]+)\)`)
	match := pkRe.FindStringSubmatch(constraint)
	if len(match) > 1 {
		cols := strings.Split(match[1], ",")
		for _, col := range cols {
			colName := strings.TrimSpace(col)
			table.PrimaryKeys = append(table.PrimaryKeys, colName)
			// Mark the column as primary
			for i := range table.Columns {
				if table.Columns[i].Name == colName {
					table.Columns[i].IsPrimary = true
					break
				}
			}
		}
	}
}

// parseForeignKeyConstraint parses FOREIGN KEY constraint
func (p *Parser) parseForeignKeyConstraint(table *Table, constraint string) {
	// FOREIGN KEY (col) REFERENCES table(col)
	fkRe := regexp.MustCompile(`(?i)FOREIGN\s+KEY\s*\((\w+)\)\s+REFERENCES\s+(?:public\.)?(\w+)\s*\((\w+)\)`)
	match := fkRe.FindStringSubmatch(constraint)
	if len(match) > 3 {
		fk := ForeignKey{
			Column:              match[1],
			ReferencedTableName: match[2],
			ReferencedColumn:    match[3],
		}
		table.ForeignKeys = append(table.ForeignKeys, fk)
	}
}

// parseAlterTable parses ALTER TABLE statements
func (p *Parser) parseAlterTable(stmt string) error {
	// Extract table name
	tableRe := regexp.MustCompile(`(?i)ALTER\s+TABLE\s+(?:ONLY\s+)?(?:public\.)?(\w+)`)
	match := tableRe.FindStringSubmatch(stmt)
	if len(match) < 2 {
		return nil
	}

	tableName := match[1]
	table, exists := p.tables[tableName]
	if !exists {
		return nil // Table not found, skip
	}

	// Check for ADD CONSTRAINT PRIMARY KEY
	if matched, _ := regexp.MatchString(`(?i)ADD\s+CONSTRAINT\s+\w+\s+PRIMARY\s+KEY`, stmt); matched {
		p.parsePrimaryKeyConstraint(table, stmt)
	}

	// Check for ADD CONSTRAINT FOREIGN KEY
	if matched, _ := regexp.MatchString(`(?i)ADD\s+CONSTRAINT\s+\w+\s+FOREIGN\s+KEY`, stmt); matched {
		p.parseForeignKeyConstraint(table, stmt)
	}

	// Check for ADD CONSTRAINT UNIQUE
	if matched, _ := regexp.MatchString(`(?i)ADD\s+CONSTRAINT\s+\w+\s+UNIQUE`, stmt); matched {
		uniqueRe := regexp.MustCompile(`(?i)UNIQUE\s*\(([^)]+)\)`)
		if uniqueMatch := uniqueRe.FindStringSubmatch(stmt); len(uniqueMatch) > 1 {
			cols := strings.Split(uniqueMatch[1], ",")
			for _, col := range cols {
				colName := strings.TrimSpace(col)
				for i := range table.Columns {
					if table.Columns[i].Name == colName {
						table.Columns[i].IsUnique = true
						break
					}
				}
			}
		}
	}

	// Check for ALTER COLUMN SET DEFAULT
	setDefaultRe := regexp.MustCompile(`(?i)ALTER\s+COLUMN\s+(\w+)\s+SET\s+DEFAULT\s+(.+)`)
	if defaultMatch := setDefaultRe.FindStringSubmatch(stmt); len(defaultMatch) > 2 {
		colName := defaultMatch[1]
		defaultVal := strings.TrimSpace(defaultMatch[2])
		for i := range table.Columns {
			if table.Columns[i].Name == colName {
				table.Columns[i].HasDefault = true
				table.Columns[i].DefaultValue = defaultVal
				break
			}
		}
	}

	return nil
}

// validateForeignKeys validates all foreign keys after parsing
func (p *Parser) validateForeignKeys() error {
	for tableName, table := range p.tables {
		for i := range table.ForeignKeys {
			fk := &table.ForeignKeys[i]

			foundLocalColumn := false
			for _, col := range table.Columns {
				if col.Name == fk.Column {
					foundLocalColumn = true
					break
				}
			}
			if !foundLocalColumn {
				return fmt.Errorf(
					"foreign key in table '%s' uses non-existent column '%s'",
					tableName,
					fk.Column,
				)
			}

			// Verify referenced table exists
			refTable, ok := p.tables[fk.ReferencedTableName]
			if !ok {
				return fmt.Errorf(
					"foreign key in table '%s' references non-existent table '%s'",
					tableName,
					fk.ReferencedTableName,
				)
			}

			// Verify referenced column exists in referenced table
			foundColumn := false
			for _, col := range refTable.Columns {
				if col.Name == fk.ReferencedColumn {
					foundColumn = true
					break
				}
			}

			if !foundColumn {
				return fmt.Errorf(
					"foreign key in table '%s' references non-existent column '%s' in table '%s'",
					tableName,
					fk.ReferencedColumn,
					fk.ReferencedTableName,
				)
			}
		}
	}

	return nil
}

// validateSkipRelConfigs validates that tables mentioned in @skip-rel exist
func (p *Parser) validateSkipRelConfigs() error {
	for tableName, table := range p.tables {
		for _, skipTable := range table.SkipRelConfigs {
			if _, exists := p.tables[skipTable]; !exists {
				return fmt.Errorf(
					"Table '%s' has @skip-rel for '%s', but table '%s' does not exist",
					tableName, skipTable, skipTable,
				)
			}
		}
	}
	return nil
}

// detectJunctionTables detects junction tables based on strict criteria
func (p *Parser) detectJunctionTables() error {
	for _, table := range p.tables {
		// Skip if explicitly opted out
		if table.NoJunction {
			continue
		}

		// Strict criteria: exactly 2 columns
		if len(table.Columns) != 2 {
			continue
		}

		// Check if both columns are foreign keys
		if len(table.ForeignKeys) != 2 {
			continue
		}

		// Check if both columns are part of primary key
		if len(table.PrimaryKeys) != 2 {
			continue
		}

		// Verify that the FK columns match the PK columns
		fkColumns := make(map[string]bool)
		for _, fk := range table.ForeignKeys {
			fkColumns[fk.Column] = true
		}

		pkColumns := make(map[string]bool)
		for _, pk := range table.PrimaryKeys {
			pkColumns[pk] = true
		}

		// Both FK columns must be in PK
		allMatch := true
		for fkCol := range fkColumns {
			if !pkColumns[fkCol] {
				allMatch = false
				break
			}
		}

		if !allMatch {
			continue
		}

		// This is a junction table!
		table.JunctionInfo = &JunctionTableInfo{
			LeftTable:     p.tables[table.ForeignKeys[0].ReferencedTableName],
			LeftFKColumn:  table.ForeignKeys[0].Column,
			RightTable:    p.tables[table.ForeignKeys[1].ReferencedTableName],
			RightFKColumn: table.ForeignKeys[1].Column,
		}
	}

	return nil
}

// buildReverseRelationships builds reverse relationship metadata for one-to-many relationships
func (p *Parser) buildReverseRelationships() error {
	// For each table, find all tables that reference it
	for tableName, table := range p.tables {
		for otherTableName, otherTable := range p.tables {
			// Skip self-references
			if tableName == otherTableName {
				continue
			}

			// Skip junction tables (they're handled via M2M)
			if otherTable.JunctionInfo != nil {
				continue
			}

			// Check if otherTable is in skip list
			skipTable := false
			for _, skip := range table.SkipRelConfigs {
				if skip == otherTableName {
					skipTable = true
					break
				}
			}
			if skipTable {
				continue
			}

			// Check if otherTable has a FK pointing to this table
			for _, fk := range otherTable.ForeignKeys {
				if fk.ReferencedTableName == tableName {
					// Create reverse relationship
					fieldName := p.reverseRelationFieldName(otherTableName, fk.Column)

					// Check for field name conflict
					if p.fieldNameExists(table, fieldName) {
						fieldName = fieldName + "List"
					}

					reverseRel := ReverseRelation{
						FromTable: otherTable,
						FKColumn:  fk.Column,
						FieldName: fieldName,
					}

					table.ReverseRelationships = append(table.ReverseRelationships, reverseRel)
				}
			}
		}
	}

	return nil
}

// buildManyToManyRelationships builds many-to-many relationship metadata
func (p *Parser) buildManyToManyRelationships() error {
	// For each junction table, add M2M relationships to both sides
	for _, junctionTable := range p.tables {
		if junctionTable.JunctionInfo == nil {
			continue
		}

		info := junctionTable.JunctionInfo
		// Add M2M relationship to left table
		p.buildLeftOrRightRelation(junctionTable, relationConfig{
			tableName:        info.LeftTable.Name,
			relatedTableName: info.RightTable.Name,
			fkColumn:         info.LeftFKColumn,
			relatedFKColumn:  info.RightFKColumn,
		})
		// Add M2M relationship to right table
		p.buildLeftOrRightRelation(junctionTable, relationConfig{
			tableName:        info.RightTable.Name,
			relatedTableName: info.LeftTable.Name,
			fkColumn:         info.RightFKColumn,
			relatedFKColumn:  info.LeftFKColumn,
		})
	}

	return nil
}

type relationConfig struct {
	tableName        string
	relatedTableName string
	fkColumn         string
	relatedFKColumn  string
}

func (p *Parser) buildLeftOrRightRelation(junctionTable *Table, cfg relationConfig) {
	if thisTable, exists := p.tables[cfg.tableName]; exists {
		// Check if other table is in skip list
		skipRelation := false
		for _, skip := range thisTable.SkipRelConfigs {
			if skip == cfg.relatedTableName {
				skipRelation = true
				break
			}
		}

		if !skipRelation {
			fieldName := p.capitalizeTableName(cfg.relatedTableName)

			// Check for field name conflict
			if p.fieldNameExists(thisTable, fieldName) {
				fieldName = fieldName + "List"
			}

			m2m := ManyToManyRelation{
				FieldName:       fieldName,
				JunctionTable:   junctionTable,
				LeftFKColumn:    cfg.fkColumn,        // Swapped for right table
				RightFKColumn:   cfg.relatedFKColumn, // Swapped for right table
				ReferencedTable: p.tables[cfg.relatedTableName],
			}

			thisTable.ManyToManyRels = append(thisTable.ManyToManyRels, m2m)
		}
	}
}

// capitalizeTableName capitalizes the first letter of each word in a table name
func (p *Parser) capitalizeTableName(tableName string) string {
	if tableName == "" {
		return ""
	}

	// Simple capitalization: just capitalize the first character
	runes := []rune(tableName)
	runes[0] = rune(strings.ToUpper(string(runes[0]))[0])
	return string(runes)
}

func (p *Parser) reverseRelationFieldName(fromTableName, fkColumn string) string {
	return p.toPascalName(fromTableName) + "By" + p.toPascalName(fkColumn)
}

func (p *Parser) joinedRelationFieldName(fkColumn string) string {
	return p.toPascalName(fkColumn) + "Ref"
}

func (p *Parser) toPascalName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
	}
	return strings.Join(parts, "")
}

// fieldNameExists checks if a field name conflicts with existing fields or joined fields
func (p *Parser) fieldNameExists(table *Table, fieldName string) bool {
	// Check against column names (converted to PascalCase)
	for _, col := range table.Columns {
		if p.toPascalName(col.Name) == fieldName {
			return true
		}
	}

	// Check against joined fields
	for _, fk := range table.ForeignKeys {
		if p.joinedRelationFieldName(fk.Column) == fieldName {
			return true
		}
	}

	// Check against already-added reverse relationships
	for _, rev := range table.ReverseRelationships {
		if rev.FieldName == fieldName {
			return true
		}
	}

	// Check against already-added M2M relationships
	for _, m2m := range table.ManyToManyRels {
		if m2m.FieldName == fieldName {
			return true
		}
	}

	return false
}

// mapSQLTypeToGo maps SQL types to Go types
func (p *Parser) mapSQLTypeToGo(sqlType string) string {
	sqlType = strings.ToUpper(sqlType)

	// Remove size specifications
	if idx := strings.Index(sqlType, "("); idx != -1 {
		sqlType = sqlType[:idx]
	}

	switch sqlType {
	case "INTEGER", "INT", "SMALLINT", "BIGINT":
		return "int64"
	case "SERIAL", "BIGSERIAL":
		return "int64"
	case "DECIMAL", "NUMERIC":
		return "pgtype.Numeric"
	case "REAL", "DOUBLE", "FLOAT":
		return "float64"
	case "BOOLEAN", "BOOL":
		return "bool"
	case "VARCHAR", "CHAR", "TEXT", "CHARACTER":
		return "string"
	case "DATE", "TIME", "TIMESTAMP", "TIME WITH TIME ZONE", "TIMESTAMP WITH TIME ZONE":
		return "time.Time"
	case "BYTEA", "BLOB":
		return "[]byte"
	case "JSON", "JSONB":
		return "json.RawMessage"
	case "UUID":
		return "uuid.UUID"
	default:
		return "any"
	}
}

// toSnakeCase converts a table name to snake case
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if notFirstOrLast(i, s) && isUpper(r) && (!isUpper(rune(s[i-1])) || !isUpper(rune(s[i+1]))) {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func notFirstOrLast(index int, s string) bool {
	return index > 0 && index < len(s)-1
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
