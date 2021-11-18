package leetcode_go

func countNodes(root *TreeNode) int {
	cnt := 0

	q := make([]*TreeNode, 0)

	if root != nil {
		q = append(q, root)
		cnt++
	}

	for len(q) > 0 {
		i := q[0]
		if i.Left != nil {
			q = append(q, i.Left)
			cnt++
		}

		if i.Right != nil {
			q = append(q, i.Right)
			cnt++
		}

		q = q[1:]
	}

	return cnt
}
