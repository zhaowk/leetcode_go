package _0xx

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestIsValidBST(t *testing.T) {
	assert.True(t, isValidBST(BuildTree(2, 1, 3)))
	//assert.True(t, isValidBST(buildTree(2,1,3)))
	assert.False(t, isValidBST(BuildTree(10, 5, 15, nil, nil, 6, 20)))
	assert.False(t, isValidBST(BuildTree(3, 1, 5, 0, 2, 4, 6, nil, nil, nil, 3)))
	assert.True(t, isValidBST(BuildTree(0, math.MinInt32, math.MaxInt32)))

}
