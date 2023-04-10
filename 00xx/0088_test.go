package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge1(t *testing.T) {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge1(a, 3, b, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, a)

	a = []int{4, 7, 9, 0, 0, 0}
	b = []int{2, 5, 6}
	merge1(a, 3, b, 3)
	assert.Equal(t, []int{2, 4, 5, 6, 7, 9}, a)

	a = []int{4, 7, 9, 0, 0, 0}
	b = []int{1, 2, 3}
	merge1(a, 3, b, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 7, 9}, a)
}
