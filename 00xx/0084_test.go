package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
