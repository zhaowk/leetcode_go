package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverage(t *testing.T) {
	assert.Equal(t, 2000.000, average([]int{1000, 2000, 3000}))
}
