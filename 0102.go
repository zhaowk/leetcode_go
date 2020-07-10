package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
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

	return result
}
