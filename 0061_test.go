package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateRight(t *testing.T) {
	assert.Equal(t, buildList(4, 5, 1, 2, 3), rotateRight(buildList(1, 2, 3, 4, 5), 2))

	assert.Equal(t, buildList(0, 1, 2), rotateRight(buildList(0, 1, 2), 0))
	assert.Equal(t, buildList(2, 0, 1), rotateRight(buildList(0, 1, 2), 1))
	assert.Equal(t, buildList(1, 2, 0), rotateRight(buildList(0, 1, 2), 2))
	assert.Equal(t, buildList(0, 1, 2), rotateRight(buildList(0, 1, 2), 3))
	assert.Equal(t, buildList(2, 0, 1), rotateRight(buildList(0, 1, 2), 4))
	assert.Equal(t, buildList(1, 2, 0), rotateRight(buildList(0, 1, 2), 5))

	assert.Equal(t, buildList(1), rotateRight(buildList(1), 1))
	assert.Equal(t, buildList(1), rotateRight(buildList(1), 2))
	assert.Equal(t, buildList(1), rotateRight(buildList(1), 3))
	assert.Equal(t, buildList(1), rotateRight(buildList(1), 4))
	assert.Equal(t, buildList(1), rotateRight(buildList(1), 5))
}
