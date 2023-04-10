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

var mxDepthResult int

func maxDepth(root *TreeNode) int {
	mxDepthResult = 0
	depth(root, 0)
	return mxDepthResult
}

func depth(root *TreeNode, dep int) {
	if root != nil {
		mxDepthResult = MyMax(mxDepthResult, dep+1)
		depth(root.Left, dep+1)
		depth(root.Right, dep+1)
	}
}
