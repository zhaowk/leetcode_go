package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadderLength(t *testing.T) {

	assert.Equal(t, 2, ladderLength("a", "c", []string{
		"a", "b", "c",
	}))

	assert.Equal(t, 5, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
