package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "a good example", reverseWords("example   good a"))
}
