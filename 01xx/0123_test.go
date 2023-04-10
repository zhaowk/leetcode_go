package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 6, maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(t, 4, maxProfit3([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit3([]int{7, 6, 4, 3, 1}))
}
