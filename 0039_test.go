package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	assert.Equal(t, [][]int{{7, 7, 3}}, combinationSum([]int{7, 3}, 17))
	assert.Equal(t, [][]int{
		{4, 4},
		{4, 2, 2},
		{2, 2, 2, 2},
		{8},
	}, combinationSum([]int{4, 2, 8}, 8))
}
