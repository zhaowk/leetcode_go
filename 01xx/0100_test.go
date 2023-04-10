package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestIsSameTree(t *testing.T) {
	assert.True(t, isSameTree(BuildTree(1, 2, 3), BuildTree(1, 2, 3)))
	assert.False(t, isSameTree(BuildTree(1, 2, 3), BuildTree(1, 2, nil, 3)))
}
