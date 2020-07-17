package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	assert.True(t, isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
	assert.False(t, isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	assert.True(t, isBipartite([][]int{{4}, {}, {4}, {4}, {0, 2, 3}}))
	assert.True(t, isBipartite([][]int{{1}, {0}, {4}, {4}, {2, 3}}))
}
