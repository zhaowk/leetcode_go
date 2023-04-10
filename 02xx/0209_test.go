package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	assert.Equal(t, 0, minSubArrayLen(7, []int{}))
	assert.Equal(t, 1, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3, 7}))
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3, 2, 3, 1, 5, 4}))
	assert.Equal(t, 3, minSubArrayLen(8, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 3, minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
