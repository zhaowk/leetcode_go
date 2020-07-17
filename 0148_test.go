package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortList(t *testing.T) {
	assert.Equal(t, buildList(1, 2, 3, 4), sortList(buildList(1, 2, 3, 4)))
	assert.Equal(t, buildList(1, 2, 3, 4), sortList(buildList(4, 2, 3, 1)))
	assert.Equal(t, buildList(1, 2, 3, 4), sortList(buildList(4, 3, 2, 1)))
	assert.Equal(t, buildList(-1, 0, 3, 4, 5), sortList(buildList(-1, 5, 3, 4, 0)))
}
