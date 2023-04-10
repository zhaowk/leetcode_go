package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	var p, q *ListNode = nil, nil
	p = head
	q = head

	for i := 0; i < n; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}

	if p == nil {
		return head.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	q.Next = q.Next.Next

	return head
}
