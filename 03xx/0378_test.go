package _3xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	a := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}

	assert.Equal(t, 1, kthSmallest(a, 1))
	assert.Equal(t, 5, kthSmallest(a, 2))
	assert.Equal(t, 9, kthSmallest(a, 3))

	assert.Equal(t, 10, kthSmallest(a, 4))
	assert.Equal(t, 11, kthSmallest(a, 5))
	assert.Equal(t, 12, kthSmallest(a, 6))
	assert.Equal(t, 13, kthSmallest(a, 7))
	assert.Equal(t, 13, kthSmallest(a, 8))
	assert.Equal(t, 15, kthSmallest(a, 9))

	a = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	assert.Equal(t, 5, kthSmallest(a, 5))

	for i := 1; i < 20; i++ {
		assert.Equal(t, i, kthSmallest(a, i))
	}

	a = [][]int{{0, 0, 0}, {2, 7, 9}, {7, 8, 11}}

	assert.Equal(t, 0, kthSmallest(a, 1))
	assert.Equal(t, 0, kthSmallest(a, 2))
	assert.Equal(t, 0, kthSmallest(a, 3))
	assert.Equal(t, 2, kthSmallest(a, 4))
	assert.Equal(t, 7, kthSmallest(a, 5))
	assert.Equal(t, 7, kthSmallest(a, 6))
	assert.Equal(t, 8, kthSmallest(a, 7))
	assert.Equal(t, 9, kthSmallest(a, 8))
	assert.Equal(t, 11, kthSmallest(a, 9))

}
