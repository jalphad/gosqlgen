package gosqlgen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jalphad/gosqlgen/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	// Arrange
	tests := []struct {
		name    string
		sql     string
		want    map[string][]string // table name -> column names
		wantErr bool
	}{
		{
			name: "simple table",
			sql: `CREATE TABLE users (
                id INTEGER PRIMARY KEY,
                name VARCHAR(255) NOT NULL,
                email VARCHAR(255) UNIQUE
            );`,
			want: map[string][]string{
				"users": {"id", "name", "email"},
			},
			wantErr: false,
		},
		{
			name: "table with foreign key",
			sql: `
                CREATE TABLE users (
                    id INTEGER PRIMARY KEY
                );
                CREATE TABLE posts (
                    id INTEGER PRIMARY KEY,
                    user_id INTEGER NOT NULL,
                    title VARCHAR(255),
                    FOREIGN KEY (user_id) REFERENCES users(id)
                );
            `,
			want: map[string][]string{
				"users": {"id"},
				"posts": {"id", "user_id", "title"},
			},
			wantErr: false,
		},
		{
			name: "multiple tables",
			sql: `
                CREATE TABLE users (
                    id INTEGER PRIMARY KEY,
                    name VARCHAR(255)
                );
                CREATE TABLE posts (
                    id INTEGER PRIMARY KEY,
                    user_id INTEGER,
                    title TEXT,
                    FOREIGN KEY (user_id) REFERENCES users(id)
                );
            `,
			want: map[string][]string{
				"users": {"id", "name"},
				"posts": {"id", "user_id", "title"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			p := parser.NewParser()

			// Act
			err := p.Parse(tt.sql)

			// Assert
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			tables := p.GetTables()
			assert.Len(t, tables, len(tt.want))
			for tableName, expectedCols := range tt.want {
				table, ok := tables[tableName]
				require.Truef(t, ok, "Table %s not found", tableName)

				actualCols := make([]string, 0, len(table.Columns))
				for _, col := range table.Columns {
					actualCols = append(actualCols, col.Name)
				}
				assert.Equal(t, expectedCols, actualCols)
			}
		})
	}
}

func TestParser_ForeignKeys(t *testing.T) {
	// Arrange
	sql := `
        CREATE TABLE users (
            id INTEGER PRIMARY KEY
        );
        CREATE TABLE posts (
            id INTEGER PRIMARY KEY,
            user_id INTEGER,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
        CREATE TABLE comments (
            id INTEGER PRIMARY KEY,
            post_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(id),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `

	p := parser.NewParser()

	// Act
	err := p.Parse(sql)

	// Assert
	require.NoError(t, err)

	posts, ok := p.GetTable("posts")
	require.True(t, ok, "Posts table not found")
	require.Len(t, posts.ForeignKeys, 1)
	assert.Equal(t, "user_id", posts.ForeignKeys[0].Column)
	assert.Equal(t, "users", posts.ForeignKeys[0].ReferencedTableName)

	comments, ok := p.GetTable("comments")
	require.True(t, ok, "Comments table not found")
	assert.Len(t, comments.ForeignKeys, 2)
}

func TestParser_ForeignKeysAllowNonConformingColumnNames(t *testing.T) {
	// Arrange
	sql := `
        CREATE TABLE users (
            id UUID PRIMARY KEY
        );
        CREATE TABLE accounts (
            uuid UUID PRIMARY KEY
        );
        CREATE TABLE posts (
            id INTEGER PRIMARY KEY,
            created_by UUID,
            updated_by UUID,
            owner UUID REFERENCES users(id),
            account_ref UUID,
            FOREIGN KEY (created_by) REFERENCES users(id),
            FOREIGN KEY (updated_by) REFERENCES users(id),
            FOREIGN KEY (account_ref) REFERENCES accounts(uuid)
        );
    `

	p := parser.NewParser()

	// Act
	err := p.Parse(sql)

	// Assert
	require.NoError(t, err)

	posts, ok := p.GetTable("posts")
	require.True(t, ok, "Posts table not found")
	require.Len(t, posts.ForeignKeys, 4)

	fksByColumn := make(map[string]parser.ForeignKey, len(posts.ForeignKeys))
	for _, fk := range posts.ForeignKeys {
		fksByColumn[fk.Column] = fk
	}

	assert.Equal(t, parser.ForeignKey{
		Column:              "created_by",
		ReferencedTableName: "users",
		ReferencedColumn:    "id",
	}, fksByColumn["created_by"])
	assert.Equal(t, parser.ForeignKey{
		Column:              "owner",
		ReferencedTableName: "users",
		ReferencedColumn:    "id",
	}, fksByColumn["owner"])
	assert.Equal(t, parser.ForeignKey{
		Column:              "account_ref",
		ReferencedTableName: "accounts",
		ReferencedColumn:    "uuid",
	}, fksByColumn["account_ref"])
}

func TestSQLGen_Integration(t *testing.T) {
	// Arrange
	schema := `
        CREATE TABLE departments (
            id INTEGER PRIMARY KEY,
            name VARCHAR(100) NOT NULL
        );
        
        CREATE TABLE employees (
            id INTEGER PRIMARY KEY,
            department_id INTEGER,
            first_name VARCHAR(50) NOT NULL,
            last_name VARCHAR(50) NOT NULL,
            email VARCHAR(255) UNIQUE,
            hired_at DATE,
            salary DECIMAL(10,2),
            is_active BOOLEAN DEFAULT true,
            FOREIGN KEY (department_id) REFERENCES departments(id)
        );
    `

	outputPath := t.TempDir()
	gen := New().
		WithPackagePath("example.local/app").
		WithPackageName("hr").
		WithOutputPath(outputPath)

	// Act
	err := gen.Parse(schema)

	// Assert
	require.NoError(t, err)
	tables := gen.GetTables()
	assert.Len(t, tables, 2)

	dept, ok := gen.GetTable("departments")
	require.True(t, ok, "Departments table not found")
	assert.Len(t, dept.Columns, 2)

	emp, ok := gen.GetTable("employees")
	require.True(t, ok, "Employees table not found")
	assert.Len(t, emp.Columns, 8)
	assert.Len(t, emp.ForeignKeys, 1)

	// Act
	err = gen.GenerateFiles()

	// Assert
	require.NoError(t, err)
	dbFile := readGeneratedFile(t, outputPath, "hr", "db.gen.go")
	assert.Contains(t, dbFile, "package hr")
	assert.Contains(t, dbFile, `"example.local/app/hr/models"`)
	assert.Contains(t, dbFile, "func (db *DB) Departments()")
	assert.Contains(t, dbFile, "func (db *DB) Employees()")

	departmentsFile := readGeneratedFile(t, outputPath, "hr", "models", "departments.gen.go")
	assert.Contains(t, departmentsFile, "type DepartmentsDto struct")

	employeesFile := readGeneratedFile(t, outputPath, "hr", "models", "employees.gen.go")
	assert.Contains(t, employeesFile, "type EmployeesDto struct")
	assert.Contains(t, employeesFile, "DepartmentIdRef *DepartmentsDto")
}

func readGeneratedFile(t *testing.T, outputPath string, elem ...string) string {
	t.Helper()

	pathParts := append([]string{outputPath}, elem...)
	content, err := os.ReadFile(filepath.Join(pathParts...))
	require.NoErrorf(t, err, "Failed to read generated file %s", filepath.Join(elem...))

	return string(content)
}
