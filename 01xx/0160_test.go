package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func build2List(t1, t2, t []int) []*ListNode {
	l1 := BuildList(t1...)
	l2 := BuildList(t2...)
	l := BuildList(t...)

	if l1 == nil {
		l1 = l
	} else {
		p := l1
		for p.Next != nil {
			p = p.Next
		}
		p.Next = l
	}

	if l2 == nil {
		l2 = l
	} else {
		p := l2
		for p.Next != nil {
			p = p.Next
		}
		p.Next = l
	}

	return []*ListNode{l1, l2}
}

func TestGetIntersectionNode(t *testing.T) {

	a := build2List([]int{1, 2}, []int{3, 4}, []int{5, 6})
	assert.Equal(t, BuildList(5, 6), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1}, []int{3, 4}, []int{5, 6})
	assert.Equal(t, BuildList(5, 6), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1, 2}, []int{3}, []int{5, 6})
	assert.Equal(t, BuildList(5, 6), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{}, []int{3, 4}, []int{5, 6})
	assert.Equal(t, BuildList(5, 6), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1, 2}, []int{}, []int{5, 6})
	assert.Equal(t, BuildList(5, 6), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1, 2}, []int{3, 4}, []int{})
	assert.Equal(t, BuildList(), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1, 2}, []int{3}, []int{})
	assert.Equal(t, BuildList(), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1}, []int{3, 4}, []int{})
	assert.Equal(t, BuildList(), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{}, []int{3, 4}, []int{})
	assert.Equal(t, BuildList(), getIntersectionNode(a[0], a[1]))

	a = build2List([]int{1}, []int{}, []int{})
	assert.Equal(t, BuildList(), getIntersectionNode(a[0], a[1]))
}
