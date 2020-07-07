package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxValueOfEquation(t *testing.T) {
	assert.Equal(t, 4, findMaxValueOfEquation([][]int{
		{1, 3}, {2, 0}, {5, 10}, {6, -10},
	}, 1))

	assert.Equal(t, 3, findMaxValueOfEquation([][]int{
		{0, 0}, {3, 0}, {9, 2},
	}, 3))

	assert.Equal(t, 35, findMaxValueOfEquation([][]int{
		{-19, -12}, {-13, -18}, {-12, 18}, {-11, -8}, {-8, 2}, {-7, 12}, {-5, 16}, {-3, 9}, {1, -7}, {5, -4}, {6, -20}, {10, 4}, {16, 4}, {19, -9}, {20, 19},
	}, 6))

	assert.Equal(t, 35, findMaxValueOfEquation([][]int{
		{-17, -6}, {-4, 0}, {-2, -16}, {-1, 2}, {0, 11}, {6, 18},
	}, 13))
}
