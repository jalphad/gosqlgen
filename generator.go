package gosqlgen

import (
    "bytes"
    "fmt"
    "strings"
    "text/template"
)

// Generator handles code generation with type-safe queries
type Generator struct {
    parser      *Parser
    packageName string
    imports     map[string]bool
}

// NewGenerator creates a new code generator
func NewGenerator(parser *Parser) *Generator {
    return &Generator{
        parser:      parser,
        packageName: "models",
        imports: map[string]bool{
            "database/sql": true,
            "fmt":          true,
            "time":         true,
            "strings":      true,
        },
    }
}

// SetPackageName sets the package name for generated code
func (g *Generator) SetPackageName(name string) {
    g.packageName = name
}

// Generate generates Go code for all parsed tables
func (g *Generator) Generate() (string, error) {
    var buf bytes.Buffer
    
    // Write package declaration
    fmt.Fprintf(&buf, "package %s\n\n", g.packageName)
    
    // Write imports
    g.writeImports(&buf)
    
    // Generate common types
    g.generateCommonTypes(&buf)
    
    // Generate structs for each table
    for _, table := range g.parser.GetTables() {
        if err := g.generateTableStruct(&buf, table); err != nil {
            return "", err
        }
    }
    
    // Generate field references for type safety
    for _, table := range g.parser.GetTables() {
        if err := g.generateFieldReferences(&buf, table); err != nil {
            return "", err
        }
    }
    
    // Generate query builders with type-safe methods
    for _, table := range g.parser.GetTables() {
        if err := g.generateTypeSafeQueryBuilder(&buf, table); err != nil {
            return "", err
        }
    }
    
    // Generate join builders
    for _, table := range g.parser.GetTables() {
        if err := g.generateJoinBuilders(&buf, table); err != nil {
            return "", err
        }
    }
    
    // Generate database wrapper
    if err := g.generateDatabaseWrapper(&buf); err != nil {
        return "", err
    }
    
    return buf.String(), nil
}

// writeImports writes import statements
func (g *Generator) writeImports(buf *bytes.Buffer) {
    buf.WriteString("import (\n")
    for imp := range g.imports {
        fmt.Fprintf(buf, "\t\"%s\"\n", imp)
    }
    buf.WriteString(")\n\n")
}

// generateCommonTypes generates common types used across queries
func (g *Generator) generateCommonTypes(buf *bytes.Buffer) {
    buf.WriteString(`// OrderDirection represents sort order
type OrderDirection string

const ASC OrderDirection = "ASC"
const DESC OrderDirection = "DESC"

// ComparisonOp represents comparison operators
type ComparisonOp string

const EQ ComparisonOp = "="
const NEQ ComparisonOp = "!="
const GT ComparisonOp = ">"
const GTE ComparisonOp = ">="
const LT ComparisonOp = "<"
const LTE ComparisonOp = "<="
const LIKE ComparisonOp = "LIKE"
const IN ComparisonOp = "IN"

// LogicalOp represents logical operators
type LogicalOp string

const AND LogicalOp = "AND"
const OR LogicalOp = "OR"

// FieldRef represents a type-safe field reference
type FieldRef struct {
    Table  string
    Column string
    GoType string
    Alias  string
}

// As creates an aliased field reference
func (f *FieldRef) As(alias string) *FieldRef {
    return &FieldRef{
        Table:  f.Table,
        Column: f.Column,
        GoType: f.GoType,
        Alias:  alias,
    }
}

// String returns the SQL representation
func (f *FieldRef) String() string {
    if f.Alias != "" {
        return fmt.Sprintf("%s.%s AS %s", f.Table, f.Column, f.Alias)
    }
    return fmt.Sprintf("%s.%s", f.Table, f.Column)
}

// Condition represents a WHERE condition
type Condition struct {
    Field    *FieldRef
    Op       ComparisonOp
    Value    interface{}
    Logical  LogicalOp
    SubConds []Condition
}

// JoinClause represents a JOIN operation
type JoinClause struct {
    Type       string
    Table      string
    LeftField  *FieldRef
    RightField *FieldRef
}

// OrderClause represents an ORDER BY clause
type OrderClause struct {
    Field     *FieldRef
    Direction OrderDirection
}

`)
}

