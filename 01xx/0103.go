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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	prev := make([]*TreeNode, 0)

	prev = append(prev, root)
	left := true

	for len(prev) > 0 {
		tmp := make([]int, 0)
		curr := make([]*TreeNode, 0)
		for idx := range prev {
			value := prev[len(prev)-1-idx]
			tmp = append(tmp, value.Val)
			if left {
				if value.Left != nil {
					curr = append(curr, value.Left)
				}

				if value.Right != nil {
					curr = append(curr, value.Right)
				}
			} else {
				if value.Right != nil {
					curr = append(curr, value.Right)
				}
				if value.Left != nil {
					curr = append(curr, value.Left)
				}
			}
		}
		result = append(result, tmp)
		prev = curr
		left = !left
	}

	return result
}
