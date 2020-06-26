package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	assert.Equal(t, [][]int{{1, 7}}, insert([][]int{{1, 5}, {2, 7}}, []int{}))
	assert.Equal(t, [][]int{{1, 7}}, insert([][]int{{1, 5}}, []int{2, 7}))
	assert.Equal(t, [][]int{{1, 5}}, insert([][]int{{1, 5}}, []int{2, 3}))
	assert.Equal(t, [][]int{{0, 0}, {1, 5}}, insert([][]int{{1, 5}}, []int{0, 0}))
}
