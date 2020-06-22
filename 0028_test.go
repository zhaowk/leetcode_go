package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, -1, strStr("123456", "13"))
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, 0, strStr("123456", ""))
	assert.Equal(t, -1, strStr("", "13"))
	assert.Equal(t, -1, strStr("aaaaa", "baa"))
}
