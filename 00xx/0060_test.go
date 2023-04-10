package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	assert.Equal(t, "1", getPermutation(1, 1))
	assert.Equal(t, "12", getPermutation(2, 1))
	assert.Equal(t, "21", getPermutation(2, 2))
	assert.Equal(t, "123", getPermutation(3, 1))
	assert.Equal(t, "132", getPermutation(3, 2))
	assert.Equal(t, "213", getPermutation(3, 3))
	assert.Equal(t, "231", getPermutation(3, 4))
	assert.Equal(t, "312", getPermutation(3, 5))
	assert.Equal(t, "321", getPermutation(3, 6))
	assert.Equal(t, "123456789", getPermutation(9, 1))
	assert.Equal(t, "123456798", getPermutation(9, 2))
}
