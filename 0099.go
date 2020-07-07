package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := make([]*TreeNode, 0)

	inorder(root, &tmp)

	//fmt.Println(tmp)

	var prev, next *TreeNode = nil, nil

	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i].Val > tmp[i+1].Val {

			next = tmp[i+1]

			if prev == nil {
				prev = tmp[i]
			} else {
				break
			}
		}
	}
	if prev != nil && next != nil {
		prev.Val, next.Val = next.Val, prev.Val
	}
}

func inorder(node *TreeNode, r *[]*TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left, r)
	*r = append(*r, node)
	inorder(node.Right, r)
}
