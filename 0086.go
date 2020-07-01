package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return head
	}

	blow := &ListNode{}
	above := &ListNode{}
	p, q := blow, above

	for head != nil {
		if head.Val < x {
			p.Next = head
			head = head.Next
			p = p.Next
			p.Next = nil
		} else {
			q.Next = head
			head = head.Next
			q = q.Next
			q.Next = nil
		}
	}

	if p == blow {
		return above.Next
	}

	p.Next = above.Next
	return blow.Next
}
