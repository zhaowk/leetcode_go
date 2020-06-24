package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{}, 0))
	assert.Equal(t, 0, searchInsert([]int{0, 1}, 0))
	assert.Equal(t, 1, searchInsert([]int{0, 1}, 1))
	assert.Equal(t, 2, searchInsert([]int{0, 1}, 2))

	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
}
