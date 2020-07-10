package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	assert.True(t, isSymmetric(buildTree()))
	assert.True(t, isSymmetric(buildTree(1, 2, 2, 3, 4, 4, 3)))
	assert.False(t, isSymmetric(buildTree(1, 2, 2, nil, 3, nil, 3)))
}
