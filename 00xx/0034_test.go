package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {
	assert.Equal(t, []int{5, 5}, searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 9))
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	assert.Equal(t, []int{1, 2}, searchRange([]int{5, 7, 7, 8, 8, 10}, 7))
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	assert.Equal(t, []int{0, 0}, searchRange([]int{5, 7, 7, 8, 8, 10}, 5))
	assert.Equal(t, []int{0, 5}, searchRange([]int{5, 5, 5, 5, 5, 5}, 5))
}
