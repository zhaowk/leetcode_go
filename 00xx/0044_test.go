package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch2(t *testing.T) {
	assert.True(t, isMatch2("", ""))
	assert.False(t, isMatch2("aa", "a"))
	assert.True(t, isMatch2("aa", "*"))
	assert.False(t, isMatch2("cb", "?a"))
	assert.True(t, isMatch2("adceb", "*a*b"))
	assert.False(t, isMatch2("acdcb", "a*c?b"))
	assert.False(t, isMatch2("mississippi", "m??*ss*?i*pi"))

	assert.False(t, isMatch2(
		"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b"))

}

func TestIsMatch3(t *testing.T) {
	assert.True(t, isMatch3("", ""))
	assert.True(t, isMatch3("a", "a*"))
	assert.False(t, isMatch3("aa", "a"))
	assert.True(t, isMatch3("aa", "*"))
	assert.False(t, isMatch3("cb", "?a"))
	assert.True(t, isMatch3("adceb", "*a*b"))
	assert.False(t, isMatch3("acdcb", "a*c?b"))
	assert.False(t, isMatch3("mississippi", "m??*ss*?i*pi"))
	assert.False(t, isMatch3(
		"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b"))

	s := "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb"
	p := "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"
	assert.False(t, isMatch3(s, p))
}
