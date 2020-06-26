package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, [][]int{}, merge([][]int{}))
	assert.Equal(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	assert.Equal(t, [][]int{{1, 4}}, merge([][]int{{1, 4}, {2, 3}}))
}
