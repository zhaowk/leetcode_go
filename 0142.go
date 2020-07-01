package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
