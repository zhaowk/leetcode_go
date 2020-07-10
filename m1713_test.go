package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRespace(t *testing.T) {
	assert.Equal(t, 7, respace([]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"))
	assert.Equal(t, 7, respace([]string{"sssjjs", "hschjf", "hhh", "fhjchfcfshhfjhs", "sfh", "jsf",
		"cjschjfscscscsfjcjfcfcfh", "hccccjjfchcffjjshccsjscsc", "chcfjcsshjj", "jh", "h", "f", "s", "jcshs", "jfjssjhsscfc"},
		"sssjjssfshscfjjshsjjsjchffffs"))

	assert.Equal(t, 7, respace([]string{"sssjjs", "h", "f", "s"},
		"sssjjssfshscfjjshsjjsjchffffs"))
}
