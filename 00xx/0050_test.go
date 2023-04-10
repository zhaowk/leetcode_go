package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPower(t *testing.T) {
	assert.Equal(t, 1024.00000, myPow(2.0, 10))
	//assert.Equal(t, 9.26100, myPow(2.1,3))
	assert.Equal(t, 0.2500, myPow(2.0, -2))
}
