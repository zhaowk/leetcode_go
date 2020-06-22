package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	assert.Equal(t, buildList(1, 2, 3, 4, 5, 6),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 1))
	assert.Equal(t, buildList(2, 1, 4, 3, 6, 5),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 2))
	assert.Equal(t, buildList(3, 2, 1, 6, 5, 4),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 3))
	assert.Equal(t, buildList(4, 3, 2, 1, 5, 6),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 4))
	assert.Equal(t, buildList(5, 4, 3, 2, 1, 6),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 5))
	assert.Equal(t, buildList(6, 5, 4, 3, 2, 1),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 6))
	assert.Equal(t, buildList(1, 2, 3, 4, 5, 6),
		reverseKGroup(buildList(1, 2, 3, 4, 5, 6), 7))
	assert.Equal(t, buildList(), reverseKGroup(buildList(), 7))
	assert.Equal(t, buildList(1), reverseKGroup(buildList(1), 7))
}
