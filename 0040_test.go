package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.Equal(t, [][]int{
		{1, 2, 2},
		{5},
	}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))

}
