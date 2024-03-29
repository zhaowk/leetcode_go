package _1xx

import "math"
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

func maxPathSum2(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := MyMax(maxGain(node.Left), 0)
		rightGain := MyMax(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = MyMax(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + MyMax(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

var maximum = 0

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}
	maximum = root.Val
	pathSum(root)
	return maximum
}

func pathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lSum := MyMax(0, pathSum(node.Left))
	rSum := MyMax(0, pathSum(node.Right))

	maximum = MyMax(maximum, lSum+rSum+node.Val)
	return node.Val + MyMax(lSum, rSum)
}
