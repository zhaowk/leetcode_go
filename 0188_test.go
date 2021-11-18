package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit1(t *testing.T) {
	assert.Equal(t, 2, maxProfit1(2, []int{2, 4, 1}))
}
