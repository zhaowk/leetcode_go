package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}
