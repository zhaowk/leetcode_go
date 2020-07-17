package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandy(t *testing.T) {
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
	assert.Equal(t, 4, candy([]int{1, 2, 2}))
	assert.Equal(t, 9, candy([]int{1, 2, 3, 2, 1}))
	assert.Equal(t, 16, candy([]int{1, 2, 3, 1, 4, 1, 5, 10, 1}))
}
