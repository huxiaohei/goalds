package deque

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
	for i := 0; i < 10000; i++ {
		assert.Equal(t, i, a.At(i))
	}
	for i := 9999; i >= 0; i-- {
		assert.Equal(t, i, a.Back())
		assert.Equal(t, i, a.PopBack())
	}

	for i := 9999; i >= 0; i-- {
		a.PushFront(i)
	}
	for i := 0; i < 10000; i++ {
		assert.Equal(t, i, a.At(i))
	}
	for i := 0; i >= 10000; i++ {
		assert.Equal(t, i, a.Front())
		assert.Equal(t, i, a.PopFront())
	}

}

func TestAt(t *testing.T) {
	a := New[int]()
	for i := 0; i < 2000; i++ {
		a.PushBack(i)
	}
	for i := 0; i < 2000; i++ {
		assert.Equal(t, i, a.At(i))
	}
	for i := 1999; i >= 0; i-- {
		a.PushFront(i)
	}
	for i := 0; i < 4000; i++ {
		assert.Equal(t, i%2000, a.At(i))
	}
}

func TestSet(t *testing.T) {
	a := New[int]()
	b := New[int]()
	for i := 0; i < 2000; i++ {
		a.PushBack(i)
		b.PushBack(0)
	}
	for i := 0; i < 2000; i++ {
		b.Set(i, i)
		assert.Equal(t, a.At(i), b.At(i))
	}
	for i := 0; i < 2000; i++ {
		a.Set(i, i+2000)
		assert.Equal(t, i+2000, a.At(i))

	}
}

func TestPop(t *testing.T) {
	a := New[int]()
	b := New[int]()
	for i := 0; i < 10000; i++ {
		a.PushBack(i)
		b.PushFront(i)
	}
	assert.Equal(t, 10000, a.Size())
	assert.Equal(t, 10000, b.Size())
	for i := 0; i < 10000; i++ {
		assert.Equal(t, i, a.PopFront())
		assert.Equal(t, i, b.PopBack())
	}
	assert.Equal(t, 0, a.Size())
	assert.Equal(t, 0, b.Size())
}

func TestInsert(t *testing.T) {
	a := New[int]()
	for i := 999; i >= 0; i-- {
		a.InsertAt(0, i)
	}
	assert.Equal(t, 1000, a.Size())

	for i := 0; i < 1000; i++ {
		a.InsertAt(a.Size(), i)
	}
	assert.Equal(t, 2000, a.Size())

	for i := 0; i < 1000; i++ {
		a.InsertAt(1000+i, i)
	}
	assert.Equal(t, 3000, a.Size())

	for i := 0; i < 1000; i++ {
		a.InsertAt(1000+i, i)
	}
	assert.Equal(t, 4000, a.Size())

	for i := 0; i < 4000; i++ {
		assert.Equal(t, i%1000, a.At(i))
	}

	a.Clear()

	for i := 0; i < 10000; i++ {
		a.PushBack(i)
	}
	for i := 0; i < 10000; i++ {
		a.InsertAt(128, 1)
	}

	for i := 0; i < 128; i++ {
		assert.Equal(t, i, a.At(i))
	}

	a.Clear()

	for i := 9999; i >= 0; i-- {
		a.PushFront(i)
	}
	insetIndex := 10000%128 + 1
	for i := 0; i < 10000; i++ {
		a.InsertAt(insetIndex, 1)
	}

	assert.Equal(t, 20000, a.Size())

	for i := 0; i < insetIndex; i++ {
		assert.Equal(t, i, a.At(i))
	}
	for i := insetIndex; i < 10000+insetIndex; i++ {
		assert.Equal(t, 1, a.At(i))
	}
	for i := 10000 + insetIndex; i < 20000; i++ {
		assert.Equal(t, i-10000, a.At(i))
	}
}

