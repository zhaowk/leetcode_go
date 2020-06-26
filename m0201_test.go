package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	assert.Equal(t, buildList(), removeDuplicateNodes(buildList()))
	assert.Equal(t, buildList(1), removeDuplicateNodes(buildList(1)))
	assert.Equal(t, buildList(1, 2), removeDuplicateNodes(buildList(1, 2)))
	assert.Equal(t, buildList(1, 2, 3), removeDuplicateNodes(buildList(1, 2, 3, 2, 1)))
}
