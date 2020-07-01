package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	assert.Equal(t, 0, numDecodings("100"))
	assert.Equal(t, 1, numDecodings("110"))
	assert.Equal(t, 0, numDecodings("0"))
	assert.Equal(t, 1, numDecodings("10"))
	assert.Equal(t, 1, numDecodings("101"))
	assert.Equal(t, 1, numDecodings("1010"))
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("123"))
	assert.Equal(t, 5, numDecodings("1111"))
	assert.Equal(t, 3, numDecodings("226"))

	assert.Equal(t, 1, numDecodings("27"))

}
