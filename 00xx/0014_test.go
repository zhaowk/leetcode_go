package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
