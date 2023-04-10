package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	//assert.False(t, canJump([]int{}))
	//assert.True(t, canJump([]int{0}))
	assert.True(t, canJump([]int{2, 0, 0}))
	assert.True(t, canJump([]int{2, 3, 1, 1, 4}))
	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))
	assert.True(t, canJump([]int{3, 0, 8, 2, 0, 0, 1}))
}
