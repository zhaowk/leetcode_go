package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestSortList(t *testing.T) {
	assert.Equal(t, BuildList(1, 2, 3, 4), sortList(BuildList(1, 2, 3, 4)))
	assert.Equal(t, BuildList(1, 2, 3, 4), sortList(BuildList(4, 2, 3, 1)))
	assert.Equal(t, BuildList(1, 2, 3, 4), sortList(BuildList(4, 3, 2, 1)))
	assert.Equal(t, BuildList(-1, 0, 3, 4, 5), sortList(BuildList(-1, 5, 3, 4, 0)))
}
