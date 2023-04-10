package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Equal(t, BuildList(1, 2), deleteDuplicates(BuildList(1, 2)))
	assert.Equal(t, BuildList(1, 2), deleteDuplicates(BuildList(1, 2, 2)))
	assert.Equal(t, BuildList(1, 2), deleteDuplicates(BuildList(1, 1, 2, 2)))
	assert.Equal(t, BuildList(1, 2, 3), deleteDuplicates(BuildList(1, 1, 2, 2, 3, 3, 3)))
}
