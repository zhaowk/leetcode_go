package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCut(t *testing.T) {
	assert.Equal(t, 1, minCut("aab"))
	assert.Equal(t, 1, minCut("aabb"))
	assert.Equal(t, 0, minCut("aabbaa"))
	assert.Equal(t, 4, minCut("coder"))
}
