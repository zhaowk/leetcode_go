package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "", countAndSay(0))
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "11", countAndSay(2))
	assert.Equal(t, "21", countAndSay(3))
	assert.Equal(t, "1211", countAndSay(4))
	assert.Equal(t, "111221", countAndSay(5))
	assert.Equal(t, "312211", countAndSay(6))
}
