package ast

type QueryContext struct {
	// PrimaryTable is set before SQL generation
	PrimaryTable string

	// JoinedTables is populated during SQL generation
	JoinedTables map[string]JoinType

	// AllTables contains all tables (primary + joined)
	AllTables []string

	// ColumnReferences tracks all column references
	ColumnReferences []ColumnReference

	// Error tracks any errors that occur during SQL generation
	Error error
}

type ColumnReference struct {
	Table  string
	Column string
	Alias  string
}
