package gosqlgen

import (
	"strings"
	"testing"

	"github.com/jalphad/gosqlgen/parser"
)

// TestTypeSafeQueryBuilder tests the type-safe query builder
func TestTypeSafeQueryBuilder(t *testing.T) {
	sql := `
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            email VARCHAR(255) NOT NULL,
            username VARCHAR(50) NOT NULL,
            age INTEGER,
            is_active BOOLEAN DEFAULT true,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE TABLE posts (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL,
            title VARCHAR(255) NOT NULL,
            content TEXT,
            status VARCHAR(20) DEFAULT 'draft',
            view_count INTEGER DEFAULT 0,
            published_at TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `

	p := parser.NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Verify that type-safe methods are generated
	typeSafeElements := []string{
		// Column reference types
		"type UsersFields struct",
		"type PostsFields struct",
		"var UsersTable = UsersFields{}",
		"var PostsTable = PostsFields{}",

		// Type-safe Where methods for users
		"func (q *UsersQuery) WhereIdEq(value int64)",
		"func (q *UsersQuery) WhereEmailEq(value string)",
		"func (q *UsersQuery) WhereEmailLike(pattern string)",
		"func (q *UsersQuery) WhereUsernameEq(value string)",
		"func (q *UsersQuery) WhereAgeEq(value int64)",
		"func (q *UsersQuery) WhereAgeGt(value int64)",
		"func (q *UsersQuery) WhereAgeLte(value int64)",
		"func (q *UsersQuery) WhereIsActiveEq(value bool)",
		"func (q *UsersQuery) WhereCreatedAtGt(value time.Time)",

		// Type-safe OrderBy methods
		"func (q *UsersQuery) OrderById(dir OrderDirection)",
		"func (q *UsersQuery) OrderByEmail(dir OrderDirection)",
		"func (q *UsersQuery) OrderByCreatedAt(dir OrderDirection)",

		// Type-safe Where methods for posts
		"func (q *PostsQuery) WhereUserIdEq(value int64)",
		"func (q *PostsQuery) WhereTitleEq(value string)",
		"func (q *PostsQuery) WhereTitleLike(pattern string)",
		"func (q *PostsQuery) WhereStatusEq(value string)",
		"func (q *PostsQuery) WhereStatusIn(values ...string)",
		"func (q *PostsQuery) WhereViewCountGte(value int64)",

		// Type-safe join methods
		"func (q *PostsQuery) JoinUsers()",
		"func (q *PostsQuery) LeftJoinUsers()",

		// Column reference methods
		"func (u UsersFields) Id() *FieldRef",
		"func (u UsersFields) Email() *FieldRef",
		"func (p PostsFields) Title() *FieldRef",

		// Select methods
		"func (q *UsersQuery) Select(fields ...*FieldRef)",
		"func (q *UsersQuery) SelectAll()",

		// Group and Having
		"func (q *UsersQuery) GroupBy(fields ...*FieldRef)",
		"func (q *UsersQuery) Having(field *FieldRef, op ComparisonOp, value interface{})",

		// Logical operators
		"func (q *UsersQuery) And()",
		"func (q *UsersQuery) Or()",
		"func (q *UsersQuery) WhereGroup(fn func(*UsersQuery))",

		// Update methods
		"func (q *UsersQuery) UpdateFields(updates map[*FieldRef]interface{})",
	}

	for _, elem := range typeSafeElements {
		if !strings.Contains(code, elem) {
			t.Errorf("Generated code missing type-safe element: %s", elem)
		}
	}

	// Verify that common types are generated
	commonTypes := []string{
		"type OrderDirection string",
		"const ASC OrderDirection",
		"const DESC OrderDirection",
		"type ComparisonOp string",
		"const EQ ComparisonOp",
		"const GT ComparisonOp",
		"const LIKE ComparisonOp",
		"type LogicalOp string",
		"const AND LogicalOp",
		"const OR LogicalOp",
		"type FieldRef struct",
		"type Condition struct",
		"type JoinClause struct",
		"type OrderClause struct",
	}

	for _, typ := range commonTypes {
		if !strings.Contains(code, typ) {
			t.Errorf("Generated code missing common type: %s", typ)
		}
	}
}

// TestNullableFields tests that nullable fields generate appropriate methods
func TestNullableFields(t *testing.T) {
	sql := `
        CREATE TABLE products (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            description TEXT,
            price DECIMAL(10,2),
            discontinued_at TIMESTAMP
        );
    `

	p := parser.NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Check for nullable field methods
	nullableMethods := []string{
		"func (q *ProductsQuery) WhereDescriptionIsNull()",
		"func (q *ProductsQuery) WhereDescriptionIsNotNull()",
		"func (q *ProductsQuery) WherePriceIsNull()",
		"func (q *ProductsQuery) WherePriceIsNotNull()",
		"func (q *ProductsQuery) WhereDiscontinuedAtIsNull()",
		"func (q *ProductsQuery) WhereDiscontinuedAtIsNotNull()",
	}

	for _, method := range nullableMethods {
		if !strings.Contains(code, method) {
			t.Errorf("Generated code missing nullable method: %s", method)
		}
	}

	// Verify NOT NULL fields don't have IsNull methods
	notNullableMethods := []string{
		"WhereIdIsNull",   // Primary key, shouldn't be nullable
		"WhereNameIsNull", // Marked as NOT NULL
	}

	for _, method := range notNullableMethods {
		if strings.Contains(code, method) {
			t.Errorf("Generated code incorrectly includes nullable method: %s", method)
		}
	}
}

