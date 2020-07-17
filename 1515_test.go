package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinDistSum(t *testing.T) {

	assert.True(t, floatEqN5(4.0, getMinDistSum([][]int{
		{0, 1}, {1, 0}, {1, 2}, {2, 1},
	})))

	assert.True(t, floatEqN5(2.8284271, getMinDistSum([][]int{
		{1, 1}, {3, 3},
	})))

	assert.True(t, floatEqN5(0, getMinDistSum([][]int{
		{1, 1},
	})))

	assert.True(t, floatEqN5(2.7320508, getMinDistSum([][]int{
		{1, 1}, {0, 0}, {2, 0},
	})))

	assert.True(t, floatEqN5(32.940362, getMinDistSum([][]int{
		{0, 1}, {3, 2}, {4, 5}, {7, 6}, {8, 9}, {11, 1}, {2, 12},
	})))
}
