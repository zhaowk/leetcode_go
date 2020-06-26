package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {

	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	b := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	rotate(a)
	assert.Equal(t, b, a)

	a = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	b = [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	rotate(a)
	assert.Equal(t, b, a)
}
