package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func maxAncestorDiff(root *TreeNode) int {
	var result int
	maxAncestorDiffRec(root, root.Val, root.Val, &result)
	return result
}

func maxAncestorDiffRec(ancestor *TreeNode, mini, maxi int, result *int) {
	if ancestor == nil {
		return
	}

	if ancestor.Left != nil {
		mi, mx := MyMin(mini, ancestor.Left.Val), MyMax(maxi, ancestor.Left.Val)
		*result = MyMax(*result, MyAbs(ancestor.Left.Val-mi), MyAbs(ancestor.Left.Val-mx))
		maxAncestorDiffRec(ancestor.Left, mi, mx, result)
	}

	if ancestor.Right != nil {
		mi, mx := MyMin(mini, ancestor.Right.Val), MyMax(maxi, ancestor.Right.Val)
		*result = MyMax(*result, MyAbs(ancestor.Right.Val-mi), MyAbs(ancestor.Right.Val-mx))
		maxAncestorDiffRec(ancestor.Right, mi, mx, result)
	}
}
