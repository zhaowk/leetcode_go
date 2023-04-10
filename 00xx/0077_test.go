package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {1, 3}, {2, 3}}, combine(3, 2))
	assert.Equal(t, [][]int{{1, 2, 3, 4, 5}}, combine(5, 5))
	assert.Equal(t, [][]int{{1}, {2}, {3}, {4}, {5}}, combine(5, 1))
}
