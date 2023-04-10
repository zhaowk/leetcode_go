package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrieTree(t *testing.T) {
	trie := Constructor1()

	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))   // 返回 true
	assert.False(t, trie.Search("app"))    // 返回 false
	assert.True(t, trie.StartsWith("app")) // 返回 true

	trie.Insert("app")
	assert.True(t, trie.Search("app"))     // 返回 true
	assert.False(t, trie.Search("apple2")) // 返回 true

}
