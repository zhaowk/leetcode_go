package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	assert.Equal(t, [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 3},
		{2},
		{2, 3},
		{3},
	}, subsets([]int{1, 2, 3}))
	assert.Equal(t, [][]int{
		{},
		{1},
		{1, 3},
		{3},
	}, subsets([]int{1, 3}))

	//fmt.Println(subsets([]int{1,2,3}))
	//fmt.Println(subsets([]int{1,3}))
	//fmt.Println(subsets([]int{1,2,3,4}))
}
