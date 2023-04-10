package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	node := &ListNode{}

	p, q, r := head, head.Next, node
	prev := false

	for p != nil {
		if q == nil { // last node
			if !prev {
				r.Next = p
			}
			break
		}

		if p.Val != q.Val {
			if !prev {
				r.Next = p
				p, q, r = p.Next, q.Next, r.Next
				r.Next = nil
			} else {
				p, q = p.Next, q.Next
			}
			prev = false
		} else {
			prev = true
			p, q = p.Next, q.Next
		}
	}

	return node.Next
}
