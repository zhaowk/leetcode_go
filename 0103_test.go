package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	assert.Equal(t, [][]int{{3}, {20, 9}, {15, 7}}, zigzagLevelOrder(buildTree(3, 9, 20, nil, nil, 15, 7)))
}
