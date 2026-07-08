package docs

#Example: {
	id: =~"^[a-z][a-z0-9-]*$"
	title: string & != ""
	source: =~"^docs/examples/.+\\.go$"
	expectedSQL: string & != ""
	expectedParams: [..._]
}

examples: [...#Example] & [
	{
		id:    "getting-started-select-active-users"
		title: "Select active users"
		source: "docs/examples/getting_started_test.go"
		expectedSQL: "SELECT users.id, users.email FROM users WHERE users.email LIKE $1 AND users.is_active = $2 ORDER BY users.email ASC LIMIT 10"
		expectedParams: ["%@corp.com", true]
	},
]
