package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	if k == 0 {
		return head
	}

	lh := 0
	p, q := head, head
	for i := 0; i < k; i++ {
		if p.Next == nil {
			lh = i + 1
			break
		} else {
			p = p.Next
		}
	}

	if lh != 0 {
		k = k % lh
		if k == 0 {
			return head
		}
		p = head
		for i := 0; i < k; i++ {
			if p.Next == nil {
				lh = i
				break
			} else {
				p = p.Next
			}
		}
	}

	for p != nil && p.Next != nil {
		p = p.Next
		q = q.Next
	}
	node := &ListNode{}
	node.Next = q.Next
	q.Next = nil
	p.Next = head

	return node.Next
}
