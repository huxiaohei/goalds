package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	assert.True(t, q.Empty())
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	assert.Equal(t, 10, q.Size())
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, q.Front())
		assert.Equal(t, i, q.Pop())
	}
	assert.True(t, q.Empty())
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	fmt.Println(q.String())
	q.Clear()
	assert.True(t, q.Empty())
}
