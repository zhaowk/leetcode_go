package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	assert.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, partitions("aab"))
	assert.Equal(t, [][]string{{"a", "a", "b", "b"}, {"a", "a", "bb"}, {"aa", "b", "b"}, {"aa", "bb"}}, partitions("aabb"))
}
