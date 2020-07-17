package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	pHead := &ListNode{}

	p := head

	for p != nil {
		q := pHead

		for q.Next != nil && p.Val > q.Next.Val {
			q = q.Next
		}

		t := p.Next
		p.Next = q.Next
		q.Next = p
		p = t
	}

	return pHead.Next
}
