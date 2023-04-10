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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}

	node := &ListNode{}
	node.Next = head
	p, q := node, head

	for i := 1; i < m; i++ {
		p, q = p.Next, q.Next
		if q == nil {
			return head
		}
	}

	p.Next = nil
	r := q

	for i := 0; i <= n-m; i++ {
		head = q.Next
		q.Next = p.Next
		p.Next = q
		q = head
		if q == nil {
			break
		}
	}

	r.Next = q

	return node.Next
}
