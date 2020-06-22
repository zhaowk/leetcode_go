package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	//[
	//	1->4->5,
	//	1->3->4,
	//	2->6
	//]
	//输出: 1->1->2->3->4->4->5->6
	//fmt.Println(mergeKLists([]*ListNode{buildList(1, 4, 5), buildList(1, 3, 4), buildList(2, 6)}))

	assert.Equal(t,
		buildList(1, 1, 2, 3, 4, 4, 5, 6),
		mergeKLists([]*ListNode{
			buildList(1, 4, 5),
			buildList(1, 3, 4),
			buildList(2, 6)}))
}
