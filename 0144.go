package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 非递归版本
func preorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	if root == nil {
		return result
	}

	stack = append(stack, root)
	result = append(result, root.Val)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]

		for curr.Left != nil {
			curr = curr.Left
			result = append(result, curr.Val)
			stack = append(stack, curr)
		}

		for len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if curr.Right != nil {
				curr = curr.Right
				result = append(result, curr.Val)
				stack = append(stack, curr)
				break
			}
		}
	}

	return result
}

// 递归版本
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	preorderTraversalF(root, &result)
	return result
}

func preorderTraversalF(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preorderTraversalF(root.Left, result)
		preorderTraversalF(root.Right, result)
	}
}
