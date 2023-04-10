package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func getNthNode(head *ListNode, n int) *ListNode {
	for i := 1; i < n; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

func TestDetectCycle(t *testing.T) {
	a := buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 1)
	assert.True(t, detectCycle(a) == getNthNode(a, 1))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 2)
	assert.True(t, detectCycle(a) == getNthNode(a, 2))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 3)
	assert.True(t, detectCycle(a) == getNthNode(a, 3))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 4)
	assert.True(t, detectCycle(a) == getNthNode(a, 4))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 5)
	assert.True(t, detectCycle(a) == getNthNode(a, 5))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 6)
	assert.True(t, detectCycle(a) == getNthNode(a, 6))

	a = buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 7)
	assert.True(t, detectCycle(a) == getNthNode(a, 7))
}
