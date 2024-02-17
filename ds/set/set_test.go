package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDS(t *testing.T) {
	s := New[string](WithGoroutineSafe())

	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Size())

	s.Insert("Hello")
	assert.False(t, s.Empty())
	assert.Equal(t, 1, s.Size())

	assert.True(t, s.Has("Hello"))

	s.Erase("Hello")
	assert.True(t, s.Empty())

	s.Insert("Hello")
	s.Insert("World")
	s.Insert("Go")
	s.Insert("Rust")
	s.Insert("Python")
	s.Insert("Java")
	s.Insert("C++")
	s.Insert("C#")
	s.Insert("JavaScript")
	s.Insert("TypeScript")
	s.Insert("Swift")
	s.Insert("Kotlin")
	s.Insert("Dart")
	s.Insert("Ruby")

	s.Traversal(func(k string) bool {
		return k == "Swift"
	})

	s.Clear()
	assert.True(t, s.Empty())
}

func TestSetAL(t *testing.T) {
	s1 := New[string](WithGoroutineSafe())
	s2 := New[string](WithGoroutineSafe())

	s1.Insert("C")
	s1.Insert("C++")
	s1.Insert("Rust")
	s1.Insert("Go")

	s2.Insert("JavaScript")
	s2.Insert("TypeScript")
	s2.Insert("Python")
	s2.Insert("Lua")

	s3 := s1.Intersect(s2)
	assert.True(t, s3.Empty())

	s1.Insert("Python")

	s3 = s1.Intersect(s2)
	assert.Equal(t, 1, s3.Size())
	assert.True(t, s3.Has("Python"))

	s3 = s2.Union(s1)
	assert.Equal(t, 8, s3.Size())

	assert.True(t, s3.IsSuperset(s1))
	assert.True(t, s2.IsSubset(s3))
	assert.False(t, s3.IsSubset(s1))

	s3 = s2.Difference(s1)
	assert.Equal(t, 3, s3.Size())
}
