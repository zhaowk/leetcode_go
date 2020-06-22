package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	assert.Equal(t, buildList(), swapPairs(buildList()))
	assert.Equal(t, buildList(1), swapPairs(buildList(1)))
	assert.Equal(t, buildList(2, 1), swapPairs(buildList(1, 2)))
	assert.Equal(t, buildList(2, 1, 3), swapPairs(buildList(1, 2, 3)))
	assert.Equal(t, buildList(2, 1, 4, 3), swapPairs(buildList(1, 2, 3, 4)))
}
