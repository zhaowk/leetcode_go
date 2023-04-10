package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestReverseKGroup(t *testing.T) {
	assert.Equal(t, BuildList(1, 2, 3, 4, 5, 6),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 1))
	assert.Equal(t, BuildList(2, 1, 4, 3, 6, 5),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 2))
	assert.Equal(t, BuildList(3, 2, 1, 6, 5, 4),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 3))
	assert.Equal(t, BuildList(4, 3, 2, 1, 5, 6),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 4))
	assert.Equal(t, BuildList(5, 4, 3, 2, 1, 6),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 5))
	assert.Equal(t, BuildList(6, 5, 4, 3, 2, 1),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 6))
	assert.Equal(t, BuildList(1, 2, 3, 4, 5, 6),
		reverseKGroup(BuildList(1, 2, 3, 4, 5, 6), 7))
	assert.Equal(t, BuildList(), reverseKGroup(BuildList(), 7))
	assert.Equal(t, BuildList(1), reverseKGroup(BuildList(1), 7))
}
