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
	Column           string
	ReferencedTable  string
	ReferencedColumn string
	Prefix           string // Derived from column name (e.g., "user" from "user_id")
	Suffix           string // Derived from column name (e.g., "id" from "user_id")
}

// Table represents a database table
type Table struct {
	Name        string
	Columns     []Column
	PrimaryKeys []string
	ForeignKeys []ForeignKey
	Indexes     []string
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
	// Remove comments and normalize whitespace
	sql = p.normalizeSQL(sql)

	// Find all CREATE TABLE statements
	tableRegex := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)\s*\((.*?)\);`)
	matches := tableRegex.FindAllStringSubmatch(sql, -1)

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

		if err := p.parseTableBody(table, tableBody); err != nil {
			return fmt.Errorf("error parsing table %s: %w", tableName, err)
		}

		p.tables[tableName] = table
	}

	// Validate foreign keys after all tables are parsed
	if err := p.validateForeignKeys(); err != nil {
		return err
	}

	return nil
}

// parseTableBody parses the body of a CREATE TABLE statement
func (p *Parser) parseTableBody(table *Table, body string) error {
	lines := strings.Split(body, ",")

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

// parseColumn parses a column definition
func (p *Parser) parseColumn(table *Table, def string) {
	parts := strings.Fields(def)
	if len(parts) < 2 {
		return
	}

	column := Column{
		Name:       parts[0],
		SQLType:    parts[1],
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
			Column:           matches[1],
			ReferencedTable:  matches[2],
			ReferencedColumn: matches[3],
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
					fk.ReferencedTable,
					fk.ReferencedColumn,
					fk.Column,
					suffix,
					fk.ReferencedColumn,
					prefix,
					fk.ReferencedColumn,
				)
			}

			// Verify referenced table exists
			refTable, ok := p.tables[fk.ReferencedTable]
			if !ok {
				return fmt.Errorf(
					"foreign key in table '%s' references non-existent table '%s'",
					tableName,
					fk.ReferencedTable,
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
					fk.ReferencedTable,
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
	case "DATE", "TIME", "TIMESTAMP", "DATETIME":
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
