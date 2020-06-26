package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
