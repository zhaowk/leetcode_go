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
var found bool

func hasPathSum(root *TreeNode, sum int) bool {
	found = false

	if root == nil {
		return false
	}

	hasPathSums(root, 0, sum)
	return found
}

func hasPathSums(n *TreeNode, t, sum int) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			if t+n.Val == sum {
				found = true
				return
			}
		} else {
			hasPathSums(n.Left, t+n.Val, sum)
			hasPathSums(n.Right, t+n.Val, sum)
		}
	}
}
