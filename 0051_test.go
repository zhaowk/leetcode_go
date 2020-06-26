package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, [][]string{{"Q"}}, solveNQueens(1))
	assert.Equal(t, [][]string{}, solveNQueens(2))
	assert.Equal(t, [][]string{}, solveNQueens(3))
	assert.Equal(t, [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}, solveNQueens(4))
}
