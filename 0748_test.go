package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	// x >= y
	if x % y == 0 {
		return y
	}
	return gcd(x % y, y)
}

func TestGcd(t *testing.T)  {
	assert.Equal(t, 29, gcd(319, 377))
	assert.Equal(t, 7, gcd(98, 63))
}
