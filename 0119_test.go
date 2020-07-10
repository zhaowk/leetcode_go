package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{1, 1}, getRow(1))
}
