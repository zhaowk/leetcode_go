package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Equal(t, buildList(1, 2), deleteDuplicates(buildList(1, 2)))
	assert.Equal(t, buildList(1, 2), deleteDuplicates(buildList(1, 2, 2)))
	assert.Equal(t, buildList(1, 2), deleteDuplicates(buildList(1, 1, 2, 2)))
	assert.Equal(t, buildList(1, 2, 3), deleteDuplicates(buildList(1, 1, 2, 2, 3, 3, 3)))
}
