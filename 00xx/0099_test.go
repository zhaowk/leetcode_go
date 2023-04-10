package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestRecoverTree(t *testing.T) {

	tree := BuildTree(nil)
	recoverTree(tree)
	assert.Equal(t, BuildTree(), tree)

	tree = BuildTree(1, 3, nil, nil, 2)
	recoverTree(tree)
	//fmt.Println(tree)
	assert.Equal(t, BuildTree(3, 1, nil, nil, 2), tree)

	tree = BuildTree(3, 1, 4, nil, nil, 2)
	recoverTree(tree)
	//fmt.Println(tree)
	assert.Equal(t, BuildTree(2, 1, 4, nil, nil, 3), tree)
}
