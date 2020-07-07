package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	assert.True(t, isInterleave("", "", ""))
	assert.True(t, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.False(t, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}
