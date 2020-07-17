package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	a := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	b := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(a)
	assert.Equal(t, b, a)
	//fmt.Println(a)

	a = [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'},
	}

	b = [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'},
	}

	solve(a)
	assert.Equal(t, b, a)

	//for _, v := range a {
	//	fmt.Println(string(v))
	//}
}
