package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestBuildTreeIP(t *testing.T) {
	assert.Equal(t, BuildTree(1, nil, 2), buildTreeIP([]int{1, 2}, []int{2, 1}))
	assert.Equal(t, BuildTree(3, 9, 20, nil, nil, 15, 7), buildTreeIP([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
