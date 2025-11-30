package gosqlgen

import (
	"strings"
	"testing"

	"github.com/jalphad/gosqlgen/parser"
)

func TestParser_Parse(t *testing.T) {
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
                CREATE TABLE posts (
                    id INTEGER PRIMARY KEY,
                    user_id INTEGER NOT NULL,
                    title VARCHAR(255),
                    FOREIGN KEY (user_id) REFERENCES users(id)
                );
            `,
			want: map[string][]string{
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
			p := parser.NewParser()
			err := p.Parse(tt.sql)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				tables := p.GetTables()

				// Check that we have the right number of tables
				if len(tables) != len(tt.want) {
					t.Errorf("Expected %d tables, got %d", len(tt.want), len(tables))
				}

				// Check each table
				for tableName, expectedCols := range tt.want {
					table, ok := tables[tableName]
					if !ok {
						t.Errorf("Table %s not found", tableName)
						continue
					}

					// Check columns
					if len(table.Columns) != len(expectedCols) {
						t.Errorf("Table %s: expected %d columns, got %d",
							tableName, len(expectedCols), len(table.Columns))
						continue
					}

					for i, col := range table.Columns {
						if col.Name != expectedCols[i] {
							t.Errorf("Table %s: column %d: expected %s, got %s",
								tableName, i, expectedCols[i], col.Name)
						}
					}
				}
			}
		})
	}
}

func TestParser_ForeignKeys(t *testing.T) {
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
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	// Check posts foreign keys
	posts, ok := p.GetTable("posts")
	if !ok {
		t.Fatal("Posts table not found")
	}

	if len(posts.ForeignKeys) != 1 {
		t.Errorf("Expected 1 foreign key in posts, got %d", len(posts.ForeignKeys))
	}

	if posts.ForeignKeys[0].Column != "user_id" {
		t.Errorf("Expected foreign key column 'user_id', got %s", posts.ForeignKeys[0].Column)
	}

	if posts.ForeignKeys[0].ReferencedTableName != "users" {
		t.Errorf("Expected referenced table 'users', got %s", posts.ForeignKeys[0].ReferencedTableName)
	}

	// Check comments foreign keys
	comments, ok := p.GetTable("comments")
	if !ok {
		t.Fatal("Comments table not found")
	}

	if len(comments.ForeignKeys) != 2 {
		t.Errorf("Expected 2 foreign keys in comments, got %d", len(comments.ForeignKeys))
	}
}

func TestSQLGen_Integration(t *testing.T) {
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

	gen := New().WithPackageName("hr")

	if err := gen.Parse(schema); err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	tables := gen.GetTables()
	if len(tables) != 2 {
		t.Errorf("Expected 2 tables, got %d", len(tables))
	}

	// Check departments table
	dept, ok := gen.GetTable("departments")
	if !ok {
		t.Fatal("Departments table not found")
	}

	if len(dept.Columns) != 2 {
		t.Errorf("Expected 2 columns in departments, got %d", len(dept.Columns))
	}

	// Check employees table
	emp, ok := gen.GetTable("employees")
	if !ok {
		t.Fatal("Employees table not found")
	}

	if len(emp.Columns) != 8 {
		t.Errorf("Expected 8 columns in employees, got %d", len(emp.Columns))
	}

	if len(emp.ForeignKeys) != 1 {
		t.Errorf("Expected 1 foreign key in employees, got %d", len(emp.ForeignKeys))
	}

	// Generate code
	code, err := gen.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Verify package name
	if !strings.Contains(code, "package hr") {
		t.Error("Generated code does not have correct package name")
	}

	// Verify structs
	if !strings.Contains(code, "type Departments struct") {
		t.Error("Generated code does not contain Departments struct")
	}

	if !strings.Contains(code, "type Employees struct") {
		t.Error("Generated code does not contain Employees struct")
	}

	// Verify query builders
	if !strings.Contains(code, "type DepartmentsQuery struct") {
		t.Error("Generated code does not contain DepartmentsQuery")
	}

	if !strings.Contains(code, "type EmployeesQuery struct") {
		t.Error("Generated code does not contain EmployeesQuery")
	}

	// Verify join functionality
	if !strings.Contains(code, "JoinDepartments") {
		t.Error("Generated code does not contain JoinDepartments method")
	}
}
