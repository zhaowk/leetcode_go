package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, -1, search([]int{3, 4, 5, 1, 2}, 0))
	assert.Equal(t, 3, search([]int{3, 4, 5, 1, 2}, 1))
	assert.Equal(t, 0, search([]int{3, 4, 5, 1, 2}, 3))
}
