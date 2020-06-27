package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, 7, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
