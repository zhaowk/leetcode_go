package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerTwoLists(t *testing.T) {

	//[
	//	1->4->5,
	//	1->3->4,
	//
	//fmt.Println(mergeTwoLists(buildList(1, 4, 5), buildList(1, 3, 4)))

	assert.Equal(t,
		buildList(1, 1, 3, 4, 4, 5),
		mergeTwoLists(buildList(1, 4, 5), buildList(1, 3, 4)))
}
