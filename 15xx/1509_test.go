package _5xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDifference(t *testing.T) {
	assert.Equal(t, 0, minDifference([]int{5, 3, 2, 4}))
	assert.Equal(t, 1, minDifference([]int{1, 5, 0, 10, 14}))
	assert.Equal(t, 2, minDifference([]int{6, 6, 0, 1, 1, 4, 6}))
	assert.Equal(t, 1, minDifference([]int{1, 5, 6, 14, 15}))
}
