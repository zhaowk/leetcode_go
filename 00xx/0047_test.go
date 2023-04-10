package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
}
