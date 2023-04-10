package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWords(t *testing.T) {
	cases := []struct {
		board  [][]byte
		words  []string
		result []string
	}{
		{
			[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"oath", "eat"},
		},
		{
			[][]byte{{'a', 'b'}, {'c', 'd'}},
			[]string{"abcb"},
			[]string{},
		},
		{
			[][]byte{{'a', 'a'}},
			[]string{"aa"},
			[]string{"aa"},
		},
		{
			[][]byte{{'a', 'a'}},
			[]string{"aaa"},
			[]string{},
		},
		{
			[][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}},
			[]string{"eaafgdcba", "eaabcdgfa"},
			[]string{"eaafgdcba", "eaabcdgfa"},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.result, findWords(c.board, c.words))
	}
}
