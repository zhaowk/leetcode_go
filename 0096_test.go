package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTrees(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 14, numTrees(4))
	assert.Equal(t, 42, numTrees(5))
	assert.Equal(t, 132, numTrees(6))
	assert.Equal(t, 16796, numTrees(10))
}
