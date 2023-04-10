package m

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	assert.Equal(t, BuildList(), removeDuplicateNodes(BuildList()))
	assert.Equal(t, BuildList(1), removeDuplicateNodes(BuildList(1)))
	assert.Equal(t, BuildList(1, 2), removeDuplicateNodes(BuildList(1, 2)))
	assert.Equal(t, BuildList(1, 2, 3), removeDuplicateNodes(BuildList(1, 2, 3, 2, 1)))
}
