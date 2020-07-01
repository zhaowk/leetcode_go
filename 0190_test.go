package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBits(t *testing.T) {
	assert.Equal(t, uint32(964176192), reverseBits(43261596))
}
