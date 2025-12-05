package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_TypeMapping(t *testing.T) {
	tests := []struct {
		sqlType string
		goType  string
	}{
		{"INTEGER", "int64"},
		{"INT", "int64"},
		{"BIGINT", "int64"},
		{"VARCHAR(255)", "string"},
		{"TEXT", "string"},
		{"BOOLEAN", "bool"},
		{"TIMESTAMP", "time.Time"},
		{"DATE", "time.Time"},
		{"FLOAT", "float64"},
		{"DECIMAL(10,2)", "float64"},
		{"JSON", "json.RawMessage"},
		{"BYTEA", "[]byte"},
	}

	p := NewParser()

	for _, tt := range tests {
		t.Run(tt.sqlType, func(t *testing.T) {
			got := p.mapSQLTypeToGo(tt.sqlType)
			if got != tt.goType {
				t.Errorf("mapSQLTypeToGo(%s) = %s, want %s", tt.sqlType, got, tt.goType)
			}
		})
	}
}

func TestToSnakeCase_CorrectlyModifiesString(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"Base case": {
			input:    "Test",
			expected: "test",
		},
		"Insert underscore for upper case": {
			input:    "TestCase",
			expected: "test_case",
		},
		"Do not insert underscores when repeating upper case": {
			input:    "TEST",
			expected: "test",
		},
		"Insert underscore before last repeating upper case": {
			input:    "TESTCase",
			expected: "test_case",
		},
		"Combination: AAA": {
			input:    "AAA",
			expected: "aaa",
		},
		"Combination: aAA": {
			input:    "aAA",
			expected: "a_aa",
		},
		"Combination: AAa": {
			input:    "AAa",
			expected: "a_aa",
		},
		"Combination: AaA": {
			input:    "AaA",
			expected: "aaa",
		},
		"Combination: AaAa": {
			input:    "AaAa",
			expected: "aa_aa",
		},
		"Something something": {
			input:    "AAAaaAAaAaAAaAAAaaA",
			expected: "aa_aaa_a_aa_aa_a_aa_aa_aaaa",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Act
			out := toSnakeCase(tc.input)

			// Assert
			assert.Equal(t, tc.expected, out)
		})
	}
}
