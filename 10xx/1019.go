package _0xx

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
func nextLargerNodes(head *ListNode) []int {
	type idxVal struct{ idx, val int }
	stack := make([]idxVal, 0)
	m := make(map[int]int)
	num := 0

	if head == nil {
		return []int{}
	}

	for ; head != nil; head = head.Next {
		for len(stack) > 0 && stack[len(stack)-1].val < head.Val {
			m[stack[len(stack)-1].idx] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idxVal{num, head.Val})
		num++
	}

	r := make([]int, num)
	for k, v := range m {
		if k < num {
			r[k] = v
		}
	}

	return r
}
