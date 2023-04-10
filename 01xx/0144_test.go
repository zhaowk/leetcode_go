package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestPreorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(BuildTree(1, 2, 3)))
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal1(BuildTree(1, 2, 3)))
	assert.Equal(t, []int{1, 2, 4, 5, 3}, preorderTraversal1(BuildTree(1, 2, 3, 4, 5)))
	assert.Equal(t, []int{1, 2, 4, 3, 5}, preorderTraversal1(BuildTree(1, 2, 3, nil, 4, 5)))
	assert.Equal(t, []int{1, 2, 4, 3, 5}, preorderTraversal1(BuildTree(1, 2, 3, 4, nil, 5)))
	assert.Equal(t, []int{1, 2, 3, 4}, preorderTraversal1(BuildTree(1, 2, 3, nil, nil, 4)))
	assert.Equal(t, []int{1, 2, 3, 5, 4}, preorderTraversal1(BuildTree(1, 2, nil, 3, 4, 5)))
}
