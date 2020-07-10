package leetcode_go

import "math"

var miDepthResult int

func minDepth(root *TreeNode) int {
	miDepthResult = math.MaxInt32

	if root == nil {
		return 0
	}
	depthT(root, 0)
	return miDepthResult
}

func depthT(n *TreeNode, dep int) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			miDepthResult = min(miDepthResult, dep+1)
		} else {
			depthT(n.Left, dep+1)
			depthT(n.Right, dep+1)
		}
	}
}
