package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDS(t *testing.T) {
	a := New[int](5)
	b := NewFrom[int](1, 2, 3, 4, 5)
	c := Clone[int](b)

	assert.False(t, false, a.Empty())
	assert.Equal(t, b.At(2), c.At(2))

	for i := 0; i < b.Size(); i += 1 {
		a.Set(i, i+1)
	}
	assert.Equal(t, a.Back(), b.Back())
	assert.Equal(t, a.Front(), b.Front())

	d := New[int](5)
	a.Swap(d)
	for i := 0; i < a.Size(); i += 1 {
		assert.Equal(t, a.At(i), 0)
		assert.Equal(t, d.At(i), i+1)
	}

	assert.Equal(t, d.String(), "[1 2 3 4 5]")
}

func TestTraversal(t *testing.T) {
	a := New[int](5)
	for i := 0; i < a.Size(); i += 1 {
		a.Set(i, i+1)
	}
	a.Traversal(func(i int, v int) bool {
		assert.Equal(t, v, i+1)
		return v < 3
	})
}

func TestError(t *testing.T) {
	a := NewFrom[int](1, 2, 3, 4, 5)
	assert.PanicsWithValue(t, "array: out of range index: 10 size: 5", func() { a.At(10) })
	assert.PanicsWithValue(t, "array: out of range index: 10 size: 5", func() { a.Set(10, 3) })

	b := NewFrom[int](1, 2, 3)
	assert.PanicsWithValue(t, "array: two arrays have different lengths len: 5 len: 3", func() { a.Swap(b) })
}
