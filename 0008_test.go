package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 123, myAtoi("123"))
	assert.Equal(t, 2147483647, myAtoi("12337249875843"))
	assert.Equal(t, -123, myAtoi("-123"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))

	assert.Equal(t, 123, myAtoi("123abcab"))
	assert.Equal(t, 0, myAtoi("ab123"))

	assert.Equal(t, 0, myAtoi("  ab123"))
	assert.Equal(t, 123, myAtoi("  123a"))
	assert.Equal(t, 0, myAtoi("  "))
	assert.Equal(t, 0, myAtoi(""))
}
