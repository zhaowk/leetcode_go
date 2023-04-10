package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
