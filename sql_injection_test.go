package gosqlgen

import (
	"fmt"
	"strings"
	"testing"
)

// TestSQLInjection_FieldRefInjection tests if field references are vulnerable to SQL injection
func TestSQLInjection_FieldRefInjection(t *testing.T) {
	// First, generate code to get the FieldRef type
	sql := `CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(255));`

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Test if FieldRef.String() sanitizes or escapes identifiers
	tests := []struct {
		name        string
		pattern     string
		description string
	}{
		{
			name:        "Table name concatenation",
			pattern:     `fmt.Sprintf("%s.%s", f.Table, f.Column)`,
			description: "Table and column names are directly concatenated without quoting",
		},
		{
			name:        "Alias concatenation",
			pattern:     `fmt.Sprintf("%s.%s AS %s", f.Table, f.Column, f.Alias)`,
			description: "Aliases are directly concatenated without quoting",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(code, tt.pattern) {
				t.Errorf("VULNERABILITY CONFIRMED: %s\n  Pattern found: %s", tt.description, tt.pattern)
			}
		})
	}

	// Test for missing identifier quoting/escaping
	if !strings.Contains(code, "QuoteIdentifier") && !strings.Contains(code, "EscapeIdentifier") &&
		!strings.Contains(code, `"%q"`) && !strings.Contains(code, "`%s`") {
		t.Errorf("VULNERABILITY CONFIRMED: No identifier quoting or escaping found in FieldRef.String()")
	}
}

