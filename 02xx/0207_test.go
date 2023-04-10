package _2xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanFinish(t *testing.T) {
	assert.True(t, canFinish(3, [][]int{{1, 0}, {2, 0}}))
}
