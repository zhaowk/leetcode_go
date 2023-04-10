package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	assert.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	assert.Equal(t, []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8}, spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}))
	assert.Equal(t, []int{1, 2, 3, 6, 9, 12, 15, 14, 13, 10, 7, 4, 5, 8, 11}, spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
		{13, 14, 15},
	}))
	assert.Equal(t, []int{2, 5, 4, -1, 0, 8}, spiralOrder([][]int{
		{2, 5},
		{8, 4},
		{0, -1},
	}))
	assert.Equal(t, []int{1, 2, 3, 4, 5,
		10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13}, spiralOrder([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}))
}
