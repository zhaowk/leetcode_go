package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("hello world"))
}
