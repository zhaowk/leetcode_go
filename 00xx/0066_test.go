package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1, 2, 5}, plusOne([]int{1, 2, 4}))
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{9, 9, 9}, plusOne([]int{9, 9, 8}))
	assert.Equal(t, []int{9, 9, 0}, plusOne([]int{9, 8, 9}))
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
}
