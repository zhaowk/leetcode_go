package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	assert.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
	assert.Equal(t, 5, numDistinct("babgbag", "bag"))
	assert.Equal(t, 4, numDistinct("aabb", "ab"))
}
