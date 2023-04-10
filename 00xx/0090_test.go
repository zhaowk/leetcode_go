package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	assert.Equal(t, [][]int{{}}, subsetsWithDup([]int{}))
	assert.Equal(t, [][]int{{}, {1}}, subsetsWithDup([]int{1}))
	assert.Equal(t, [][]int{{}, {1}, {1, 2}, {2}}, subsetsWithDup([]int{1, 2}))
	assert.Equal(t, [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 2}, {2}}, subsetsWithDup([]int{1, 1, 2}))
	assert.Equal(t, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetsWithDup([]int{1, 2, 2}))
	//fmt.Println(subsetsWithDup([]int{1,2,2,2}))
}
