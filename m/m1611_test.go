package m

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivingBoard(t *testing.T) {
	assert.Equal(t, []int{}, divingBoard(1, 1, 0))
	assert.Equal(t, []int{1}, divingBoard(1, 1, 1))
	assert.Equal(t, []int{2}, divingBoard(1, 1, 2))
	assert.Equal(t, []int{1, 2}, divingBoard(1, 2, 1))
	assert.Equal(t, []int{2, 3, 4}, divingBoard(1, 2, 2))
	assert.Equal(t, []int{2, 4, 6}, divingBoard(1, 3, 2))
}
