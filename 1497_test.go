package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanArrange(t *testing.T) {
	assert.True(t, canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	assert.True(t, canArrange([]int{-1, 2, 3, 4, 5, 10, 6, 7, 8, 11}, 5))
}
