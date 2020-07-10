package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}

	root := &TreeNode{}
	root.Val = postorder[len(postorder)-1]
	pos := -1

	for idx := range inorder {
		if inorder[idx] == postorder[len(postorder)-1] {
			pos = idx
			break
		}
	}

	if pos > 0 {
		root.Left = buildTreeIP(inorder[:pos], postorder[:pos])
	}

	if pos < len(inorder) {
		root.Right = buildTreeIP(inorder[pos+1:], postorder[pos:len(postorder)-1])
	}
	return root
}
