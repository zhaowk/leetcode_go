package _5xx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReformatDate(t *testing.T) {
	assert.Equal(t, "2052-10-20", reformatDate("20th Oct 2052"))
	assert.Equal(t, "1933-06-06", reformatDate("6th Jun 1933"))
	assert.Equal(t, "1960-05-26", reformatDate("26th May 1960"))

}
