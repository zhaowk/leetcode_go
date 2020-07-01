package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
	assert.Equal(t, 10, minDistance("zoologicoarchaeologist", "zoogeologist"))
}
