package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func printBytes(board [][]byte) {
//	for _,b := range board {
//		for _, v := range b {
//			fmt.Print(string(v))
//		}
//		fmt.Println()
//	}
//}

func TestSolveSudoku(t *testing.T) {
	var result = [][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	assert.Equal(t, result, board)
	//printBytes(board)

	board2 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '.'},
	}
	solveSudoku(board2)
	assert.Equal(t, result, board)

	//printBytes(board2)
}
