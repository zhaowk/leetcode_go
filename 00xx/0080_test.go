package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, []int{1, 1, 2, 2, 3}, a[:removeDuplicates2(a)])
}
