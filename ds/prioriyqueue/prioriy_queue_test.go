package prioriyqueue

import (
	"goalds/utils/comparator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrioriyQueue(t *testing.T) {
	q := New[int](comparator.OrderedTypeCmp[int], WithGoroutineSafe())
	assert.True(t, q.Empty())
	q.Push(1)
	q.Push(10)
	q.Push(5)
	q.Push(3)
	q.Push(2)
	q.Push(4)
	q.Push(6)
	q.Push(7)
	q.Push(8)
	q.Push(9)
	assert.False(t, q.Empty())
	assert.Equal(t, 10, q.Len())

	for i := q.Len(); i > 0; i-- {
		assert.Equal(t, i, q.Top())
		assert.Equal(t, i, q.Pop())
	}

	q = New[int](comparator.Reverse[int](comparator.OrderedTypeCmp[int]), WithGoroutineSafe())
	assert.True(t, q.Empty())
	q.Push(1)
	q.Push(10)
	q.Push(5)
	q.Push(3)
	q.Push(2)
	q.Push(4)
	q.Push(6)
	q.Push(7)
	q.Push(8)
	q.Push(9)
	assert.False(t, q.Empty())

	for i := 1; i <= q.Len(); i++ {
		assert.Equal(t, i, q.Top())
		assert.Equal(t, i, q.Pop())
	}

}
