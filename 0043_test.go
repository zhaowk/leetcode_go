package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, "0", multiply("0", "0"))
	assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "56088", multiply("123", "456"))
	assert.Equal(t, "1000000", multiply("1000", "1000"))
}
