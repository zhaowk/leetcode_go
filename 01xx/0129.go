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
var sumNumbersRes = 0

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumNumbersRes = 0

	sum(root, 0)
	return sumNumbersRes
}

func sum(root *TreeNode, tmp int) {
	if root.Left == nil && root.Right == nil {
		sumNumbersRes += tmp*10 + root.Val
	}
	if root.Left != nil {
		sum(root.Left, tmp*10+root.Val)
	}
	if root.Right != nil {
		sum(root.Right, tmp*10+root.Val)
	}
}
