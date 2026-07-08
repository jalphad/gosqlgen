package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateExamplesRejectsDuplicateIDs(t *testing.T) {
	// Arrange
	examples := []example{
		{
			ID:          "duplicate",
			Kind:        "query",
			Source:      "docs/examples/getting_started_test.go",
			ExpectedSQL: "SELECT 1",
		},
		{
			ID:          "duplicate",
			Kind:        "query",
			Source:      "docs/examples/getting_started_test.go",
			ExpectedSQL: "SELECT 1",
		},
	}

	// Act
	err := validateExampleMetadata("../..", examples)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), `duplicate docs example ID "duplicate"`)
}
