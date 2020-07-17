package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	a := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	assert.Equal(t, 7, calculateMinimumHP(a))
}
