package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmapDS(t *testing.T) {
	a := New(1024)
	assert.Equal(t, a.Size(), uint64(1024))

	for i := 0; i < 1024; i++ {
		if i%2 == 0 {
			a.Set(uint64(i))
		}
	}
	for i := 0; i < int(a.Size()); i++ {
		assert.Equal(t, a.IsSet(uint64(i)), i%2 == 0)
	}

	for i := 0; i < int(a.Size()); i++ {
		if i%2 == 0 {
			a.Unset(uint64(i))
		} else {
			a.Set(uint64(i))
		}
	}
	for i := 0; i < int(a.Size()); i++ {
		assert.Equal(t, a.IsSet(uint64(i)), i%2 == 1)
	}

	a.Clear()

	for i := 0; i < int(a.Size()); i++ {
		assert.False(t, a.IsSet(uint64(i)))
	}

	b := []byte{255, 255, 255}
	c := NewFromBits(b)

	assert.Equal(t, c.Data(), b)

	for i := 0; i < int(c.Size()); i++ {
		assert.True(t, c.IsSet(uint64(i)))
	}

	c.Resize(1025)
	assert.Equal(t, c.Size(), uint64(1032))

}

func TestError(t *testing.T) {
	a := New(1024)
	assert.False(t, a.Set(1024))
	assert.False(t, a.Unset(1024))
	assert.False(t, a.IsSet(1024))
	assert.False(t, a.Resize(1024))
}
