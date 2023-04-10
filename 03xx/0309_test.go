package _3xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit4(t *testing.T) {
	assert.Equal(t, 3, maxProfit4([]int{1, 2, 3, 0, 2}))
}
