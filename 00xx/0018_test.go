package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFourSum(t *testing.T) {
	expected := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	assert.Equal(t, expected, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
