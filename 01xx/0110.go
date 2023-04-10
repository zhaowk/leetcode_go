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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	} else if MyAbs(height(root.Left)-height(root.Right)) < 2 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func height(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		return MyMax(height(n.Left), height(n.Right)) + 1
	}
}
