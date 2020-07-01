package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates2(t *testing.T) {
	assert.Equal(t, buildList(1, 2, 5), deleteDuplicates2(buildList(1, 2, 3, 3, 4, 4, 5)))
	assert.Equal(t, buildList(2, 3), deleteDuplicates2(buildList(1, 1, 1, 2, 3)))
}
