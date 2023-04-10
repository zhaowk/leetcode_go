package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.False(t, isPalindrome(-1))
	assert.False(t, isPalindrome(112321))
	assert.True(t, isPalindrome(12321))
	assert.True(t, isPalindrome(123321))
	assert.True(t, isPalindrome(1234554321))
}
