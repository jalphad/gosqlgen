package templates

import (
	"bytes"
	_ "embed"
	"regexp"
	"strings"
	"text/template"
)

//go:embed models/table_struct.tmpl
var tableStructTemplate string

//go:embed query/per_table/table_query_expressions.tmpl
var tableQueryExpressionsTemplate string

//go:embed query_builder.tmpl
var queryBuilderTemplate string

//go:embed common.tmpl
var commonTypesTemplate string

//go:embed field_references.tmpl
var fieldReferencesTemplate string

//go:embed join_builders.tmpl
var joinBuildersTemplate string

//go:embed db_wrapper.tmpl
var dbWrapperTemplate string

//go:embed collection_loaders.tmpl
var collectionLoadersTemplate string

// QueryBuilderData contains the data for rendering the query builder template
type QueryBuilderData struct {
	BaseName           string
	BuilderName        string
	StructName         string
	TableName          string
	Columns            []Column
	NonPrimaryColumns  []Column
	NonSequenceColumns []Column
	PrimaryKey         string
	PrimaryKeyField    string
	PrimaryKeyType     string
	ForeignKeys        []ForeignKeyData        // FK information for join handling
	ReverseRelations   []ReverseRelLoaderField // 1-to-many relationships
	ManyToManyRels     []ManyToManyLoaderField // Many-to-many relationships
}

// Column represents template column data
type Column struct {
	ColumnName string
	FieldName  string
	GoType     string
	IsNullable bool
	IsPointer  bool // true if the Go field is a pointer type
}

// ForeignKeyData contains FK information for query building
type ForeignKeyData struct {
	JoinedFieldName      string   // e.g., "User" (field name in struct)
	ReferencedTable      string   // e.g., "users"
	ReferencedStructName string   // e.g., "UsersDto"
	ReferencedColumns    []Column // columns in the referenced table
}

// TableStructData contains the data for rendering the table struct template
type TableStructData struct {
	StructName         string
	TableName          string
	ReceiverName       string
	Fields             []StructField
	JoinedFields       []JoinedField
	ReverseRelFields   []ReverseRelField
	ManyToManyFields   []ManyToManyField
	NonSequenceColumns []StructField
	PrimaryKeys        []string
	PrimaryKeyFields   []string
}

// StructField represents a field in the generated struct
type StructField struct {
	FieldName  string
	ColumnName string
	GoType     string
	SQLType    string
	StructTags string
	IsPointer  bool
}

// JoinedField represents a joined relationship field
type JoinedField struct {
	FieldName         string   // e.g., "User", "Author", "Sender"
	GoType            string   // e.g., "*UsersDto"
	StructName        string   // e.g., "UsersDto"
	ReferencedTable   string   // e.g., "users"
	FKColumn          string   // e.g., "user_id"
	ReferencedColumn  string   // e.g., "id"
	ReferencedColumns []Column // columns in the referenced table
}

// ReverseRelLoaderField contains data for reverse relationship loader
type ReverseRelLoaderField struct {
	FieldName       string // "Posts"
	FromTable       string // "posts"
	FromStructName  string // "PostsDto"
	FromTableMethod string // "Posts" (DB wrapper method name)
	FKFieldName     string // "UserId"
}

// ReverseRelField represents a reverse one-to-many relationship field
type ReverseRelField struct {
	FieldName   string // e.g., "Posts"
	GoType      string // e.g., "[]*PostsDto"
	StructName  string // "PostsDto"
	FromTable   string // e.g., "posts"
	FKColumn    string // e.g., "user_id"
	FromPKField string // e.g., "Id"
}

// ManyToManyField represents a many-to-many relationship field
type ManyToManyField struct {
	FieldName         string // e.g., "Tags"
	GoType            string // e.g., "[]*TagsDto"
	StructName        string // "PostsDto"
	JunctionTable     string // e.g., "post_tags"
	LeftFKColumn      string // e.g., "post_id"
	RightFKColumn     string // e.g., "tag_id"
	ReferencedTable   string // e.g., "tags"
	ReferencedPKField string // e.g., "Id"
}

