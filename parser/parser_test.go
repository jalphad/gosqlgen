package parser

import "testing"

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
