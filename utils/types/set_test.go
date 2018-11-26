package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFlow(t *testing.T) {
	set := Set()

	set.Add("value")

	assert.Equal(t, 1, set.Length())
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
}

func TestSetConstructor(t *testing.T) {
	set := Set("first", "second")

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

func TestSetKeys(t *testing.T) {
	expected := []string{"first", "second", "third"}
	set := Set()
	for _, key := range expected {
		set.Add(key)
	}
	keys := set.Slice()

	assert.ElementsMatch(t, expected, keys)
}