// FieldReferencesData contains the data for rendering field references template
type FieldReferencesData struct {
	BaseName         string
	StructName       string
	TableName        string
	FieldsTypeName   string
	ReceiverName     string
	Fields           []FieldRefData
	ReverseRelations []ReverseRelField
	ManyToManyRels   []ManyToManyField
}

// FieldRefData represents a field reference
type FieldRefData struct {
	FieldName  string
	ColumnName string
	GoType     string
}

// JoinBuildersData contains the data for rendering join builders template
type JoinBuildersData struct {
	BaseName    string
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
	ReferencedBaseName   string
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
	StructName  string // "TagsDto
}

// CollectionLoaderData contains data for rendering collection loader methods
type CollectionLoaderData struct {
	StructName       string
	TableName        string
	ReceiverName     string
	HasPrimaryKey    bool
	PrimaryKeyField  string
	PrimaryKeyType   string
	ReverseRelFields []ReverseRelLoaderField
	ManyToManyFields []ManyToManyLoaderField
}

// ManyToManyLoaderField contains data for M2M loader
type ManyToManyLoaderField struct {
	FieldName             string // "Tags"
	JunctionTable         string // "post_tags"
	LeftFKColumn          string // "post_id"
	RightFKColumn         string // "tag_id"
	ReferencedTable       string // "tags"
	ReferencedStructName  string // "TagsDto"
	ReferencedTableMethod string // "Tags" (DB wrapper method name)
	ReferencedPKField     string // "Id"
	ReferencedPKType      string // "string"
}

// RenderTableStruct renders the table struct template with the given data
func RenderTableStruct(data TableStructData) (string, error) {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"isPtr": func(goType string) bool {
			return strings.HasPrefix(goType, "*")
		},
		"hasSuffix":    strings.HasSuffix,
		"toLower":      strings.ToLower,
		"toPascalCase": ToPascalCase,
		"join":         strings.Join,
		"needsCustomUnmarshal": func(fields []StructField) bool {
			for _, field := range fields {
				re := regexp.MustCompile(`(?i)^(date|time(stamp)?)( with(out)? time zone)?$`)
				if strings.HasSuffix(field.GoType, "time.Time") &&
					re.MatchString(field.SQLType) {
					return true
				}
			}

			return false
		},
	}

	t, err := template.New("tableStruct").Funcs(funcMap).Parse(tableStructTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderColumnExpressions(data TableStructData) (string, error) {
	funcMap := template.FuncMap{
		"toTypeExpression": ToTypeExpression,
	}

	t, err := template.New("columnExpressions").Funcs(funcMap).Parse(tableQueryExpressionsTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderQueryBuilder renders the query builder template with the given data
func RenderQueryBuilder(data QueryBuilderData) (string, error) {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"add":     func(a, b int) int { return a + b },
		"toLower": strings.ToLower,
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

// RenderFieldReferences renders the field references template with the given data
func RenderFieldReferences(data FieldReferencesData) (string, error) {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"toLower":      strings.ToLower,
		"toPascalCase": ToPascalCase,
	}

	t, err := template.New("fieldReferences").Funcs(funcMap).Parse(fieldReferencesTemplate)
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

// RenderCollectionLoaders renders the collection loaders template with the given data
func RenderCollectionLoaders(data CollectionLoaderData) (string, error) {
	t, err := template.New("collectionLoaders").Parse(collectionLoadersTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ToPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[0:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

func ToTypeExpression(goType string) string {
	switch goType {
	case "int64":
		return "IntColumnExpression"
	case "float64":
		return "FloatColumnExpression"
	case "pgtype.Numeric":
		return "NumericType"
	case "bool":
		return "BoolType"
	case "string":
		return "StringType"
	case "time.Time":
		return "TimeType"
	case "[]byte":
		return "BytesType"
	case "json.Rawmessage":
		return "JsonType"
	case "uuid.UUID":
		return "UUIDType"
	default:
		return "<UnknownAstType>"
	}
}