// generateTableStruct generates a struct for a table
func (g *Generator) generateTableStruct(buf *bytes.Buffer, table *Table) error {
    structName := g.toPascalCase(table.Name)
    
    fmt.Fprintf(buf, "// %s represents the %s table\n", structName, table.Name)
    fmt.Fprintf(buf, "type %s struct {\n", structName)
    
    for _, col := range table.Columns {
        fieldName := g.toPascalCase(col.Name)
        goType := col.GoType
        
        if col.IsNullable && !col.IsPrimary {
            goType = "*" + goType
        }
        
        tags := g.buildStructTags(col)
        fmt.Fprintf(buf, "\t%s %s %s\n", fieldName, goType, tags)
    }
    
    buf.WriteString("}\n\n")
    
    fmt.Fprintf(buf, "// TableName returns the table name for %s\n", structName)
    fmt.Fprintf(buf, "func (%s *%s) TableName() string {\n", strings.ToLower(structName[0:1]), structName)
    fmt.Fprintf(buf, "\treturn \"%s\"\n", table.Name)
    buf.WriteString("}\n\n")
    
    return nil
}

// generateFieldReferences generates type-safe field references
func (g *Generator) generateFieldReferences(buf *bytes.Buffer, table *Table) error {
    structName := g.toPascalCase(table.Name)
    fieldsTypeName := structName + "Fields"
    
    // Generate fields struct
    fmt.Fprintf(buf, "// %s provides type-safe field references for %s\n", fieldsTypeName, structName)
    fmt.Fprintf(buf, "type %s struct{}\n\n", fieldsTypeName)
    
    fmt.Fprintf(buf, "// Fields returns field references for %s\n", structName)
    fmt.Fprintf(buf, "var %sTable = %s{}\n\n", structName, fieldsTypeName)
    
    // Generate field accessor methods
    for _, col := range table.Columns {
        fieldName := g.toPascalCase(col.Name)

        // Generate field reference type
        fmt.Fprintf(buf, "// %s returns a field reference for %s.%s\n", fieldName, structName, fieldName)
        fmt.Fprintf(buf, "func (%s %s) %s() *FieldRef {\n", strings.ToLower(fieldsTypeName[0:1]), fieldsTypeName, fieldName)
        fmt.Fprintf(buf, "\treturn &FieldRef{\n")
        fmt.Fprintf(buf, "\t\tTable:  \"%s\",\n", table.Name)
        fmt.Fprintf(buf, "\t\tColumn: \"%s\",\n", col.Name)
        fmt.Fprintf(buf, "\t\tGoType: \"%s\",\n", col.GoType)
        fmt.Fprintf(buf, "\t}\n")
        buf.WriteString("}\n\n")
    }

    return nil
}

