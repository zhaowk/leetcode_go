package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	assert.Equal(t, 25, sumNumbers(buildTree(1, 2, 3)))
	assert.Equal(t, 1026, sumNumbers(buildTree(4, 9, 0, 5, 1)))
	assert.Equal(t, 896, sumNumbers(buildTree(4, 9, 0, 5, nil, 1)))
}
