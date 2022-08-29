package leetcode_go

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if findNode(p, q) {
		return p
	} else if findNode(q, p) {
		return q
	}

	if findNode(root.Left, p) && findNode(root.Left, q) {
		return lowestCommonAncestor(root.Left, p, q)
	} else if findNode(root.Right, p) && findNode(root.Right, q) {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func findNode(root, p *TreeNode) bool {
	if root == nil {
		return false
	}

	if root == p {
		return true
	}

	return findNode(root.Left, p) || findNode(root.Right, p)
}
