package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	//assert.Equal(t, "2", fractionToDecimal(2,1))
	//assert.Equal(t, "0.5", fractionToDecimal(1,2))
	//assert.Equal(t, "0.(3)", fractionToDecimal(1,3))
	//assert.Equal(t, "0.(6)", fractionToDecimal(2,3))
	//assert.Equal(t, "0.(285714)", fractionToDecimal(2,7))
	//assert.Equal(t, "0.(18)", fractionToDecimal(2,11))
	assert.Equal(t, "0.1(6)", fractionToDecimal(1, 6))
}
