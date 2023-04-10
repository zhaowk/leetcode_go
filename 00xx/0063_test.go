package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	//assert.Equal(t, 2, uniquePathsWithObstacles([][]int{
	//	{0,0,0},
	//	{0,1,0},
	//	{0,0,0},
	//}))
	//assert.Equal(t, 1, uniquePathsWithObstacles([][]int{
	//	{0,1,0},
	//	{0,1,0},
	//	{0,0,0},
	//}))

	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}))

	assert.Equal(t, 1, uniquePathsWithObstacles([][]int{
		{0, 0},
	}))
}
