package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	assert.Equal(t, 1, totalNQueens(1))
	assert.Equal(t, 0, totalNQueens(2))
	assert.Equal(t, 0, totalNQueens(3))
	assert.Equal(t, 2, totalNQueens(4))
	assert.Equal(t, 10, totalNQueens(5))
	assert.Equal(t, 4, totalNQueens(6))
	assert.Equal(t, 40, totalNQueens(7))
	assert.Equal(t, 92, totalNQueens(8))
	assert.Equal(t, 352, totalNQueens(9))
	assert.Equal(t, 724, totalNQueens(10))
	assert.Equal(t, 2680, totalNQueens(11))
	assert.Equal(t, 14200, totalNQueens(12))

}
