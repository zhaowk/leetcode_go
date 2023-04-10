package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestDeleteDuplicates2(t *testing.T) {
	assert.Equal(t, BuildList(1, 2, 5), deleteDuplicates2(BuildList(1, 2, 3, 3, 4, 4, 5)))
	assert.Equal(t, BuildList(2, 3), deleteDuplicates2(BuildList(1, 1, 1, 2, 3)))
}
