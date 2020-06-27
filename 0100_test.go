package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	assert.True(t, isSameTree(buildTree(1, 2, 3), buildTree(1, 2, 3)))
	assert.False(t, isSameTree(buildTree(1, 2, 3), buildTree(1, 2, nil, 3)))
}
