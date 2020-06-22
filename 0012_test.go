package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "IV", intToRoman(4))
	assert.Equal(t, "VI", intToRoman(6))
	assert.Equal(t, "VII", intToRoman(7))
	assert.Equal(t, "IX", intToRoman(9))

	for i := 1; i < 3999; i++ {
		assert.Equal(t, i, romanToInt(intToRoman(i)))
	}
}
