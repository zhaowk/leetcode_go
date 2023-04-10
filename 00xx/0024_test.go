package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestSwapPairs(t *testing.T) {
	assert.Equal(t, BuildList(), swapPairs(BuildList()))
	assert.Equal(t, BuildList(1), swapPairs(BuildList(1)))
	assert.Equal(t, BuildList(2, 1), swapPairs(BuildList(1, 2)))
	assert.Equal(t, BuildList(2, 1, 3), swapPairs(BuildList(1, 2, 3)))
	assert.Equal(t, BuildList(2, 1, 4, 3), swapPairs(BuildList(1, 2, 3, 4)))
}
