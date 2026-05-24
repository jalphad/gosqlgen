package ast

// QueryType represents the type of SQL query being generated
type QueryType string

const (
	QueryTypeSelect QueryType = "SELECT"
	QueryTypeInsert QueryType = "INSERT"
	QueryTypeUpdate QueryType = "UPDATE"
	QueryTypeDelete QueryType = "DELETE"
)

// QueryPart represents the current part of the query being generated
type QueryPart string

const (
	QueryPartSelectList QueryPart = "SELECT_LIST"
	QueryPartFrom       QueryPart = "FROM"
	QueryPartJoin       QueryPart = "JOIN"
	QueryPartWhere      QueryPart = "WHERE"
	QueryPartGroupBy    QueryPart = "GROUP_BY"
	QueryPartHaving     QueryPart = "HAVING"
	QueryPartOrderBy    QueryPart = "ORDER_BY"
	QueryPartLimit      QueryPart = "LIMIT"
	QueryPartReturning  QueryPart = "RETURNING"
	QueryPartInto       QueryPart = "INTO"
	QueryPartValues     QueryPart = "VALUES"
	QueryPartSet        QueryPart = "SET"
	QueryPartUsing      QueryPart = "USING"
)

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

	// Type indicates the type of query (SELECT, INSERT, UPDATE, DELETE)
	Type QueryType

	// CurrentPart indicates which part of the query is currently being generated
	CurrentPart QueryPart
}

type ColumnReference struct {
	Table  string
	Column string
	Alias  string
}
