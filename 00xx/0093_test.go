package _0xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))
	assert.Equal(t, []string{"0.10.0.10", "0.100.1.0"}, restoreIpAddresses("010010"))
}
