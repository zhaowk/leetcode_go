package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullJustify(t *testing.T) {
	assert.Equal(t, []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}, fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))

	assert.Equal(t, []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}, fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))

	assert.Equal(t, []string{
		"Listen",
		"to    ",
		"many, ",
		"speak ",
		"to   a",
		"few.  ",
	}, fullJustify([]string{"Listen", "to", "many,", "speak", "to", "a", "few."}, 6))

}
