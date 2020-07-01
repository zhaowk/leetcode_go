package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))
	assert.Equal(t, 1, missingNumber([]int{2, 0, 3}))
	assert.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
