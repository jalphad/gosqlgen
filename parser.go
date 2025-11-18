package gosqlgen

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
	Prefix              string // Derived from column name (e.g., "user" from "user_id")
	Suffix              string // Derived from column name (e.g., "id" from "user_id")
}

// ReverseRelation represents a one-to-many relationship from the "one" side
type ReverseRelation struct {
	FromTable *Table // Table with the FK
	FKColumn  string // FK column name (e.g., "user_id")
	Prefix    string // Semantic name from FK (e.g., "user")
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
	tables map[string]*Table
}

// NewParser creates a new SQL parser
func NewParser() *Parser {
	return &Parser{
		tables: make(map[string]*Table),
	}
}

// Parse parses CREATE TABLE statements
func (p *Parser) Parse(sql string) error {
	// Parse comment directives BEFORE normalizing (which removes comments)
	tableComments := p.parseTableComments(sql)

	// Remove comments and normalize whitespace
	normalizedSQL := p.normalizeSQL(sql)

	// Find all CREATE TABLE statements
	tableRegex := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)\s*\((.*?)\);`)
	matches := tableRegex.FindAllStringSubmatch(normalizedSQL, -1)

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		tableName := strings.ToLower(match[1])
		tableBody := match[2]

		table := &Table{
			Name:    tableName,
			Columns: []Column{},
		}

		// Merge comment directives if they exist
		if commentData, exists := tableComments[tableName]; exists {
			table.SkipRelConfigs = commentData.SkipRelConfigs
			table.NoJunction = commentData.NoJunction
		}

		if err := p.parseTableBody(table, tableBody); err != nil {
			return fmt.Errorf("error parsing table %s: %w", tableName, err)
		}

		p.tables[tableName] = table
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

// parseTableBody parses the body of a CREATE TABLE statement
func (p *Parser) parseTableBody(table *Table, body string) error {
	lines := p.splitRespectingParentheses(body)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(strings.ToUpper(line), "PRIMARY KEY") {
			// Parse primary key constraint
			p.parsePrimaryKey(table, line)
		} else if strings.HasPrefix(strings.ToUpper(line), "FOREIGN KEY") {
			// Parse foreign key constraint
			p.parseForeignKey(table, line)
		} else if strings.HasPrefix(strings.ToUpper(line), "INDEX") ||
			strings.HasPrefix(strings.ToUpper(line), "KEY") {
			// Parse index
			p.parseIndex(table, line)
		} else {
			// Parse column definition
			p.parseColumn(table, line)
		}
	}

	return nil
}

// splitRespectingParentheses splits a string by commas but ignores commas inside parentheses
func (p *Parser) splitRespectingParentheses(s string) []string {
	var result []string
	var current strings.Builder
	parenDepth := 0

	for _, ch := range s {
		if ch == '(' {
			parenDepth++
			current.WriteRune(ch)
		} else if ch == ')' {
			parenDepth--
			current.WriteRune(ch)
		} else if ch == ',' && parenDepth == 0 {
			// This is a top-level comma, use it as a delimiter
			result = append(result, current.String())
			current.Reset()
		} else {
			current.WriteRune(ch)
		}
	}

	// Don't forget the last piece
	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

// parseColumn parses a column definition
func (p *Parser) parseColumn(table *Table, def string) {
	parts := strings.Fields(def)
	if len(parts) < 2 {
		return
	}

	fullType := parts[1]
	if (fullType == "TIMESTAMP" || fullType == "TIME") && len(parts) > 4 &&
		strings.ToUpper(parts[2]) == "WITH" &&
		strings.ToUpper(parts[3]) == "TIME" &&
		strings.ToUpper(parts[4]) == "ZONE" {
		fullType += " WITH TIME ZONE"
	}

	column := Column{
		Name:       parts[0],
		SQLType:    fullType,
		IsNullable: true,
		Tags:       make(map[string]string),
	}

	// Map SQL type to Go type
	column.GoType = p.mapSQLTypeToGo(parts[1])

	// Check if column is a sequence type
	column.IsSequence = p.isSequenceType(parts[1])

	// Check if column has a DEFAULT clause
	column.HasDefault = p.hasDefaultClause(def)

	// Parse column constraints
	defUpper := strings.ToUpper(def)
	if strings.Contains(defUpper, "PRIMARY KEY") {
		column.IsPrimary = true
		column.IsNullable = false // Primary keys are never nullable
		table.PrimaryKeys = append(table.PrimaryKeys, column.Name)
	}
	if strings.Contains(defUpper, "NOT NULL") {
		column.IsNullable = false
	}
	if strings.Contains(defUpper, "UNIQUE") {
		column.IsUnique = true
	}
	if idx := strings.Index(defUpper, "DEFAULT"); idx != -1 {
		// Extract default value
		remaining := def[idx+7:]
		column.DefaultValue = strings.TrimSpace(strings.Split(remaining, " ")[0])
	}

	// Set struct tags
	column.Tags["db"] = column.Name
	column.Tags["json"] = p.toSnakeCase(column.Name)

	table.Columns = append(table.Columns, column)
}

// parseForeignKey parses a foreign key constraint
func (p *Parser) parseForeignKey(table *Table, def string) {
	// Example: FOREIGN KEY (user_id) REFERENCES users(id)
	fkRegex := regexp.MustCompile(`(?i)FOREIGN\s+KEY\s*\((\w+)\)\s*REFERENCES\s+(\w+)\s*\((\w+)\)`)
	matches := fkRegex.FindStringSubmatch(def)

	if len(matches) >= 4 {
		fk := ForeignKey{
			Column:              matches[1],
			ReferencedTableName: matches[2],
			ReferencedColumn:    matches[3],
		}
		table.ForeignKeys = append(table.ForeignKeys, fk)
	}
}

// parseFKColumnName parses a foreign key column name into prefix and suffix
// Returns prefix, suffix, and error if parsing fails
func parseFKColumnName(fkColumn string) (prefix, suffix string, err error) {
	parts := strings.Split(fkColumn, "_")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("FK column '%s' does not follow <prefix>_<suffix> pattern", fkColumn)
	}

	suffix = parts[len(parts)-1]
	prefix = strings.Join(parts[:len(parts)-1], "_")
	return prefix, suffix, nil
}

// validateForeignKeys validates all foreign keys after parsing
func (p *Parser) validateForeignKeys() error {
	for tableName, table := range p.tables {
		for i := range table.ForeignKeys {
			fk := &table.ForeignKeys[i]

			// Parse FK column name
			prefix, suffix, err := parseFKColumnName(fk.Column)
			if err != nil {
				return fmt.Errorf("invalid FK column name in table '%s': %w", tableName, err)
			}

			// Store parsed values
			fk.Prefix = prefix
			fk.Suffix = suffix

			// Validate suffix matches referenced column
			if suffix != fk.ReferencedColumn {
				return fmt.Errorf(
					"Invalid foreign key in table '%s'\n"+
						"  FK column: '%s'\n"+
						"  References: %s.%s\n"+
						"\n"+
						"  Foreign key column name '%s' has suffix '%s' but references column '%s'.\n"+
						"\n"+
						"  Expected FK column to be named: '%s_%s' (following pattern <prefix>_<referenced_column>)\n"+
						"\n"+
						"  Please rename the column to follow the convention: <semantic_prefix>_<referenced_column_name>\n"+
						"\n"+
						"  Examples of valid FK names:\n"+
						"    - user_id (references users.id)\n"+
						"    - author_id (references users.id)\n"+
						"    - sender_id (references users.id)",
					tableName,
					fk.Column,
					fk.ReferencedTableName,
					fk.ReferencedColumn,
					fk.Column,
					suffix,
					fk.ReferencedColumn,
					prefix,
					fk.ReferencedColumn,
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

// parsePrimaryKey parses a primary key constraint
func (p *Parser) parsePrimaryKey(table *Table, def string) {
	// Example: PRIMARY KEY (id)
	pkRegex := regexp.MustCompile(`(?i)PRIMARY\s+KEY\s*\((.*?)\)`)
	matches := pkRegex.FindStringSubmatch(def)

	if len(matches) >= 2 {
		keys := strings.Split(matches[1], ",")
		for _, key := range keys {
			key = strings.TrimSpace(key)
			table.PrimaryKeys = append(table.PrimaryKeys, key)

			// Update column to mark as primary
			for i := range table.Columns {
				if table.Columns[i].Name == key {
					table.Columns[i].IsPrimary = true
					table.Columns[i].IsNullable = false // Primary keys are never nullable
					break
				}
			}
		}
	}
}

// parseIndex parses an index definition
func (p *Parser) parseIndex(table *Table, def string) {
	// Simple index parsing - can be extended
	idxRegex := regexp.MustCompile(`(?i)(?:INDEX|KEY)\s+\w+\s*\((.*?)\)`)
	matches := idxRegex.FindStringSubmatch(def)

	if len(matches) >= 2 {
		table.Indexes = append(table.Indexes, matches[1])
	}
}

// isSequenceType checks if a SQL type is a sequence (auto-increment)
func (p *Parser) isSequenceType(sqlType string) bool {
	sqlType = strings.ToUpper(sqlType)

	// Remove size specifications
	if idx := strings.Index(sqlType, "("); idx != -1 {
		sqlType = sqlType[:idx]
	}

	return sqlType == "SERIAL" || sqlType == "BIGSERIAL"
}

// hasDefaultClause checks if a column definition has a DEFAULT clause
func (p *Parser) hasDefaultClause(definition string) bool {
	defUpper := strings.ToUpper(definition)
	return strings.Contains(defUpper, "DEFAULT")
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
	case "DECIMAL", "NUMERIC", "REAL", "DOUBLE", "FLOAT":
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
		return "string"
	default:
		return "interface{}"
	}
}

// normalizeSQL removes comments and normalizes whitespace
func (p *Parser) normalizeSQL(sql string) string {
	// Remove single-line comments
	re := regexp.MustCompile(`--.*$`)
	sql = re.ReplaceAllString(sql, "")

	// Remove multi-line comments
	re = regexp.MustCompile(`/\*.*?\*/`)
	sql = re.ReplaceAllString(sql, "")

	// Normalize whitespace
	sql = strings.ReplaceAll(sql, "\n", " ")
	sql = strings.ReplaceAll(sql, "\t", " ")
	re = regexp.MustCompile(`\s+`)
	sql = re.ReplaceAllString(sql, " ")

	return strings.TrimSpace(sql)
}

// parseTableComments extracts comment directives before CREATE TABLE statements
func (p *Parser) parseTableComments(sql string) map[string]*Table {
	tableComments := make(map[string]*Table)

	// Find all table definitions with preceding comments
	lines := strings.Split(sql, "\n")
	var currentComments []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Collect comments
		if strings.HasPrefix(trimmed, "--") {
			currentComments = append(currentComments, trimmed)
		} else if strings.HasPrefix(strings.ToUpper(trimmed), "CREATE TABLE") {
			// Extract table name
			tableRegex := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)`)
			match := tableRegex.FindStringSubmatch(trimmed)
			if len(match) > 1 {
				tableName := strings.ToLower(match[1])

				table := &Table{
					Name:           tableName,
					SkipRelConfigs: []string{},
					NoJunction:     false,
				}

				// Parse comment directives
				for _, comment := range currentComments {
					p.parseCommentDirective(table, comment)
				}

				tableComments[tableName] = table
			}

			// Reset comments after table
			currentComments = nil
		} else if trimmed != "" && !strings.HasPrefix(trimmed, "--") {
			// Non-comment, non-CREATE TABLE line - reset comments
			currentComments = nil
		}
	}

	return tableComments
}

