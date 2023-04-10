package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestRotateRight(t *testing.T) {
	assert.Equal(t, BuildList(4, 5, 1, 2, 3), rotateRight(BuildList(1, 2, 3, 4, 5), 2))

	assert.Equal(t, BuildList(0, 1, 2), rotateRight(BuildList(0, 1, 2), 0))
	assert.Equal(t, BuildList(2, 0, 1), rotateRight(BuildList(0, 1, 2), 1))
	assert.Equal(t, BuildList(1, 2, 0), rotateRight(BuildList(0, 1, 2), 2))
	assert.Equal(t, BuildList(0, 1, 2), rotateRight(BuildList(0, 1, 2), 3))
	assert.Equal(t, BuildList(2, 0, 1), rotateRight(BuildList(0, 1, 2), 4))
	assert.Equal(t, BuildList(1, 2, 0), rotateRight(BuildList(0, 1, 2), 5))

	assert.Equal(t, BuildList(1), rotateRight(BuildList(1), 1))
	assert.Equal(t, BuildList(1), rotateRight(BuildList(1), 2))
	assert.Equal(t, BuildList(1), rotateRight(BuildList(1), 3))
	assert.Equal(t, BuildList(1), rotateRight(BuildList(1), 4))
	assert.Equal(t, BuildList(1), rotateRight(BuildList(1), 5))
}
