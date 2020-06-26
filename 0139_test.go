package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreak(t *testing.T) {
	assert.True(t, wordBreak("", []string{}))
	assert.True(t, wordBreak("leetcode", []string{"leet", "code"}))
	assert.True(t, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.True(t, wordBreak("cars", []string{"car", "ca", "rs"}))
	assert.False(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	//a := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//b := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}
	//assert.True(t, wordBreak(a,b))
}

func TestWordBreak2(t *testing.T) {
	assert.True(t, wordBreak2("", []string{}))
	assert.True(t, wordBreak2("leetcode", []string{"leet", "code"}))
	assert.True(t, wordBreak2("applepenapple", []string{"apple", "pen"}))
	assert.True(t, wordBreak2("cars", []string{"car", "ca", "rs"}))
	assert.False(t, wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	a := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	b := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	assert.False(t, wordBreak2(a, b))
}
