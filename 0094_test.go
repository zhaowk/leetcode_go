package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	assert.Equal(t, []int{}, inorderTraversal(nil))
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(buildTree(1, nil, 2, 3)))
}
