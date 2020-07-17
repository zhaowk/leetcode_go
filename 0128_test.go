package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
