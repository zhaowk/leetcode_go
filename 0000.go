package leetcode_go

import (
	"bytes"
	"reflect"
	"strconv"
)

func max(i ...int) int {
	m := -2147483648
	for _, v := range i {
		if v > m {
			m = v
		}
	}
	return m
}

func min(i ...int) int {
	m := 2147483647
	for _, v := range i {
		if v < m {
			m = v
		}
	}
	return m
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	r := l
	s := bytes.Buffer{}
	first := true
	s.WriteString("[")
	for r != nil {
		if first {
			first = false
		} else {
			s.WriteString(",")
		}
		s.WriteString(strconv.Itoa(r.Val))
		r = r.Next
	}

	s.WriteString("]")
	return s.String()
}

func buildList(n ...int) *ListNode {

	head := &ListNode{Val: 0, Next: nil}
	for i := len(n) - 1; i >= 0; i-- {
		p := head.Next
		head.Next = &ListNode{Val: n[i], Next: p}
	}
	return head.Next
}

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

func buildTree(nodes ...interface{}) (root *TreeNode) {

	ln := len(nodes)
	if ln == 0 {
		return nil
	}

	if nodes[0] != nil && reflect.TypeOf(nodes[0]).Kind() == reflect.Int {
		root = &TreeNode{int(reflect.ValueOf(nodes[0]).Int()), nil, nil}
	}

	items := make([]*TreeNode, 0)
	items = append(items, root)
	for i := 1; i < ln; i++ {
		if i%2 == 1 { // left
			if nodes[i] == nil {
				items[0].Left = nil
			} else {
				if reflect.TypeOf(nodes[i]).Kind() == reflect.Int {
					item := &TreeNode{int(reflect.ValueOf(nodes[i]).Int()), nil, nil}
					items[0].Left = item
					items = append(items, item)
				}
			}
		} else { // right
			if nodes[i] == nil {
				items[0].Right = nil
			} else {
				if reflect.TypeOf(nodes[i]).Kind() == reflect.Int {
					item := &TreeNode{int(reflect.ValueOf(nodes[i]).Int()), nil, nil}
					items[0].Right = item
					items = append(items, item)
				}
			}
			items = items[1:]
		}
	}

	return root
}

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
