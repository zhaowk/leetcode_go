package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{0, nil}
	p := head
	heads := make([]*ListNode, 0)

	for _, l := range lists {
		heads = append(heads, l)
	}

	for {
		minimum := 2147483647
		merging := -1
		for idx, h := range heads {
			if h != nil {
				if h.Val <= minimum {
					merging = idx
					minimum = h.Val
				}
			}
		}

		if merging == -1 {
			return head.Next
		}

		p.Next = heads[merging]
		p = p.Next
		heads[merging] = heads[merging].Next
	}
}
