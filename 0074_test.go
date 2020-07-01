package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	a := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	assert.True(t, searchMatrix(a, 1))
	assert.False(t, searchMatrix(a, 2))
	assert.True(t, searchMatrix(a, 3))
	assert.True(t, searchMatrix(a, 5))
	assert.True(t, searchMatrix(a, 7))

	assert.True(t, searchMatrix(a, 10))
	assert.False(t, searchMatrix(a, 15))
	assert.True(t, searchMatrix(a, 20))
	assert.False(t, searchMatrix(a, 21))
	assert.True(t, searchMatrix(a, 23))
	assert.True(t, searchMatrix(a, 30))
	assert.False(t, searchMatrix(a, 31))
	assert.True(t, searchMatrix(a, 50))
	assert.False(t, searchMatrix(a, 55))
	assert.False(t, searchMatrix(a, 100))

	assert.True(t, searchMatrix([][]int{{1}}, 1))
	assert.False(t, searchMatrix([][]int{{1}}, 2))
}
