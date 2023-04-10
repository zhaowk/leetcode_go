package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestMaxPathSum(t *testing.T) {
	a := BuildTree(1, 2, 3)
	assert.Equal(t, 6, maxPathSum(a))
	assert.Equal(t, 6, maxPathSum2(a))

	b := BuildTree(-10, 9, 20, nil, nil, 15, 7)
	assert.Equal(t, 42, maxPathSum(b))
	assert.Equal(t, 42, maxPathSum2(b))

	c := BuildTree(1, 2, nil, 3, nil, 4, nil, 5)
	assert.Equal(t, 15, maxPathSum(c))
	assert.Equal(t, 15, maxPathSum2(c))

	d := BuildTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1)
	assert.Equal(t, 48, maxPathSum(d))
	assert.Equal(t, 48, maxPathSum2(d))
}
