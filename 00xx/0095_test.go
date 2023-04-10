package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	//fmt.Println(generateTrees(2))
	//fmt.Println(generateTrees(3))
	//fmt.Println(generateTrees(4))
	//fmt.Println(generateTrees(5))

	assert.Equal(t, 2, len(generateTrees(2)))
	assert.Equal(t, 5, len(generateTrees(3)))
	assert.Equal(t, 14, len(generateTrees(4)))
	assert.Equal(t, 42, len(generateTrees(5)))
	assert.Equal(t, 132, len(generateTrees(6)))
	assert.Equal(t, 1430, len(generateTrees(8)))
}
