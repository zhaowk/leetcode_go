package leetcode_go

func kthSmallest1(root *TreeNode, k int) int {
	i, r := 0, 0
	var findKth func (root *TreeNode)

	findKth = func(root *TreeNode) {
		if root == nil {
			return
		}

		findKth(root.Left)
		i++
		if i == k {
			r = root.Val
			return
		}
		findKth(root.Right)
	}

	findKth(root)
	return r
}


