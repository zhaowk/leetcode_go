package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func buildCycle(a []int, cycle int) *ListNode {

	if cycle <= 0 {
		return BuildList(a...)
	}

	head := &ListNode{}
	var pos *ListNode = nil
	p := head

	for idx, v := range a {
		p.Next = &ListNode{v, nil}
		p = p.Next
		if idx == cycle-1 {
			pos = p
		}
	}
	if pos != nil {
		p.Next = pos
	}
	return head.Next
}

func TestHasCycle(t *testing.T) {
	assert.False(t, hasCycle(BuildList()))
	assert.False(t, hasCycle(BuildList(1)))
	assert.False(t, hasCycle(BuildList(1, 2)))
	assert.False(t, hasCycle(BuildList(1, 2, 3)))
	assert.False(t, hasCycle(BuildList(1, 2, 3, 4)))
	assert.False(t, hasCycle(BuildList(1, 2, 3, 4, 3, 2, 1)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 1)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 2)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 3)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 4)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 5)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 6)))
	assert.True(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 7)))
	assert.False(t, hasCycle(buildCycle([]int{1, 2, 3, 4, 3, 2, 1}, 8)))
}
