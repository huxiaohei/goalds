package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New[int]()
	assert.True(t, a.Empty())
	assert.True(t, a.Size() == 0)
}

func TestPushAndPop(t *testing.T) {
	a := New[int]()
	for i := 0; i < 10000; i++ {
		a.PushBack(i)
	}
	assert.Equal(t, 10000, a.Size())
	assert.Equal(t, 0, a.Front().Val)
	assert.Equal(t, 9999, a.Back().Val)
	for i := 0; i < 10000; i++ {
		assert.Equal(t, i, a.PopFront().Val)
	}

	for i := 0; i < 10000; i++ {
		a.PushFront(i)
	}
	for i := 0; i < 10000; i++ {
		assert.Equal(t, i, a.PopBack().Val)
	}
}

func TestInsertAt(t *testing.T) {
	a := New[int]()
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	// [0 1 2 4 4 4 4 4 3 4 5 6 7 8 9]
	for i := 0; i < 5; i++ {
		a.InsertAt(3, 4)
	}
	fmt.Println(a.String())
	assert.Equal(t, 15, a.Size())

	// [0 1 2 4 4 4 4 4 3 4 4 4 4 4 4 5 6 7 8 9]
	for i := 0; i < 5; i++ {
		a.InsertAt(10, 4)
	}
	fmt.Println(a.String())
	assert.Equal(t, 20, a.Size())

	a.InsertAt(0, 100)
	assert.Equal(t, 100, a.Front().Val)
	a.InsertAt(a.Size(), 200)
	assert.Equal(t, 200, a.Back().Val)
}

func TestRemoveAt(t *testing.T) {
	a := New[int]()
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	assert.Equal(t, 10, a.Size())
	assert.Equal(t, 2, a.RemoveAt(2).Val)
	assert.Equal(t, 8, a.RemoveAt(7).Val)
	assert.Equal(t, 0, a.RemoveAt(0).Val)
	assert.Equal(t, 9, a.RemoveAt(a.Size()-1).Val)
}

func TestRemoveRange(t *testing.T) {
	a := New[int]()
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	node := a.RemoveRange(2, 5)
	assert.Equal(t, 2, node.Val)
	node = node.Next()
	assert.Equal(t, 3, node.Val)
	node = node.Next()
	assert.Equal(t, 4, node.Val)
	assert.Nil(t, node.Next())

	// [0 1 5 6 7 8 9]
	fmt.Println(a.String())
	node = a.RemoveRange(4, a.Size())
	assert.Equal(t, 7, node.Val)
	node = node.Next()
	assert.Equal(t, 8, node.Val)
	node = node.Next()
	assert.Equal(t, 9, node.Val)
	assert.Nil(t, node.Next())

	// [0 1 5 6]
	fmt.Println(a.String())

	node = a.RemoveRange(0, 2)
	assert.Nil(t, node.Prev())
	assert.Equal(t, 0, node.Val)
	node = node.Next()
	assert.Equal(t, 1, node.Val)
	assert.Nil(t, node.Next())

	// [5 6]
	fmt.Println(a.String())

	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	a.RemoveRange(0, a.Size())
	assert.True(t, a.Empty())

	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	a.Clear()
	assert.True(t, a.Empty())
}

func TestError(t *testing.T) {
	a := New[int]()
	assert.Panics(t, func() { a.PopFront() })
	assert.Panics(t, func() { a.PopBack() })
	assert.Panics(t, func() { a.At(10) })
	assert.Panics(t, func() { a.RemoveAt(10) })
	assert.Panics(t, func() { a.RemoveRange(10, 2) })
}
