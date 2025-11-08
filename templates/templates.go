package templates

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed query_builder.tmpl
var queryBuilderTemplate string

// QueryBuilderData contains the data for rendering the query builder template
type QueryBuilderData struct {
	BuilderName         string
	StructName          string
	TableName           string
	Columns             []Column
	NonPrimaryColumns   []Column
	NonSequenceColumns  []Column
	PrimaryKey          string
	PrimaryKeyField     string
	PrimaryKeyType      string
}

// Column represents template column data
type Column struct {
	Name       string
	FieldName  string
	GoType     string
	IsNullable bool
	IsPointer  bool // true if the Go field is a pointer type
}

// RenderQueryBuilder renders the query builder template with the given data
func RenderQueryBuilder(data QueryBuilderData) (string, error) {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	t, err := template.New("queryBuilder").Funcs(funcMap).Parse(queryBuilderTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}