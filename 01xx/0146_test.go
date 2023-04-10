package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDConstructor(t *testing.T) {
	obj := DConstructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	assert.Equal(t, 1, obj.Get(1))
	obj.Put(3, 3)
	assert.Equal(t, -1, obj.Get(2))
	obj.Put(4, 4)
	assert.Equal(t, -1, obj.Get(1))
	assert.Equal(t, 3, obj.Get(3))
	assert.Equal(t, 4, obj.Get(4))

}
