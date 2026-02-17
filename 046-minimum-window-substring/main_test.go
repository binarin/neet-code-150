package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
	assert.Equal(t, "", minWindow("a", "aa"))
	assert.Equal(t, "", minWindow("ab", "A"))
	assert.Equal(t, "cwae", minWindow("cabwefgewcwaefgcf", "cae"))
}
