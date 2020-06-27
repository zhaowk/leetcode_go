package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	node := &ListNode{0, nil}
	node.Next = head

	prev, curr := head, head.Next

	for curr != nil {
		if curr.Val == prev.Val {
			prev.Next = curr.Next
			curr = prev.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return node.Next
}
