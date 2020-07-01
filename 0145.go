package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal1(root *TreeNode) []int {

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	if root == nil {
		return result
	}

	stack = append(stack, root)
	//result = append(result, root.Val)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]

		for curr.Left != nil {
			curr = curr.Left
			//result = append(result, curr.Val)
			stack = append(stack, curr)
		}

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		} else { // left ,right  null
			result = append(result, curr.Val)
			stack = stack[:len(stack)-1] // pop
			for len(stack) > 0 {
				prev := stack[len(stack)-1]
				if prev.Right == nil || prev.Right == curr {
					result = append(result, prev.Val)
					stack = stack[:len(stack)-1] // pop
					curr = prev
				} else {
					stack = append(stack, prev.Right)
					break
				}
			}
		}
	}

	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	postorderTraversalF(root, &result)
	return result
}

func postorderTraversalF(root *TreeNode, r *[]int) {
	if root != nil {
		postorderTraversalF(root.Left, r)
		postorderTraversalF(root.Right, r)
		*r = append(*r, root.Val)
	}
}
