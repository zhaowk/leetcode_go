package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeStr(t *testing.T) {
	assert.True(t, isPalindromeStr("A man, a plan, a canal: Panama"))
	assert.False(t, isPalindromeStr("race a car"))
	assert.True(t, isPalindromeStr(""))
	assert.True(t, isPalindromeStr(" "))
}
