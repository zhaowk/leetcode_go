package leetcode_go

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil { // at lease two elements
		return head
	}
	h := &ListNode{0, nil}

	h.Next = head
	prev := h
	mid := prev.Next
	last := mid.Next

	for mid != nil && last != nil {
		next := last.Next
		prev.Next = last
		last.Next = mid
		mid.Next = next
		prev = mid
		mid = prev.Next

		if mid != nil {
			last = mid.Next
		}
	}

	return h.Next
}
