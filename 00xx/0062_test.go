package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 3, uniquePaths(2, 3))
	assert.Equal(t, 6, uniquePaths(3, 3))
	assert.Equal(t, 10, uniquePaths(3, 4))
	assert.Equal(t, 15, uniquePaths(3, 5))
	assert.Equal(t, 21, uniquePaths(3, 6))
	assert.Equal(t, 28, uniquePaths(3, 7))

	assert.Equal(t, 20, uniquePaths(4, 4))
	assert.Equal(t, 35, uniquePaths(4, 5))
	assert.Equal(t, 56, uniquePaths(4, 6))
}
