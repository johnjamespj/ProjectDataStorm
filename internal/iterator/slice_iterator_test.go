package iterator_test

import (
	"strings"
	"testing"

	"github.com/johnjamespj/project_datastorm/internal/iterator"
	"github.com/stretchr/testify/assert"
)

func TestSliceIterator(t *testing.T) {
	// create a simple slice of strings
	slice := []string{"a", "b", "c", "d", "e"}

	// create a new slice iterator
	it := iterator.NewSliceIterator(slice)

	// move to the next item for all items
	assert.True(t, it.MoveNext())
	assert.Equal(t, "a", it.Current())
	assert.True(t, it.MoveNext())
	assert.Equal(t, "b", it.Current())
	assert.True(t, it.MoveNext())
	assert.Equal(t, "c", it.Current())
	assert.True(t, it.MoveNext())
	assert.Equal(t, "d", it.Current())
	assert.True(t, it.MoveNext())
	assert.Equal(t, "e", it.Current())
	assert.False(t, it.MoveNext())

	iterator.Contains(it, strings.Compare, "c")
}
