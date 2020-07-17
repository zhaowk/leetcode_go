package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPoint(t *testing.T) {
	assert.Equal(t, 0, maxPoints([][]int{}))
	assert.Equal(t, 1, maxPoints([][]int{{0, 0}}))
	assert.Equal(t, 2, maxPoints([][]int{{0, 0}, {1, 2}}))
	assert.Equal(t, 3, maxPoints([][]int{{0, 0}, {1, 2}, {0, 0}}))
	assert.Equal(t, 3, maxPoints([][]int{{0, 0}, {0, 0}, {0, 0}}))
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	assert.Equal(t, 4, maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
	assert.Equal(t, 5, maxPoints([][]int{{2, 1}, {3, 1}, {0, 1}, {1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
	assert.Equal(t, 6, maxPoints([][]int{{1, 2}, {1, 3}, {1, 5}, {1, 0}, {1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
	//[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]
}
