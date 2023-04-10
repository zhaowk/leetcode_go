package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	//fmt.Println(generateParenthesis(2))
	//fmt.Println(generateParenthesis(3))
	//fmt.Println(generateParenthesis(4))

	assert.Equal(t, []string{
		"(())", "()()",
	}, generateParenthesis(2))
	assert.Equal(t, []string{
		"((()))", "(()())", "(())()", "()(())", "()()()",
	}, generateParenthesis(3))
}
