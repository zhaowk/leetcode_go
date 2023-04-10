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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return genBst(nums, 0, len(nums)-1)
}

func genBst(nums []int, left, right int) *TreeNode {
	if left == right {
		return &TreeNode{nums[left], nil, nil}
	} else if left > right {
		return nil
	}

	mid := (left + right) / 2

	n := &TreeNode{nums[mid], nil, nil}

	n.Left = genBst(nums, left, mid-1)
	n.Right = genBst(nums, mid+1, right)
	return n
}
