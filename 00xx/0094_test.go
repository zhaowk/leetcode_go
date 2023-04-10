package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestInorderTraversal(t *testing.T) {
	assert.Equal(t, []int{}, inorderTraversal(nil))
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(BuildTree(1, nil, 2, 3)))
}
