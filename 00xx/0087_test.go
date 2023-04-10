package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsScramble(t *testing.T) {
	assert.True(t, isScramble("great", "great"))
	assert.True(t, isScramble("great", "taerg"))
	assert.True(t, isScramble("great", "rgeat"))
	assert.True(t, isScramble("great", "rgtae"))
	assert.True(t, isScramble("great", "taerg"))
	assert.False(t, isScramble("abcde", "caebd"))
	assert.True(t, isScramble("abc", "bca"))
}
