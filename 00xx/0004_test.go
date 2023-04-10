package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2, 2}))
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3, 5}, []int{2, 2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 2}))
}
