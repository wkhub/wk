package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSetFlow(t *testing.T) {
	set := NewStringSet()
	assert.True(t, set.IsEmpty())

	set.Add("value")

	assert.Equal(t, 1, set.Length())
	assert.False(t, set.IsEmpty())
	assert.True(t, set.Has("value"))
	assert.False(t, set.Has("missing"))

	set.Add("value")
	assert.Equal(t, 1, set.Length())
	assert.True(t, set.Has("value"))

	set.Remove("value")
	assert.Equal(t, 0, set.Length())
	assert.False(t, set.Has("value"))

	set.Remove("value")
	assert.Equal(t, 0, set.Length())
	assert.False(t, set.Has("value"))
	assert.True(t, set.IsEmpty())
}

func TestStringSetConstructor(t *testing.T) {
	set := NewStringSet("first", "second")

	assert.Equal(t, 2, set.Length())
	assert.True(t, set.Has("first"))
	assert.True(t, set.Has("second"))
	assert.False(t, set.Has("value"))

	set.Add("value")
	assert.Equal(t, 3, set.Length())
	assert.True(t, set.Has("first"))
	assert.True(t, set.Has("second"))
	assert.True(t, set.Has("value"))
}

func TestStringSetConstructorWithDuplicates(t *testing.T) {
	set := NewStringSet("first", "second", "second")

	assert.Equal(t, 2, set.Length())
	assert.True(t, set.Has("first"))
	assert.True(t, set.Has("second"))
}

func TestStringSetKeys(t *testing.T) {
	expected := []string{"first", "second", "third"}
	set := NewStringSet()
	for _, key := range expected {
		set.Add(key)
	}
	keys := set.Slice()

	assert.ElementsMatch(t, expected, keys)
}
