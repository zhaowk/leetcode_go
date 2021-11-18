package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	assert.Equal(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
	assert.Equal(t, "0", largestNumber([]int{0, 0}))
	assert.Equal(t, "88811", largestNumber([]int{8, 88, 11}))
}
