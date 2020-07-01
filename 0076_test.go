package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "", minWindow("", ""))
	assert.Equal(t, "", minWindow("", ""))
	assert.Equal(t, "A", minWindow("A", "A"))
	assert.Equal(t, "AA", minWindow("AA", "AA"))
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "BAC"))

	assert.Equal(t, "", minWindow("ABCD", "BR"))
	assert.Equal(t, "baa", minWindow("bbaa", "aba"))
	assert.Equal(t, "cwae", minWindow("cabwefgewcwaefgcf", "cae"))

}
