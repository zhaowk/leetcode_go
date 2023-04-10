package _0xx

import (
	"bytes"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	//var depth = 0
	b := bytes.Buffer{}
	b.WriteString("[")
	parent := make([]*TreeNode, 1)
	parent[0] = t
	first := true

	for {
		child := make([]*TreeNode, 0)
		for _, v := range parent {
			if first {
				first = false
			} else {
				b.WriteString(",")
			}
			if nil == v {
				b.WriteString("null")
			} else {
				b.WriteString(strconv.Itoa(v.Val))
				child = append(child, v.Left)
				child = append(child, v.Right)
			}
		}

		parent = child
		hasNext := false
		for _, v := range parent {
			if v != nil {
				hasNext = true
				break
			}
		}
		if !hasNext {
			break
		}
	}

	b.WriteString("]")
	return string(b.Bytes())
}

func BuildTree(nodes ...interface{}) (root *TreeNode) {
	ln := len(nodes)
	if ln == 0 {
		return nil
	}

	if v, ok := nodes[0].(int); ok {
		root = &TreeNode{v, nil, nil}
	} else {
		return nil // no root
	}

	items := make([]*TreeNode, 0)
	items = append(items, root)
	for i := 1; i < ln; i++ {
		if i%2 == 1 { // left
			if nodes[i] == nil {
				items[0].Left = nil
			} else {
				if v, ok := nodes[i].(int); ok {
					item := &TreeNode{v, nil, nil}
					items[0].Left = item
					items = append(items, item)
				}
			}
		} else { // right
			if nodes[i] == nil {
				items[0].Right = nil
			} else {
				if v, ok := nodes[i].(int); ok {
					item := &TreeNode{v, nil, nil}
					items[0].Right = item
					items = append(items, item)
				}
			}
			items = items[1:]
		}
	}

	return root
}

func TreeEq(i, j *TreeNode) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return TreeEq(i.Left, j.Left) && TreeEq(i.Right, j.Right)
}