// generateTypeSafeQueryBuilder generates type-safe query builder
func (g *Generator) generateTypeSafeQueryBuilder(buf *bytes.Buffer, table *Table) error {
    structName := g.toPascalCase(table.Name)
    builderName := structName + "Query"
    
    tmpl := `// {{ .BuilderName }} is a type-safe query builder for {{ .StructName }}
type {{ .BuilderName }} struct {
    db         *sql.DB
    tx         *sql.Tx
    selectFields []*FieldRef
    conditions []Condition
    joins      []JoinClause
    orderBy    []OrderClause
    groupBy    []*FieldRef
    having     []Condition
    limit      *int
    offset     *int
    logical    LogicalOp
}

// New{{ .BuilderName }} creates a new query builder
func New{{ .BuilderName }}(db *sql.DB) *{{ .BuilderName }} {
    return &{{ .BuilderName }}{
        db:      db,
        logical: AND,
    }
}

// WithTx sets the transaction
func (q *{{ .BuilderName }}) WithTx(tx *sql.Tx) *{{ .BuilderName }} {
    q.tx = tx
    return q
}

// Select specifies fields to select using type-safe field references
func (q *{{ .BuilderName }}) Select(fields ...*FieldRef) *{{ .BuilderName }} {
    q.selectFields = fields
    return q
}

// SelectAll selects all fields
func (q *{{ .BuilderName }}) SelectAll() *{{ .BuilderName }} {
    q.selectFields = []*FieldRef{
        {{- range .Columns }}
        {{ $.StructName }}Table.{{ .FieldName }}(),
        {{- end }}
    }
    return q
}

{{/* Generate type-safe Where methods for each field */}}
{{- range .Columns }}

// Where{{ .FieldName }} adds a condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}(op ComparisonOp, value {{ .GoType }}) *{{ $.BuilderName }} {
    q.conditions = append(q.conditions, Condition{
        Field:   {{ $.StructName }}Table.{{ .FieldName }}(),
        Op:      op,
        Value:   value,
        Logical: q.logical,
    })
    return q
}

// Where{{ .FieldName }}Eq adds an equality condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Eq(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(EQ, value)
}

// Where{{ .FieldName }}NotEq adds a not-equal condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}NotEq(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(NEQ, value)
}

{{- if or (eq .GoType "int64") (eq .GoType "float64") (eq .GoType "time.Time") }}

// Where{{ .FieldName }}Gt adds a greater-than condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Gt(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(GT, value)
}

// Where{{ .FieldName }}Gte adds a greater-than-or-equal condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Gte(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(GTE, value)
}

// Where{{ .FieldName }}Lt adds a less-than condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Lt(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(LT, value)
}

// Where{{ .FieldName }}Lte adds a less-than-or-equal condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Lte(value {{ .GoType }}) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(LTE, value)
}
{{- end }}

{{- if eq .GoType "string" }}

// Where{{ .FieldName }}Like adds a LIKE condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}Like(pattern string) *{{ $.BuilderName }} {
    return q.Where{{ .FieldName }}(LIKE, pattern)
}
{{- end }}

// Where{{ .FieldName }}In adds an IN condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}In(values ...{{ .GoType }}) *{{ $.BuilderName }} {
    q.conditions = append(q.conditions, Condition{
        Field:   {{ $.StructName }}Table.{{ .FieldName }}(),
        Op:      IN,
        Value:   values,
        Logical: q.logical,
    })
    return q
}

{{- if .IsNullable }}

// Where{{ .FieldName }}IsNull adds an IS NULL condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}IsNull() *{{ $.BuilderName }} {
    q.conditions = append(q.conditions, Condition{
        Field:   {{ $.StructName }}Table.{{ .FieldName }}(),
        Op:      EQ,
        Value:   nil,
        Logical: q.logical,
    })
    return q
}

// Where{{ .FieldName }}IsNotNull adds an IS NOT NULL condition for {{ .FieldName }}
func (q *{{ $.BuilderName }}) Where{{ .FieldName }}IsNotNull() *{{ $.BuilderName }} {
    q.conditions = append(q.conditions, Condition{
        Field:   {{ $.StructName }}Table.{{ .FieldName }}(),
        Op:      NEQ,
        Value:   nil,
        Logical: q.logical,
    })
    return q
}
{{- end }}

// OrderBy{{ .FieldName }} adds ORDER BY for {{ .FieldName }}
func (q *{{ $.BuilderName }}) OrderBy{{ .FieldName }}(dir OrderDirection) *{{ $.BuilderName }} {
    q.orderBy = append(q.orderBy, OrderClause{
        Field:     {{ $.StructName }}Table.{{ .FieldName }}(),
        Direction: dir,
    })
    return q
}
{{- end }}

// And sets the logical operator to AND for subsequent conditions
func (q *{{ .BuilderName }}) And() *{{ .BuilderName }} {
    q.logical = AND
    return q
}

// Or sets the logical operator to OR for subsequent conditions
func (q *{{ .BuilderName }}) Or() *{{ .BuilderName }} {
    q.logical = OR
    return q
}

// WhereGroup adds a group of conditions
func (q *{{ .BuilderName }}) WhereGroup(fn func(*{{ .BuilderName }})) *{{ .BuilderName }} {
    subQuery := &{{ .BuilderName }}{logical: AND}
    fn(subQuery)
    
    q.conditions = append(q.conditions, Condition{
        Logical:  q.logical,
        SubConds: subQuery.conditions,
    })
    return q
}

// GroupBy adds GROUP BY clause
func (q *{{ .BuilderName }}) GroupBy(fields ...*FieldRef) *{{ .BuilderName }} {
    q.groupBy = append(q.groupBy, fields...)
    return q
}

// Having adds HAVING clause
func (q *{{ .BuilderName }}) Having(field *FieldRef, op ComparisonOp, value interface{}) *{{ .BuilderName }} {
    q.having = append(q.having, Condition{
        Field: field,
        Op:    op,
        Value: value,
    })
    return q
}

// Limit sets the LIMIT
func (q *{{ .BuilderName }}) Limit(limit int) *{{ .BuilderName }} {
    q.limit = &limit
    return q
}

// Offset sets the OFFSET
func (q *{{ .BuilderName }}) Offset(offset int) *{{ .BuilderName }} {
    q.offset = &offset
    return q
}

// buildQuery builds the SQL query
func (q *{{ .BuilderName }}) buildQuery() (string, []interface{}) {
    var query strings.Builder
    var args []interface{}
    argIndex := 1
    
    // SELECT clause
    query.WriteString("SELECT ")
    if len(q.selectFields) == 0 {
        query.WriteString("{{ .TableName }}.*")
    } else {
        fields := make([]string, len(q.selectFields))
        for i, field := range q.selectFields {
            fields[i] = field.String()
        }
        query.WriteString(strings.Join(fields, ", "))
    }
    
    query.WriteString(" FROM {{ .TableName }}")
    
    // JOIN clauses
    for _, join := range q.joins {
        query.WriteString(" ")
        query.WriteString(join.Type)
        query.WriteString(" ")
        query.WriteString(join.Table)
        query.WriteString(" ON ")
        query.WriteString(join.LeftField.String())
        query.WriteString(" = ")
        query.WriteString(join.RightField.String())
    }
    
    // WHERE clause
    if len(q.conditions) > 0 {
        query.WriteString(" WHERE ")
        whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
        query.WriteString(whereStr)
        args = append(args, whereArgs...)
    }
    
    // GROUP BY clause
    if len(q.groupBy) > 0 {
        query.WriteString(" GROUP BY ")
        groupFields := make([]string, len(q.groupBy))
        for i, field := range q.groupBy {
            groupFields[i] = field.String()
        }
        query.WriteString(strings.Join(groupFields, ", "))
    }
    
    // HAVING clause
    if len(q.having) > 0 {
        query.WriteString(" HAVING ")
        havingStr, havingArgs := q.buildConditions(q.having, &argIndex)
        query.WriteString(havingStr)
        args = append(args, havingArgs...)
    }
    
    // ORDER BY clause
    if len(q.orderBy) > 0 {
        query.WriteString(" ORDER BY ")
        orderClauses := make([]string, len(q.orderBy))
        for i, order := range q.orderBy {
            orderClauses[i] = fmt.Sprintf("%s %s", order.Field.String(), order.Direction)
        }
        query.WriteString(strings.Join(orderClauses, ", "))
    }
    
    // LIMIT clause
    if q.limit != nil {
        query.WriteString(fmt.Sprintf(" LIMIT %d", *q.limit))
    }
    
    // OFFSET clause
    if q.offset != nil {
        query.WriteString(fmt.Sprintf(" OFFSET %d", *q.offset))
    }
    
    return query.String(), args
}

// buildConditions builds WHERE/HAVING conditions
func (q *{{ .BuilderName }}) buildConditions(conditions []Condition, argIndex *int) (string, []interface{}) {
    var parts []string
    var args []interface{}

    for i, cond := range conditions {
        if len(cond.SubConds) > 0 {
            subStr, subArgs := q.buildConditions(cond.SubConds, argIndex)
            if i > 0 {
                parts = append(parts, fmt.Sprintf("%s (%s)", cond.Logical, subStr))
            } else {
                parts = append(parts, fmt.Sprintf("(%s)", subStr))
            }
            args = append(args, subArgs...)
        } else {
            var condStr string
            if cond.Value == nil {
                if cond.Op == EQ {
                    condStr = fmt.Sprintf("%s IS NULL", cond.Field.String())
                } else {
                    condStr = fmt.Sprintf("%s IS NOT NULL", cond.Field.String())
                }
            } else if cond.Op == IN {
                // Handle IN clause with reflection to support any slice type
                v := cond.Value
                switch vals := v.(type) {
                case []int64:
                    placeholders := make([]string, len(vals))
                    for j, val := range vals {
                        placeholders[j] = fmt.Sprintf("$%d", *argIndex)
                        *argIndex++
                        args = append(args, val)
                    }
                    condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
                case []string:
                    placeholders := make([]string, len(vals))
                    for j, val := range vals {
                        placeholders[j] = fmt.Sprintf("$%d", *argIndex)
                        *argIndex++
                        args = append(args, val)
                    }
                    condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
                case []float64:
                    placeholders := make([]string, len(vals))
                    for j, val := range vals {
                        placeholders[j] = fmt.Sprintf("$%d", *argIndex)
                        *argIndex++
                        args = append(args, val)
                    }
                    condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
                case []bool:
                    placeholders := make([]string, len(vals))
                    for j, val := range vals {
                        placeholders[j] = fmt.Sprintf("$%d", *argIndex)
                        *argIndex++
                        args = append(args, val)
                    }
                    condStr = fmt.Sprintf("%s IN (%s)", cond.Field.String(), strings.Join(placeholders, ", "))
                default:
                    // Fallback for other types - this shouldn't normally happen
                    condStr = fmt.Sprintf("%s IN (?)", cond.Field.String())
                }
            } else {
                condStr = fmt.Sprintf("%s %s $%d", cond.Field.String(), cond.Op, *argIndex)
                *argIndex++
                args = append(args, cond.Value)
            }

            if i > 0 {
                parts = append(parts, fmt.Sprintf("%s %s", cond.Logical, condStr))
            } else {
                parts = append(parts, condStr)
            }
        }
    }

    return strings.Join(parts, " "), args
}

// Find executes the query and returns results
func (q *{{ .BuilderName }}) Find() ([]*{{ .StructName }}, error) {
    query, args := q.buildQuery()
    
    var rows *sql.Rows
    var err error
    
    if q.tx != nil {
        rows, err = q.tx.Query(query, args...)
    } else {
        rows, err = q.db.Query(query, args...)
    }
    
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var results []*{{ .StructName }}
    for rows.Next() {
        result := &{{ .StructName }}{}
        if err := q.scanInto(rows, result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }
    
    return results, rows.Err()
}

// FindOne returns a single result
func (q *{{ .BuilderName }}) FindOne() (*{{ .StructName }}, error) {
    q.Limit(1)
    results, err := q.Find()
    if err != nil {
        return nil, err
    }
    if len(results) == 0 {
        return nil, sql.ErrNoRows
    }
    return results[0], nil
}

// Count returns the count
func (q *{{ .BuilderName }}) Count() (int64, error) {
    // Save and restore select fields
    originalSelect := q.selectFields
    countField := &FieldRef{Table: "", Column: "COUNT(*)", GoType: "int64"}
    q.selectFields = []*FieldRef{countField}
    defer func() { q.selectFields = originalSelect }()

    query, args := q.buildQuery()

    var count int64
    if q.tx != nil {
        err := q.tx.QueryRow(query, args...).Scan(&count)
        return count, err
    }
    err := q.db.QueryRow(query, args...).Scan(&count)
    return count, err
}

// scanInto scans a row into a struct
func (q *{{ .BuilderName }}) scanInto(rows *sql.Rows, dest *{{ .StructName }}) error {
    // If custom fields selected, use dynamic scanning
    if len(q.selectFields) > 0 && q.selectFields[0].Table != "" {
        // This would need more complex implementation for custom field scanning
        // For now, scan all fields
    }
    
    return rows.Scan(
        {{- range .Columns }}
        &dest.{{ .FieldName }},
        {{- end }}
    )
}

// Insert inserts a new record
func (q *{{ .BuilderName }}) Insert(record *{{ .StructName }}) error {
    query := "INSERT INTO {{ .TableName }} (" +
        "{{- range $i, $col := .Columns }}{{if $i}}, {{end}}{{ $col.Name }}{{- end }}" +
        ") VALUES (" +
        "{{- range $i, $col := .Columns }}{{if $i}}, {{end}}${{ add $i 1 }}{{- end }}" +
        "){{ if .PrimaryKey }} RETURNING {{ .PrimaryKey }}{{ end }}"

    {{ if .PrimaryKey -}}
    var err error
    if q.tx != nil {
        err = q.tx.QueryRow(query,
            {{- range .Columns }}
            record.{{ .FieldName }},
            {{- end }}
        ).Scan(&record.{{ .PrimaryKeyField }})
    } else {
        err = q.db.QueryRow(query,
            {{- range .Columns }}
            record.{{ .FieldName }},
            {{- end }}
        ).Scan(&record.{{ .PrimaryKeyField }})
    }
    return err
    {{- else -}}
    var err error
    if q.tx != nil {
        _, err = q.tx.Exec(query,
            {{- range .Columns }}
            record.{{ .FieldName }},
            {{- end }}
        )
    } else {
        _, err = q.db.Exec(query,
            {{- range .Columns }}
            record.{{ .FieldName }},
            {{- end }}
        )
    }
    return err
    {{- end }}
}

// Update updates a record using its primary key
func (q *{{ .BuilderName }}) Update(record *{{ .StructName }}) error {
    {{ if not .PrimaryKey -}}
    return fmt.Errorf("table {{ .TableName }} has no primary key")
    {{- else -}}
    query := "UPDATE {{ .TableName }} SET " +
        "{{- range $i, $col := .NonPrimaryColumns }}{{if $i}}, {{end}}{{ $col.Name }} = ${{ add $i 2 }}{{- end }}" +
        " WHERE {{ .PrimaryKey }} = $1"

    var result sql.Result
    var err error

    if q.tx != nil {
        result, err = q.tx.Exec(query,
            record.{{ .PrimaryKeyField }},
            {{- range .NonPrimaryColumns }}
            record.{{ .FieldName }},
            {{- end }}
        )
    } else {
        result, err = q.db.Exec(query,
            record.{{ .PrimaryKeyField }},
            {{- range .NonPrimaryColumns }}
            record.{{ .FieldName }},
            {{- end }}
        )
    }

    if err != nil {
        return err
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if affected == 0 {
        return sql.ErrNoRows
    }

    return nil
    {{- end }}
}

// UpdateFields updates specific fields for matching records
func (q *{{ .BuilderName }}) UpdateFields(updates map[*FieldRef]interface{}) (int64, error) {
    if len(updates) == 0 {
        return 0, fmt.Errorf("no fields to update")
    }
    
    var query strings.Builder
    var args []interface{}
    argIndex := 1
    
    query.WriteString("UPDATE {{ .TableName }} SET ")
    
    setClauses := make([]string, 0, len(updates))
    for field, value := range updates {
        setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field.Column, argIndex))
        args = append(args, value)
        argIndex++
    }
    query.WriteString(strings.Join(setClauses, ", "))
    
    // Add WHERE conditions
    if len(q.conditions) > 0 {
        query.WriteString(" WHERE ")
        whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
        query.WriteString(whereStr)
        args = append(args, whereArgs...)
    }
    
    var result sql.Result
    var err error
    
    if q.tx != nil {
        result, err = q.tx.Exec(query.String(), args...)
    } else {
        result, err = q.db.Exec(query.String(), args...)
    }
    
    if err != nil {
        return 0, err
    }
    
    return result.RowsAffected()
}

// Delete deletes matching records
func (q *{{ .BuilderName }}) Delete() (int64, error) {
    var query strings.Builder
    var args []interface{}
    argIndex := 1
    
    query.WriteString("DELETE FROM {{ .TableName }}")
    
    // Add WHERE conditions
    if len(q.conditions) > 0 {
        query.WriteString(" WHERE ")
        whereStr, whereArgs := q.buildConditions(q.conditions, &argIndex)
        query.WriteString(whereStr)
        args = append(args, whereArgs...)
    }
    
    var result sql.Result
    var err error
    
    if q.tx != nil {
        result, err = q.tx.Exec(query.String(), args...)
    } else {
        result, err = q.db.Exec(query.String(), args...)
    }
    
    if err != nil {
        return 0, err
    }
    
    return result.RowsAffected()
}
`
    
    // Get primary key info
    var primaryKey string
    var primaryKeyField string
    var primaryKeyType string = "int64" // default
    
    for _, col := range table.Columns {
        if col.IsPrimary {
            primaryKey = col.Name
            primaryKeyField = g.toPascalCase(col.Name)
            primaryKeyType = col.GoType
            break
        }
    }
    
    // Prepare template data
    data := struct {
        BuilderName       string
        StructName        string
        TableName         string
        Columns           []templateColumn
        NonPrimaryColumns []templateColumn
        PrimaryKey        string
        PrimaryKeyField   string
        PrimaryKeyType    string
    }{
        BuilderName:     builderName,
        StructName:      structName,
        TableName:       table.Name,
        PrimaryKey:      primaryKey,
        PrimaryKeyField: primaryKeyField,
        PrimaryKeyType:  primaryKeyType,
    }
    
    // Convert columns for template
    for _, col := range table.Columns {
        tc := templateColumn{
            Name:       col.Name,
            FieldName:  g.toPascalCase(col.Name),
            GoType:     col.GoType,
            IsNullable: col.IsNullable,
        }
        data.Columns = append(data.Columns, tc)
        
        if !col.IsPrimary {
            data.NonPrimaryColumns = append(data.NonPrimaryColumns, tc)
        }
    }
    
    // Create template with custom functions
    funcMap := template.FuncMap{
        "add": func(a, b int) int { return a + b },
    }
    
    t, err := template.New("queryBuilder").Funcs(funcMap).Parse(tmpl)
    if err != nil {
        return err
    }
    
    return t.Execute(buf, data)
}

