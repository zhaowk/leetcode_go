package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeSum(t *testing.T) {
	assert.Equal(t, 13, rangeSum([]int{
		1, 2, 3, 4,
	}, 4, 1, 5))
}
