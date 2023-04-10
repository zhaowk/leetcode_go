package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
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
		BuildList(1, 1, 2, 3, 4, 4, 5, 6),
		mergeKLists([]*ListNode{
			BuildList(1, 4, 5),
			BuildList(1, 3, 4),
			BuildList(2, 6)}))
}