type templateColumn struct {
    Name       string
    FieldName  string
    GoType     string
    IsNullable bool
}

// generateJoinBuilders generates type-safe join builders
func (g *Generator) generateJoinBuilders(buf *bytes.Buffer, table *Table) error {
    structName := g.toPascalCase(table.Name)
    builderName := structName + "Query"
    
    // Generate join methods for foreign keys
    for _, fk := range table.ForeignKeys {
        referencedTable, ok := g.parser.GetTable(fk.ReferencedTable)
        if !ok {
            continue
        }
        
        referencedStructName := g.toPascalCase(referencedTable.Name)
        joinMethodName := "Join" + referencedStructName
        leftJoinMethodName := "LeftJoin" + referencedStructName
        
        // InnerJoin method
        fmt.Fprintf(buf, "// %s performs a type-safe inner join with %s\n", joinMethodName, referencedTable.Name)
        fmt.Fprintf(buf, "func (q *%s) %s() *%s {\n", builderName, joinMethodName, builderName)
        fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
        fmt.Fprintf(buf, "\t\tType:       \"INNER JOIN\",\n")
        fmt.Fprintf(buf, "\t\tTable:      \"%s\",\n", referencedTable.Name)
        fmt.Fprintf(buf, "\t\tLeftField:  %sTable.%s(),\n", structName, g.toPascalCase(fk.Column))
        fmt.Fprintf(buf, "\t\tRightField: %sTable.%s(),\n", referencedStructName, g.toPascalCase(fk.ReferencedColumn))
        fmt.Fprintf(buf, "\t})\n")
        fmt.Fprintf(buf, "\treturn q\n")
        fmt.Fprintf(buf, "}\n\n")
        
        // LeftJoin method
        fmt.Fprintf(buf, "// %s performs a type-safe left join with %s\n", leftJoinMethodName, referencedTable.Name)
        fmt.Fprintf(buf, "func (q *%s) %s() *%s {\n", builderName, leftJoinMethodName, builderName)
        fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
        fmt.Fprintf(buf, "\t\tType:       \"LEFT JOIN\",\n")
        fmt.Fprintf(buf, "\t\tTable:      \"%s\",\n", referencedTable.Name)
        fmt.Fprintf(buf, "\t\tLeftField:  %sTable.%s(),\n", structName, g.toPascalCase(fk.Column))
        fmt.Fprintf(buf, "\t\tRightField: %sTable.%s(),\n", referencedStructName, g.toPascalCase(fk.ReferencedColumn))
        fmt.Fprintf(buf, "\t})\n")
        fmt.Fprintf(buf, "\treturn q\n")
        fmt.Fprintf(buf, "}\n\n")
        
        // Generate join result struct
        joinStructName := structName + referencedStructName + "Join"
        fmt.Fprintf(buf, "// %s represents a join result between %s and %s\n", 
            joinStructName, table.Name, referencedTable.Name)
        fmt.Fprintf(buf, "type %s struct {\n", joinStructName)
        fmt.Fprintf(buf, "\t%s *%s\n", structName, structName)
        fmt.Fprintf(buf, "\t%s *%s\n", referencedStructName, referencedStructName)
        fmt.Fprintf(buf, "}\n\n")
    }
    
    // Generate custom join method for arbitrary tables
    fmt.Fprintf(buf, "// JoinOn performs a custom join with type-safe field references\n")
    fmt.Fprintf(buf, "func (q *%s) JoinOn(joinType string, table string, leftField, rightField *FieldRef) *%s {\n", builderName, builderName)
    fmt.Fprintf(buf, "\tq.joins = append(q.joins, JoinClause{\n")
    fmt.Fprintf(buf, "\t\tType:       joinType,\n")
    fmt.Fprintf(buf, "\t\tTable:      table,\n")
    fmt.Fprintf(buf, "\t\tLeftField:  leftField,\n")
    fmt.Fprintf(buf, "\t\tRightField: rightField,\n")
    fmt.Fprintf(buf, "\t})\n")
    fmt.Fprintf(buf, "\treturn q\n")
    fmt.Fprintf(buf, "}\n\n")
    
    return nil
}

