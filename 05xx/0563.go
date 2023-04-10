package _5xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func findTilt(root *TreeNode) int {
	result := 0
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var findTiltNode func(node *TreeNode) int
	findTiltNode = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := findTiltNode(node.Left)
		right := findTiltNode(node.Right)
		result += abs(left - right)
		return left + right + node.Val
	}

	findTiltNode(root)

	return result
}
