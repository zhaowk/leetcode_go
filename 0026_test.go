package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	//a := []int{1,1,2}
	//b := []int{1,2}
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b := []int{0, 1, 2, 3, 4}

	c := removeDuplicates(a)
	assert.Equal(t, b, a[:c])

	assert.Equal(t, 0, removeDuplicates([]int{}))
}
