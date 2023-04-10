package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestIsSymmetric(t *testing.T) {
	assert.True(t, isSymmetric(BuildTree()))
	assert.True(t, isSymmetric(BuildTree(1, 2, 2, 3, 4, 4, 3)))
	assert.False(t, isSymmetric(BuildTree(1, 2, 2, nil, 3, nil, 3)))
}
