package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateExamplesRejectsDuplicateIDs(t *testing.T) {
	// Arrange
	examples := []example{
		{ID: "duplicate", Source: "docs/examples/getting_started_test.go"},
		{ID: "duplicate", Source: "docs/examples/getting_started_test.go"},
	}

	// Act
	err := validateExamples("../..", examples)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), `duplicate docs example ID "duplicate"`)
}
