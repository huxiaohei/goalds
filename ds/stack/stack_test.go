package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())
	assert.False(t, s.Empty())
	assert.Equal(t, 3, s.Top())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Top())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Top())
	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.Empty())

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	assert.Equal(t, 10, s.Size())
	fmt.Println(s.String())
	s.Clear()
	assert.True(t, s.Empty())
}
