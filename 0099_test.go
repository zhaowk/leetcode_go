package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecoverTree(t *testing.T) {

	tree := buildTree(nil)
	recoverTree(tree)
	assert.Equal(t, buildTree(), tree)

	tree = buildTree(1, 3, nil, nil, 2)
	recoverTree(tree)
	//fmt.Println(tree)
	assert.Equal(t, buildTree(3, 1, nil, nil, 2), tree)

	tree = buildTree(3, 1, 4, nil, nil, 2)
	recoverTree(tree)
	//fmt.Println(tree)
	assert.Equal(t, buildTree(2, 1, 4, nil, nil, 3), tree)
}