// generateDatabaseWrapper generates the main database wrapper
func (g *Generator) generateDatabaseWrapper(buf *bytes.Buffer) error {
    buf.WriteString("// DB wraps the database connection\n")
    buf.WriteString("type DB struct {\n")
    buf.WriteString("\tconn *sql.DB\n")
    buf.WriteString("}\n\n")
    
    buf.WriteString("// NewDB creates a new database wrapper\n")
    buf.WriteString("func NewDB(conn *sql.DB) *DB {\n")
    buf.WriteString("\treturn &DB{conn: conn}\n")
    buf.WriteString("}\n\n")
    
    // Generate methods for each table
    for _, table := range g.parser.GetTables() {
        structName := g.toPascalCase(table.Name)
        methodName := structName
        builderName := structName + "Query"
        
        fmt.Fprintf(buf, "// %s returns a query builder for %s\n", methodName, table.Name)
        fmt.Fprintf(buf, "func (db *DB) %s() *%s {\n", methodName, builderName)
        fmt.Fprintf(buf, "\treturn New%s(db.conn)\n", builderName)
        buf.WriteString("}\n\n")
    }
    
    // Transaction support
    buf.WriteString("// Transaction executes a function within a transaction\n")
    buf.WriteString("func (db *DB) Transaction(fn func(*Tx) error) error {\n")
    buf.WriteString("\ttx, err := db.conn.Begin()\n")
    buf.WriteString("\tif err != nil {\n")
    buf.WriteString("\t\treturn err\n")
    buf.WriteString("\t}\n")
    buf.WriteString("\tdefer tx.Rollback()\n\n")
    buf.WriteString("\ttxWrapper := &Tx{tx: tx, db: db}\n")
    buf.WriteString("\tif err := fn(txWrapper); err != nil {\n")
    buf.WriteString("\t\treturn err\n")
    buf.WriteString("\t}\n\n")
    buf.WriteString("\treturn tx.Commit()\n")
    buf.WriteString("}\n\n")
    
    // Transaction wrapper
    buf.WriteString("// Tx wraps a database transaction\n")
    buf.WriteString("type Tx struct {\n")
    buf.WriteString("\ttx *sql.Tx\n")
    buf.WriteString("\tdb *DB\n")
    buf.WriteString("}\n\n")
    
    // Transaction methods for each table
    for _, table := range g.parser.GetTables() {
        structName := g.toPascalCase(table.Name)
        methodName := structName
        builderName := structName + "Query"
        
        fmt.Fprintf(buf, "// %s returns a query builder within the transaction\n", methodName)
        fmt.Fprintf(buf, "func (tx *Tx) %s() *%s {\n", methodName, builderName)
        fmt.Fprintf(buf, "\tq := New%s(tx.db.conn)\n", builderName)
        fmt.Fprintf(buf, "\tq.WithTx(tx.tx)\n")
        fmt.Fprintf(buf, "\treturn q\n")
        buf.WriteString("}\n\n")
    }
    
    return nil
}

// Helper methods remain the same
func (g *Generator) buildStructTags(col Column) string {
    var tags []string
    for key, value := range col.Tags {
        tags = append(tags, fmt.Sprintf(`%s:"%s"`, key, value))
    }
    if len(tags) == 0 {
        return ""
    }
    return "`" + strings.Join(tags, " ") + "`"
}

func (g *Generator) toPascalCase(s string) string {
    parts := strings.Split(s, "_")
    for i, part := range parts {
        if len(part) > 0 {
            parts[i] = strings.ToUpper(part[0:1]) + strings.ToLower(part[1:])
        }
    }
    return strings.Join(parts, "")
}
