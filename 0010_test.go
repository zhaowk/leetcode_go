package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	assert.True(t, isMatch("", ""))
	assert.False(t, isMatch("", "a"))
	assert.False(t, isMatch("aa", "a"))
	assert.True(t, isMatch("aa", "a*"))
	assert.True(t, isMatch("ab", ".*"))
	assert.True(t, isMatch("aab", "c*a*b"))
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
}
