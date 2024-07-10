package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomfilter(t *testing.T) {
	a := NewWithEstimates(1024, 0.01, WithGoroutineSafe())
	assert.False(t, a.Test("hello"))
	a.Add("hello")
	assert.True(t, a.Test("hello"))

	b := NewFromData(a.Data(), WithGoroutineSafe())
	assert.True(t, b.Test("hello"))
}
