package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{}, generateMatrix(0))
	assert.Equal(t, [][]int{{1}}, generateMatrix(1))
	assert.Equal(t, [][]int{{1, 2}, {4, 3}}, generateMatrix(2))
	assert.Equal(t, [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}, generateMatrix(3))
	assert.Equal(t, [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}, generateMatrix(4))
	assert.Equal(t, [][]int{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}, generateMatrix(5))
}
