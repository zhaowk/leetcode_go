package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreePI(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[0]
	pos := -1

	for idx := range inorder {
		if inorder[idx] == preorder[0] {
			pos = idx
		}
	}

	if pos > 0 {
		root.Left = buildTreePI(preorder[1:], inorder[0:pos])
	}

	if pos < len(inorder)-1 {
		root.Right = buildTreePI(preorder[pos+1:], inorder[pos+1:])
	}
	return root
}
