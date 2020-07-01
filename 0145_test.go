package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	assert.Equal(t, []int{2, 3, 1}, postorderTraversal(buildTree(1, 2, 3)))
	assert.Equal(t, []int{4, 5, 2, 3, 1}, postorderTraversal(buildTree(1, 2, 3, 4, 5)))
	assert.Equal(t, []int{4, 2, 5, 3, 1}, postorderTraversal(buildTree(1, 2, 3, nil, 4, 5)))
	assert.Equal(t, []int{4, 2, 5, 3, 1}, postorderTraversal(buildTree(1, 2, 3, 4, nil, 5)))
	assert.Equal(t, []int{2, 4, 3, 1}, postorderTraversal(buildTree(1, 2, 3, nil, nil, 4)))
	assert.Equal(t, []int{5, 3, 4, 2, 1}, postorderTraversal(buildTree(1, 2, nil, 3, 4, 5)))

	assert.Equal(t, []int{2, 3, 1}, postorderTraversal1(buildTree(1, 2, 3)))
	assert.Equal(t, []int{4, 5, 2, 3, 1}, postorderTraversal1(buildTree(1, 2, 3, 4, 5)))
	assert.Equal(t, []int{4, 2, 5, 3, 1}, postorderTraversal1(buildTree(1, 2, 3, nil, 4, 5)))
	assert.Equal(t, []int{4, 2, 5, 3, 1}, postorderTraversal1(buildTree(1, 2, 3, 4, nil, 5)))
	assert.Equal(t, []int{2, 4, 3, 1}, postorderTraversal1(buildTree(1, 2, 3, nil, nil, 4)))
	assert.Equal(t, []int{5, 3, 4, 2, 1}, postorderTraversal1(buildTree(1, 2, nil, 3, 4, 5)))
}
