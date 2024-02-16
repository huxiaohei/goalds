package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorNew(t *testing.T) {
	v := New[int](WithCapacity(100), WithGoroutineSafe())
	assert.True(t, v.Empty())
	assert.Equal(t, 100, v.Capacity())

	for i := 0; i < 100; i++ {
		v.PushBack(i)
	}
	assert.Equal(t, 100, v.Size())

	vv := NewFromVector[int](v, WithGoroutineSafe())
	assert.Equal(t, 100, vv.Size())
	assert.Equal(t, 100, vv.Capacity())

	for i := 0; i < 100; i++ {
		assert.Equal(t, v.At(i), vv.At(i))
	}
}

func TestVectorShrink(t *testing.T) {
	v := New[int](WithCapacity(100))
	assert.Equal(t, 100, v.Capacity())
	for i := 0; i < 100; i++ {
		v.PushBack(i)
	}
	for i := 0; i < 50; i++ {
		v.PopBack()
	}
	v.ShrinkToFit()
	assert.Equal(t, 50, v.Capacity())

	v.Clear()
	assert.Equal(t, 50, v.Capacity())
	assert.True(t, v.Empty())
}

func TestVectorPushBack(t *testing.T) {
	v := New[int](WithCapacity(100), WithGoroutineSafe())
	for i := 0; i < 100; i++ {
		v.PushBack(i)
	}
	assert.Equal(t, 100, v.Size())

	for i := 0; i < 100; i++ {
		assert.Equal(t, v.Back(), v.PopBack())
	}
}
func TestVectorInsert(t *testing.T) {
	v := New[int](WithCapacity(100), WithGoroutineSafe())
	for i := 0; i < 100; i++ {
		v.InsertAt(0, i)
	}
	assert.Equal(t, 100, v.Size())
	for i := 0; i < 100; i++ {
		assert.Equal(t, v.Front(), 99-i)
		v.EraseAt(0)
	}
	assert.True(t, v.Empty())
}

func TestVectorString(t *testing.T) {
	v := New[int](WithCapacity(100), WithGoroutineSafe())
	for i := 0; i < 10; i++ {
		v.InsertAt(0, i)
	}
	assert.Equal(t, "vector: [9 8 7 6 5 4 3 2 1 0]", v.String())
}

func TestTraversal(t *testing.T) {
	v := New[int](WithCapacity(100), WithGoroutineSafe())
	for i := 0; i < 100; i++ {
		v.PushBack(i)
	}
	v.Traversal(func(i int, v int) bool {
		assert.Equal(t, i, v)
		return i != 50
	})

}

func TestPanic(t *testing.T) {
	v := New[int](WithCapacity(10))
	v.ShrinkToFit()
	assert.Empty(t, 0, v.Capacity())
	v.ShrinkToFit()
	assert.PanicsWithValue(t, "vector: pop back from empty array", func() { v.PopBack() })
	assert.PanicsWithValue(t, "vector: out of range index: 10 size: 0", func() { v.At(10) })
	assert.PanicsWithValue(t, "vector: out of range index: 10 size: 0", func() { v.InsertAt(10, 0) })
	assert.PanicsWithValue(t, "vector: out of range index: 1 4 size: 0", func() { v.EraseRange(1, 4) })
}
