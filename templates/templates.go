package templates

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed query_builder.tmpl
var queryBuilderTemplate string

//go:embed common.tmpl
var commonTypesTemplate string

//go:embed table_struct.tmpl
var tableStructTemplate string

//go:embed field_references.tmpl
var fieldReferencesTemplate string

//go:embed join_builders.tmpl
var joinBuildersTemplate string

//go:embed db_wrapper.tmpl
var dbWrapperTemplate string

// QueryBuilderData contains the data for rendering the query builder template
type QueryBuilderData struct {
	BuilderName        string
	StructName         string
	TableName          string
	Columns            []Column
	NonPrimaryColumns  []Column
	NonSequenceColumns []Column
	PrimaryKey         string
	PrimaryKeyField    string
	PrimaryKeyType     string
}

// Column represents template column data
type Column struct {
	Name       string
	FieldName  string
	GoType     string
	IsNullable bool
	IsPointer  bool // true if the Go field is a pointer type
}

// TableStructData contains the data for rendering the table struct template
type TableStructData struct {
	StructName   string
	TableName    string
	ReceiverName string
	Fields       []StructField
}

// StructField represents a field in the generated struct
type StructField struct {
	FieldName  string
	GoType     string
	StructTags string
}

// FieldReferencesData contains the data for rendering field references template
type FieldReferencesData struct {
	StructName     string
	TableName      string
	FieldsTypeName string
	ReceiverName   string
	Fields         []FieldRefData
}

// FieldRefData represents a field reference
type FieldRefData struct {
	FieldName  string
	ColumnName string
	GoType     string
}

// JoinBuildersData contains the data for rendering join builders template
type JoinBuildersData struct {
	BuilderName string
	StructName  string
	TableName   string
	Joins       []JoinData
}

// JoinData represents a foreign key join
type JoinData struct {
	JoinMethodName       string
	LeftJoinMethodName   string
	JoinStructName       string
	ReferencedTableName  string
	ReferencedStructName string
	LeftFieldName        string
	RightFieldName       string
}

// DBWrapperData contains the data for rendering the DB wrapper template
type DBWrapperData struct {
	Tables []TableMethod
}

// TableMethod represents a table method in the DB wrapper
type TableMethod struct {
	MethodName  string
	BuilderName string
	TableName   string
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

// RenderCommonTypes renders the common types template
func RenderCommonTypes() (string, error) {
	t, err := template.New("commonTypes").Parse(commonTypesTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, nil); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderTableStruct renders the table struct template with the given data
func RenderTableStruct(data TableStructData) (string, error) {
	t, err := template.New("tableStruct").Parse(tableStructTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderFieldReferences renders the field references template with the given data
func RenderFieldReferences(data FieldReferencesData) (string, error) {
	t, err := template.New("fieldReferences").Parse(fieldReferencesTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderJoinBuilders renders the join builders template with the given data
func RenderJoinBuilders(data JoinBuildersData) (string, error) {
	t, err := template.New("joinBuilders").Parse(joinBuildersTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderDBWrapper renders the DB wrapper template with the given data
func RenderDBWrapper(data DBWrapperData) (string, error) {
	t, err := template.New("dbWrapper").Parse(dbWrapperTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
