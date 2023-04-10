package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 10/1, divide(10, 1))
	assert.Equal(t, 10/2, divide(10, 2))
	assert.Equal(t, 10/3, divide(10, 3))
	assert.Equal(t, 10/4, divide(10, 4))
	assert.Equal(t, 10/5, divide(10, 5))
	assert.Equal(t, 10/6, divide(10, 6))
	assert.Equal(t, 10/7, divide(10, 7))
	assert.Equal(t, 10/8, divide(10, 8))
	assert.Equal(t, 10/9, divide(10, 9))
	assert.Equal(t, 10/10, divide(10, 10))
	assert.Equal(t, 10/11, divide(10, 11))
	assert.Equal(t, 10/12, divide(10, 12))

	assert.Equal(t, -10/3, divide(-10, 3))
	assert.Equal(t, 10/-3, divide(10, -3))
	assert.Equal(t, 0/4, divide(0, 4))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
