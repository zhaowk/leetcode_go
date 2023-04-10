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
func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	prev, curr := root, root
	stack := make([]*TreeNode, 0)

	for curr != nil {
		for curr.Left == nil {
			if curr.Right == nil {
				if len(stack) == 0 {
					return
				} else {
					curr = stack[len(stack)-1]
					stack = stack[:len(stack)-1] // pop
					prev.Left = nil
					prev.Right = curr
					prev = curr
				}
				break
			}
			curr = curr.Right
			prev.Left = nil
			prev.Right = curr
			prev = curr
		}

		for curr.Left != nil {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			curr = curr.Left
			prev.Left = nil
			prev.Right = curr
			prev = curr
		}
	}
}
