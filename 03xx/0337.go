package _3xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func rob2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return MyMax(robInternal2(root, true), robInternal2(root, false))
}

func robInternal2(node *TreeNode, useNode bool) int {
	if node == nil {
		return 0
	}

	if useNode {
		return node.Val + robInternal2(node.Left, false) + robInternal2(node.Right, false)
	} else {
		return MyMax(robInternal2(node.Left, false), robInternal2(node.Left, true)) +
			MyMax(robInternal2(node.Right, false), robInternal2(node.Right, true))
	}
}
