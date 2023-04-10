package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(2))
	assert.Equal(t, 1, mySqrt(3))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(5))
	assert.Equal(t, 2, mySqrt(6))
	assert.Equal(t, 2, mySqrt(7))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 3, mySqrt(9))
	assert.Equal(t, 3, mySqrt(10))
	assert.Equal(t, 3, mySqrt(11))

	assert.Equal(t, 9, mySqrt(99))
	assert.Equal(t, 10, mySqrt(100))
	assert.Equal(t, 10, mySqrt(101))

}