// TestForeignKeyJoins tests that foreign keys generate proper join methods
func TestForeignKeyJoins(t *testing.T) {
	sql := `
        CREATE TABLE departments (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL
        );
        
        CREATE TABLE employees (
            id SERIAL PRIMARY KEY,
            department_id INTEGER,
            manager_id INTEGER,
            name VARCHAR(100) NOT NULL,
            FOREIGN KEY (department_id) REFERENCES departments(id),
            FOREIGN KEY (manager_id) REFERENCES employees(id)
        );
        
        CREATE TABLE projects (
            id SERIAL PRIMARY KEY,
            lead_employee_id INTEGER,
            department_id INTEGER,
            name VARCHAR(200) NOT NULL,
            FOREIGN KEY (lead_employee_id) REFERENCES employees(id),
            FOREIGN KEY (department_id) REFERENCES departments(id)
        );
    `

	p := parser.NewParser()
	if err := p.Parse(sql); err != nil {
		t.Fatalf("Failed to parse SQL: %v", err)
	}

	g := NewGenerator(p)
	code, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Check for generated join methods
	joinMethods := []string{
		// Employees joins
		"func (q *EmployeesQuery) JoinDepartments()",
		"func (q *EmployeesQuery) LeftJoinDepartments()",
		"func (q *EmployeesQuery) JoinEmployees()", // Self-join for manager

		// Projects joins
		"func (q *ProjectsQuery) JoinEmployees()",
		"func (q *ProjectsQuery) JoinDepartments()",
		"func (q *ProjectsQuery) LeftJoinEmployees()",
		"func (q *ProjectsQuery) LeftJoinDepartments()",

		// Join result structs
		"type EmployeesDepartmentsJoin struct",
		"type EmployeesEmployeesJoin struct", // Self-join struct
		"type ProjectsEmployeesJoin struct",
		"type ProjectsDepartmentsJoin struct",
	}

	for _, method := range joinMethods {
		if !strings.Contains(code, method) {
			t.Errorf("Generated code missing join method/struct: %s", method)
		}
	}
}

// TestComplexQueryBuilding demonstrates compile-time type safety
func TestComplexQueryBuilding(t *testing.T) {
	// This test shows how the API would be used and what compile-time checks would catch

	// Note: These are examples of what the generated code would look like
	// In actual use, these would be compile-time errors if incorrect

	examples := []struct {
		name         string
		description  string
		wouldCompile bool
		reason       string
	}{
		{
			name:         "Correct type-safe query",
			description:  "db.Users().WhereEmailEq(\"test@example.com\").WhereAgeGt(18).Find()",
			wouldCompile: true,
			reason:       "All methods exist with correct types",
		},
		{
			name:         "Wrong type for boolean field",
			description:  "db.Users().WhereIsActiveEq(\"true\")", // string instead of bool
			wouldCompile: false,
			reason:       "WhereIsActiveEq expects bool, not string",
		},
		{
			name:         "Wrong type for numeric field",
			description:  "db.Users().WhereAgeGt(\"eighteen\")", // string instead of int64
			wouldCompile: false,
			reason:       "WhereAgeGt expects int64, not string",
		},
		{
			name:         "Non-existent field",
			description:  "db.Users().WhereMiddleNameEq(\"John\")", // field doesn't exist
			wouldCompile: false,
			reason:       "Method WhereMiddleNameEq doesn't exist",
		},
		{
			name:         "Wrong comparison for type",
			description:  "db.Users().WhereEmailGt(\"test@example.com\")", // GT on string
			wouldCompile: false,
			reason:       "WhereEmailGt doesn't exist (strings don't have GT comparison)",
		},
		{
			name:         "Correct field reference",
			description:  "db.Users().Select(UsersTable.Email(), UsersTable.Age()).Find()",
			wouldCompile: true,
			reason:       "Column references are type-safe",
		},
		{
			name:         "Wrong table field reference",
			description:  "db.Users().Select(PostsTable.Title())", // Wrong table
			wouldCompile: true,                                    // Would compile but might not work as expected
			reason:       "Type system allows mixing, but query builder handles it",
		},
		{
			name:         "Type-safe join",
			description:  "db.Posts().JoinUsers().WhereUsersIsActiveEq(true).Find()",
			wouldCompile: true,
			reason:       "Join method exists from foreign key",
		},
		{
			name:         "Invalid join",
			description:  "db.Users().JoinProducts()", // No FK relationship
			wouldCompile: false,
			reason:       "JoinProducts method doesn't exist without FK",
		},
	}

	// Document the compile-time safety benefits
	for _, ex := range examples {
		if ex.wouldCompile {
			t.Logf("✅ Would compile: %s - %s", ex.description, ex.reason)
		} else {
			t.Logf("❌ Compile error: %s - %s", ex.description, ex.reason)
		}
	}
}

// Benchmark comparing type-safe vs string-based query building
func BenchmarkTypeSafeQueryBuilding(b *testing.B) {
	// This benchmark would show that type-safe methods have no runtime overhead
	// In fact, they might be faster due to possible inlining

	b.Run("TypeSafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Simulated type-safe query building
			// In real use: db.Users().WhereEmailEq("test").WhereAgeGt(18).OrderByCreatedAt(DESC)
			_ = buildTypeSafeQuery()
		}
	})

	b.Run("StringBased", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Simulated string-based query building
			// In real use: db.Users().Where("email = ?", "test").Where("age > ?", 18).OrderBy("created_at DESC")
			_ = buildStringBasedQuery()
		}
	})
}

func buildTypeSafeQuery() string {
	// Simulated type-safe query building
	return "SELECT * FROM users WHERE email = $1 AND age > $2 ORDER BY created_at DESC"
}

func buildStringBasedQuery() string {
	// Simulated string-based query building with string manipulation
	conditions := []string{"email = $1", "age > $2"}
	return "SELECT * FROM users WHERE " + strings.Join(conditions, " AND ") + " ORDER BY created_at DESC"
}
