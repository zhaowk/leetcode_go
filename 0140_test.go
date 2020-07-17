package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestWordBreak1(t *testing.T) {
	assert.Equal(t, []string{}, wordBreak1("", []string{}))
	assert.Equal(t, []string{"leet code"}, wordBreak1("leetcode", []string{"leet", "code"}))
	assert.Equal(t, []string{"apple pen apple"}, wordBreak1("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, []string{"ca rs"}, wordBreak1("cars", []string{"car", "ca", "rs"}))
	assert.Equal(t, []string{}, wordBreak1("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	assert.Equal(t, []string{"cat sand dog", "cats and dog"}, wordBreak1("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}))

	a := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	b := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	assert.Equal(t, []string{}, wordBreak1(a, b))
}
