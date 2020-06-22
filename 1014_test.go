package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScoreSightseeingPair(t *testing.T) {
	assert.Equal(t, 11, MaxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	assert.Equal(t, 13, MaxScoreSightseeingPair([]int{4, 7, 5, 8}))
	assert.Equal(t, 13, MaxScoreSightseeingPair([]int{6, 3, 7, 4, 7, 6, 6, 4, 9}))
}
