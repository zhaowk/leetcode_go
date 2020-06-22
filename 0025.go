package leetcode_go

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	h := &ListNode{0, nil}
	h.Next = head
	nextHead := h

	for nextHead != nil {
		needReverse := true
		p := nextHead
		head = p.Next
		for i := 0; i < k; i++ {
			if p == nil || p.Next == nil {
				needReverse = false
				break
			}
			p = p.Next
		}

		if !needReverse {
			nextHead.Next = head
			break
		} else {
			current := nextHead.Next
			nextHead.Next = p.Next
			p = current

			for i := 0; i < k; i++ {
				q := nextHead.Next
				nextHead.Next = current
				current = current.Next
				nextHead.Next.Next = q
			}
			nextHead = p
		}
	}
	return h.Next
}
