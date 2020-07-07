package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthFactor(t *testing.T) {
	assert.Equal(t, 3, kthFactor(12, 3))
	assert.Equal(t, 7, kthFactor(7, 2))
	assert.Equal(t, -1, kthFactor(4, 4))
	assert.Equal(t, 1, kthFactor(1, 1))
	assert.Equal(t, 4, kthFactor(1000, 3))

}
