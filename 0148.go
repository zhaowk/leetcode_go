package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	qHead := &ListNode{}
	qHead.Next = head
	qSort(qHead, nil)
	return qHead.Next
}

func qSort(start, end *ListNode) {

	if start == nil || start.Next == end {
		return
	}

	h, p := start.Next, start.Next
	target := h.Val

	for p.Next != end {
		if p.Next.Val < target {
			q := start.Next
			start.Next = p.Next
			p.Next = p.Next.Next
			start.Next.Next = q
		} else {
			p = p.Next
		}
	}

	qSort(start, h)
	qSort(h, end)
}
