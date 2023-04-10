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
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	prev := make([]*TreeNode, 0)

	prev = append(prev, root)

	for len(prev) > 0 {
		tmp := make([]int, 0)
		curr := make([]*TreeNode, 0)
		for _, value := range prev {
			tmp = append(tmp, value.Val)
			if value.Left != nil {
				curr = append(curr, value.Left)
			}

			if value.Right != nil {
				curr = append(curr, value.Right)
			}
		}
		result = append(result, tmp)
		prev = curr
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}
