package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pathSumResult [][]int

func pathSum1(root *TreeNode, sum int) [][]int {
	pathSumResult = make([][]int, 0)
	if root == nil {
		return pathSumResult
	}

	p := make([]int, 0)
	pathSums(root, 0, sum, p)
	return pathSumResult
}

func pathSums(n *TreeNode, t, sum int, path []int) {
	if n != nil {
		if n.Left == nil && n.Right == nil {
			if t+n.Val == sum {
				tmp := make([]int, len(path)+1)
				copy(tmp, path)
				tmp[len(tmp)-1] = n.Val
				pathSumResult = append(pathSumResult, tmp)
			}
		} else {
			pathSums(n.Left, t+n.Val, sum, append(path, n.Val))
			pathSums(n.Right, t+n.Val, sum, append(path, n.Val))
		}
	}
}
