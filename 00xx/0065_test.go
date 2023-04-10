package _0xx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var isNumberTestcaseMaps = map[string]bool{
	"0":         true,
	" 0.1 ":     true,
	"abc":       false,
	"1 a":       false,
	"2e10":      true,
	" -90e3   ": true,
	" 1e":       false,
	"e3":        false,
	" 6e-1":     true,
	" 99e2.5 ":  false,
	"53.5e93":   true,
	" --6 ":     false,
	"-+3":       false,
	"95a54e53":  false,
	".1":        true,
	"3.":        true,
}

func TestIsNumber(t *testing.T) {

	assert.True(t, isNumber("6e-1"))

	for k, v := range isNumberTestcaseMaps {
		assert.Equalf(t, v, isNumber(k), fmt.Sprintf("%t %s", v, k))
	}
}
