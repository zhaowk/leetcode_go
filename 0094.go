package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var inorderTraversalResult []int

func inorderTraversal(root *TreeNode) []int {
	inorderTraversalResult = make([]int, 0)

	if root == nil {
		return inorderTraversalResult
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		for p.Left != nil {
			stack = append(stack, p.Left)
			p = p.Left
		}

		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		inorderTraversalResult = append(inorderTraversalResult, p.Val)

		for len(stack) > 0 && p.Right == nil {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inorderTraversalResult = append(inorderTraversalResult, p.Val)
		}

		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}

	// inorderTraversalF(root)
	return inorderTraversalResult
}

// func inorderTraversalF(root *TreeNode) {
//     if root != nil {
//         inorderTraversalF(root.Left)
//         inorderTraversalResult = append(inorderTraversalResult, root.Val)
//         inorderTraversalF(root.Right)
//     }
// }
