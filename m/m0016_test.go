package m

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatternMatching(t *testing.T) {
	assert.True(t, patternMatching("", ""))
	assert.True(t, patternMatching("a", ""))
	assert.False(t, patternMatching("", "a"))
	assert.False(t, patternMatching("ab", ""))
	assert.True(t, patternMatching("abba", "dogcatcatdog"))
	assert.False(t, patternMatching("abba", "dogcatcatfish"))
	assert.False(t, patternMatching("aaaa", "dogcatcatdog"))
	assert.True(t, patternMatching("abba", "dogdogdogdog"))
}
