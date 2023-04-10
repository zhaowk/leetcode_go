package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestMerTwoLists(t *testing.T) {

	//[
	//	1->4->5,
	//	1->3->4,
	//
	//fmt.Println(mergeTwoLists(buildList(1, 4, 5), buildList(1, 3, 4)))

	assert.Equal(t,
		BuildList(1, 1, 3, 4, 4, 5),
		mergeTwoLists(BuildList(1, 4, 5), BuildList(1, 3, 4)))
}
