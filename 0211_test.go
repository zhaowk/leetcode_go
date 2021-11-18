package leetcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructorW(t *testing.T) {
	wd := ConstructorW()

	wd.AddWord("bag")

	assert.True(t, wd.Search("bag"))
	assert.False(t, wd.Search("ba"))
	assert.False(t, wd.Search("a"))

	assert.True(t, wd.Search("b.g"))
	assert.False(t, wd.Search(".bg"))
	assert.True(t, wd.Search("ba."))
	assert.True(t, wd.Search(".a."))
	assert.True(t, wd.Search("..."))
}
