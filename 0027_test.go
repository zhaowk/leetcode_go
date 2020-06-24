package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	a := []int{3, 2, 2, 3}
	b := []int{2, 2}

	c := removeElement(a, 3)
	assert.Equal(t, b, a[:c])

	assert.Equal(t, 0, removeElement([]int{}, 0))
}
