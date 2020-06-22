package leetcode_go

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{0, nil}
	p := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
		return head.Next
	}

	if l2 == nil {
		p.Next = l1
		return head.Next
	}

	return head.Next
}
