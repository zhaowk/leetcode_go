package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	la, lb := 0, 0
	pa, pb := headA, headB
	for pa != nil && pb != nil {
		pa = pa.Next
		pb = pb.Next
	}

	for pa != nil {
		la++
		pa = pa.Next
	}

	for pb != nil {
		lb++
		pb = pb.Next
	}

	pa, pb = headA, headB
	for i := 0; i < la && pa != nil; i++ {
		pa = pa.Next
	}

	for i := 0; i < lb && pb != nil; i++ {
		pb = pb.Next
	}

	if pa == nil || pb == nil {
		return nil
	}

	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}
