package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanArrange(t *testing.T) {
	//assert.True(t, canArrange([]int{1,2,3,4,5,6,7,8,9},5))
	//assert.True(t, canArrange([]int{1,2,3,4,5,6},7))
	//assert.False(t, canArrange([]int{1,2,3,4,5,6},10))
	//assert.True(t, canArrange([]int{-10,10},2))
	assert.False(t, canArrange([]int{3, 8, 17, 2, 5, 6}, 10))
}
