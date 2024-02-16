package gomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := New[string, string](WithGoroutineSafe())

	assert.True(t, m.Empty())
	assert.Equal(t, 0, m.Size())

	m.Set("Hello", "world")
	assert.False(t, m.Empty())
	assert.Equal(t, 1, m.Size())

	assert.True(t, m.Has("Hello"))

	v, ok := m.Get("Hello")
	assert.True(t, ok)
	assert.Equal(t, "world", v)

	m.Erase("Hello")
	assert.True(t, m.Empty())

	m.Set("Hello", "Hello")
	m.Set("World", "World")
	m.Set("Go", "Go")
	m.Set("Rust", "Rust")
	m.Set("Python", "Python")
	m.Set("Java", "Java")
	m.Set("C++", "C++")
	m.Set("C#", "C#")
	m.Set("JavaScript", "JavaScript")
	m.Set("TypeScript", "TypeScript")
	m.Set("Swift", "Swift")
	m.Set("Kotlin", "Kotlin")
	m.Set("Dart", "Dart")
	m.Set("Ruby", "Ruby")

	m.Traversal(func(k string, v string) bool {
		assert.Equal(t, k, v)
		return k == "Swift"
	})

	m.Clear()
	assert.True(t, m.Empty())
}
