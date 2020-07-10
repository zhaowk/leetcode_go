package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	assert.Equal(t, 11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