func TestErase(t *testing.T) {
	q := New[int]()
	assert.True(t, q.Empty())
	for i := 0; i < 5; i++ {
		q.PushBack(i + 1)
	}
	assert.False(t, q.Empty())

	q.EraseAt(1)
	assert.Equal(t, "[1 3 4 5]", q.String())

	q.EraseAt(0)
	assert.Equal(t, "[3 4 5]", q.String())

	q.PushFront(6)
	q.PushBack(7)
	q.PushFront(8)
	assert.Equal(t, "[8 6 3 4 5 7]", q.String())

	q.EraseRange(3, 5)
	assert.Equal(t, "[8 6 3 7]", q.String())

	q.Clear()
	assert.True(t, q.Empty())

	for i := 0; i < 10000; i++ {
		q.PushBack(i)
		assert.Equal(t, i, q.At(i))
	}

	for i := 8000; i < 10000; i++ {
		assert.Equal(t, i, q.EraseAt(8000))
	}

	for i := 1000; i < 3000; i++ {
		assert.Equal(t, i, q.EraseAt(1000))
	}
	assert.Equal(t, 6000, q.Size())

}

func TestEraseRange(t *testing.T) {
	a := New[int]()
	for i := 0; i < 10000; i++ {
		a.PushBack(i)
	}
	a.EraseRange(0, 2000)
	assert.Equal(t, 8000, a.Size())
	for i := 2000; i < 10000; i++ {
		assert.Equal(t, i, a.At(i-2000))
	}

	a.EraseRange(6000, 8000)
	assert.Equal(t, 6000, a.Size())
	for i := 2000; i < 8000; i++ {
		assert.Equal(t, i, a.At(i-2000))
	}

	a.EraseRange(2000, 4000)
	assert.Equal(t, 4000, a.Size())
	for i := 2000; i < 4000; i++ {
		assert.Equal(t, i, a.At(i-2000))
		assert.Equal(t, i+4000, a.At(i))
	}

	a.Clear()
	assert.Equal(t, 0, a.Size())
	assert.Equal(t, true, a.Empty())

	for i := 0; i < 10000; i++ {
		a.PushBack(i)
	}
	a.EraseRange(5000, 10000)
	assert.Equal(t, 5000, a.Size())
	for i := 0; i < 5000; i++ {
		assert.Equal(t, i, a.At(i))
	}

	a.EraseRange(4000, 4200)
	for i := 0; i < 4000; i++ {
		assert.Equal(t, i, a.At(i))
	}
	for i := 4200; i < 5000; i++ {
		assert.Equal(t, i, a.At(i-200))
	}

	assert.False(t, a.EraseRange(0, 0))
}

func TestString(t *testing.T) {
	a := New[int]()
	b := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a.PushBack(i)
		b[i] = i
	}
	assert.Equal(t, fmt.Sprintf("%v", b), a.String())
}

func TestError(t *testing.T) {
	a := New[int]()
	assert.Panics(t, func() { a.Front() })
	assert.Panics(t, func() { a.Back() })
	assert.Panics(t, func() { a.PopFront() })
	assert.Panics(t, func() { a.PopBack() })
	assert.Panics(t, func() { a.At(-1) })
	assert.Panics(t, func() { a.At(10) })
	assert.Panics(t, func() { a.Set(-1, 10) })
	assert.Panics(t, func() { a.Set(10, 10) })
	assert.Panics(t, func() { a.EraseAt(-1) })
	assert.Panics(t, func() { a.EraseAt(10) })

	// a.PushBack(1)
	// a.PushBack(2)
	// a.PushBack(3)
	// assert.Panics(t, func() { a.At(-1) })
	// assert.Panics(t, func() { a.At(10) })
	// assert.Panics(t, func() { a.Set(-1, 10) })
	// assert.Panics(t, func() { a.Set(10, 10) })
	// assert.Panics(t, func() { a.InsertAt(-1, 10) })
	// assert.Panics(t, func() { a.InsertAt(10, 10) })
	// assert.Panics(t, func() { a.EraseAt(-1) })
	// assert.Panics(t, func() { a.EraseAt(10) })
}