// parseCommentDirective parses a single comment line for directives
func (p *Parser) parseCommentDirective(table *Table, comment string) {
	comment = strings.TrimPrefix(strings.TrimSpace(comment), "--")
	comment = strings.TrimSpace(comment)

	// Parse @skip-rel directive
	if strings.HasPrefix(comment, "@skip-rel:") {
		skipTables := strings.TrimPrefix(comment, "@skip-rel:")
		skipTables = strings.TrimSpace(skipTables)

		// Split by comma and clean up
		for _, tableName := range strings.Split(skipTables, ",") {
			tableName = strings.TrimSpace(tableName)
			if tableName != "" {
				table.SkipRelConfigs = append(table.SkipRelConfigs, tableName)
			}
		}
	}

	// Parse @no-junction directive
	if strings.HasPrefix(comment, "@no-junction") {
		table.NoJunction = true
	}
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
					fieldName := p.capitalizeTableName(otherTableName)

					// Check for field name conflict
					if p.fieldNameExists(table, fieldName) {
						fieldName = fieldName + "List"
					}

					reverseRel := ReverseRelation{
						FromTable: otherTable,
						FKColumn:  fk.Column,
						Prefix:    fk.Prefix,
						FieldName: fieldName,
					}

					table.ReverseRelationships = append(table.ReverseRelationships, reverseRel)
				}
			}
		}
	}

	return nil
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

// fieldNameExists checks if a field name conflicts with existing fields or joined fields
func (p *Parser) fieldNameExists(table *Table, fieldName string) bool {
	// Check against column names (converted to PascalCase)
	for _, col := range table.Columns {
		if p.capitalizeTableName(col.Name) == fieldName {
			return true
		}
	}

	// Check against joined fields (which use FK prefix)
	for _, fk := range table.ForeignKeys {
		if p.capitalizeTableName(fk.Prefix) == fieldName {
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

// toSnakeCase converts a string to snake_case
func (p *Parser) toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// GetTables returns all parsed tables
func (p *Parser) GetTables() map[string]*Table {
	return p.tables
}

// GetTable returns a specific table
func (p *Parser) GetTable(name string) (*Table, bool) {
	table, ok := p.tables[strings.ToLower(name)]
	return table, ok
}
