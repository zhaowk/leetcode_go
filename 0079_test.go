package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {

	a := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	assert.True(t, exist(a, "ABCCED"))
	assert.True(t, exist(a, "SEE"))
	assert.False(t, exist(a, "ABCB"))

	assert.True(t, exist([][]byte{{'a'}}, "a"))
}
