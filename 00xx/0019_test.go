package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	. "github.com/zhaowk/leetcode_go"
)

func TestRemoveNthFromEnd(t *testing.T) {
	a := BuildList(1, 2, 3, 4, 5)
	assert.Equal(t, BuildList(2, 3, 4, 5), removeNthFromEnd(a, 5))
	assert.Equal(t, BuildList(1, 2, 4, 5), removeNthFromEnd(a, 3))
	//fmt.Println(a)
	//fmt.Println(removeNthFromEnd(a, 5))

	a = BuildList(1)
	assert.Equal(t, BuildList(), removeNthFromEnd(a, 1))

	//fmt.Println(a)
	//fmt.Println(removeNthFromEnd(a, 1))
}
