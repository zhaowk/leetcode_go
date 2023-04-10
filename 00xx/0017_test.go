package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	a := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
		"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
		"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}

	assert.Equal(t, a, letterCombinations("234"))
}
