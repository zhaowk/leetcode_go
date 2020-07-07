package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	assert.Equal(t, 3, longestSubarray([]int{1, 1, 0, 1}))
	assert.Equal(t, 2, longestSubarray([]int{1, 1, 0, 0, 1}))
	assert.Equal(t, 5, longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	assert.Equal(t, 2, longestSubarray([]int{1, 1, 1}))
	assert.Equal(t, 4, longestSubarray([]int{1, 1, 0, 0, 1, 1, 1, 0, 1}))
	assert.Equal(t, 0, longestSubarray([]int{0, 0, 0}))

	assert.Equal(t, 11, longestSubarray([]int{1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}))

	a := []int{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	assert.Equal(t, 18, longestSubarray(a))

	a = []int{1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	assert.Equal(t, 15, longestSubarray(a))

	a = []int{1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//a = []int{1,0,1,1,1,1,1,1,1,1,1,1}
	assert.Equal(t, 11, longestSubarray(a))

}
