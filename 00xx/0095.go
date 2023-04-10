package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func generateTrees(n int) []*TreeNode {
	treesResult := make([]*TreeNode, 0)
	if n == 0 {
		return treesResult
	} else if n == 1 {
		treesResult = append(treesResult, &TreeNode{1, nil, nil})
		return treesResult
	}

	for i := 1; i <= n; i++ { // left
		left, right := make([]*TreeNode, 0), make([]*TreeNode, 0)
		genLeft(n, i-1, 0, nil, &left)
		genRight(n, i+1, 0, nil, &right)

		root := &TreeNode{i, nil, nil}
		for _, l := range left {
			for _, r := range right {
				root.Left = l
				root.Right = r
				node := &TreeNode{i, nil, nil}
				copyTree(root, node)
				treesResult = append(treesResult, node)
			}
		}
	}

	return treesResult
}

func copyTree(root, node *TreeNode) {
	node.Val = root.Val
	if root.Left != nil {
		node.Left = &TreeNode{root.Left.Val, nil, nil}
		copyTree(root.Left, node.Left)
	}

	if root.Right != nil {
		node.Right = &TreeNode{root.Right.Val, nil, nil}
		copyTree(root.Right, node.Right)
	}
}

func appendTree(r []*TreeNode, i *TreeNode) []*TreeNode {
	for _, v := range r {
		if TreeEq(v, i) {
			return r
		}
	}
	if i == nil {
		return append(r, nil)
	}
	node := &TreeNode{i.Val, nil, nil}
	copyTree(i, node)
	return append(r, node)
}

func genLeft(n, k, step int, root *TreeNode, r *[]*TreeNode) {
	if k <= step {
		*r = appendTree(*r, root)
		return
	}

	for i := 1; i <= k; i++ {
		if step == 0 {
			root = nil
		}

		if root == nil {
			root = &TreeNode{i, nil, nil}
			genLeft(n, k, step+1, root, r)
		} else {
			node := root
			for node != nil {
				val := node.Val
				if i == val {
					break
				} else if i < val {
					if node.Left == nil {
						node.Left = &TreeNode{i, nil, nil}
						genLeft(n, k, step+1, root, r)
						node.Left = nil
						break
					} else {
						node = node.Left
					}
				} else {
					if node.Right == nil {
						node.Right = &TreeNode{i, nil, nil}
						genLeft(n, k, step+1, root, r)
						node.Right = nil
						break
					} else {
						node = node.Right
					}
				}
			}
		}
	}
}

func genRight(n, k, step int, root *TreeNode, r *[]*TreeNode) {

	if n-k < step {
		*r = appendTree(*r, root)
		return
	}

	for i := k; i <= n; i++ {
		if step == 0 {
			root = nil
		}

		if root == nil {
			root = &TreeNode{i, nil, nil}
			genRight(n, k, step+1, root, r)
		} else {
			node := root
			for node != nil {
				val := node.Val
				if i == val {
					break
				} else if i < val {
					if node.Left == nil {
						node.Left = &TreeNode{i, nil, nil}
						genRight(n, k, step+1, root, r)
						node.Left = nil
						break
					} else {
						node = node.Left
					}
				} else {
					if node.Right == nil {
						node.Right = &TreeNode{i, nil, nil}
						genRight(n, k, step+1, root, r)
						node.Right = nil
						break
					} else {
						node = node.Right
					}
				}
			}
		}
	}
}
