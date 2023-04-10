package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrayCode(t *testing.T) {
	assert.Equal(t, []int{0, 1, 3, 2, 6, 7, 5, 4}, grayCode(3))
	assert.Equal(t, []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}, grayCode(4))
}
