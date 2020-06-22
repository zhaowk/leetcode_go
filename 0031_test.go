package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	a := []int{1, 1, 5}
	nextPermutation(a)
	assert.Equal(t, []int{1, 5, 1}, a)

	b := []int{1, 2, 3}
	nextPermutation(b)
	assert.Equal(t, []int{1, 3, 2}, b)

	c := []int{3, 2, 1}
	nextPermutation(c)
	assert.Equal(t, []int{1, 2, 3}, c)

	d := []int{1, 3, 2}
	nextPermutation(d)
	assert.Equal(t, []int{2, 1, 3}, d)

	e := []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
	nextPermutation(e)
	assert.Equal(t, []int{2, 3, 1, 2, 2, 2, 4, 5, 7}, e)
}
