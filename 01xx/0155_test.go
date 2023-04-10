package _1xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Pop()
	assert.Equal(t, 3, obj.Top())
	assert.Equal(t, 1, obj.GetMin())
}
