package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestZigzagLevelOrder(t *testing.T) {
	assert.Equal(t, [][]int{{3}, {20, 9}, {15, 7}}, zigzagLevelOrder(BuildTree(3, 9, 20, nil, nil, 15, 7)))
}
