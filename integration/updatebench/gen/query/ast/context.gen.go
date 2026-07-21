package ast

import "fmt"

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

	parent              *QueryContext
	inheritParentRanges bool
	namedRelations      map[string]*TableAlias
	rangeVariables      map[string]*TableAlias
}

func newChildQueryContext(parent *QueryContext, inheritParentRanges bool) *QueryContext {
	return &QueryContext{
		parent:              parent,
		inheritParentRanges: inheritParentRanges,
	}
}

func (c *QueryContext) finishChild(child *QueryContext) {
	if c != nil && child != nil && c.Error == nil {
		c.Error = child.Error
	}
}

func (c *QueryContext) registerNamedRelation(alias *TableAlias) {
	if c == nil || alias == nil || c.Error != nil {
		return
	}
	if c.namedRelations == nil {
		c.namedRelations = make(map[string]*TableAlias)
	}
	if _, ok := c.namedRelations[alias.GetName()]; ok {
		c.Error = fmt.Errorf("relation %q is already declared in this query scope", alias.GetName())
		return
	}
	c.namedRelations[alias.GetName()] = alias
}

func (c *QueryContext) lookupNamedRelation(alias *TableAlias) bool {
	if c == nil || alias == nil {
		return false
	}
	for scope := c; scope != nil; scope = scope.parent {
		if declared, ok := scope.namedRelations[alias.GetName()]; ok {
			return declared == alias
		}
	}
	return false
}

func (c *QueryContext) registerRangeVariable(alias *TableAlias) {
	if c == nil || alias == nil || c.Error != nil {
		return
	}
	if c.rangeVariables == nil {
		c.rangeVariables = make(map[string]*TableAlias)
	}
	if _, ok := c.rangeVariables[alias.GetName()]; ok {
		c.Error = fmt.Errorf("table alias %q is already declared in this query scope", alias.GetName())
		return
	}
	c.rangeVariables[alias.GetName()] = alias
}

func (c *QueryContext) lookupRangeVariable(alias *TableAlias) bool {
	if c == nil || alias == nil {
		return false
	}
	if declared, ok := c.rangeVariables[alias.GetName()]; ok {
		return declared == alias
	}
	if c.inheritParentRanges {
		return c.parent.lookupRangeVariable(alias)
	}
	return false
}

type ColumnReference struct {
	Table  string
	Column string
	Alias  string
}
