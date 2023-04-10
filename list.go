package _0xx

import (
	"bytes"
	"strconv"
)

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

func BuildList(n ...int) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	for i := len(n) - 1; i >= 0; i-- {
		p := head.Next
		head.Next = &ListNode{Val: n[i], Next: p}
	}
	return head.Next
}
