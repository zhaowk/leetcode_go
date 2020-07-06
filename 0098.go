package leetcode_go

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidLeft(root.Left, root.Val, math.MinInt64, int64(root.Val)) &&
		isValidRight(root.Right, root.Val, int64(root.Val), math.MaxInt64)
}

func isValidLeft(n *TreeNode, val int, mi, mx int64) bool {
	if n == nil {
		return true
	}

	if n.Val >= val || int64(n.Val) <= mi || int64(n.Val) >= mx {
		return false
	}

	if n.Val >= val || !isValidLeft(n.Left, n.Val, mi, int64(n.Val)) || !isValidRight(n.Right, n.Val, int64(n.Val), mx) {
		return false
	}

	return true
}

func isValidRight(n *TreeNode, val int, mi, mx int64) bool {
	if n == nil {
		return true
	}

	if n.Val <= val || int64(n.Val) <= mi || int64(n.Val) >= mx {
		return false
	}

	if !isValidLeft(n.Left, n.Val, mi, int64(n.Val)) || !isValidRight(n.Right, n.Val, int64(n.Val), mx) {
		return false
	}

	return true
}
