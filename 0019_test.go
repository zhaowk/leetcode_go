package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	a := buildList(1, 2, 3, 4, 5)
	assert.Equal(t, buildList(2, 3, 4, 5), removeNthFromEnd(a, 5))
	assert.Equal(t, buildList(1, 2, 4, 5), removeNthFromEnd(a, 3))
	//fmt.Println(a)
	//fmt.Println(removeNthFromEnd(a, 5))

	a = buildList(1)
	assert.Equal(t, buildList(), removeNthFromEnd(a, 1))

	//fmt.Println(a)
	//fmt.Println(removeNthFromEnd(a, 1))
}
