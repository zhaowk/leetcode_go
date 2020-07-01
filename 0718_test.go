package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLength(t *testing.T) {
	assert.Equal(t, 3, findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	assert.Equal(t, 4, findLength([]int{1, 2, 3, 2, 1}, []int{2, 3, 2, 1, 4, 7}))

	assert.Equal(t, 4, findLength([]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5}))
	assert.Equal(t, 0, findLength([]int{1, 2, 3, 2, 1}, []int{}))
	assert.Equal(t, 0, findLength([]int{}, []int{2, 3, 2, 1, 4, 7}))
	assert.Equal(t, 2, findLength([]int{1, 2, 3}, []int{2, 3, 4}))
	assert.Equal(t, 3, findLength([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.Equal(t, 3, findLength([]int{2, 3, 4}, []int{1, 2, 3, 4}))
	assert.Equal(t, 9, findLength([]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}))
}