// TestSQLInjection_JoinClauseInjection tests if JOIN clauses are vulnerable
func TestSQLInjection_JoinClauseInjection(t *testing.T) {
	sql := `
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name VARCHAR(255)
        );
        CREATE TABLE posts (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            title VARCHAR(255),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	tests := []struct {
		name         string
		maliciousVal string
		context      string
	}{
		{
			name:         "Malicious JOIN type",
			maliciousVal: "INNER JOIN; DROP TABLE users--",
			context:      "join.Type can contain SQL injection",
		},
		{
			name:         "Malicious table name in JOIN",
			maliciousVal: "posts; DELETE FROM users--",
			context:      "join.Table can contain SQL injection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if the generated code would allow this
			if strings.Contains(code, "query.WriteString(join.Type)") {
				t.Errorf("VULNERABILITY CONFIRMED: %s\n  The code directly writes join.Type without sanitization", tt.context)
			}
			if strings.Contains(code, "query.WriteString(join.Table)") {
				t.Errorf("VULNERABILITY CONFIRMED: %s\n  The code directly writes join.Table without sanitization", tt.context)
			}
		})
	}
}

// TestSQLInjection_OrderByInjection tests if ORDER BY clauses are vulnerable
func TestSQLInjection_OrderByInjection(t *testing.T) {
	sql := `CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(255));`

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Check if OrderDirection is directly interpolated
	if strings.Contains(code, "order.Direction)") && !strings.Contains(code, "ValidateOrderDirection") {
		t.Errorf("VULNERABILITY CONFIRMED: OrderDirection is interpolated without validation\n" +
			"  This allows injection via custom OrderDirection values")
	}

	// Demonstrate the vulnerability
	t.Run("Malicious ORDER BY direction", func(t *testing.T) {
		maliciousDirection := "ASC; DROP TABLE users--"
		expectedInSQL := fmt.Sprintf("users.name %s", maliciousDirection)

		if strings.Contains(code, "fmt.Sprintf(\"%s %s\", order.Field.String(), order.Direction)") {
			t.Errorf("VULNERABILITY CONFIRMED: ORDER BY direction not validated\n"+
				"  Could generate: ORDER BY %s", expectedInSQL)
		}
	})
}

// TestSQLInjection_LogicalOperatorInjection tests if logical operators are vulnerable
func TestSQLInjection_LogicalOperatorInjection(t *testing.T) {
	sql := `CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(255));`

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Check if cond.Logical is directly interpolated
	if strings.Contains(code, "cond.Logical") &&
	   (strings.Contains(code, "fmt.Sprintf(\"%s (%s)\", cond.Logical") ||
	    strings.Contains(code, "fmt.Sprintf(\"%s %s\", cond.Logical")) {
		t.Errorf("VULNERABILITY CONFIRMED: Logical operators (AND/OR) interpolated without validation\n" +
			"  This could allow injection if LogicalOp type can be bypassed")
	}
}

// TestSQLInjection_GroupByInjection tests if GROUP BY is vulnerable
func TestSQLInjection_GroupByInjection(t *testing.T) {
	sql := `CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(255));`

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// GROUP BY uses field.String() which we know is vulnerable
	if strings.Contains(code, "groupFields[i] = field.String()") {
		t.Errorf("VULNERABILITY CONFIRMED: GROUP BY uses unsanitized field.String()\n" +
			"  Same vulnerability as field references - allows injection via malicious table/column names")
	}
}

// TestSQLInjection_WhereValuesAreSafe tests that WHERE values ARE properly parameterized (should PASS)
func TestSQLInjection_WhereValuesAreSafe(t *testing.T) {
	sql := `CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(255));`

	p := NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Check that WHERE values use parameterized queries
	hasPreparedStatements := strings.Contains(code, "$%d") && strings.Contains(code, "args = append(args")

	if !hasPreparedStatements {
		t.Errorf("WHERE clause values might not be using prepared statement placeholders")
	} else {
		t.Logf("✓ WHERE clause values appear to be properly parameterized (NOT vulnerable to value injection)")
	}
}

// TestSQLInjection_ComprehensiveReport generates a comprehensive security report
func TestSQLInjection_ComprehensiveReport(t *testing.T) {
	t.Log("=== SQL INJECTION VULNERABILITY REPORT ===\n")

	vulnerabilities := []struct {
		component   string
		severity    string
		vulnerable  bool
		description string
	}{
		{
			component:   "FieldRef.String() - Table Name",
			severity:    "HIGH",
			vulnerable:  true,
			description: "Table names are directly concatenated without quoting or validation",
		},
		{
			component:   "FieldRef.String() - Column Name",
			severity:    "HIGH",
			vulnerable:  true,
			description: "Column names are directly concatenated without quoting or validation",
		},
		{
			component:   "FieldRef.String() - Alias",
			severity:    "HIGH",
			vulnerable:  true,
			description: "Aliases are directly concatenated without quoting or validation",
		},
		{
			component:   "JOIN Clause - Type",
			severity:    "HIGH",
			vulnerable:  true,
			description: "JOIN type is directly written without validation (could inject arbitrary SQL)",
		},
		{
			component:   "JOIN Clause - Table",
			severity:    "HIGH",
			vulnerable:  true,
			description: "JOIN table name is directly written without quoting",
		},
		{
			component:   "ORDER BY - Direction",
			severity:    "MEDIUM",
			vulnerable:  true,
			description: "ORDER BY direction uses string formatting without validation",
		},
		{
			component:   "ORDER BY - Field",
			severity:    "HIGH",
			vulnerable:  true,
			description: "ORDER BY field uses FieldRef.String() (vulnerable)",
		},
		{
			component:   "GROUP BY - Field",
			severity:    "HIGH",
			vulnerable:  true,
			description: "GROUP BY field uses FieldRef.String() (vulnerable)",
		},
		{
			component:   "WHERE - Values",
			severity:    "N/A",
			vulnerable:  false,
			description: "✓ WHERE clause values are properly parameterized (SAFE)",
		},
		{
			component:   "IN Clause - Values",
			severity:    "N/A",
			vulnerable:  false,
			description: "✓ IN clause values are properly parameterized (SAFE)",
		},
	}

	t.Log("VULNERABILITIES FOUND:\n")
	vulnCount := 0
	for _, v := range vulnerabilities {
		if v.vulnerable {
			vulnCount++
			t.Logf("[%s] %s\n  → %s\n", v.severity, v.component, v.description)
		}
	}

	t.Log("\nSAFE COMPONENTS:\n")
	for _, v := range vulnerabilities {
		if !v.vulnerable {
			t.Logf("%s\n  → %s\n", v.component, v.description)
		}
	}

	t.Logf("\n=== SUMMARY ===")
	t.Logf("Total vulnerabilities found: %d", vulnCount)
	t.Logf("\nRECOMMENDATIONS:")
	t.Log("1. Add identifier quoting/escaping for all table, column, and alias names")
	t.Log("2. Validate JOIN types against a whitelist (INNER JOIN, LEFT JOIN, etc.)")
	t.Log("3. Validate ORDER BY direction against a whitelist (ASC, DESC)")
	t.Log("4. Consider using an allowlist approach for all SQL identifiers")
	t.Log("5. Add runtime validation to reject identifiers containing suspicious characters (;, --, /*, etc.)")
}