package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestSumNumbers(t *testing.T) {
	assert.Equal(t, 25, sumNumbers(BuildTree(1, 2, 3)))
	assert.Equal(t, 1026, sumNumbers(BuildTree(4, 9, 0, 5, 1)))
	assert.Equal(t, 896, sumNumbers(BuildTree(4, 9, 0, 5, nil, 1)))
}
