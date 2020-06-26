package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]bool, 0)

	result := &ListNode{0, nil}
	p := result

	for head != nil {
		if !m[head.Val] {
			m[head.Val] = true
			p.Next = &ListNode{head.Val, nil}
			p = p.Next
		}
		head = head.Next
	}

	return result.Next
}
