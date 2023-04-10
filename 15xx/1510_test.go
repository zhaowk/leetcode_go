package _5xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinnerSquareGame(t *testing.T) {
	assert.True(t, winnerSquareGame(1))
	assert.True(t, winnerSquareGame(4))
	assert.False(t, winnerSquareGame(7))
	assert.True(t, winnerSquareGame(8))
}
