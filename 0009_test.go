package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(12321))
	assert.True(t, isPalindrome(123321))
	assert.True(t, isPalindrome(1234554321))
}
