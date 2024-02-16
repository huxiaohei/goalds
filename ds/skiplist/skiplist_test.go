package skiplist

import (
	"goalds/utils/comparator"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	sl := New[int, int](comparator.OrderedTypeCmp[int], WithMaxLevel(8), WithGoroutineSafe())
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := rand.Int() % 100
		sl.Insert(key, i)
		m[key] = i
	}
	for key, v := range m {
		ret, _ := sl.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), sl.Len())

	keys := sl.Keys()
	pre := keys[0]
	for i := 1; i < len(keys); i++ {
		assert.True(t, pre <= keys[i])
		pre = keys[i]
	}

	_, err := sl.Get(10000)
	assert.NotNil(t, err)
}

func TestRemove(t *testing.T) {
	sl := New[int, int](comparator.OrderedTypeCmp[int], WithGoroutineSafe())
	assert.False(t, sl.Remove(10000))

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		key := rand.Int() % 1000
		sl.Insert(key, i)
		m[key] = i
	}
	assert.Equal(t, len(m), sl.Len())

	for i := 0; i < 300; i++ {
		key := rand.Int() % 1000
		sl.Remove(key)
		delete(m, key)
		key2 := rand.Int() % 10440
		sl.Insert(key2, key)
		m[key2] = key
	}

	for key, v := range m {
		ret, _ := sl.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), sl.Len())
}

func TestTraversal(t *testing.T) {
	sl := New[int, int](comparator.OrderedTypeCmp[int])
	for i := 0; i < 10; i++ {
		sl.Insert(i, i*10)
	}
	keys := sl.Keys()
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, keys[i])
	}
	i := 0
	sl.Traversal(func(key, value int) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i*10, value)
		i++
		return true
	})
	keys = make([]int, 0)
	sl.Traversal(func(key, value int) bool {
		keys = append(keys, key)
		return len(keys) < 5
	})
	assert.Equal(t, 5, len(keys))
}
