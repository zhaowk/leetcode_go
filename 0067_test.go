package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "", addBinary("", ""))
	assert.Equal(t, "0", addBinary("0", "0"))
	assert.Equal(t, "1", addBinary("0", "1"))
	assert.Equal(t, "1", addBinary("1", "0"))
	assert.Equal(t, "10", addBinary("1", "1"))
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
