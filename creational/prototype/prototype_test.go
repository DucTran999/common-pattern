package prototype_test

import (
	"patterns/creational/prototype"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DocumentClone(t *testing.T) {
	t.Parallel()

	doc := &prototype.Document{
		Title: "Original Title",
		Body:  "This is the body of the original document.",
	}

	// Clone the document
	clone := doc.Clone()

	// Verify that the clone has the same content
	require.Equal(t, doc.Title, clone.Title, "Clone title should match original")

	// Verify that the clone has the same body
	require.Equal(t, doc.Body, clone.Body, "Clone body should match original")

	// Verify that the clone is a different instance
	assert.NotSame(t, doc, clone, "Clone should not be the same instance as the original")
}
